package jsonx

import "encoding/json"

func String(value any) string {
	if value == nil {
		return ""
	}

	data, err := json.Marshal(value)
	if err != nil {
		return ""
	}

	return string(data)
}
