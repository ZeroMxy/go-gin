package user

import (
	"go-gin/app/model"
	"go-gin/app/service/adminService"
	"go-gin/core/session"

	"github.com/gin-gonic/gin"
)

func GetUser (context *gin.Context) *model.AdminRole {

	token := context.Query("token")
	id := session.Get(token).(int)

	adminRole := adminService.AdminDetail(id, "")

	return adminRole
}