package adminController

import (
	"go-gin/app/model"
	"go-gin/app/service/adminService"
	"go-gin/app/tool/cipher"
	"go-gin/app/tool/token"
	"go-gin/app/tool/user"
	"go-gin/core/response"
	"go-gin/core/session"
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

	id, _ := strconv.Atoi(context.Query("id"))
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

	var adminRole model.AdminRole
	adminRole.Id = id
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

	_, err := adminService.UpdateAdmin(&adminRole)
	if err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
	return
}

// 删除后台用户
func (*AdminController) DelAdmin (context *gin.Context) {

	id, _ := strconv.Atoi(context.Query("id"))

	if _, err := adminService.DelAdmin(id); err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
	return
}

// 后台用户登录
func (*AdminController) AdminLogin (context *gin.Context) {

	username := context.Query("username")
	password := context.Query("password")
	captcha := context.Query("captcha")

	if username == "" {
		response.Fail(context, "请输入用户名")
		return
	}

	if password == "" {
		response.Fail(context, "请输入密码")
		return
	}

	if captcha == "" {
		response.Fail(context, "请输入图形验证码")
		return
	}

	if captcha != session.Get("captcha") {
		response.Fail(context, "图形验证码错误")
		return
	}

	admin := adminService.AdminDetail(0, username)
	if admin.Id <= 0 {
		response.Fail(context, "用户不存在")
		return
	}

	if !cipher.Verify(admin.Password, password) {
		response.Fail(context, "密码错误")
		return
	}
	// 创建 token
	token := token.Create(strconv.Itoa(admin.Id))
	// 缓存用户信息
	session.Set(token, admin.Id)

	data := map[string] interface {} {
		"admin": admin,
		"token": token,
	}

	response.Success(context, data)
	return
}

// 后台用户拥有的菜单权限列表
func (*AdminController) AdminMenus (context *gin.Context) {

	adminRole := user.GetUser(context)
	menus := adminService.MenuListByRoleId(adminRole.RoleId)
	menuChildrenList := adminService.MenuToTree(*menus, 0)

	response.Success(context, menuChildrenList)
	return
}