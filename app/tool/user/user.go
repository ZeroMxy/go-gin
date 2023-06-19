package user

import (
	"encoding/json"
	"go-gin/app/model"
	"go-gin/core/log"
	"go-gin/core/session"

	"github.com/gin-gonic/gin"
)

func GetAdmin (context *gin.Context) *model.AdminRole {

	token := context.Query("token")
	data := session.Get(token)

	dataByte, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		return nil
	}
	
	var admin model.AdminRole
	// TODO: 反序列化时间失败
	if err := json.Unmarshal(dataByte, &admin); err != nil {
		log.Error(err)
		return nil
	}

	return &admin
}