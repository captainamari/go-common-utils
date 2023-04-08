package utils

import (
	"encoding/binary"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ToInt64 转换为int64
func ToInt64(i interface{}) int64 {
	if i == nil {
		return 0
	}
	intTemp, retBool := forceToint64(i)
	if retBool {
		return intTemp
	}
	switch i.(type) {
	case map[string]interface{}:
		return 0
	case string:
		num, err := strconv.Atoi(i.(string))
		if err != nil {
			return 0
		}
		return int64(num)
	case []byte:
		bits := i.([]byte)
		if len(bits) == 8 {
			return int64(binary.LittleEndian.Uint64(bits))
		} else if len(bits) <= 4 {
			num, err := strconv.Atoi(string(bits))
			if err != nil {
				return 0
			}
			return int64(num)
		}
	}
	return 0
}

func forceToint64(i interface{}) (int64, bool) {
	switch i.(type) {
	case uint:
		return int64(i.(uint)), true
	case uint8:
		return int64(i.(uint8)), true
	case uint16:
		return int64(i.(uint16)), true
	case uint32:
		return int64(i.(uint32)), true
	case uint64:
		return int64(i.(uint64)), true
	case int:
		return int64(i.(int)), true
	case int8:
		return int64(i.(int8)), true
	case int16:
		return int64(i.(int16)), true
	case int32:
		return int64(i.(int32)), true
	case int64:
		return i.(int64), true
	case float32:
		return int64(i.(float32)), true
	case float64:
		return int64(i.(float64)), true
	}
	return 0, false
}

// ToString 转换为string
func ToString(str interface{}) string {
	if str == nil {
		return ""
	}
	//如果是指针的话，则需要转换一下
	strValue := reflect.ValueOf(str)
	if strValue.Kind() == reflect.Ptr {
		if strValue.IsNil() {
			return ""
		}
		str = strValue.Elem().Interface()
	}
	switch str.(type) {
	case string:
		return str.(string)
	case []byte:
		return string(str.([]byte))
	case int:
		return strconv.Itoa(str.(int))
	case error:
		err, _ := str.(error)
		return err.Error()
	case int64:
		return strconv.FormatInt(str.(int64), 10)
	//case float64:
	//	return strconv.FormatFloat(str.(float64), 'g', -1, 64)
	case time.Time:
		{
			oneTime := str.(time.Time)
			return oneTime.Format(fullTimeForm)
		}
	}
	json, err := jsoniter.Marshal(str)
	if err == nil {
		str := string(json)
		if len(str) >= 2 { //解决返回字符串首位带"的问题
			match, _ := regexp.MatchString(`^".*"$`, str)
			if match {
				str = str[1 : len(str)-1]
			}
		}
		return str
	}
	return fmt.Sprintf("%s", str)
}

// ToTime 转换为Time
func ToTime(val interface{}) (time.Time, bool) {
	timeRet := time.Time{}
	if val == nil {
		return timeRet, true
	}
	reValue := reflect.ValueOf(val)
	for reValue.Kind() == reflect.Ptr {
		reValue = reValue.Elem()
		if !reValue.IsValid() {
			return timeRet, true
		}
		val = reValue.Interface()
		if val == nil {
			return timeRet, true
		}
		reValue = reflect.ValueOf(val)
	}
	if val == "" {
		return timeRet, true
	}
	if v, ok := val.(time.Time); ok {
		return v, true
	}
	valTemp := ToString(val)
	if timeTemp, ok := toTimeFromString(valTemp); ok {
		return timeTemp, ok
	}
	return timeRet, true
}

func toTimeFromString(v string) (time.Time, bool) {
	tlen := len(v)
	var t time.Time
	var err error
	mcStr := ToString(MicroTime())
	switch tlen {
	case 0:
		t, err = time.Time{}, nil
	case 8:
		t, err = time.ParseInLocation(shortDateForm, v, time.Local)
	case 10:
		if IsNumeric(v) {
			mcInt := ToInt64(v)
			t = time.Unix(mcInt, 0)
			err = nil
			return t, true
		}
		t, err = time.ParseInLocation(fullDateForm, v, time.Local)
	case len(mcStr): //毫秒
		if IsNumeric(v) {
			mcTempStr := v[0 : len(v)-3]
			mcInt := ToInt64(mcTempStr)
			t = time.Unix(mcInt, 0)
			err = nil
			return t, true
		}
	case 19:
		t, err = time.ParseInLocation(fullTimeForm, v, time.Local)
		if err != nil {
			t, err = time.Parse(time.RFC822, v)
		}
	case len("2019-12-10T11:18:18.979878"), len("2019-12-10T11:18:18.9798786"):
		tempArr := strings.Split(v, ".")
		if len(tempArr) == 2 {
			timeTemp := tempArr[0]
			timeTemp = strings.Replace(timeTemp, "T", " ", 1)
			t, err = time.ParseInLocation(fullTimeForm, timeTemp, time.Local)
			if err != nil {
				t, err = time.Parse(time.RFC822, v)
			}
		}
	case len(time.ANSIC):
		t, err = time.Parse(time.ANSIC, v)
	case len(time.UnixDate):
		t, err = time.Parse(time.UnixDate, v)
	case len(time.RubyDate):
		t, err = time.Parse(time.RFC850, v)
		if err != nil {
			t, err = time.Parse(time.RubyDate, v)
		}
	case len(time.RFC822Z):
		t, err = time.Parse(time.RFC822Z, v)
	case len(time.RFC1123):
		t, err = time.Parse(time.RFC1123, v)
	case len(time.RFC1123Z):
		t, err = time.Parse(time.RFC1123Z, v)
	case len(time.RFC3339):
		t, err = time.Parse(time.RFC3339, v)
	case len(time.RFC3339Nano):
		t, err = time.Parse(time.RFC3339Nano, v)
	default:
		if tlen > 19 {
			tempArr := strings.Split(v, ".")
			if len(tempArr) == 2 {
				timeTemp := tempArr[0]
				timeTemp = strings.Replace(timeTemp, "T", " ", 1)
				t, err = time.ParseInLocation(fullTimeForm, timeTemp, time.Local)
				if err == nil {
					break
				}
			}
		}
		t, err = time.Parse(time.RFC1123, v)
	}
	if err != nil {
		return t, false
	}
	return t, true
}

// IsNil 判断是否为空
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	vi := reflect.ValueOf(i)
	kind := vi.Kind()
	if kind == reflect.Ptr ||
		kind == reflect.Chan ||
		kind == reflect.Func ||
		kind == reflect.UnsafePointer ||
		kind == reflect.Map ||
		kind == reflect.Interface ||
		kind == reflect.Slice {
		return vi.IsNil()
	}
	return false
}

// IsTime 是否是时间格式
func IsTime(dateTime string) bool {
	regPattern := "((([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-"
	regPattern += "(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8]))))|"
	regPattern += "((([0-9]{2})(0[48]|[2468][048]|[13579][26])|((0[48]|[2468][048]|[3579][26])00))-02-29))\\s"
	regPattern += "([0-1][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$"
	matched, err := regexp.Match(regPattern, []byte(dateTime))
	if err == nil {
		return matched
	}
	return false
}

// IsEmptyTime 是否为空时间
func IsEmptyTime(timeParam *time.Time) bool {
	if timeParam == nil {
		return true
	}
	nilTime := time.Time{}       //赋零值
	return *timeParam == nilTime //此处即为零值
}

// IsNumeric 是否是数字
func IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		// Trim any whitespace
		str = strings.Trim(str, " \\t\\n\\r")
		if str == "" {
			return false
		}
		return toNumberFromString(str)
	}
	//其它类型的全部转换为字符串来判断
	return IsNumeric(ToString(val))
}

func toNumberFromString(str string) bool {
	if str[0] == '-' || str[0] == '+' {
		if len(str) == 1 {
			return false
		}
		str = str[1:]
	}
	// hex
	if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
		for _, h := range str[2:] {
			if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
				return false
			}
		}
		return true
	}
	// 0-9,Point,Scientific
	p, s, l := 0, 0, len(str)
	for i, v := range str {
		if v == '.' { // Point
			if p > 0 || s > 0 || i+1 == l {
				return false
			}
			p = i
		} else if v == 'e' || v == 'E' { // Scientific
			if i == 0 || s > 0 || i+1 == l {
				return false
			}
			s = i
		} else if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

// UInt32ToIP 将uint32类型转化为ipv4地址
func UInt32ToIP(val uint32) string {
	ipData := net.IPv4(byte(val>>24), byte(val>>16&0xFF), byte(val>>8)&0xFF, byte(val&0xFF))
	return ipData.String()
}

// IPToUInt32 ip转数字
func IPToUInt32(ipnr string) uint32 {
	bits := strings.Split(ipnr, ".")
	if len(bits) == 4 {
		b0, _ := strconv.Atoi(bits[0])
		b1, _ := strconv.Atoi(bits[1])
		b2, _ := strconv.Atoi(bits[2])
		b3, _ := strconv.Atoi(bits[3])
		var sum uint32
		sum += uint32(b0) << 24
		sum += uint32(b1) << 16
		sum += uint32(b2) << 8
		sum += uint32(b3)
		return sum
	}
	return 0
}
