package master

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	peerstore "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	maddr "github.com/multiformats/go-multiaddr"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	. "yu/common"
	"yu/config"
	. "yu/node"
	"yu/tripod"
	. "yu/txn"
	. "yu/yerror"
)

func makeP2pHost(ctx context.Context, cfg *config.MasterConf) (host.Host, error) {
	r, err := loadNodeKeyReader(cfg)
	if err != nil {
		return nil, err
	}
	priv, _, err := crypto.GenerateKeyPairWithReader(cfg.NodeKeyType, cfg.NodeKeyBits, r)
	if err != nil {
		return nil, err
	}
	p2pHost, err := libp2p.New(
		ctx,
		libp2p.Identity(priv),
		libp2p.ListenAddrStrings(cfg.P2pListenAddrs...),
	)
	if err != nil {
		return nil, err
	}

	hostAddr, err := maddr.NewMultiaddr(fmt.Sprintf("/p2p/%s", p2pHost.ID().Pretty()))
	if err != nil {
		return nil, err
	}
	addr := p2pHost.Addrs()[0]
	fullAddr := addr.Encapsulate(hostAddr)
	logrus.Infof("I am %s", fullAddr)

	return p2pHost, nil
}

func loadNodeKeyReader(cfg *config.MasterConf) (io.Reader, error) {
	if cfg.NodeKey != "" {
		return bytes.NewBufferString(cfg.NodeKey), nil
	}
	if cfg.NodeKeyFile != "" {
		return os.Open(cfg.NodeKeyFile)
	}
	return nil, nil
}

func (m *Master) ConnectP2PNetwork(cfg *config.MasterConf) error {
	m.host.SetStreamHandler(protocol.ID(cfg.ProtocolID), m.AcceptShakeHand)

	for _, addrStr := range cfg.ConnectAddrs {
		addr, err := maddr.NewMultiaddr(addrStr)
		if err != nil {
			return err
		}
		peer, err := peerstore.AddrInfoFromP2pAddr(addr)
		if err != nil {
			return err
		}
		err = m.host.Connect(context.Background(), *peer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Master) AcceptShakeHand(s network.Stream) {
	blockRange, err := m.compareNewNodeInfo(s)
	if err != nil {
		logrus.Errorf("compare new node info error: %s", err.Error())
	}
	hsResp := &HandShakeResp{
		Br:  blockRange,
		Err: err,
	}
	byt, err := hsResp.Encode()
	if err != nil {
		logrus.Errorf("encode handshake response error: %s", err.Error())
		return
	}
	_, err = s.Write(byt)
	if err != nil {
		logrus.Errorf("write handshake response error: %s", err.Error())
	}
}

func (m *Master) compareNewNodeInfo(s network.Stream) (*BlocksRange, error) {
	buf := bufio.NewReader(s)
	byt, err := buf.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	remoteInfo, err := DecodeHsInfo(byt)
	if err != nil {
		return nil, err
	}

	localInfo, err := m.NewHsInfo()
	if err != nil {
		return nil, err
	}

	return localInfo.Compare(m.chain.ConvergeType(), remoteInfo)
}

func (m *Master) SendHandShake(rw *bufio.ReadWriter) error {
	hs, err := m.NewHsInfo()
	if err != nil {
		return err
	}
	byt, err := hs.Encode()
	if err != nil {
		return err
	}
	_, err = rw.Write(byt)
	return err
}

func (m *Master) AcceptBlocks() error {
	block, err := m.subBlock()
	if err != nil {
		return err
	}

	switch m.RunMode {
	case MasterWorker:
		// todo: switch MasterWorker Mode
	case LocalNode:
		err = m.land.RangeList(func(tri tripod.Tripod) error {
			if tri.ValidateBlock(m.chain, block) {
				return nil
			}
			return BlockIllegal(block)
		})
		if err != nil {
			return err
		}
	}

	return m.chain.InsertBlockFromP2P(block)
}

func (m *Master) AcceptUnpkgTxns() error {
	txns, err := m.subUnpackedTxns()
	if err != nil {
		return err
	}

	switch m.RunMode {
	case MasterWorker:
		// key: workerIP
		forwardMap := make(map[string]*TxnsAndWorkerName)
		for _, txn := range txns {
			ecall := txn.GetRaw().Ecall()
			tripodName := ecall.TripodName
			execName := ecall.ExecName
			workerIP, workerName, err := m.findWorkerIpAndName(tripodName, execName, ExecCall)
			if err != nil {
				return err
			}
			oldTxns := forwardMap[workerIP].Txns
			forwardMap[workerIP] = &TxnsAndWorkerName{
				Txns:       append(oldTxns, txn),
				WorkerName: workerName,
			}
		}

		err := m.forwardTxnsForCheck(forwardMap)
		if err != nil {
			return err
		}

		for _, twn := range forwardMap {
			err = m.txPool.BatchInsert(twn.WorkerName, twn.Txns)
			if err != nil {
				return err
			}
		}

	case LocalNode:
		err = m.txPool.BatchInsert("", txns)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Master) forwardTxnsForCheck(forwardMap map[string]*TxnsAndWorkerName) error {
	for workerIP, txns := range forwardMap {
		byt, err := txns.Txns.Encode()
		if err != nil {
			return err
		}
		_, err = PostRequest(workerIP+CheckTxnsPath, byt)
		if err != nil {
			return err
		}
	}

	return nil
}

type TxnsAndWorkerName struct {
	Txns       SignedTxns
	WorkerName string
}
