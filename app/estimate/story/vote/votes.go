// Content managed by Project Forge, see [projectforge.md] for details.
package vote

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Votes []*Vote

func (v Votes) Get(storyID uuid.UUID, userID uuid.UUID) *Vote {
	for _, x := range v {
		if x.StoryID == storyID && x.UserID == userID {
			return x
		}
	}
	return nil
}

func (v Votes) StoryIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(v)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range v {
		ret = append(ret, x.StoryID.String())
	}
	return ret
}

func (v Votes) UserIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(v)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range v {
		ret = append(ret, x.UserID.String())
	}
	return ret
}

func (v Votes) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(v)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range v {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (v Votes) Clone() Votes {
	return slices.Clone(v)
}
