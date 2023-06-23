package router

import (
	"go-gin/app/http/controller/toolController"
	"go-gin/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func ToolRoute (app *gin.Engine) {
	
	route := app.Group("tool", (&middleware.Middleware{}).OptLogHandler)

	{
		route.GET("captcha", (&toolController.ToolController{}).Captcha)
	}
	return
}