// Content managed by Project Forge, see [projectforge.md] for details.
package retro

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Retros []*Retro

func (r Retros) Get(id uuid.UUID) *Retro {
	for _, x := range r {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (r Retros) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range r {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (r Retros) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range r {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (r Retros) Clone() Retros {
	return slices.Clone(r)
}
