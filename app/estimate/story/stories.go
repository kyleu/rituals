// Content managed by Project Forge, see [projectforge.md] for details.
package story

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Stories []*Story

func (s Stories) Get(id uuid.UUID) *Story {
	for _, x := range s {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (s Stories) GetByIDs(ids ...uuid.UUID) Stories {
	var ret Stories
	for _, x := range s {
		if slices.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (s Stories) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.ID)
	}
	return ret
}

func (s Stories) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (s Stories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range s {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (s Stories) Clone() Stories {
	return slices.Clone(s)
}
