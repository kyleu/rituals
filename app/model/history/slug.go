package history

import (
	"regexp"
	"strings"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/database/query"

	"emperror.dev/errors"
)

func (s *Service) NewSlugFor(str string) (string, error) {
	randomStrLength := 4
	if len(str) == 0 {
		str = strings.ToLower(util.RandomString(randomStrLength))
	}
	slug := slugify(str)

	q := query.SQLSelectSimple(util.KeyID, s.svc.Key, "slug = $1")
	x, err := s.db.Query(q, nil, slug)
	if err != nil {
		return slug, errors.Wrap(err, "error fetching existing slug")
	}

	q2 := query.SQLSelectSimple(util.WithDBID(util.KeyModel), s.tableName, "slug = $1")
	y, err := s.db.Query(q2, nil, slug)
	if err != nil {
		return slug, errors.Wrap(err, "error fetching historical slug")
	}

	if x.Next() || y.Next() {
		junk := strings.ToLower(util.RandomString(randomStrLength))
		slug, err = s.NewSlugFor(slug + "-" + junk)
		if err != nil {
			return slug, errors.Wrap(err, "error finding slug for new "+s.svc.Key+" session")
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
