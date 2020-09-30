package router

import (
	"github.com/gin-contrib/cors"
	"github.com/project-zero-233/userapisrv/pkg/handlers"
	"time"

	"github.com/gin-gonic/gin"
)

// InitEngine 初始化engine
func InitEngine(mode string) (engine *gin.Engine, err error) {
	gin.SetMode(mode)

	// 初始化路由
	engine = gin.New()
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"x-xq5-jwt", "Content-Type", "Origin", "Content-Length"},
		ExposeHeaders:    []string{"x-xq5-jwt", "Download-Status"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 初始化中间件
	//userMiddleware(engine, config.GlobalConfig)
	// 添加路由
	addHandler(engine)
	return engine, nil
}

// addHandler 添加路由
func addHandler(engine *gin.Engine) {
	engine.GET("/", handlers.Index)
	engine.HEAD("/", handlers.Index)

	// 客户端分组
	clientGroup(engine)
	// 运营平台分组
	omsGroup(engine)
}

// userMiddleware 设置中间件
func userMiddleware(engine *gin.Engine) {
	//engine.Use(myMiddlewares.Config(config.GlobalConfig))
	//
	//// MiddlewareUserCache
	//usercacheCfg, ok := config.GlobalConfig.RedisServerConf(define.MiddlewareUserCache)
	//if !ok {
	//	panic(fmt.Sprintf("initMiddleware: %v配置不存在\n", define.MiddlewareUserCache))
	//}
	//usercache := usercacheCfg.NewPool(10)
	//engine.Use(middlewares.SetMiddleware(define.MiddlewareUserCache, usercache))
	//
	//// 新增接口限制，限制频繁的请求接口
	//limit := utils.InitLimitClick()
	//engine.Use(myMiddlewares.SetLimitClick(utils.Limit, limit))
}
