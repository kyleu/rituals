// Content managed by Project Forge, see [projectforge.md] for details.
package email

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Emails []*Email

func (e Emails) Get(id uuid.UUID) *Email {
	return lo.FindOrElse(e, nil, func(x *Email) bool {
		return x.ID == id
	})
}

func (e Emails) GetByIDs(ids ...uuid.UUID) Emails {
	return lo.Filter(e, func(x *Email, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (e Emails) IDs() []uuid.UUID {
	return lo.Map(e, func(x *Email, _ int) uuid.UUID {
		return x.ID
	})
}

func (e Emails) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(e, func(x *Email, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (e Emails) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(e)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(e, func(x *Email, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (e Emails) Clone() Emails {
	return slices.Clone(e)
}
