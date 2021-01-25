package verify

import "strings"

func IsEmpty(str string, ignoreSpaces bool) bool {
	if ignoreSpaces {
		str = strings.TrimSpace(str)
	}
	return str == "" || len(str) == 0
}

func IsNotEmpty(str string, ignoreSpaces bool) bool {
	return !IsEmpty(str, ignoreSpaces)
}
