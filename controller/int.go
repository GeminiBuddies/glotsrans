package controller

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"glotsrans/utils/gbinary"
)

var r = gbinary.NewReader(rand.Reader)

func sendAsText(ctx *gin.Context, data interface{}) {
	ctx.String(200, "%v", data)
}

func GetInt32S(ctx *gin.Context) {
	sendAsText(ctx, r.ReadInt32())
}

func GetInt32U(ctx *gin.Context) {
	sendAsText(ctx, r.ReadUInt32())
}
