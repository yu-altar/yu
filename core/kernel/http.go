package kernel

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	. "github.com/yu-org/yu/common"
	. "github.com/yu-org/yu/core"
	"github.com/yu-org/yu/core/context"
	"github.com/yu-org/yu/core/types"
	"net/http"
)

func (k *Kernel) HandleHttp() {
	r := gin.Default()

	// POST request
	r.POST(WrApiPath, func(c *gin.Context) {
		k.handleHttpWr(c)
	})
	// POST request
	r.POST(RdApiPath, func(c *gin.Context) {
		k.handleHttpRd(c)
	})

	err := r.Run(k.httpPort)
	if err != nil {
		logrus.Fatal("serve http failed: ", err)
	}
}

func (k *Kernel) handleHttpWr(c *gin.Context) {
	rawWrCall, err := GetRawWrCall(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err = k.land.GetWriting(rawWrCall.Call.TripodName, rawWrCall.Call.FuncName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	stxn, err := types.NewSignedTxn(rawWrCall.Call, rawWrCall.Pubkey, rawWrCall.Signature)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if k.Pool.Exist(stxn) {
		return
	}

	err = k.Pool.CheckTxn(stxn)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	go func() {
		err = k.pubUnpackedTxns(types.FromArray(stxn))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}()

	err = k.Pool.Insert(stxn)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func (k *Kernel) handleHttpRd(c *gin.Context) {
	rdCall, err := GetRdCall(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	switch k.RunMode {
	case LocalNode:
		ctx, err := context.NewReadContext(c)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		rd, err := k.land.GetReading(rdCall.TripodName, rdCall.FuncName)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		rd(ctx)
	}
}
