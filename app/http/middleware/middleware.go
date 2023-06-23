package middleware

import (
	"bytes"
	"encoding/json"
	"go-gin/app/model"
	"go-gin/app/service/adminService"
	"go-gin/app/tool/user"
	"go-gin/core/response"
	"go-gin/core/session"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct{}

func (*Middleware) UserHandler (context *gin.Context) {

	token := context.Query("token")
	if token == "" {
		context.Abort()
		response.TokenFail(context)
		return
	}

	if session.Get(token) == nil {
		context.Abort()
		response.TokenFail(context)
		return
	}

	context.Next()
}

type responseWriter struct {
	gin.ResponseWriter
	bodyBuffer *bytes.Buffer
}

func (resp responseWriter) Write (b []byte) (int, error) {

	resp.bodyBuffer.Write(b)
	return resp.ResponseWriter.Write(b)
}

func (*Middleware) OptLogHandler (context *gin.Context) {

	// 排除上传文件
	if !strings.Contains(context.GetHeader("Content-Type"), "multipart/form-data") {
		// 请求
		reqJson, _ := json.Marshal(context.Request.URL.Query())
		req := string(reqJson)
		// 响应
		respWriter := responseWriter {
			ResponseWriter: context.Writer,
			bodyBuffer: &bytes.Buffer{},
		}
		context.Writer = respWriter
		context.Next()

		// var req = ""
		// switch context.Request.Method {
		// 	case "GET", "DELETE":
		// 		reqJson, _ := json.Marshal(context.Request.URL.Query())
		// 		req = string(reqJson)
		// 	case "POST", "PUT":
		// 		reqJson, _ := json.Marshal(context.Request.URL.Query())
		// 		req = string(reqJson)
		// }
		go func (context *gin.Context)  {
			adminService.AddOptLog(&model.OptLog {
				UserId: user.GetUser(context).Id,
				Ip: context.ClientIP(),
				Method: context.Request.Method,
				Status: strconv.Itoa(context.Writer.Status()),
				Path: context.Request.URL.Path,
				Agent: context.Request.UserAgent(),
				Req: strings.ReplaceAll(strings.ReplaceAll(req, "[", ""), "]", ""),
				Resp: respWriter.bodyBuffer.String(),
			})
		}(context)
	}
	
}