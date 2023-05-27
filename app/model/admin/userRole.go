package admin

import "go-gin/app/model"

type UserRole struct {
	model.Model `xorm:"extends"`
	UserId int `json:"userId"`
	RoleId int `json:"roleId"`
}