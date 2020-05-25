package member

import (
	"github.com/kyleu/rituals.dev/app/database"
	"github.com/kyleu/rituals.dev/app/util"
	"regexp"
	"strings"

	"github.com/kyleu/rituals.dev/app/query"

	"emperror.dev/errors"
)

func NewSlugFor(db *database.Service, svc string, str string) (string, error) {
	randomStrLength := 4
	if len(str) == 0 {
		str = strings.ToLower(util.RandomString(randomStrLength))
	}
	slug := slugify(str)
	q := query.SQLSelect(util.KeyID, svc, "slug = $1", "", 0, 0)

	x, err := db.Query(q, nil, slug)
	if err != nil {
		return slug, errors.Wrap(err, "error fetching existing slug")
	}

	if x.Next() {
		junk := strings.ToLower(util.RandomString(randomStrLength))
		slug, err = NewSlugFor(db, svc, slug+"-"+junk)
		if err != nil {
			return slug, errors.Wrap(err, "error finding slug for new "+svc+" session")
		}
	}

	return slug, nil
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
