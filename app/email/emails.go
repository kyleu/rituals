// Content managed by Project Forge, see [projectforge.md] for details.
package email

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Emails []*Email

func (e Emails) Get(id uuid.UUID) *Email {
	for _, x := range e {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (e Emails) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range e {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (e Emails) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(e)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range e {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (e Emails) Clone() Emails {
	return slices.Clone(e)
}
