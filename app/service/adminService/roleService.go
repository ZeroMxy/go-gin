package adminService

import (
	"go-gin/app/model"

	"xorm.io/xorm"
)

// 角色列表
func RoleList (name string, status int) *xorm.Session {

	sql := model.DB().Table("role").Asc("sort")

	if name != "" {
		sql = sql.Where("name like ?", "%" + name + "%")
	}

	if status >= 0 {
		sql = sql.Where("status = ?", status)
	}

	return sql
}

// 角色详情
func RoleDetail (id int, mark string) *model.Role {

	var role model.Role

	sql := model.DB().Table("role")

	if id > 0 {
		sql = sql.Where("id = ?", id)
	}

	if mark != "" {
		sql = sql.Where("mark = ?", mark)
	}

	sql.Get(&role)

	return &role
}

// 角色下拉
func RoleSelect (name string) *[]model.Role {

	var roleList []model.Role
	sql := model.DB().Table("role").Asc("sort").Cols("id", "name")

	if name != "" {
		sql = sql.Where("name like ?", "%" + name + "%")
	}

	sql.Find(&roleList)

	return &roleList
}

// 新增角色
func AddRole (role *model.Role, roleHasMenus *[]model.RoleHasMenu) (bool, error) {
	session := model.DB().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		session.Rollback()
		return false, err
	}

	_, err := session.Table("role").InsertOne(&role)
	if err != nil {
		session.Rollback()
		return false, err
	}

	if len(*roleHasMenus) > 0 {
		_, err := session.Table("roleHasMenu").Insert(&roleHasMenus)
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

// 更新角色
func UpdateRole (role *model.Role, roleHasMenus *[]model.RoleHasMenu) (bool, error) {

	session := model.DB().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		session.Rollback()
		return false, err
	}

	_, err := session.Table("role").Where("id = ?", role.Id).Update(&role)
	if err != nil {
		session.Rollback()
		return false, err
	}

	// 清空旧权限关联
	_, err = session.Table("roleHasMenu").Where("roleId = ?", role.Id).Delete(&model.RoleHasMenu {})
	if err != nil {
		session.Rollback()
		return false, err
	}

	// 添加新权限关联
	if len(*roleHasMenus) > 0 {
		_, err := session.Table("roleHasMenu").Insert(&roleHasMenus)
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

// 删除角色
func DelRole (id int) (bool, error) {

	session := model.DB().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		session.Rollback()
		return false, err
	}

	_, err := session.Table("role").Where("id = ?", id).Delete(&model.Role {})
	if err != nil {
		session.Rollback()
		return false, err
	}

	// 清空旧权限关联
	_, err = session.Table("roleHasMenu").Where("roleId = ?", id).Delete(&model.RoleHasMenu {})
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