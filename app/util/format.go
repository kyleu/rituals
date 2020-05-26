package util

import (
	"regexp"
	"strings"
	"time"

	"emperror.dev/errors"

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

const YMD = "2006-01-02"

func ToYMD(d *time.Time) string {
	if d == nil {
		return ""
	}
	return d.Format(YMD)
}
func FromYMD(s string) (*time.Time, error) {
	ret, err := time.Parse(YMD, s)
	if err != nil {
		return nil, errors.New("invalid date [" + s + "] (expected 2020-01-15)")
	}
	return &ret, nil
}

const DateFull = "2006-01-02 15:04:05"

func ToDateString(d *time.Time) string {
	if d == nil {
		return ""
	}
	return d.Format(DateFull)
}

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
