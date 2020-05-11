package util

import (
	"fmt"
	"regexp"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func MicrosToMillis(l language.Tag, i int) string {
	ms := i / 1000
	if ms >= 20 {
		return FormatInteger(l, ms) + "ms"
	}
	x := float64(ms) + (float64(i%1000) / 1000)
	p := message.NewPrinter(l)
	return p.Sprintf("%.3f", x) + "ms"
}

func FormatInteger(l language.Tag, v int) string {
	p := message.NewPrinter(l)
	return p.Sprintf("%d", v)
}

func PluralChoice(single string, plural string, v int) string {
	if v == 1 || v == -1 {
		return fmt.Sprint(v, " ", single)
	}
	return fmt.Sprint(v, " ", plural)
}

func BoolUnicode(b bool) string {
	if b {
		return "✓"
	}
	return "✗"
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
