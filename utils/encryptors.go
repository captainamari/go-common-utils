package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// Md5 计算出md5的值
func Md5(s string) string {
	d := md5.Sum([]byte(s))
	return hex.EncodeToString(d[:])
}

// HmacSha256 hmac
func HmacSha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(hashed.Sum(nil))
}

// Sha256Hex 转换为sha256
func Sha256Hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}
