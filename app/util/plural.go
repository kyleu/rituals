package util

import (
	"github.com/jinzhu/inflection"
)

func Plural(s string) string {
	return inflection.Plural(s)
}

func PluralChoice(i int, s string) string {
	if i == 1 {
		return s
	}
	return Plural(s)
}
