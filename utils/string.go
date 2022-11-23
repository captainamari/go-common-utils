package utils

import (
	"math/rand"
	"time"
)

// GetRandomString 生成随机字符串
func GetRandomString(length int) string {
	asciiLetters := "abcdefghijklmnopqrstuvwxyxABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(asciiLetters)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
