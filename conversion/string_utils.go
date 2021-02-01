package conversion

import (
	"encoding/json"
	"strconv"
	"strings"
)

/*
	ToString 获取变量的字符串值
	浮点型 3.0将会转换成字符串3, "3"
	非数值或字符类型的变量将会被转换成JSON格式字符串
*/
func ToString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		return strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		return strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		return strconv.Itoa(it)
	case uint:
		it := value.(uint)
		return strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		return strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		return strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		return strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		return strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		return strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		return strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		return strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		return strconv.FormatUint(it, 10)
	case string:
		return value.(string)
	case []byte:
		return string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		return string(newValue)
	}
}

//将任意数组转换成 string 以指定字符进行分割
func ArrayToString(array []interface{}, sep string) (result string) {
	if len(array) <= 0 {
		return result
	}
	for i := 0; i < len(array); i++ {
		if i == 0 {
			result += ToString(array[i])
		} else {
			result += sep + ToString(array[i])
		}
	}
	return
}

/**
下划线转大驼峰
*/
func UnderscoreToBigHump(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

/**
下划线转小驼峰
*/
func UnderscoreToSmallHump(name string) string {
	if len(name) < 2 {
		return name
	}
	result := UnderscoreToBigHump(name)
	index := result[:1]
	return strings.ToLower(index) + result[1:]
}

/**
驼峰转下划线
*/
func HumpToUnderscore(name string) string {
	newName := ""
	for index, item := range name {
		if index > 0 &&
			(item >= 'A' && item <= 'Z') &&
			(index < len(name)) {
			newName += "_"
		}
		newName += name[index : index+1]
	}
	return strings.ToLower(newName)
}
