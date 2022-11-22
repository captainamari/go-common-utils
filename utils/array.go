package utils

import (
	"reflect"
	"strconv"
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

// Join 任意数组类型，通过分隔符连接成字符串
func Join(i interface{}, s string) string {
	if i == nil {
		return ""
	}
	iType := reflect.TypeOf(i)
	if iType.Kind() == reflect.Ptr {
		i = reflect.ValueOf(i).Elem().Interface()
	}
	r := make([]string, 0)
	switch i.(type) {
	case [][]byte:
		temp := i.([][]byte)
		for _, val := range temp {
			r = append(r, string(val))
		}
		return strings.Join(r, s)
	case []string:
		temp := i.([]string)
		return strings.Join(temp, s)

	case []int:
		temp := i.([]int)
		for _, val := range temp {
			r = append(r, strconv.Itoa(val))
		}
		return strings.Join(r, s)
	case []int64:
		temp := i.([]int64)
		for _, val := range temp {
			r = append(r, strconv.FormatInt(val, 10))
		}
		return strings.Join(r, s)
	}
	return ""
}

// AddUniInArray 给数组添加不重复的对象
func AddUniInArray(arrList []string, oneStr ...string) []string {
	if len(oneStr) == 0 {
		return arrList
	}
	if len(arrList) > 0 {
		hasSame := false
		newArrayList := make([]string, 0)
		for _, str := range arrList {
			if ok, _ := InArray(str, newArrayList); ok {
				hasSame = true
				continue
			}
			newArrayList = append(newArrayList, str)
		}
		if hasSame {
			arrList = newArrayList
		}
	}
	for _, str := range oneStr {
		if ok, _ := InArray(str, arrList); ok {
			continue
		}
		arrList = append(arrList, str)
	}
	return arrList
}

// AddUniInArrayInt64 给int64数组添加不重复的数字
func AddUniInArrayInt64(arrList []int64, oneStr ...int64) []int64 {
	if len(arrList) == 0 {
		return arrList
	}
	for _, str := range oneStr {
		if ok, _ := InArray(str, arrList); ok {
			continue
		}
		arrList = append(arrList, str)
	}
	return arrList
}
