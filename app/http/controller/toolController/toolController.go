package toolController

import (
	"go-gin/app/tool/captcha"
	"go-gin/core/log"
	"go-gin/core/response"

	"github.com/gin-gonic/gin"
)

type ToolController struct {}

func (*ToolController) Captcha (context *gin.Context) {

	base64Captcha, captcha := captcha.Create(context, "storage/font/Zaio.ttf")

	log.Info(captcha)

	response.Success(context, map[string] string {
		"captcha": base64Captcha,
	})

	return
}

