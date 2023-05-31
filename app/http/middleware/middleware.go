package middleware

import (
	"go-gin/core/response"
	"go-gin/core/session"

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