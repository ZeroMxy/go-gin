package router

import (
	admincontroller "go-gin/app/http/controller/adminController"

	"github.com/gin-gonic/gin"
)

func AdminRoute(app *gin.Engine) {

	route := app.Group("admin")

	{
		// 菜单管理
		route.GET("menu/list", (&admincontroller.MenuController{}).MenuList) // 菜单列表
		route.GET("menu/detail", (&admincontroller.MenuController{}).MenuDetail) // 菜单详情
		route.POST("menu/add", (&admincontroller.MenuController{}).AddMenu) // 添加菜单
		route.POST("menu/update", (&admincontroller.MenuController{}).UpdateMenu) // 更新菜单
		route.GET("menu/del", (&admincontroller.MenuController{}).DelMenu) // 删除菜单
	}
	return
}