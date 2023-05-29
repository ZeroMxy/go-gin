package model

type AdminHasRole struct {
	Model   `xorm:"extends"`
	AdminId int `json:"adminId"`
	RoleId  int `json:"roleId"`
}
