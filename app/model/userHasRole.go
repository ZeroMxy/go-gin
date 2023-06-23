package model

type UserHasRole struct {
	PKIncrModel `xorm:"extends"`
	UserId      int `json:"userId"`
	RoleId      int `json:"roleId"`
}
