package model

import (
	"go-gin/core/database"
	"time"

	"xorm.io/xorm"
)

type PKIncrModel struct {
	Id        int        `json:"id" xorm:"pk autoincr"`
	CreatedAt DateTime `json:"createdAt" xorm:"created"`
	UpdatedAt DateTime `json:"updatedAt" xorm:"updated"`
	DeletedAt DateTime `json:"deletedAt" xorm:"deleted"`
}

type Model struct {
	Id        int        `json:"id"`
	CreatedAt DateTime `json:"createdAt" xorm:"created"`
	UpdatedAt DateTime `json:"updatedAt" xorm:"updated"`
	DeletedAt DateTime `json:"deletedAt" xorm:"deleted"`
}

type DateTime time.Time

// Model time format conversion
// 模型时间格式转换
func (this DateTime) MarshalJSON () ([]byte, error) {

	// Special handling is required when the return time is null
	// 当返回时间为空时，需要进行特殊处理
	if time.Time(this).IsZero() {
		return []byte(`""`), nil
	}
	
	return []byte(`"` + time.Time(this).Format(time.DateTime) + `"`), nil
}

// Database instance
// 数据库实例
func DB () *xorm.Engine {
	
	conn := database.Conn()
	return conn
}
