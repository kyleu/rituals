package util

import (
	"fmt"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func GetDomain(email string) string {
	var idx = strings.LastIndex(email, "@")
	if idx == -1 {
		return email
	}
	return email[idx:]
}

func OxfordComma(names []string, clause string) string {
	ret := ""
	for idx, name := range names {
		if idx > 0 {
			if len(clause) > 0 && idx == (len(names) - 1) {
				if idx > 1 {
					ret += ","
				}
				ret += " " + clause + " "
			} else {
				ret += ", "
			}
		}
		ret += name
	}
	return ret
}

func ValueStrings(values []interface{}) string {
	ret := make([]string, 0, len(values))
	for _, v := range values {
		ret = append(ret, fmt.Sprintf("\"%v\"", v))
	}
	return strings.Join(ret, ", ")
}
