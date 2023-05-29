package admin

import "go-gin/app/model"

type UserHasRole struct {
	model.Model `xorm:"extends"`
	UserId int `json:"userId"`
	RoleId int `json:"roleId"`
}