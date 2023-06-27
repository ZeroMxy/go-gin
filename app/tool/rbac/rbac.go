package rbac

import (
	"sync"

	"github.com/mikespook/gorbac"
)

var (
	rbac *gorbac.RBAC
	once sync.Once
)

func New () *gorbac.RBAC {
	
	once.Do(func() {
		new()
	})
	return rbac
}

// 创建 rbac 实例
func new () {
	rbac = gorbac.New()
	return
}

// 检查权限
func CheckApi (roleName, api string) bool {
	return rbac.IsGranted(roleName, gorbac.NewStdPermission(api), nil)
}