// Content managed by Project Forge, see [projectforge.md] for details.
package comment

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Comments []*Comment

func (c Comments) Get(id uuid.UUID) *Comment {
	for _, x := range c {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (c Comments) GetByIDs(ids ...uuid.UUID) Comments {
	var ret Comments
	for _, x := range c {
		if lo.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (c Comments) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(c)+1)
	for _, x := range c {
		ret = append(ret, x.ID)
	}
	return ret
}

func (c Comments) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(c)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range c {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (c Comments) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(c)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range c {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (c Comments) Clone() Comments {
	return slices.Clone(c)
}
