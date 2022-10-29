// Content managed by Project Forge, see [projectforge.md] for details.
package comment

import (
	"github.com/google/uuid"
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
