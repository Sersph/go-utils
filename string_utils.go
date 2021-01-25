package go_utils

import (
	"regexp"
	"strings"
)

// ignoreSpaces 是否去掉前后空格再判断
// isBlank("")         = true
// isBlank(" ")        = false
// isBlank("bob")      = false
// isBlank("  bob  ")  = false
// isBlank("\r\n\t\f") = false
func IsEmpty(str string, ignoreSpaces bool) bool {
	if ignoreSpaces {
		str = strings.TrimSpace(str)
	}
	return str == "" || len(str) == 0
}

func IsNotEmpty(str string, ignoreSpaces bool) bool {
	return !IsEmpty(str, ignoreSpaces)
}

// isBlank("")         = true
// isBlank(" ")        = true
// isBlank("bob")      = false
// isBlank("  bob  ")  = false
// isBlank("\r\n\t\f") = true
func IsBlank(str string) bool {
	result, _ := regexp.MatchString("^\\s*$", str)
	return result
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}
