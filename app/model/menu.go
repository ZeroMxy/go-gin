package model

type Menu struct {
	PKIncrModel `xorm:"extends"`
	ParentId    int    `json:"parentId"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	Key         string `json:"key"`
	Icon        string `json:"icon"`
	Path        string `json:"path"`
	Redirect    string `json:"redirect"`
	Component   string `json:"component"`
	Remark      string `json:"remark"`
	Status      int    `json:"status"`
}

type MenuChildren struct {
	Menu
	Children []MenuChildren `json:"children"`
}
