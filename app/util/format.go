package util

import (
	"regexp"

	"github.com/gofrs/uuid"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func MicrosToMillis(l language.Tag, i int) string {
	div := 1000
	min := 20

	ms := i / div
	if ms >= min {
		return FormatInteger(l, ms) + "ms"
	}

	x := float64(ms) + (float64(i%div) / float64(div))
	p := message.NewPrinter(l)

	return p.Sprintf("%.3f", x) + "ms"
}

func FormatInteger(l language.Tag, v int) string {
	p := message.NewPrinter(l)
	return p.Sprintf("%d", v)
}

var re *regexp.Regexp

func PathParams(s string) []string {
	if re == nil {
		re = regexp.MustCompile("{([^}]*)}")
	}

	matches := re.FindAll([]byte(s), -1)

	ret := make([]string, 0, len(matches))
	for _, m := range matches {
		ret = append(ret, string(m))
	}

	return ret
}

func GetUUIDFromString(s string) *uuid.UUID {
	var retID *uuid.UUID

	if len(s) > 0 {
		s, err := uuid.FromString(s)

		if err == nil {
			retID = &s
		}
	}

	return retID
}

func GetUUIDPointer(m map[string]string, key string) *uuid.UUID {
	retOut, ok := m[key]

	if !ok {
		return nil
	}

	return GetUUIDFromString(retOut)
}
