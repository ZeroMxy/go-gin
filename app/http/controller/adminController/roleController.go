package admincontroller

import (
	"go-gin/app/model"
	"go-gin/app/service/adminService"
	"go-gin/core/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

// 角色列表
func (*RoleController) RoleList (context *gin.Context) {

	current, _ := strconv.Atoi(context.DefaultQuery("current", "1"))
	size, _ := strconv.Atoi(context.DefaultQuery("size", "10"))
	name := context.Query("name")
	status, _ := strconv.Atoi(context.DefaultQuery("status", "-1"))

	var roles []model.Role

	total, _ := adminService.RoleList(name, status).Limit(size, (current - 1) * size).FindAndCount(&roles)

	response.Pager(context, roles, int(total), current, size)
	return
}

// 角色详情
func (*RoleController) RoleDetail (context *gin.Context) {

}

// 新增角色
func (*RoleController) AddRole (context *gin.Context) {
	
}

// 修改角色
func (*RoleController) UpdateRole (context *gin.Context) {
	
}

// 删除角色
func (*RoleController) DelRole (context *gin.Context) {
	
}
