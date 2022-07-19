package history

import (
	peerstore "github.com/libp2p/go-libp2p-core/peer"
	"github.com/sirupsen/logrus"
	. "github.com/yu-org/yu/common"
	. "github.com/yu-org/yu/common/yerror"
	. "github.com/yu-org/yu/core/tripod"
	. "github.com/yu-org/yu/core/types"
)

func (h *History) SyncFullHistory() error {
	logrus.Info("start to sync history from other node")

	resp, err := h.requestBlocks(nil)
	if err != nil {
		return err
	}
	if resp.Err != nil {
		return resp.Err
	}

	for resp.MissingRange != nil {
		// todo: the missing range maybe very huge and we need fetch them multiple times
		// the remote node will return new Missing blocks-range in this response.
		resp, err = h.requestBlocks(resp.MissingRange)
		if err != nil {
			return err
		}

		if resp.Err != nil {
			return resp.Err
		}

		if resp.BlocksByt != nil {
			blocks, err := DecodeBlocks(resp.BlocksByt)
			if err != nil {
				return err
			}

			err = h.syncHistoryBlocks(blocks)
			if err != nil {
				return err
			}

			resp.MissingRange = nil
		}
	}

	return nil
}

func (h *History) syncHistoryBlocks(blocks []*Block) error {
	for _, block := range blocks {
		logrus.Trace("sync history block is ", block.Hash.String())

		err := h.RangeList(func(tri *Tripod) error {
			if tri.VerifyBlock(block) {
				return nil
			}
			return BlockIllegal(block.Hash)
		})
		if err != nil {
			return err
		}

		// todo: sync state trie
		err = h.Chain.AppendBlock(block)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *History) handleHsReq(byt []byte) ([]byte, error) {
	remoteReq, err := DecodeHsRequest(byt)
	if err != nil {
		return nil, err
	}

	var (
		blocksByt []byte
	)
	if remoteReq.FetchRange != nil {
		blocksByt, err = h.getMissingBlocks(remoteReq)
		if err != nil {
			return nil, err
		}
	}

	missingRange, err := h.compareMissingRange(remoteReq.Info)

	if missingRange != nil {
		logrus.Debugf("missing range start-height is %d,  end-height is %d", missingRange.StartHeight, missingRange.EndHeight)
	}

	hsResp := &HandShakeResp{
		MissingRange: missingRange,
		BlocksByt:    blocksByt,
		Err:          err,
	}
	return hsResp.Encode()
}

func (h *History) requestBlocks(fetchRange *BlocksRange) (*HandShakeResp, error) {
	hs, err := h.NewHsReq(fetchRange)
	if err != nil {
		return nil, err
	}

	if hs.FetchRange != nil {
		logrus.Infof("fetch history blocks from (%d) to (%d)", hs.FetchRange.StartHeight, hs.FetchRange.EndHeight)
	}

	byt, err := hs.Encode()
	if err != nil {
		return nil, err
	}

	respByt, err := h.P2pNetwork.RequestPeer(h.P2pNetwork.GetBootNodes()[0], HandshakeCode, byt)
	if err != nil {
		return nil, err
	}
	return DecodeHsResp(respByt)
}

func (h *History) compareMissingRange(remoteInfo *HandShakeInfo) (*BlocksRange, error) {
	localInfo, err := h.NewHsInfo()
	if err != nil {
		return nil, err
	}
	return localInfo.Compare(remoteInfo)
}

func (h *History) getMissingBlocks(remoteReq *HandShakeRequest) ([]byte, error) {
	fetchRange := remoteReq.FetchRange
	blocks, err := h.Chain.GetRangeBlocks(fetchRange.StartHeight, fetchRange.EndHeight)
	if err != nil {
		return nil, err
	}
	return EncodeBlocks(blocks)
}

func (h *History) handleSyncTxnsReq(byt []byte) ([]byte, error) {
	txnsReq, err := DecodeTxnsRequest(byt)
	if err != nil {
		return nil, err
	}
	var (
		txns             SignedTxns
		missingTxnHashes []Hash
	)
	for _, hash := range txnsReq.Hashes {
		stxn, err := h.Pool.GetTxn(hash)
		if err != nil {
			return nil, err
		}

		if stxn != nil {
			txns = append(txns, stxn)
		} else {
			missingTxnHashes = append(missingTxnHashes, hash)
		}
	}

	// request the node of block-producer for missingTxnHashes
	if txnsReq.BlockProducer != h.P2pNetwork.LocalID() {
		stxns, err := h.requestTxns(txnsReq.BlockProducer, txnsReq.BlockProducer, missingTxnHashes)
		if err != nil {
			return nil, err
		}

		txns = append(txns, stxns...)
	}

	var txnsByt []byte
	if txns != nil {
		txnsByt, err = txns.Encode()
		if err != nil {
			return nil, err
		}
	}

	return txnsByt, nil
}

func (h *History) requestTxns(connectPeer, blockProducer peerstore.ID, txnHashes []Hash) (SignedTxns, error) {
	txnsRequest := TxnsRequest{
		Hashes:        txnHashes,
		BlockProducer: blockProducer,
	}
	reqByt, err := txnsRequest.Encode()
	if err != nil {
		return nil, err
	}

	respByt, err := h.P2pNetwork.RequestPeer(connectPeer, SyncTxnsCode, reqByt)
	if err != nil {
		return nil, err
	}
	return DecodeSignedTxns(respByt)
}
