package utils

import (
	"math/rand"
	"strings"
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

// CamelToUnderline AbAb to ab_ab , AbBB to aa_bb
func CamelToUnderline(s string) string {
	//先将连续的大写转换为第一个大写，后面小写，如果后面大写接的小写，则将最近的小写转大写
	num := len(s)
	tempData := make([]string, 0, len(s))
	for i := 0; i < num; i++ {
		d := s[i]
		tempData = append(tempData, string(d))
		if d >= 'A' && d <= 'Z' {
			n := 0
			i = i + 1
			for ; i < num; i++ {
				d2 := s[i]
				if d2 >= 'A' && d2 <= 'Z' {
					n = n + 1
					tempData = append(tempData, strings.ToLower(string(d2)))
				} else {
					//表示是多个大写后的第一个小写
					if n > 0 {
						tempData = append(tempData, strings.ToUpper(string(d2)))
					} else {
						tempData = append(tempData, string(d2))
					}
					break
				}
			}
		}
	}
	s = strings.Join(tempData, "")
	data := make([]byte, 0, len(s)*2)
	j := false
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	res := strings.ToLower(string(data[:]))
	return res
}

// UnderlineToCamel ab_ab to AbAb
func UnderlineToCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
