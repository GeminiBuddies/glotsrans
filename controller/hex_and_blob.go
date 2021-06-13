package controller

import (
	"github.com/gin-gonic/gin"
	"glotsrans/utils/gformat"
)

type HexQuery struct {
	Bytes int  `form:"bytes"`
	Upper bool `form:"upper"`
}

func GetHex(ctx *gin.Context) {
	var q HexQuery
	q.Bytes = 16

	if err := ctx.ShouldBindQuery(&q); err != nil {
		_ = ctx.AbortWithError(400, err)
	} else if q.Bytes > 1048576 {
		ctx.String(429, "too greedy")
	} else {
		ctx.Data(200, "text/plain", []byte(gformat.ToHexString(r.ReadByteSlice(q.Bytes), q.Upper)))
	}
}

type BlobQuery struct {
	Bytes int `form:"bytes"`
}

func GetBlob(ctx *gin.Context) {
	var q BlobQuery
	q.Bytes = 16

	if err := ctx.ShouldBindQuery(&q); err != nil {
		_ = ctx.AbortWithError(400, err)
	} else if q.Bytes > 1048576 {
		ctx.String(429, "too greedy")
	} else {
		ctx.Data(200, "application/octet-stream", r.ReadByteSlice(q.Bytes))
	}
}
