package vote

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type Votes []*Vote

func (v Votes) Get(storyID uuid.UUID, userID uuid.UUID) *Vote {
	return lo.FindOrElse(v, nil, func(x *Vote) bool {
		return x.StoryID == storyID && x.UserID == userID
	})
}

func (v Votes) StoryIDs() []uuid.UUID {
	return lo.Map(v, func(xx *Vote, _ int) uuid.UUID {
		return xx.StoryID
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
	return lo.Map(v, func(xx *Vote, _ int) uuid.UUID {
		return xx.UserID
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

func (v Votes) ToPKs() []*PK {
	return lo.Map(v, func(x *Vote, _ int) *PK {
		return x.ToPK()
	})
}

func (v Votes) GetByStoryID(storyID uuid.UUID) Votes {
	return lo.Filter(v, func(xx *Vote, _ int) bool {
		return xx.StoryID == storyID
	})
}

func (v Votes) GetByStoryIDs(storyIDs ...uuid.UUID) Votes {
	return lo.Filter(v, func(xx *Vote, _ int) bool {
		return lo.Contains(storyIDs, xx.StoryID)
	})
}

func (v Votes) GetByUserID(userID uuid.UUID) Votes {
	return lo.Filter(v, func(xx *Vote, _ int) bool {
		return xx.UserID == userID
	})
}

func (v Votes) GetByUserIDs(userIDs ...uuid.UUID) Votes {
	return lo.Filter(v, func(xx *Vote, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (v Votes) ToMaps() []util.ValueMap {
	return lo.Map(v, func(xx *Vote, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (v Votes) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(v, func(x *Vote, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (v Votes) ToCSV() ([]string, [][]string) {
	return VoteFieldDescs.Keys(), lo.Map(v, func(x *Vote, _ int) []string {
		return x.Strings()
	})
}

func (v Votes) Random() *Vote {
	return util.RandomElement(v)
}

func (v Votes) Clone() Votes {
	return lo.Map(v, func(xx *Vote, _ int) *Vote {
		return xx.Clone()
	})
}
