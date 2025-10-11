package action

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type Actions []*Action

func (a Actions) Get(id uuid.UUID) *Action {
	return lo.FindOrElse(a, nil, func(x *Action) bool {
		return x.ID == id
	})
}

func (a Actions) IDs() []uuid.UUID {
	return lo.Map(a, func(xx *Action, _ int) uuid.UUID {
		return xx.ID
	})
}

func (a Actions) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(a)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(a, func(x *Action, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (a Actions) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(a)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(a, func(x *Action, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (a Actions) GetByID(id uuid.UUID) Actions {
	return lo.Filter(a, func(xx *Action, _ int) bool {
		return xx.ID == id
	})
}

func (a Actions) GetByIDs(ids ...uuid.UUID) Actions {
	return lo.Filter(a, func(xx *Action, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (a Actions) UserIDs() []uuid.UUID {
	return lo.Map(a, func(xx *Action, _ int) uuid.UUID {
		return xx.UserID
	})
}

func (a Actions) GetByUserID(userID uuid.UUID) Actions {
	return lo.Filter(a, func(xx *Action, _ int) bool {
		return xx.UserID == userID
	})
}

func (a Actions) GetByUserIDs(userIDs ...uuid.UUID) Actions {
	return lo.Filter(a, func(xx *Action, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (a Actions) ToMaps() []util.ValueMap {
	return lo.Map(a, func(xx *Action, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (a Actions) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(a, func(x *Action, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (a Actions) ToCSV() ([]string, [][]string) {
	return ActionFieldDescs.Keys(), lo.Map(a, func(x *Action, _ int) []string {
		return x.Strings()
	})
}

func (a Actions) Random() *Action {
	return util.RandomElement(a)
}

func (a Actions) Clone() Actions {
	return lo.Map(a, func(xx *Action, _ int) *Action {
		return xx.Clone()
	})
}
