// Content managed by Project Forge, see [projectforge.md] for details.
package team

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Teams []*Team

func (t Teams) Get(id uuid.UUID) *Team {
	for _, x := range t {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (t Teams) GetByIDs(ids ...uuid.UUID) Teams {
	var ret Teams
	for _, x := range t {
		if slices.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (t Teams) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(t)+1)
	for _, x := range t {
		ret = append(ret, x.ID)
	}
	return ret
}

func (t Teams) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range t {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (t Teams) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range t {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (t Teams) Clone() Teams {
	return slices.Clone(t)
}
