package controller

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"glotsrans/utils/gbinary"
	"glotsrans/utils/grand"
)

var sr = gbinary.NewReader(rand.Reader)
var r = gbinary.NewReader(grand.Reader(grand.NewPCGWithSeed(sr.ReadUInt64(), sr.ReadUInt64())))

func sendAsText(ctx *gin.Context, data interface{}) {
	ctx.String(200, "%v", data)
}

func GetInt32S(ctx *gin.Context) {
	sendAsText(ctx, r.ReadInt32())
}

func GetInt32U(ctx *gin.Context) {
	sendAsText(ctx, r.ReadUInt32())
}

func GetInt64S(ctx *gin.Context) {
	sendAsText(ctx, r.ReadInt64())
}

func GetInt64U(ctx *gin.Context) {
	sendAsText(ctx, r.ReadUInt64())
}
