package router

import (
	"go-gin/app/http/controller/toolController"

	"github.com/gin-gonic/gin"
)

func ToolRoute(app *gin.Engine) {

	toolApp := app.Group("")

	{
		toolApp.GET("tool/captcha", (&toolController.ToolController{}).Captcha)
	}

}