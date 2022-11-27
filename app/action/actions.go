// Content managed by Project Forge, see [projectforge.md] for details.
package action

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Actions []*Action

func (a Actions) Get(id uuid.UUID) *Action {
	for _, x := range a {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (a Actions) GetByIDs(ids ...uuid.UUID) Actions {
	var ret Actions
	for _, x := range a {
		if slices.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (a Actions) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(a)+1)
	for _, x := range a {
		ret = append(ret, x.ID)
	}
	return ret
}

func (a Actions) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(a)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range a {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (a Actions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(a)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range a {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (a Actions) Clone() Actions {
	return slices.Clone(a)
}
