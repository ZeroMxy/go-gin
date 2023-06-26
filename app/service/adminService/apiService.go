package adminService

import (
	"go-gin/app/model"

	"xorm.io/xorm"
)

// api 列表
func ApiList (path, method, group, description string) *xorm.Session {
	
	sql := model.DB().Table("api").Desc("id")

	if path != "" {
		sql = sql.Where("path like ?", "%" + path + "%")
	}

	if method != "" {
		sql = sql.Where("method = ?", method)
	}

	if group != "" {
		sql = sql.Where("group like ?", "%" + group + "%")
	}

	if description != "" {
		sql = sql.Where("description like ?", "%" + description + "%")
	}

	return sql
}

// api 详情
func ApiDetail (id int) *model.Api {
	
	var api model.Api

	sql := model.DB().Table("api")

	if id > 0 {
		sql = sql.Where("id = ?", id)
	}

	sql.Get(&api)

	return &api
}

// 新增 api
func AddApi (api *model.Api) (bool, error) {

	_, err := model.DB().Table("api").InsertOne(&api)
	if err != nil {
		return false, err
	}

	return true, nil
}

// 更新 api
func UpdateApi (api *model.Api) (bool, error) {
	
	_, err := model.DB().Table("api").Where("id = ?", api.Id).Update(&api)
	if err != nil {
		return false, err
	}

	return true, nil
}

// 删除 api
func DelApi (id int) (bool, error) {
	
	_, err := model.DB().Table("menu").Where("id = ?", id).Delete(&model.Api{})
	if err != nil {
		return false, err
	}

	return true, nil
}

// 根据角色 ID 查询 api
func ApiListByRoleId (roleId int) *[]model.Api {

	var apiList []model.Api

	sql := model.DB().Table("api").Select("api.*").
			Join("inner", "roleHasApi", "roleHasApi.apiId = api.id")
	
	if roleId > 0 {
		sql = sql.Where("roleHasApi.roleId = ?", roleId)
	}

	sql.Find(&apiList)

	return &apiList
}