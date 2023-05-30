package model

type RoleHasMenu struct {
	PKIncrModel
	RoleId    int        `json:"roleId"`
	MenuId    int        `json:"menuId"`
}
