package adminService

import (
	"go-gin/app/model"
	"go-gin/core/log"
)

type MenuChildren struct {
	model.Menu
	Children []MenuChildren `json:"children"`
}

// 菜单列表
func MenuList (name string, status int) *[]model.Menu {

	var menus []model.Menu

	menusSql := model.DB().Table("menu")

	if name != "" {
		menusSql.Where("name lile ?", "%" + name + "%")
	}

	if status > 0 {
		menusSql.Where("status = ?", status)
	}

	menusSql.Find(&menus)

	return &menus
}

// 菜单详情
func MenuDetail (id, parentId int, name string) *model.Menu {

	var menu model.Menu
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
func AddMenu (menu *model.Menu) bool {

	affected, err := model.DB().Table("menu").Insert(menu)

	if err != nil || affected <= 0 {
		log.Error(err)
		return false
	}

	return true
}

// 修改菜单
func UpdateMenu (menu *model.Menu) bool {

	affected, err := model.DB().Table("menu").Where("id = ?", menu.Id).Update(menu)
	if err != nil || affected <= 0 {
		log.Error(err)
		return false
	}

	return true
}

// 删除菜单
func DelMenu (id int) bool {

	affected, err := model.DB().Table("menu").Where("id = ?", id).Delete(&model.Menu{})
	if err != nil || affected <= 0 {
		return false
	}

	return true
}

// 无限级 tree 类型菜单
func MenuToTree (menus []model.Menu, parentId int) *[]MenuChildren {

	var menusTree []MenuChildren

	for _, value := range menus {
		// 循环中找到子级
		if value.ParentId == parentId {
			// 获取子级菜单
			var children = MenuToTree(menus, value.Id)
			var menuTree MenuChildren
			// 初始化赋值
			menuTree.Id = value.Id
			menuTree.ParentId = value.ParentId
			menuTree.Type = value.Type
			menuTree.Name = value.Name
			menuTree.Icon = value.Icon
			menuTree.Path = value.Path
			menuTree.Redirect = value.Redirect
			menuTree.Component = value.Component
			menuTree.Key = value.Key
			menuTree.Remark = value.Remark
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
