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

// Database instance
// 数据库实例
func DB () *xorm.Engine {
	
	conn := database.Conn()
	return conn
}

type DateTime time.Time

func (this DateTime) MarshalJSON () ([]byte, error) {

    if time.Time(this).IsZero() {
		return []byte(`""`), nil
	}
	
	return []byte(`"` + time.Time(this).Format(time.DateTime) + `"`), nil
}

func (this *DateTime) UnmarshalJSON (data []byte) (err error) {

    if len(data) == 2 {
        *this = DateTime(time.Time{})
        return
    }

    now, err := time.Parse(`"` + time.DateTime + `"`, string(data))
    *this = DateTime(now)

    return
}
