package adminService

import (
	"go-gin/app/model"
)

type UserRoleMenus struct {
	model.User
	RoleMenus `json:"roleMenus"`
}
