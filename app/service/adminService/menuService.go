package adminService

import (
	"go-gin/app/model"
	"go-gin/app/model/admin"
	"go-gin/core/log"
)

// 菜单列表
func MenuList (name string, status int) *[]admin.Menu {

	var menus []admin.Menu

	menusSql := model.DB().Table("menu")

	if name != "" {
		menusSql.Where("name lile ?", "%"+name+"%")
	}

	if status > -1 {
		menusSql.Where("status = ?", status)
	}

	menusSql.Find(&menus)

	return &menus
}

// 菜单详情
func MenuDetail (id, parentId int, name string) *admin.Menu {

	var menu admin.Menu
	menuSql := model.DB().Table("menu")

	if id > 0 {
		menuSql.Where("id = ?", id)
	}

	if parentId > 0 {
		menuSql.Where("parentId = ?", parentId)
	}

	if name != "" {
		menuSql.Where("name = ?", name)
	}

	result, err := menuSql.Get(&menu)
	if err != nil || !result {
		return nil
	}

	return &menu
}

// 新增菜单
func AddMenu (menu *admin.Menu) bool {

	affected, err := model.DB().Table("menu").Insert(&menu)
	log.Info(affected)
	log.Info(err)
	if err != nil || affected <= 0 {
		return false
	}

	return true
}

// 修改菜单
func UpdateMenu (menu *admin.Menu) *admin.Menu {

	affected, err := model.DB().Table("menu").Where("id = ?", menu.Id).Update(&menu)
	if err != nil || affected <= 0 {
		return nil
	}

	return menu
}

// 删除菜单
func DelMenu (id int) bool {

	var menu admin.Menu
	affected, err := model.DB().Table("menu").Where("id = ?", id).Delete(&menu)
	if err != nil || affected <= 0 {
		return false
	}

	return true
}

// 无限级 tree 类型菜单
func MenuToTree (menus []admin.Menu, parentId int) *[]admin.MenuChildren {

	var menusTree []admin.MenuChildren

	for _, value := range menus {
		// 循环中找到子级
		if value.ParentId == parentId {
			// 获取子级菜单
			var children = MenuToTree(menus, value.Id)
			var menuTree admin.MenuChildren
			// 初始化赋值
			menuTree.Id = value.Id
			menuTree.ParentId = value.ParentId
			menuTree.Name = value.Name
			menuTree.Icon = value.Icon
			menuTree.Path = value.Path
			menuTree.Redirect = value.Redirect
			menuTree.Component = value.Component
			menuTree.Key = value.Key
			menuTree.Status = value.Status
			menuTree.CreatedAt = value.CreatedAt
			menuTree.UpdatedAt = value.UpdatedAt
			menuTree.Children = *children
			// 追加至菜单列表
			menusTree = append(menusTree, menuTree)
		}
	}

	return &menusTree
}
