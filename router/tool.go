package router

import (
	"go-gin/app/http/controller/toolController"

	"github.com/gin-gonic/gin"
)

func ToolRoute (app *gin.Engine) {
	
	toolApp := app.Group("tool")

	{
		toolApp.GET("captcha", (&toolController.ToolController{}).Captcha)
	}
	return
}