package provider

import (
	"go-gin/router"

	"github.com/gin-gonic/gin"
)

type Provider struct {}

// Registered routing service
// 注册路由服务
func (*Provider) RouteServer(app *gin.Engine) {
	router.ToolRoute(app)
	return
}