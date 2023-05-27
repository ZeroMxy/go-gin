package toolController

import (
	"go-gin/app/tool/captcha"
	"go-gin/core/response"

	"github.com/gin-gonic/gin"
)

type ToolController struct {}

func (*ToolController) Captcha (context *gin.Context) {

	base64Captcha, _ := captcha.Create(context, "storage/font/Zaio.ttf")

	response.Success(context, map[string] string {
		"captcha": base64Captcha,
	})

	return
}

