package verify_method

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

type Method interface {
	Verify(value reflect.Value) bool
}

func CreateMethod(method, metadata string) Method {
	var metaJson = "{" + strings.ReplaceAll(metadata, "'", "\"") + "}"
	var result Method
	switch strings.ToLower(method) {
	case "required":
		params := RequiredMetadata{}
		json.Unmarshal([]byte(metaJson), &params)
		result = Required{Metadata: &params}
	}

	if result == nil {
		panic(errors.New("校验方式不存在 method = " + method))
	}
	return result
}

func parseTag(tag string) (method, metadata string) {
	method = ""
	metadata = ""
	if !strings.Contains(tag, "(") {
		method = strings.Trim(tag, " ")
	} else {
		items := strings.Split(tag, "(")
		method = strings.Trim(items[0], " ")
		metadata = strings.Trim(items[1], ")")
	}
	return
}

func CreateMethods(methods []string) []Method {
	results := make([]Method, 0)
	for _, method := range methods {
		results = append(results, CreateMethod(parseTag(method)))
	}
	return results
}
