package admin

import "go-gin/app/model"

type UserRole struct {
	model.Model
	UserId int `json:"userId"`
	RoleId int `json:"roleId"`
}