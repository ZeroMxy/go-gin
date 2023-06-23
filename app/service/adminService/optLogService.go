package adminService

import (
	"go-gin/app/model"

	"xorm.io/xorm"
)

// 操作日志列表
func OptLogLits (method, path, status string) *xorm.Session {

	sql := model.DB().Table("optLog").Desc("id")

	if method != "" {
		// collate utf8_general_ci 不区分大小写
		sql = sql.Where("method collate utf8_general_ci = ?", method)
	}
	
	if path != "" {
		sql = sql.Where("path like ?", "%" + path + "%")
	}

	if status != "" {
		sql = sql.Where("status = ?", status)
	}

	return sql
}

// 新增操作日志
func AddOptLog (optLog *model.OptLog) bool {

	if _, err := model.DB().Table("optLog").InsertOne(&optLog); err != nil {
		return false
	}

	return true
}