// Content managed by Project Forge, see [projectforge.md] for details.
package estimate

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Estimates []*Estimate

func (e Estimates) Get(id uuid.UUID) *Estimate {
	for _, x := range e {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (e Estimates) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(e)+1)
	for _, x := range e {
		ret = append(ret, x.ID)
	}
	return ret
}

func (e Estimates) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range e {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (e Estimates) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(e)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range e {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (e Estimates) Clone() Estimates {
	return slices.Clone(e)
}
