package router

import (
	"github.com/gin-gonic/gin"
	"glotsrans/controller"
)

func Create() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	buildMainRouter(&r.RouterGroup)
	buildSecRouter(r.Group("/s"))

	return r
}

func buildMainRouter(r *gin.RouterGroup) {
	r.GET("/int32s", controller.GetInt32S)
	r.GET("/int32u", controller.GetInt32U)
	r.GET("/hex", controller.GetHex)
	r.GET("/blob", controller.GetBlob)
}

func buildSecRouter(r *gin.RouterGroup) {
	// r.GET("/hex", controller.GetHexSec)
}
