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
func (*AdminController) UserList (context *gin.Context) {

	username := context.Query("username")
	nickname := context.Query("nickname")
	phone := context.Query("phone")
	current, _ := strconv.Atoi(context.DefaultQuery("current", "1"))
	size, _ := strconv.Atoi(context.DefaultQuery("size", "10"))

	var userList []model.UserRole
	total, _ := adminService.UserList(username, nickname, phone).Limit(size, (current - 1) * size).
				FindAndCount(&userList)


	response.Pager(context, userList, int(total), current, size)
	return
}

// 后台用户详情
func (*AdminController) UserDetail (context *gin.Context) {

	id, _ := strconv.Atoi(context.Query("id"))

	user := adminService.UserDetail(id, "")

	response.Success(context, user)
	return
}

// 新增后台用户
func (*AdminController) AddUser (context *gin.Context) {

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

	var userRole model.UserRole
	userRole.Username = username
	userRole.Password = cipher.Encrypt(password)
	userRole.Nickname = nickname
	userRole.Phone = phone
	userRole.Email = email
	userRole.Gender = gender
	userRole.Age = age
	userRole.Avatar = avatar
	userRole.Remark = remark
	userRole.RoleId = roleId

	_, err := adminService.AddUser(&userRole)
	if err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
	return
}

// 更新后台用户
func (*AdminController) UpdateUser (context *gin.Context) {

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

	var userRole model.UserRole
	userRole.Id = id
	userRole.Username = username
	userRole.Password = cipher.Encrypt(password)
	userRole.Nickname = nickname
	userRole.Phone = phone
	userRole.Email = email
	userRole.Gender = gender
	userRole.Age = age
	userRole.Avatar = avatar
	userRole.Remark = remark
	userRole.RoleId = roleId

	_, err := adminService.UpdateUser(&userRole)
	if err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
	return
}

// 删除后台用户
func (*AdminController) DelUser (context *gin.Context) {

	id, _ := strconv.Atoi(context.Query("id"))

	if _, err := adminService.DelUser(id); err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
	return
}

// 后台用户登录
func (*AdminController) Login (context *gin.Context) {

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

	user := adminService.UserDetail(0, username)
	if user.Id <= 0 {
		response.Fail(context, "用户不存在")
		return
	}

	if !cipher.Verify(user.Password, password) {
		response.Fail(context, "密码错误")
		return
	}
	// 创建 token
	token := token.Create(strconv.Itoa(user.Id))
	// 更新最后ip
	user.LastIp = context.ClientIP()
	adminService.UpdateUser(user)
	
	// 缓存用户信息
	session.Set(token, user)

	data := map[string] interface {} {
		"user": user,
		"token": token,
	}

	response.Success(context, data)
	return
}

// 后台用户拥有的菜单权限列表
func (*AdminController) UserMenus (context *gin.Context) {

	adminRole := user.GetUser(context)

	menus := adminService.MenuListByRoleId(adminRole.RoleId)
	menuChildrenList := adminService.MenuToTree(*menus, 0)

	response.Success(context, menuChildrenList)
	return
}