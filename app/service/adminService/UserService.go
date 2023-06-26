package adminService

import (
	"errors"
	"go-gin/app/model"
	"xorm.io/xorm"
)

// 用户列表
func UserList (username, nickname, phone string) *xorm.Session {

	sql := model.DB().Table("user").Desc("user.id").Omit("user.password").
			Select("user.*, role.id as roleId, role.name as roleName").
			Join("left", "userHasRole", "userHasRole.userId = user.id").
			Join("left", "role", "role.id = userHasRole.roleId").
			Where("user.type = 1")

	if username != "" {
		sql = sql.Where("user.username like ?", "%" + username + "%")
	}

	if nickname != "" {
		sql = sql.Where("user.nickname like ?", "%" + nickname + "%")
	}

	if phone != "" {
		sql = sql.Where("user.phone like ?", "%" + phone + "%")
	}

	return sql
}

// 用户详情
func UserDetail (id int, username string) *model.UserRole {

	var user model.UserRole

	sql := model.DB().Table("user").
			Select("user.*, role.id as roleId, role.name as roleName").
			Join("left", "userHasRole", "userHasRole.userId = user.id").
			Join("left", "role", "role.id = userHasRole.roleId").
			Where("user.type = 1")

	if id > 0 {
		sql = sql.Where("user.id = ?", id)
	}

	if username != "" {
		sql = sql.Where("user.username = ?", username)
	}

	sql.Get(&user)

	return &user
}

// 新增用户
func AddUser (userRole *model.UserRole) (bool, error) {

	session := model.DB().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		session.Rollback()
		return false, err
	}

	user := userRole.User
	_, err := session.Table("user").InsertOne(&user)
	if err != nil {
		session.Rollback()
		return false, err
	}
	
	if userRole.RoleId > 0 {
		_, err = session.Table("userHasRole").InsertOne(&model.UserHasRole {
			UserId: user.Id,
			RoleId: userRole.RoleId,
		})
		if err != nil {
			session.Rollback()
			return false, err
		}
	}

	if err := session.Commit(); err != nil {
		session.Rollback()
		return false, err
	}

	return true, nil
}

// 更新用户
func UpdateUser (userRole *model.UserRole) (bool, error) {

	session := model.DB().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		session.Rollback()
		return false, err
	}

	user := userRole.User
	_, err := session.Table("user").Where("id = ?", user.Id).Update(user)
	if err != nil {
		session.Rollback()
		return false, err
	}
	
	var userHasRole model.UserHasRole
	// 绑定角色
	if userRole.RoleId > 0 {
		model.DB().Table("userHasRole").Where("userId = ?", user.Id).Get(&userHasRole)
		// 未绑定新增
		if userHasRole.Id <= 0 {
			_, err = session.Table("userHasRole").InsertOne(&model.UserHasRole {
				UserId: user.Id,
				RoleId: userRole.RoleId,
			})
			// 已绑定但角色不同，则更新
		} else if userHasRole.RoleId != userRole.RoleId {
			userHasRole.RoleId = userRole.RoleId
			_, err = session.Table("userHasRole").Update(&userHasRole)
		}
	}
	
	if err != nil {
		session.Rollback()
		return false, err
	}

	if err := session.Commit(); err != nil {
		session.Rollback()
		return false, err
	}

	return true, nil
}

// 删除用户
func DelUser (id int) (bool, error) {

	user := UserDetail(id, "")
	if user == nil {
		return false, errors.New("用户不存在")
	}

	session := model.DB().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		session.Rollback()
		return false, err
	}

	_, err := session.Table("user").Where("id = ?", id).Delete(&model.User {})
	if err != nil {
		session.Rollback()
		return false, err
	}

	var userHasRole model.UserHasRole
	model.DB().Table("userHasRole").Where("userId = ?", user.Id).Get(&userHasRole)
	if userHasRole.Id > 0 {
		session.Table("userHasRole").Where("userId = ?", user.Id).Delete(&model.UserHasRole {})
	}

	if err := session.Commit(); err != nil {
		session.Rollback()
		return false, err
	}

	return true, nil
}