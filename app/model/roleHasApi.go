package model

type RoleHasApi struct {
	PKIncrModel `xorm:"extends"`
	RoleId      int `json:"roleId"`
	ApiId       int `json:"menuId"`
}
