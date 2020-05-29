package util

import (
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
