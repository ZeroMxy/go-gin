package adminController

import (
	"go-gin/app/model"
	"go-gin/app/service/adminService"
	"go-gin/core/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApiController struct{}

// api 列表
func (*ApiController) ApiList (context *gin.Context) {

	current, _ := strconv.Atoi(context.DefaultQuery("current", "1"))
	size, _ := strconv.Atoi(context.DefaultQuery("size", "10"))

	path := context.Query("path")
	method := context.Query("method")
	group := context.Query("group")
	description := context.Query("description")

	var apiList []model.Api

	total, _ := adminService.ApiList(path, method, group, description).Limit(size, (current - 1) * size).
				FindAndCount(&apiList)

	response.Pager(context, apiList, int(total), current, size)
	return
}

// api 详情
func (*ApiController) ApiDetail (context *gin.Context) {
	
	id, _ := strconv.Atoi(context.Query("id"))

	api := adminService.ApiDetail(id)

	response.Success(context, api)
	return
}

// 新增 api
func (*ApiController) AddApi (context *gin.Context) {

	path := context.Query("path")
	method := context.Query("method")
	group := context.Query("group")
	description := context.Query("description")

	if path == "" {
		response.Fail(context, "请填写路径")
		return
	}

	if method == "" {
		response.Fail(context, "请填写请求方式")
		return
	}

	if group == "" {
		response.Fail(context, "请填写分组")
		return
	}

	if description == "" {
		response.Fail(context, "请填写描述")
		return
	}

	_, err := adminService.AddApi(&model.Api {
		Path: path,
		Method: method,
		Group: group,
		Description: description,
	})

	if err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
	return
}

// 更新 api
func (*ApiController) UpdateApi (context *gin.Context) {

	id, _ := strconv.Atoi(context.Query("id"))
	path := context.Query("path")
	method := context.Query("method")
	group := context.Query("group")
	description := context.Query("description")

	if path == "" {
		response.Fail(context, "请填写路径")
		return
	}

	if method == "" {
		response.Fail(context, "请填写请求方式")
		return
	}

	if group == "" {
		response.Fail(context, "请填写分组")
		return
	}

	if description == "" {
		response.Fail(context, "请填写描述")
		return
	}

	var api = &model.Api {
		Path: path,
		Method: method,
		Group: group,
		Description: description,
	}
	api.Id = id

	_, err := adminService.UpdateApi(api)

	if err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
	return
}

// 删除 api
func (*ApiController) DelApi (context *gin.Context) {

	id, _ := strconv.Atoi(context.Query("id"))

	_, err := adminService.DelApi(id)
	if err != nil {
		response.Fail(context, err.Error())
		return
	}

	response.Success(context, nil)
	return
}
