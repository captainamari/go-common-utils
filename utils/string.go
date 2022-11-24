package utils

import (
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
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

// UnicodeDecodeString 解码unicode
func UnicodeDecodeString(s string) string {
	if s == "" {
		return s
	}
	newStr := make([]string, 0)
	for i := 0; i < len(s); {
		r, n := utf8.DecodeRuneInString(s[i:])
		newStr = append(newStr, fmt.Sprintf("%c", r))
		i += n
	}
	if len(newStr) == 0 {
		return s
	}
	return strings.Join(newStr, "")
}
