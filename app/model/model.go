package model

import (
	"go-gin/core/database"
	"time"

	"xorm.io/xorm"
)

type PKIncrModel struct {
	Id        int       `json:"id" xorm:"pk autoincr"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
	DeletedAt time.Time `json:"deletedAt" xorm:"deleted"`
}

type Model struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"updated"`
	DeletedAt time.Time `json:"deletedAt" xorm:"deleted"`
}

// Database instance
// 数据库实例
func DB () *xorm.Engine {
	conn := database.Conn()
	return conn
}
