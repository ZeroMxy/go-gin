package adminController

import (
	"go-gin/app/model"
	"go-gin/app/service/adminService"
	"go-gin/core/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OptLogController struct{}

// 操作日志列表
func (*OptLogController) OptLogList (context *gin.Context) {

	current, _ := strconv.Atoi(context.DefaultQuery("current", "1"))
	size, _ := strconv.Atoi(context.DefaultQuery("size", "10"))

	method := context.Query("method")
	path := context.Query("path")
	status := context.Query("status")
	
	var optLogList []model.OptLog
	total, _ := adminService.OptLogLits(method, path, status).Limit(size, (current - 1) * size).
				FindAndCount(&optLogList)

	response.Pager(context, optLogList, int(total), current, size)
	return
}