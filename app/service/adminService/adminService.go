package adminService

import (
	"errors"
	"go-gin/app/model"
	"xorm.io/xorm"
)

// 后台用户列表
func AdminList (username, nickname, phone string) *xorm.Session {

	sql := model.DB().Table("admin").Omit("admin.password").
			Select("admin.*, role.id as roleId, role.name as roleName").
			Join("left", "adminHasRole", "adminHasRole.adminId = admin.id").
			Join("left", "role", "role.id = adminHasRole.roleId")

	if username != "" {
		sql = sql.Where("admin.username like ?", "%" + username + "%")
	}

	if nickname != "" {
		sql = sql.Where("admin.nickname like ?", "%" + nickname + "%")
	}

	if phone != "" {
		sql = sql.Where("admin.phone like ?", "%" + phone + "%")
	}

	return sql
}

// 后台用户详情
func AdminDetail (id int, username string) *model.AdminRole {

	var admin model.AdminRole

	sql := model.DB().Table("admin").Omit("admin.password").
			Select("admin.*, role.id as roleId, role.name as roleName").
			Join("left", "adminHasRole", "adminHasRole.adminId = admin.id").
			Join("left", "role", "role.id = adminHasRole.roleId")

	if id > 0 {
		sql = sql.Where("admin.id = ?", id)
	}

	if username != "" {
		sql = sql.Where("admin.username = ?", username)
	}

	sql.Get(&admin)

	return &admin
}

// 新增后台用户
func AddAdmin (adminRole *model.AdminRole) (bool, error) {

	session := model.DB().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		session.Rollback()
		return false, err
	}

	admin := adminRole.Admin
	affecte, err := session.Table("admin").InsertOne(&admin)
	if affecte <= 0 || err != nil {
		session.Rollback()
		return false, err
	}
	
	_, err = session.Table("adminHasRole").InsertOne(&model.AdminHasRole {
		AdminId: admin.Id,
		RoleId: adminRole.RoleId,
	})
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

// 修改后台用户
func UpdateAdmin (adminRole *model.AdminRole) (bool, error) {

	session := model.DB().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		session.Rollback()
		return false, err
	}

	admin := adminRole.Admin
	_, err := session.Table("admin").Where("id = ?", admin.Id).Update(&admin)
	if err != nil {
		session.Rollback()
		return false, err
	}
	
	var adminHasRole model.AdminHasRole
	model.DB().Table("adminHasRole").Where("adminId = ?", admin.Id).Get(&adminHasRole)
	if adminHasRole.Id <= 0 {
		_, err = session.Table("adminHasRole").InsertOne(&model.AdminHasRole {
			AdminId: admin.Id,
			RoleId: adminRole.RoleId,
		})
	} else {
		adminHasRole.RoleId = adminRole.RoleId
		_, err = session.Table("adminHasRole").Update(&adminHasRole)
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

// 删除后台用户
func DelAdmin (id int) (bool, error) {

	if admin := AdminDetail(id, ""); admin == nil {
		return false, errors.New("用户不存在")
	}

	_, err := model.DB().Table("admin").Where("id = ?", id).Delete(&model.Admin {})
	if err != nil {
		return false, err
	}

	return true, nil
}