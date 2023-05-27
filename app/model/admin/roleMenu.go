package admin

import "go-gin/app/model"

type RoleMenu struct {
	model.Model `xorm:"extends"`
	RoleId int `json:"roleId"`
	MenuId int `json:"menuId"`
}