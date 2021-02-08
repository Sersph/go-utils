package reflect

import (
	"reflect"
)

func GetFieldName(structure interface{}) []string {
	t := reflect.TypeOf(structure)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("Check type error not Struct")
	}
	fieldNames := make([]string, 0)
	for index := 0; index < t.NumField(); index++ {
		fieldNames = append(fieldNames, t.Field(index).Name)
	}
	return fieldNames
}

func GetFieldAnnotation(fieldName string, structure interface{}) []string {
	t := reflect.TypeOf(structure)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		panic("Check type error not Struct")
	}
	fieldNames := make([]string, 0)
	for index := 0; index < t.NumField(); index++ {
		fieldNames = append(fieldNames, t.Field(index).Name)
	}
	return fieldNames
}
