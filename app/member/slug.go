package member

import (
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/query"

	"emperror.dev/errors"
	"github.com/jmoiron/sqlx"
)

func NewSlugFor(db *sqlx.DB, svc string, str string) (string, error) {
	randomStrLength := 4
	if len(str) == 0 {
		str = strings.ToLower(randomString(randomStrLength))
	}
	slug := slugify(str)
	q := query.SQLSelect(util.KeyID, svc, "slug = $1", "", 0, 0)

	x, err := db.Queryx(q, slug)
	if err != nil {
		return slug, errors.WithStack(errors.Wrap(err, "error fetching existing slug"))
	}

	if x.Next() {
		junk := strings.ToLower(randomString(randomStrLength))
		slug, err = NewSlugFor(db, svc, slug+"-"+junk)
		if err != nil {
			return slug, errors.WithStack(errors.Wrap(err, "error finding slug for new "+svc+" session"))
		}
	}

	return slug, nil
}

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

var regexpNonAuthorizedChars = regexp.MustCompile("[^a-zA-Z0-9-_]")
var regexpMultipleDashes = regexp.MustCompile("-+")

func slugify(s string) (slug string) {
	slug = strings.TrimSpace(s)

	slug = strings.ToLower(slug)

	slug = regexpNonAuthorizedChars.ReplaceAllString(slug, "-")
	slug = regexpMultipleDashes.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-_")

	slug = smartTruncate(slug)

	return slug
}

func smartTruncate(text string) string {
	maxLength := 256
	if len(text) < maxLength {
		return text
	}

	var truncated string
	words := strings.SplitAfter(text, "-")
	if len(words[0]) > maxLength {
		return words[0][:maxLength]
	}

	for _, word := range words {
		if len(truncated)+len(word)-1 <= maxLength {
			truncated += word
		} else {
			break
		}
	}
	return strings.Trim(truncated, "-")
}
