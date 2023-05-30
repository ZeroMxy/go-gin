package adminController

import (
	"go-gin/app/model"
	"go-gin/app/service/adminService"
	"go-gin/app/tool/cipher"
	"go-gin/core/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

// 后台用户列表
func (*AdminController) AdminList (context *gin.Context) {

	username := context.Query("username")
	nickname := context.Query("nickname")
	phone := context.Query("phone")
	current, _ := strconv.Atoi(context.DefaultQuery("current", "1"))
	size, _ := strconv.Atoi(context.DefaultQuery("size", "10"))

	var adminList []model.AdminRole
	total, _ := adminService.AdminList(username, nickname, phone).Limit(size, (current - 1) * size).
				FindAndCount(&adminList)


	response.Pager(context, adminList, int(total), current, size)
	return
}

// 后台用户详情
func (*AdminController) AdminDetail (context *gin.Context) {

	id, _ := strconv.Atoi(context.Query("id"))

	admin := adminService.AdminDetail(id, "")

	response.Success(context, admin)
	return
}

// 新增后台用户
func (*AdminController) AddAdmin (context *gin.Context) {

	username := context.Query("username")
	password := context.Query("password")
	nickname := context.Query("nickname")
	phone := context.Query("phone")
	email := context.Query("email")
	gender, _ := strconv.Atoi(context.Query("gender"))
	age, _ := strconv.Atoi(context.Query("age"))
	avatar := context.Query("avatar")
	remark := context.Query("remark")
	roleId, _ := strconv.Atoi(context.Query("roleId"))

	if username == "" {
		response.Fail(context, "请填写用户名")
		return
	}
	if password == "" {
		response.Fail(context, "请填写密码")
		return
	}

	var adminRole model.AdminRole
	adminRole.Username = username
	adminRole.Password = cipher.Encrypt(password)
	adminRole.Nickname = nickname
	adminRole.Phone = phone
	adminRole.Email = email
	adminRole.Gender = gender
	adminRole.Age = age
	adminRole.Avatar = avatar
	adminRole.Remark = remark
	adminRole.RoleId = roleId

	_, err := adminService.AddAdmin(&adminRole)
	if err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
	return
}

// 更新后台用户
func (*AdminController) UpdateAdmin (context *gin.Context) {

}

// 删除后台用户
func (*AdminController) DelAdmin (context *gin.Context) {

}

// 后台用户登录
func (*AdminController) AdminLogin (context *gin.Context) {

}