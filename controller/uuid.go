package controller

import (
	"github.com/gin-gonic/gin"
	"unsafe"
)

var lower = []byte("0123456789abcdef")

type UUID [16]byte

func (u UUID) String() string {
	data := []byte{
		lower[u[0x0]>>4], lower[u[0x0]&15], lower[u[0x1]>>4], lower[u[0x1]&15],
		lower[u[0x2]>>4], lower[u[0x2]&15], lower[u[0x3]>>4], lower[u[0x3]&15], '-',
		lower[u[0x4]>>4], lower[u[0x4]&15], lower[u[0x5]>>4], lower[u[0x5]&15], '-',
		lower[u[0x6]>>4], lower[u[0x6]&15], lower[u[0x7]>>4], lower[u[0x7]&15], '-',
		lower[u[0x8]>>4], lower[u[0x8]&15], lower[u[0x9]>>4], lower[u[0x9]&15], '-',
		lower[u[0xa]>>4], lower[u[0xa]&15], lower[u[0xb]>>4], lower[u[0xb]&15],
		lower[u[0xc]>>4], lower[u[0xc]&15], lower[u[0xd]>>4], lower[u[0xd]&15],
		lower[u[0xe]>>4], lower[u[0xe]&15], lower[u[0xf]>>4], lower[u[0xf]&15],
	}

	return *(*string)(unsafe.Pointer(&data))
}

func generateRandomUUID() (uuid UUID) {
	sr.MustRead(uuid[:])
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return
}

func SecurityGetUUID(ctx *gin.Context) {
	sendAsText(ctx, generateRandomUUID())
}
