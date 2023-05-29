package model

type UserHasRole struct {
	Model `xorm:"extends"`
	UserId int `json:"userId"`
	RoleId int `json:"roleId"`
}