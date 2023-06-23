package model

type OptLog struct {
	PKIncrModel	`xorm:"extends"`
	UserId int `json:"userId"`
	Ip	   string `json:"ip"`
	Method string `json:"method"`
	Path   string `json:"path"`
	Status string `json:"status"`
	Agent  string `json:"agent"`
	Req    string `json:"req"`
	Resp   string `json:"resp"`
}