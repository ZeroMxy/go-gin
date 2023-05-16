package response

import "github.com/gin-gonic/gin"

type Response struct {}

var (
	// Response successful
	// 响应成功
	SUCCESS    = 1
	// Response failure
	// 响应失败
	FAIL       = 0
	// Invalid token
	// 无效的令牌
	TOKEN_FAIL = -1
)

// Response successful
// 响应成功
func Success (context *gin.Context, data interface{}) {
	
	Result(context, SUCCESS, "ok", data)
	return
}

// Response failure
// 响应失败
func Fail (context *gin.Context, message string) {

	Result(context, FAIL, message, nil)
	return
}

// Invalid token
// 无效的令牌
func TokenFail (context *gin.Context) {

	Result(context, TOKEN_FAIL, "Invalid token", nil)
	return
}

func Pager (context *gin.Context, data interface{}, total, current, size int) {

	data = map[string]interface{}{
		"rows":    data,
		"total":   total,
		"current": current,
		"size":    size,
	}
	Result(context, SUCCESS, "ok", data)
	return
}

// Response result
// 响应结果
func Result (context *gin.Context, code int, message string, data interface{}) {

	context.JSON(200, map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	})
	return
}
