package model

type RoleHasMenu struct {
	PKIncrModel `xorm:"extends"`
	RoleId      int `json:"roleId"`
	MenuId      int `json:"menuId"`
}
