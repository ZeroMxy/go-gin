package model

type Api struct {
	PKIncrModel `xorm:"extends"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Group       string `json:"group"`
	Description string `json:"description"`
}
