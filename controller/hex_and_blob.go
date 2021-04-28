package controller

import (
	"github.com/gin-gonic/gin"
	"glotsrans/utils/gformat"
)

type HexOrBlobQuery struct {
	Bytes int `form:"bytes"`
	Upper bool `form:"upper"`
}

func GetHex(ctx *gin.Context) {
	var q HexOrBlobQuery
	q.Bytes = 16

	if err := ctx.ShouldBindQuery(&q); err != nil {
		_ = ctx.AbortWithError(400, err)
	} else {
		ctx.Data(200, "text/plain", []byte(gformat.ToHexString(r.ReadByteSlice(q.Bytes), q.Upper)))
	}
}

func GetBlob(ctx *gin.Context) {
	var q HexOrBlobQuery
	q.Bytes = 16

	if err := ctx.ShouldBindQuery(&q); err != nil {
		_ = ctx.AbortWithError(400, err)
	} else {
		ctx.Data(200, "application/octet-stream", r.ReadByteSlice(q.Bytes))
	}
}
