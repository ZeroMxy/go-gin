package adminService

import (
	"errors"
	"go-gin/app/model"

	"xorm.io/xorm"
)

// 后台用户列表
func AdminList (username, nickname, phone string) *xorm.Session {

	sql := model.DB().Table("admin")

	if username != "" {
		sql.Where("username like ?", "%" + username + "%")
	}

	if nickname != "" {
		sql.Where("nickname like ?", "%" + nickname + "%")
	}

	if phone != "" {
		sql.Where("phone like ?", "%" + phone + "%")
	}

	return sql
}

// 后台用户详情
func AdminDetail (id int, username string) *model.Admin {

	var admin model.Admin

	sql := model.DB().Table("admin")

	if id > 0 {
		sql.Where("id = ?", id)
	}

	if username != "" {
		sql.Where("username = ?", username)
	}

	sql.Get(&admin)

	return &admin
}

// 新增后台用户
func AddAdmin (adminRole *model.AdminRole) (bool, error) {

	session := model.DB().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return false, err
	}

	admin := adminRole.Admin
	_, err := session.Table("admin").Insert(admin)
	if err != nil {
		return false, err
	}

	_, err = session.Table("adminHasRole").Insert(&model.AdminHasRole {
		AdminId: admin.Id,
		RoleId: adminRole.RoleId,
	})
	if err != nil {
		return false, err
	}

	if err := session.Commit(); err != nil {
		return false, err
	}

	return true, nil
}

// 修改后台用户
func UpdateAdmin (adminRole *model.AdminRole) (bool, error) {

	session := model.DB().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return false, err
	}

	admin := adminRole.Admin
	_, err := session.Table("admin").Where("id = ?", admin.Id).Update(admin)
	if err != nil {
		return false, err
	}
	
	var adminHasRole model.AdminHasRole
	model.DB().Table("adminHasRole").Where("adminId = ?", admin.Id).Get(&adminHasRole)
	if adminHasRole.Id <= 0 {
		_, err = session.Table("adminHasRole").Insert(&model.AdminHasRole {
			AdminId: admin.Id,
			RoleId: adminRole.RoleId,
		})
	} else {
		adminHasRole.RoleId = adminRole.RoleId
		_, err = session.Table("adminHasRole").Update(adminHasRole)
	}

	if err != nil {
		return false, err
	}

	if err := session.Commit(); err != nil {
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