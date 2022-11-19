// Content managed by Project Forge, see [projectforge.md] for details.
package standup

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Standups []*Standup

func (s Standups) Get(id uuid.UUID) *Standup {
	for _, x := range s {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (s Standups) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(s)+1)
	for _, x := range s {
		ret = append(ret, x.ID)
	}
	return ret
}

func (s Standups) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range s {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (s Standups) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range s {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (s Standups) Clone() Standups {
	return slices.Clone(s)
}
