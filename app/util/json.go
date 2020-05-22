package util

import "encoding/json"

func ToJSON(x interface{}) string {
	b, _ := json.MarshalIndent(x, "", "  ")
	return string(b)
}
