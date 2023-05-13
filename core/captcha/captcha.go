package captcha

import (
	"bytes"
	"encoding/base64"
	"go-gin/core/session"
	"image/color"
	"image/png"
	"unsafe"

	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
)

type Captcha struct {}

// Create a graphic verification code
// 创建图形验证码
func Create(context *gin.Context, fontPath string) (string, string) {

	captchaObject := captcha.New()
	captchaObject.SetFont(fontPath)
	captchaObject.SetFrontColor(color.Black, color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	captchaImage, captcha := captchaObject.Create(4, captcha.NUM)

	base64Buff := bytes.NewBuffer(nil)
	png.Encode(base64Buff, captchaImage)

	base64Byte := make([]byte, 5000)
	base64.StdEncoding.Encode(base64Byte, base64Buff.Bytes())

	// Remove the unfilled part
	// 删除未填充的部分
	index := bytes.IndexByte(base64Byte, 0)
	base64Image := base64Byte[0:index]

	return "data:image/png;base64," + *(*string)(unsafe.Pointer(&base64Image)), captcha
}

// Verify the graphic verification code
// 验证图形验证码
func Verify(context *gin.Context, captcha string) bool {

	sessionCaptcha := session.Get("captcha")
	if sessionCaptcha == "" || sessionCaptcha != captcha {
		return false
	}

	return true
}