package go_utils

import (
	"regexp"
	"strings"
)

func IsEmpty(str string, ignoreSpaces bool) bool {
	if ignoreSpaces {
		str = strings.TrimSpace(str)
	}
	return str == "" || len(str) == 0
}

func IsNotEmpty(str string, ignoreSpaces bool) bool {
	return !IsEmpty(str, ignoreSpaces)
}

func IsBlank(str string) bool {
	result, _ := regexp.MatchString("^\\s*$", str)
	return result
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}
