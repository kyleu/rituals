package util

import (
	"emperror.dev/errors"
	"github.com/jmoiron/sqlx"
	"strings"
)

func NewSlugFor(db *sqlx.DB, svc string, str string) (string, error) {
	if len(str) == 0 {
		str = strings.ToLower(RandomString(4))
	}
	slug := strings.ReplaceAll(strings.ToLower(str), " ", "-")
	q := "select id from " + svc + " where slug = $1"
	x, err := db.Queryx(q, slug)
	if err != nil {
		return slug, errors.WithStack(errors.Wrap(err, "error fetching existing slug"))
	}
	if x.Next() {
		junk := strings.ToLower(RandomString(4))
		slug, err = NewSlugFor(db, svc, slug+"-"+junk)
		if err != nil {
			return slug, errors.WithStack(errors.Wrap(err, "error finding slug for new session"))
		}
	}
	return slug, nil
}
