package upload

import (
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"go-gin/config"

	"github.com/gin-gonic/gin"
)

type Upload struct {}

// The file is uploaded to the local PC
// 文件上传至本地
func Local (context *gin.Context, file *multipart.FileHeader) string {

	fileFolderPath := "storage/upload/"
	_, pathErr := os.Stat(fileFolderPath)

	if os.IsNotExist(pathErr) {
		os.MkdirAll(fileFolderPath, os.ModePerm)
	}

	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	filePath := fileFolderPath + fileName

	context.SaveUploadedFile(file, filePath)

	return config.App["host"] + filePath
}

