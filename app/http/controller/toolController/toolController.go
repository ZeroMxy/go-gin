package toolController

import (
	"go-gin/core/captcha"
	"go-gin/core/response"
	"go-gin/core/session"

	"github.com/gin-gonic/gin"
)

type ToolController struct {}

func (*ToolController) Captcha(context *gin.Context) {

	base64Captcha, captcha := captcha.Create(context, "storage/font/Zaio.ttf")

	session.Set("captcha", captcha)

	response.Success(context, map[string] string {
		"captcha": base64Captcha,
	})

	return

}

