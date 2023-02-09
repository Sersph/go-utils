package verify

import (
	"errors"
	"strings"

	verify_method "github.com/winily/go-utils/verify/method"
)

type Tag struct {
	Scenes  []string
	Methods []verify_method.Method
	Message string
}

func TagParse(tag string) Tag {
	tags := strings.Split(tag, ";")
	result := Tag{}
	for _, tag := range tags {
		tagitmes := strings.Split(tag, "=")
		if len(tagitmes) != 2 {
			panic(errors.New("校验 tag 格式错误 tag = " + tag))
		}
		key := tagitmes[0]
		value := tagitmes[1]
		switch strings.ToLower(key) {
		case "scene":
			result.Scenes = strings.Split(value, ",")
		case "method":
			methods := strings.Split(value, ",")
			result.Methods = verify_method.CreateMethods(methods)
		case "message":
			result.Message = value
		}
	}
	return result
}
