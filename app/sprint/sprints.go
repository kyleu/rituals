// Content managed by Project Forge, see [projectforge.md] for details.
package sprint

import (
	"github.com/google/uuid"
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
