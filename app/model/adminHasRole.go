package model

type AdminHasRole struct {
	PKIncrModel `xorm:"extends"`
	AdminId     int `json:"adminId"`
	RoleId      int `json:"roleId"`
}
