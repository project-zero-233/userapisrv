package router

import (
	"github.com/gin-gonic/gin"
	"github.com/project-zero-233/userapisrv/pkg/handlers"
)

// 客户端分组
func clientGroup(engine *gin.Engine) {
	router := engine.Group("/v1/client/vipapisrv")
	{
		// 用户任务管理接口
		{
			router.GET("/vip/getUserVip", handlers.Index)
		}
	}
}

// 运营平台分组
func omsGroup(engine *gin.Engine) {
	router := engine.Group("/v1/oms/vipapisrv")
	{
		router.GET("/vip/getUserVip", handlers.Index)
	}
}
