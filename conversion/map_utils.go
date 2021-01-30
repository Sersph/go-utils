package conversion

import "encoding/json"

func ToMap(data interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	j, _ := json.Marshal(data)

	_ = json.Unmarshal(j, &m)
	return m
}
