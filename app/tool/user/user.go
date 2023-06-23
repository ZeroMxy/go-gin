package user

import (
	"encoding/json"
	"go-gin/app/model"
	"go-gin/core/log"
	"go-gin/core/session"

	"github.com/gin-gonic/gin"
)

func GetUser (context *gin.Context) *model.UserRole {

	token := context.Query("token")
	data := session.Get(token)

	dataByte, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		return nil
	}
	
	var user model.UserRole
	// TODO: 反序列化时间失败
	if err := json.Unmarshal(dataByte, &user); err != nil {
		log.Error(err)
		return nil
	}

	return &user
}