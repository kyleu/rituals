// Package vote - Content managed by Project Forge, see [projectforge.md] for details.
package vote

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Votes []*Vote

func (v Votes) Get(storyID uuid.UUID, userID uuid.UUID) *Vote {
	return lo.FindOrElse(v, nil, func(x *Vote) bool {
		return x.StoryID == storyID && x.UserID == userID
	})
}

func (v Votes) ToPKs() []*PK {
	return lo.Map(v, func(x *Vote, _ int) *PK {
		return x.ToPK()
	})
}

func (v Votes) GetByStoryIDs(storyIDs ...uuid.UUID) Votes {
	return lo.Filter(v, func(xx *Vote, _ int) bool {
		return lo.Contains(storyIDs, xx.StoryID)
	})
}

func (v Votes) GetByStoryID(storyID uuid.UUID) Votes {
	return lo.Filter(v, func(xx *Vote, _ int) bool {
		return xx.StoryID == storyID
	})
}

func (v Votes) GetByUserIDs(userIDs ...uuid.UUID) Votes {
	return lo.Filter(v, func(xx *Vote, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (v Votes) GetByUserID(userID uuid.UUID) Votes {
	return lo.Filter(v, func(xx *Vote, _ int) bool {
		return xx.UserID == userID
	})
}

func (v Votes) StoryIDs() []uuid.UUID {
	return lo.Map(v, func(x *Vote, _ int) uuid.UUID {
		return x.StoryID
	})
}

func (v Votes) StoryIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(v)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(v, func(x *Vote, _ int) {
		ret = append(ret, x.StoryID.String())
	})
	return ret
}

func (v Votes) UserIDs() []uuid.UUID {
	return lo.Map(v, func(x *Vote, _ int) uuid.UUID {
		return x.UserID
	})
}

func (v Votes) UserIDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(v)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(v, func(x *Vote, _ int) {
		ret = append(ret, x.UserID.String())
	})
	return ret
}

func (v Votes) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(v)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(v, func(x *Vote, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (v Votes) Clone() Votes {
	return slices.Clone(v)
}
