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
	status, _ := strconv.Atoi(context.DefaultQuery("status", "-1"))
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

	var roleMenus model.RoleMenu

	role := adminService.RoleDetail(id, "", "")
	menus := adminService.MenuListByRoleId(role.Id)
	menuChildren := adminService.MenuToTree(*menus, 0)

	roleMenus.Role = *role
	roleMenus.Menus = *menuChildren

	response.Success(context, roleMenus)
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

	name := context.Query("name")
	mark := context.Query("mark")
	status, _ := strconv.Atoi(context.Query("status"))
	sort, _ := strconv.Atoi(context.Query("sort"))
	remark := context.Query("remark")
	menuIds := context.Query("menuIds")

	if name == "" {
		response.Fail(context, "请填写角色名称")
		return
	}

	if mark == "" {
		response.Fail(context, "请填写角色值")
		return
	}

	roleDetail := adminService.RoleDetail(0, name, mark)
	if roleDetail.Id > 0 {
		response.Fail(context, "角色已存在")
		return
	}

	var role model.Role
	role.Name = name
	role.Mark = mark
	role.Sort = sort
	role.Status = status
	role.Remark = remark

	if _, err := adminService.AddRole(&role, menuIds); err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
	return
}

// 更新角色
func (*RoleController) UpdateRole (context *gin.Context) {

	id, _ := strconv.Atoi(context.Query("id"))
	name := context.Query("name")
	mark := context.Query("mark")
	status, _ := strconv.Atoi(context.Query("status"))
	sort, _ := strconv.Atoi(context.Query("sort"))
	remark := context.Query("remark")
	menuIds := context.Query("menuIds")

	roleDetail := adminService.RoleDetail(0, name, mark)
	if roleDetail.Id > 0 && roleDetail.Id != id {
		response.Fail(context, "角色已存在")
		return
	}

	var role model.Role
	role.Id = id
	role.Name = name
	role.Mark = mark
	role.Sort = sort
	role.Status = status
	role.Remark = remark

	if _, err := adminService.UpdateRole(&role, menuIds); err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
	return
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