package sql

import (
	"github.com/winily/go-utils/conversion"
	"reflect"
	"strings"
)

func ColumnProcess(column string) string {
	return BACK_QUOTE + column + BACK_QUOTE
}

func ConstituteWherePhrase(column, medium string, val interface{}) string {
	valStr := conversion.ToString(val)

	valType := reflect.TypeOf(val).Kind().String()
	if !isNumber(valType) {
		valStr = "\"" + valStr + "\""
	}

	return ColumnProcess(column) + medium + valStr
}

func isNumber(typeStr string) bool {
	return strings.Contains(typeStr, "int") || strings.Contains(typeStr, "float") || strings.Contains(typeStr, "uint") || strings.Contains(typeStr, "complex")
}
