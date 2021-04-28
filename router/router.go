package router

import (
	"github.com/gin-gonic/gin"
	"glotsrans/controller"
)

func Create() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/int32s", controller.GetInt32S)
	r.GET("/int32u", controller.GetInt32U)
	r.GET("/hex", controller.GetHex)
	r.GET("/blob", controller.GetBlob)

	return r
}
