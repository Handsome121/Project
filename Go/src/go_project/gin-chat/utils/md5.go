package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 小写
func Md5Encode(str string) string {
	//创建MD5算法
	h := md5.New()
	//写入待加密数据
	h.Write([]byte(str))
	//获取哈希值字符切片
	bts := h.Sum(nil)
	//转化为16进制字符串
	return hex.EncodeToString(bts)
}

// MD5Encode 大写
func MD5Encode(str string) string {
	return strings.ToUpper(Md5Encode(str))
}

// MakePassword 加密
func MakePassword(plainPwd, salt string) string {
	return Md5Encode(plainPwd + salt)
}

// ValidPassword 解密
func ValidPassword(plainPwd, salt string, password string) bool {
	return Md5Encode(plainPwd+salt) == password
}
