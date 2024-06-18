package story

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type Stories []*Story

func (s Stories) Get(id uuid.UUID) *Story {
	return lo.FindOrElse(s, nil, func(x *Story) bool {
		return x.ID == id
	})
}

func (s Stories) IDs() []uuid.UUID {
	return lo.Map(s, func(xx *Story, _ int) uuid.UUID {
		return xx.ID
	})
}

func (s Stories) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *Story, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (s Stories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *Story, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s Stories) GetByID(id uuid.UUID) Stories {
	return lo.Filter(s, func(xx *Story, _ int) bool {
		return xx.ID == id
	})
}

func (s Stories) GetByIDs(ids ...uuid.UUID) Stories {
	return lo.Filter(s, func(xx *Story, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (s Stories) EstimateIDs() []uuid.UUID {
	return lo.Map(s, func(xx *Story, _ int) uuid.UUID {
		return xx.EstimateID
	})
}

func (s Stories) GetByEstimateID(estimateID uuid.UUID) Stories {
	return lo.Filter(s, func(xx *Story, _ int) bool {
		return xx.EstimateID == estimateID
	})
}

func (s Stories) GetByEstimateIDs(estimateIDs ...uuid.UUID) Stories {
	return lo.Filter(s, func(xx *Story, _ int) bool {
		return lo.Contains(estimateIDs, xx.EstimateID)
	})
}

func (s Stories) UserIDs() []uuid.UUID {
	return lo.Map(s, func(xx *Story, _ int) uuid.UUID {
		return xx.UserID
	})
}

func (s Stories) GetByUserID(userID uuid.UUID) Stories {
	return lo.Filter(s, func(xx *Story, _ int) bool {
		return xx.UserID == userID
	})
}

func (s Stories) GetByUserIDs(userIDs ...uuid.UUID) Stories {
	return lo.Filter(s, func(xx *Story, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (s Stories) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), lo.Map(s, func(x *Story, _ int) []string {
		return x.Strings()
	})
}

func (s Stories) Random() *Story {
	return util.RandomElement(s)
}

func (s Stories) Clone() Stories {
	return slices.Clone(s)
}
