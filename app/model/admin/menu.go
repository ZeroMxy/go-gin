package admin

import "go-gin/app/model"

type Menu struct {
	model.Model
	ParentId  int    `json:"parentId"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
	Key       string `json:"key"`
	Icon      string `json:"icon"`
	Path      string `json:"path"`
	Redirect  string `json:"redirect"`
	Component string `json:"component"`
	Status    int    `json:"status"`
}

type MenuChildren struct {
	Menu
	Children []MenuChildren `json:"children"`
}
