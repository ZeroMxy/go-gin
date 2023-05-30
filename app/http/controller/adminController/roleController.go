package adminController

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

	name := context.Query("name")
	status, _ := strconv.Atoi(context.Query("status"))
	current, _ := strconv.Atoi(context.DefaultQuery("current", "1"))
	size, _ := strconv.Atoi(context.DefaultQuery("size", "10"))

	var roleList []model.Role
	total, _ := adminService.RoleList(name, status).Limit(size, (current - 1) * size).
				FindAndCount(&roleList)

	response.Pager(context, roleList, int(total), current, size)
	return
}

// 角色详情
func (*RoleController) RoleDetail (context *gin.Context) {

	id, _ := strconv.Atoi(context.Query("id"))
	mark := context.Query("mark")

	role := adminService.RoleDetail(id, mark)

	response.Success(context, role)
	return
}

// 角色下拉
func (*RoleController) RoleSelect (context *gin.Context) {

	name := context.Query("name")

	roles := adminService.RoleSelect(name)

	response.Success(context, roles)
	return
}

// 新增角色
func (*RoleController) AddRole (context *gin.Context) {

}

// 更新角色
func (*RoleController) UpdateRole (context *gin.Context) {

}

// 删除角色
func (*RoleController) DelRole (context *gin.Context) {

	id, _ := strconv.Atoi(context.Query("id"))

	if _, err := adminService.DelRole(id); err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
	return
}