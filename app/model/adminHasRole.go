package model

type AdminHasRole struct {
	PKIncrModel
	AdminId   int        `json:"adminId"`
	RoleId    int        `json:"roleId"`
}
