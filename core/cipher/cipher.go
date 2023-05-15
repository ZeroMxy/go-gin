package cipher

import (
	"go-gin/core/log"

	"golang.org/x/crypto/bcrypt"
)

type Cipher struct{}

// Password encryption
// 密码加密
func Encrypt (param string) string {

	paramByte := []byte(param)
	paramHash, hashErr := bcrypt.GenerateFromPassword(paramByte, bcrypt.MinCost)
	if hashErr != nil {
		log.Error(hashErr)
		return ""
	}

	return string(paramHash)
}

// Authentication password
// 验证密码
func Verify (paramEncrypted, param string) bool {

	paramEncryptedByte := []byte(paramEncrypted)
	paramByte := []byte(param)

	hashErr := bcrypt.CompareHashAndPassword(paramEncryptedByte, paramByte)
	if hashErr != nil {
		return false
	}

	return true
}