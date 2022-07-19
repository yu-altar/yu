package history

import (
	"github.com/sirupsen/logrus"
	. "github.com/yu-org/yu/core/tripod"
)

const (
	Full = iota
	Snapshot
	Light
)

type History struct {
	*Tripod
	mode int
}

func NewHistory(mode int) *History {
	tri := NewTripod("full_history")
	fh := &History{Tripod: tri, mode: mode}
	tri.SetInit(fh)
	tri.SetP2pHandler(HandshakeCode, fh.handleHsReq).SetP2pHandler(SyncTxnsCode, fh.handleSyncTxnsReq)
	return fh
}

func (h *History) InitChain() {
	switch h.mode {
	case Full:
		err := h.SyncFullHistory()
		if err != nil {
			logrus.Panic("sync full history failed, err: ", err)
		}
	case Snapshot:

	case Light:

	}
}
