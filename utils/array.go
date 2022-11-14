package utils

import (
	"reflect"
	"strings"
)

// Split 字符串，可通过字符串中存在的字符进行分隔
func Split(s string, dlm []string) []string {
	if s == "" {
		return []string{}
	}
	if len(dlm) == 0 {
		return []string{s}
	}
	dlmStr := dlm[0]
	for i, v := range dlm {
		if i == 0 {
			continue
		}
		s = strings.Replace(s, v, dlmStr, -1)
	}
	return strings.Split(s, dlmStr)
}

// InArray val 是否在 arr 中，如果在的话 index 是多少
func InArray(val interface{}, arr interface{}) (exists bool, index int) {
	exists, index = false, -1
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(arr)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}
