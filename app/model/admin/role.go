package admin

import "go-gin/app/model"

type Role struct {
	model.Model `xorm:"extends"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}
