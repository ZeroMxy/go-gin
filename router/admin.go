package router

import (
	"go-gin/app/http/controller/adminController"

	"github.com/gin-gonic/gin"
)

func AdminRoute(app *gin.Engine) {

	route := app.Group("admin")

	{
		// 用户管理
		route.GET("admin/list", (&adminController.AdminController{}).AdminList) // 后台用户列表
		route.GET("admin/detail", (&adminController.AdminController{}).AdminDetail) // 后台用户详情
		route.POST("admin/add", (&adminController.AdminController{}).AddAdmin) // 后台用户详情

		// 菜单管理
		route.GET("menu/list", (&adminController.MenuController{}).MenuList) // 菜单列表
		route.GET("menu/detail", (&adminController.MenuController{}).MenuDetail) // 菜单详情
		route.POST("menu/add", (&adminController.MenuController{}).AddMenu) // 添加菜单
		route.POST("menu/update", (&adminController.MenuController{}).UpdateMenu) // 更新菜单
		route.GET("menu/del", (&adminController.MenuController{}).DelMenu) // 删除菜单
	}
	return
}