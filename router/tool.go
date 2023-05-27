package router

import (
	"go-gin/app/http/controller/toolController"

	"github.com/gin-gonic/gin"
)

func ToolRoute (app *gin.Engine) {
	
	route := app.Group("tool")

	{
		route.GET("captcha", (&toolController.ToolController{}).Captcha)
	}
	return
}