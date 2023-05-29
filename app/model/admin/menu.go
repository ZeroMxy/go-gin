package admin

import "go-gin/app/model"

type Menu struct {
	model.Model `xorm:"extends"`
	ParentId  int    `json:"parentId"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
	Key       string `json:"key"`
	Icon      string `json:"icon"`
	Path      string `json:"path"`
	Redirect  string `json:"redirect"`
	Component string `json:"component"`
	Remark    string `json:"remark"`
	Status    int    `json:"status"`
}
