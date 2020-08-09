package history

import (
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"regexp"
	"strings"

	"github.com/gofrs/uuid"

	"emperror.dev/errors"
)

func (s *Service) NewSlugFor(modelID *uuid.UUID, title string) (string, error) {
	randomStrLength := 4
	if len(title) == 0 {
		title = strings.ToLower(npncore.RandomString(randomStrLength))
	}
	slug := slugify(title)

	q := npndatabase.SQLSelectSimple(npncore.KeyID, s.svc.Key, "slug = $1")
	x, err := s.db.Query(q, nil, slug)
	if err != nil {
		return slug, errors.Wrap(err, "error fetching existing slug")
	}

	curr := s.Get(slug)

	if x.Next() {
		junk := strings.ToLower(npncore.RandomString(randomStrLength))
		slug, err = s.NewSlugFor(modelID, slug+"-"+junk)
		if err != nil {
			return slug, errors.Wrap(err, "error finding slug for new "+s.svc.Key+" session")
		}
	} else if curr != nil && modelID != nil && curr.ModelID == *modelID {
		err = s.Remove(curr.Slug)
		return slug, errors.Wrap(err, "unable to remove old history with slug ["+curr.Slug+"]")
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
