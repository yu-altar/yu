package kernel

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	. "github.com/yu-org/yu/common"
	. "github.com/yu-org/yu/core"
	"github.com/yu-org/yu/core/context"
	"github.com/yu-org/yu/core/types"
	"net/http"
)

func (k *Kernel) HandleWS() {
	http.HandleFunc(WrApiPath, func(w http.ResponseWriter, req *http.Request) {
		k.handleWS(w, req, writing)
	})

	http.HandleFunc(RdApiPath, func(w http.ResponseWriter, req *http.Request) {
		k.handleWS(w, req, reading)
	})

	http.HandleFunc(SubResultsPath, func(w http.ResponseWriter, req *http.Request) {
		k.handleWS(w, req, subscription)
	})

	logrus.Panic(http.ListenAndServe(k.wsPort, nil))
}

const (
	reading = iota
	writing
	subscription
)

func (k *Kernel) handleWS(w http.ResponseWriter, req *http.Request, typ int) {
	upgrade := websocket.Upgrader{}
	c, err := upgrade.Upgrade(w, req, nil)
	if err != nil {
		k.errorAndClose(c, err.Error())
		return
	}
	if typ == subscription {
		logrus.Debugf("Register a Subscription(%s)", c.RemoteAddr().String())
		k.sub.Register(c)
		return
	}

	_, params, err := c.ReadMessage()
	if err != nil {
		k.errorAndClose(c, fmt.Sprintf("reading websocket message from client error: %v", err))
		return
	}
	switch typ {
	case writing:
		k.handleWsWr(c, req, string(params))
	case reading:
		k.handleWsRd(c, req, string(params))
	}

}

func (k *Kernel) handleWsWr(c *websocket.Conn, req *http.Request, params string) {
	stxn, err := getWrFromHttp(req, params)
	if err != nil {
		k.errorAndClose(c, fmt.Sprintf("get Writing info from websocket error: %v", err))
		return
	}

	_, err = k.land.GetWriting(stxn.Raw.WrCall)
	if err != nil {
		k.errorAndClose(c, err.Error())
		return
	}

	if k.txPool.Exist(stxn) {
		return
	}

	err = k.txPool.CheckTxn(stxn)
	if err != nil {
		k.errorAndClose(c, err.Error())
		return
	}

	go func() {
		err = k.pubUnpackedTxns(types.FromArray(stxn))
		if err != nil {
			k.errorAndClose(c, fmt.Sprintf("publish Unpacked txn(%s) error: %v", stxn.TxnHash.String(), err))
		}
	}()

	err = k.txPool.Insert(stxn)
	if err != nil {
		k.errorAndClose(c, err.Error())
		return
	}
}

func (k *Kernel) handleWsRd(c *websocket.Conn, req *http.Request, params string) {
	qcall, err := getRdFromHttp(req, params)
	if err != nil {
		k.errorAndClose(c, fmt.Sprintf("get Reading info from websocket error: %v", err))
		return
	}

	switch k.RunMode {
	case LocalNode:
		ctx, err := context.NewReadContext(qcall.Params)
		if err != nil {
			k.errorAndClose(c, fmt.Sprintf("new context error: %s", err.Error()))
			return
		}

		err = k.land.Read(qcall, ctx)
		if err != nil {
			k.errorAndClose(c, FindNoCallStr(qcall.TripodName, qcall.ReadingName, err))
			return
		}
		err = c.WriteMessage(websocket.BinaryMessage, ctx.Response())
		if err != nil {
			logrus.Errorf("response Read result error: %s", err.Error())
		}
	}

}

func (k *Kernel) errorAndClose(c *websocket.Conn, text string) {
	// FIXEME
	c.WriteMessage(websocket.CloseMessage, []byte(text))
	c.Close()
}
