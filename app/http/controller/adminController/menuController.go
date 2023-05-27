package admincontroller

import (
	"go-gin/app/model/admin"
	"go-gin/app/service/adminService"
	"go-gin/core/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MenuController struct{}

// 菜单列表
func (*MenuController) MenuList (context *gin.Context) {

	name := context.Query("name")
	status, _ := strconv.Atoi(context.DefaultQuery("status", "-1"))

	menus := adminService.MenuList(name, status)

	menusTree := adminService.MenuToTree(*menus, 0)
	
	response.Success(context, menusTree)
	return 
}

// 菜单详情
func (*MenuController) MenuDetail (context *gin.Context) {

	id := context.GetInt("id")

	menu := adminService.MenuDetail(id, 0, "")
	
	response.Success(context, menu)
	return
}

// 新增菜单
func (*MenuController) AddMenu (context *gin.Context) {

	status, _ 		:= strconv.Atoi(context.DefaultQuery("status", "1"))
	parentId, _ 	:= strconv.Atoi(context.DefaultQuery("parentId", "0"))
	name 			:= context.Query("name")
	icon 			:= context.Query("icon")
	path 			:= context.Query("path")
	redirect 		:= context.Query("redirect")
	component 		:= context.Query("component")
	key 			:= context.Query("key")
	menuType, _ 	:= strconv.Atoi(context.Query("type"))

	if name == "" {
		response.Fail(context, "请添加菜单名")
		return
	}
	if menuType <= 0 {
		response.Fail(context, "请选择菜单类型")
		return
	}

	menuInfo := adminService.MenuDetail(0, 0, name)
	if menuInfo != nil {
		response.Fail(context, "菜单名已存在")
		return
	}

	menu := adminService.AddMenu(&admin.Menu {
		ParentId: 	parentId,
		Name: 		name,
		Type: 		menuType,
		Icon: 		icon,
		Path: 		path,
		Redirect: 	redirect,
		Component: 	component,
		Key: 		key,
		Status: 	status,
	})

	if menu {
		response.Fail(context, "添加失败")
		return
	}

	response.Success(context, nil)
	return
}

// 修改菜单
func (*MenuController) UpdateMenu (context *gin.Context) {

	status, _ 		:= strconv.Atoi(context.DefaultQuery("status", "1"))
	parentId, _ 	:= strconv.Atoi(context.DefaultQuery("parentId", "0"))
	id, _ 			:= strconv.Atoi(context.Query("id"))
	name 			:= context.Query("name")
	icon 			:= context.Query("icon")
	path 			:= context.Query("path")
	redirect 		:= context.Query("redirect")
	component 		:= context.Query("component")
	key 			:= context.Query("key")
	menuType, _ 	:= strconv.Atoi(context.Query("type"))

	if name == "" {
		response.Fail(context, "请添加菜单名")
		return
	}
	if menuType <= 0 {
		response.Fail(context, "请选择菜单类型")
		return
	}

	menuInfo := adminService.MenuDetail(0, 0, name)
	if menuInfo != nil  && menuInfo.Id != id {
		response.Fail(context, "菜单名已存在")
		return
	}

	menu := adminService.UpdateMenu(&admin.Menu {
		ParentId: 	parentId,
		Name: 		name,
		Type: 		menuType,
		Icon: 		icon,
		Path: 		path,
		Redirect: 	redirect,
		Component: 	component,
		Key: 		key,
		Status: 	status,
	})

	if menu.Id <= 0 {
		response.Fail(context, "更新失败")
		return
	}

	response.Success(context, nil)
	return
}

// 删除菜单
func (*MenuController) DelMenu (context *gin.Context) {

	id, _ := strconv.Atoi(context.Query("id"))
	if id <= 0 {
		response.Fail(context, "删除失败")
		return
	}

	var menuChildren = adminService.MenuDetail(0, id, "")
	if menuChildren != nil {
		response.Fail(context, "存在下级菜单")
		return
	}

	var result = adminService.DelMenu(id)
	if !result {
		response.Fail(context, "删除失败")
		return
	}

	response.Success(context, nil)
	return
}