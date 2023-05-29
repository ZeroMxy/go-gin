package model

type RoleHasMenu struct {
	Model `xorm:"extends"`
	RoleId int `json:"roleId"`
	MenuId int `json:"menuId"`
}