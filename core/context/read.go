package context

import (
	"github.com/gin-gonic/gin"
	"github.com/yu-org/yu/common"
	"github.com/yu-org/yu/core"
	"net/http"
)

type ReadContext struct {
	*gin.Context
}

func NewReadContext(ctx *gin.Context) (*ReadContext, error) {
	return &ReadContext{
		ctx,
	}, nil
}

func (rc *ReadContext) GetBlockHash() common.Hash {
	return common.HexToHash(rc.Query(core.BlockHashKey))
}

func (rc *ReadContext) JsonOk(v any) {
	rc.JSON(http.StatusOK, v)
}

func (rc *ReadContext) StringOk(format string, values ...any) {
	rc.String(http.StatusOK, format, values)
}
