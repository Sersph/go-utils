package verify

import (
	"errors"
	"reflect"
	"strings"

	verify_scene "github.com/winily/go-utils/verify/scene"
)

func isMathScene(tag Tag, scenes ...verify_scene.Scene) bool {
	for _, sceneName := range tag.Scenes {
		for _, scene := range scenes {
			if scene.Match(sceneName) {
				return true
			}
		}
	}
	return false
}

func ValidateValue(value reflect.Value, tagString string, scenes ...verify_scene.Scene) error {
	tag := TagParse(tagString)
	if !isMathScene(tag, scenes...) {
		// 跳过验证
		return nil
	}

	for _, method := range tag.Methods {
		if ok := method.Verify(value); !ok {
			return errors.New(tag.Message)
		}
	}

	return nil
}

func ValidateStruct(obj any, scenes ...verify_scene.Scene) error {
	objType := reflect.TypeOf(obj)
	var objValue reflect.Value
	if objType.PkgPath() == "reflect" && objType.Name() == "Value" {
		objValue = obj.(reflect.Value)
		objType = objValue.Type()
	} else {
		objValue = reflect.ValueOf(obj)
	}
	validateRet := make([]string, 0)

	for i := 0; i < objValue.NumField(); i++ {
		fieldType := objType.Field(i)
		fieldValue := objValue.FieldByName(fieldType.Name)
		tag := fieldType.Tag.Get("verify")
		var valifyErr error
		if tag != "" {
			if err := ValidateValue(fieldValue, tag, scenes...); err != nil {
				validateRet = append(validateRet, err.Error())
			}
		}

		switch fieldType.Type.Kind() {
		case reflect.Ptr:
			if fieldValue.Elem().Kind() == reflect.Struct {
				valifyErr = ParamsVerify(fieldValue.Elem().Interface(), scenes...)
			}
		case reflect.Struct:
			valifyErr = ValidateStruct(fieldValue, scenes...)
		case reflect.Slice, reflect.Array:
			valifyErr = ParamsVerify(fieldValue, scenes...)
		default:

		}

		if valifyErr != nil {
			validateRet = append(validateRet, valifyErr.Error())
		}
	}

	if len(validateRet) > 0 {
		return errors.New(strings.Join(validateRet, ";"))
	}

	return nil
}

func ParamsVerify(obj any, scenes ...verify_scene.Scene) error {
	if len(scenes) == 0 {
		scenes = append(scenes, verify_scene.DEFAULT)
	}

	if obj == nil {
		return nil
	}

	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		return ParamsVerify(value.Elem().Interface(), scenes...)
	case reflect.Struct:
		return ValidateStruct(obj, scenes...)
	case reflect.Slice, reflect.Array:
		count := value.Len()
		validateRet := make([]string, 0)
		for i := 0; i < count; i++ {
			if err := ParamsVerify(value.Index(i).Interface(), scenes...); err != nil {
				validateRet = append(validateRet, err.Error())
			}
		}
		if len(validateRet) == 0 {
			return nil
		}
		return errors.New(strings.Join(validateRet, ";"))
	default:
		return nil
	}
}
