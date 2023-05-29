package adminService

import (
	"go-gin/app/model"
)

type AdminRoleMenus struct {
	model.Admin
	RoleMenus `json:"roleMenus"`
}
