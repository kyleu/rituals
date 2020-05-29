package util

import (
	"encoding/json"
	"fmt"
	"logur.dev/logur"
)

func ToJSON(x interface{}) string {
	b, _ := json.MarshalIndent(x, "", "  ")
	return string(b)
}

func FromJSON(msg json.RawMessage, tgt interface{}, logger logur.Logger) {
	err := json.Unmarshal(msg, tgt)
	if err != nil {
		logger.Warn(fmt.Sprintf("error unmarshalling JSON [%v]: %+v", string(msg), err))
	}
}
