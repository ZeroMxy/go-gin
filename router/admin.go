package router

import (
	"go-gin/app/http/controller/adminController"
	"go-gin/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoute (app *gin.Engine) {

	// 无需登录
	noAuth := app.Group("")

	{
		noAuth.POST("admin/login", (&adminController.UserController{}).Login) // 后台用户登录
	}

	// 需要登录
	route := app.Group("admin", (&middleware.Middleware{}).OptLogHandler, (&middleware.Middleware{}).UserHandler)

	{
		// 用户管理
		route.GET("admin/list", (&adminController.UserController{}).UserList) // 后台用户列表
		route.GET("admin/detail", (&adminController.UserController{}).UserDetail) // 后台用户详情
		route.POST("admin/add", (&adminController.UserController{}).AddUser) // 新增后台用户
		route.POST("admin/update", (&adminController.UserController{}).UpdateUser) // 更新后台用户
		route.GET("admin/del", (&adminController.UserController{}).DelUser) // 删除后台用户
		route.GET("admin/menus", (&adminController.UserController{}).UserMenus) // 后台用户拥有的菜单权限列表

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

		// 操作日志管理
		route.GET("optLog/list", (&adminController.OptLogController{}).OptLogList) // 操作日志列表
	}
	return
}