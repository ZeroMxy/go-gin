package model

type Role struct {
	PKIncrModel `xorm:"extends"`
	Name        string `json:"name"`
	Mark        string `json:"mark"`
	Sort        int    `json:"sort"`
	Remark      string `json:"remark"`
	Status      int    `json:"status"`
}
