package model

type Role struct {
	Model `xorm:"extends"`
	Name        string `json:"name"`
	Mark        string `json:"mark"`
	Sort        int    `json:"sort"`
	Reamrk      string `json:"reamrk"`
	Status      int    `json:"status"`
}

