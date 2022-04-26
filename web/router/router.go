package router

import (
	"chenxi/initialize"
	"chenxi/web/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Router() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//注意 Recover 要尽量放在第一个被加载
	//如不是的话，在recover前的中间件或路由，将不能被拦截到
	//程序的原理是：
	//1.请求进来，执行recover
	//2.程序异常，抛出panic
	//3.panic被 recover捕获，返回异常信息，并Abort,终止这次请求
	router.Use(middleware.Recover, middleware.LoggerToFile())
	// 为页面提供的接口，统一走网关调用，限定请求源ip
	page_api := router.Group("/ui/chenxi", middleware.AccessIPWhitelist)
	{

	}
	// 接口api
	api := router.Group("/api/chenxi", middleware.AccessIPWhitelist)
	{

	}
	// 接口注册到网关
	router.Use(middleware.RegistGateway(router.Routes()))
	router.Run(fmt.Sprintf(":%s", initialize.Config.System.ServerPort))
}
