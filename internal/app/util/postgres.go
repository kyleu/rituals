package util

import (
	"strings"
)

func StringToArray(s string) []string {
	split := strings.Split(strings.TrimPrefix(strings.TrimSuffix(s, "}"), "{"), ",")
	ret := make([]string, 0)
	for _, x := range split {
		y := strings.TrimSpace(x)
		if len(y) > 0 {
			ret = append(ret, y)
		}
	}
	return ret
}
