package adminService

import (
	"go-gin/app/model"
	"go-gin/core/log"

	"xorm.io/xorm"
)

type RoleMenus struct {
	model.Role
	Menus []MenuChildren `json:"menus"`
}

// 角色列表
func RoleList (name string, status int) *xorm.Session {

	roleSql := model.DB().Table("role")

	if name != "" {
		roleSql.Where("name like ?", "%" + name + "%")
	}

	if status >= 0 {
		roleSql.Where("status = ?", status)
	}

	return roleSql.Asc("sort")
}

// 角色详情
func RoleDetail (id int, name string) *model.Role {

	var role model.Role
	roleSql := model.DB().Table("role")
	
	if id > 0 {
		roleSql.Where("id = ?", id)
	}

	if name != "" {
		roleSql.Where("name like ?", "%" + name + "%")
	}

	result, err := roleSql.Get(&role)
	if err != nil || !result {
		return nil
	}

	return &role
}

// 新增角色
func AddRole (role *model.Role) bool {
	
	affected, err := model.DB().Table("role").Insert(role)

	if err != nil || affected <= 0 {
		log.Error(err)
		return false
	}

	return true
}

// 修改角色
func UpdateRole (role *model.Role) bool {
	
	affected, err := model.DB().Table("role").Where("id = ?", role.Id).Update(role)
	if err != nil || affected <= 0 {
		log.Error(err)
		return false
	}

	return true
}

// 删除角色
func DelRole (id int) bool {
	
	affected, err := model.DB().Table("role").Where("id = ?", id).Delete(&model.Role{})
	if err != nil || affected <= 0 {
		return false
	}

	return true
}