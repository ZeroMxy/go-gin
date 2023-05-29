package admin

import "go-gin/app/model"

type Role struct {
	model.Model `xorm:"extends"`
	Name        string `json:"name"`
	Mark        string `json:"mark"`
	Sort        int    `json:"sort"`
	Reamrk      string `json:"reamrk"`
	Status      int    `json:"status"`
}

