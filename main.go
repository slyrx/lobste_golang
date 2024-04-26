package main

import (
	"github.com/gin-gonic/gin"
	"github.com/slyrx/lobste_golang/routes"
)

func main() {
	// 加载配置
	//config.LoadConfig()

	// 创建 Gin 引擎
	r := gin.Default()

	// 注册中间件
	registerMiddlewares(r)

	// 注册路由
	routes.RegisterRoutes(r)

	// 启动服务器
	r.Run()
}

func registerMiddlewares(r *gin.Engine) {
	// 在这里注册中间件
}
