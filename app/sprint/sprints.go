// Content managed by Project Forge, see [projectforge.md] for details.
package sprint

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Sprints []*Sprint

func (s Sprints) Get(id uuid.UUID) *Sprint {
	for _, x := range s {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (s Sprints) GetByIDs(ids ...uuid.UUID) Sprints {
	var ret Sprints
	for _, x := range s {
		if lo.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s Sprints) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.ID)
	}
	return ret
}

func (s Sprints) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (s Sprints) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range s {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (s Sprints) Clone() Sprints {
	return slices.Clone(s)
}
