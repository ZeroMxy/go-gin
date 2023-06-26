package adminService

import (
	"go-gin/app/model"
)



// 菜单列表
func MenuList (name string, status int) *[]model.Menu {

	var menus []model.Menu

	sql := model.DB().Table("menu")

	if name != "" {
		sql = sql.Where("name lile ?", "%" + name + "%")
	}

	if status > 0 {
		sql = sql.Where("status = ?", status)
	}

	sql.Find(&menus)

	return &menus
}

// 菜单详情
func MenuDetail (id, parentId int, name string) *model.Menu {

	var menu model.Menu
	sql := model.DB().Table("menu")

	if id > 0 {
		sql = sql.Where("id = ?", id)
	}

	if parentId > 0 {
		sql = sql.Where("parentId = ?", parentId)
	}

	if name != "" {
		sql = sql.Where("name = ?", name)
	}

	sql.Get(&menu)

	return &menu
}

// 新增菜单
func AddMenu (menu *model.Menu) (bool, error) {

	_, err := model.DB().Table("menu").InsertOne(&menu)
	if err != nil {
		return false, err
	}

	return true, nil
}

// 更新菜单
func UpdateMenu (menu *model.Menu) (bool, error) {

	_, err := model.DB().Table("menu").Where("id = ?", menu.Id).Update(menu)
	if err != nil {
		return false, err
	}

	return true, nil
}

// 删除菜单
func DelMenu (id int) (bool, error) {

	_, err := model.DB().Table("menu").Where("id = ?", id).Delete(&model.Menu{})
	if err != nil {
		return false, err
	}

	return true, nil
}

// 根据角色 ID 查询菜单列表
func MenuListByRoleId (roleId int) *[]model.Menu {

	var menus []model.Menu

	sql := model.DB().Table("roleHasMenu").Select("menu.*").
			Join("inner", "menu", "roleHasMenu.menuId = menu.id")

	if roleId > 0 {
		sql = sql.Where("roleHasMenu.roleId = ?", roleId)
	}

	sql.Find(&menus)

	return &menus
}

// 无限级 tree 类型菜单
func MenuToTree (menus []model.Menu, parentId int) *[]model.MenuChildren {

	var menusTree []model.MenuChildren

	for _, value := range menus {
		// 循环中找到子级
		if value.ParentId == parentId {
			// 获取子级菜单
			children := MenuToTree(menus, value.Id)
			var menuTree model.MenuChildren
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
