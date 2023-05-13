package token

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

type Token struct {}

// Generate token
// 生成 token
func Create (value string) string {
	
	timeString := []byte(strconv.FormatInt(time.Now().Unix(), 10) + value)
	tokenByte := md5.Sum(timeString)

	return hex.EncodeToString((tokenByte[:]))
}