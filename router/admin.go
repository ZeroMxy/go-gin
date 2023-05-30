package router

import (
	"go-gin/app/http/controller/adminController"

	"github.com/gin-gonic/gin"
)

func AdminRoute (app *gin.Engine) {

	// 无需登录
	noAuth := app.Group("admin")

	{
		noAuth.POST("admin/login", (&adminController.AdminController{}).AdminLogin) // 后台用户登录
	}

	// 需要登录
	route := app.Group("admin")

	{
		// 用户管理
		route.GET("admin/list", (&adminController.AdminController{}).AdminList) // 后台用户列表
		route.GET("admin/detail", (&adminController.AdminController{}).AdminDetail) // 后台用户详情
		route.POST("admin/add", (&adminController.AdminController{}).AddAdmin) // 新增后台用户
		route.POST("admin/update", (&adminController.AdminController{}).UpdateAdmin) // 更新后台用户
		route.GET("admin/del", (&adminController.AdminController{}).DelAdmin) // 删除后台用户

		// 角色管理
		route.GET("role/list", (&adminController.RoleController{}).RoleList) // 角色列表
		route.GET("role/detail", (&adminController.RoleController{}).RoleDetail) // 角色详情
		route.GET("role/select", (&adminController.RoleController{}).RoleSelect) // 角色下拉
		route.POST("role/add", (&adminController.RoleController{}).AddRole) // 新增角色
		route.POST("role/update", (&adminController.RoleController{}).UpdateRole) // 更新角色
		route.GET("role/del", (&adminController.RoleController{}).DelRole) // 删除角色

		// 菜单管理
		route.GET("menu/list", (&adminController.MenuController{}).MenuList) // 菜单列表
		route.GET("menu/detail", (&adminController.MenuController{}).MenuDetail) // 菜单详情
		route.POST("menu/add", (&adminController.MenuController{}).AddMenu) // 添加菜单
		route.POST("menu/update", (&adminController.MenuController{}).UpdateMenu) // 更新菜单
		route.GET("menu/del", (&adminController.MenuController{}).DelMenu) // 删除菜单
	}
	return
}