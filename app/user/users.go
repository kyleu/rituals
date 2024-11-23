package user

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type Users []*User

func (u Users) Get(id uuid.UUID) *User {
	return lo.FindOrElse(u, nil, func(x *User) bool {
		return x.ID == id
	})
}

func (u Users) IDs() []uuid.UUID {
	return lo.Map(u, func(xx *User, _ int) uuid.UUID {
		return xx.ID
	})
}

func (u Users) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(u)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(u, func(x *User, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (u Users) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(u)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(u, func(x *User, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (u Users) GetByID(id uuid.UUID) Users {
	return lo.Filter(u, func(xx *User, _ int) bool {
		return xx.ID == id
	})
}

func (u Users) GetByIDs(ids ...uuid.UUID) Users {
	return lo.Filter(u, func(xx *User, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (u Users) ToMaps() []util.ValueMap {
	return lo.Map(u, func(x *User, _ int) util.ValueMap {
		return x.ToMap()
	})
}

func (u Users) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(u, func(x *User, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (u Users) ToCSV() ([]string, [][]string) {
	return UserFieldDescs.Keys(), lo.Map(u, func(x *User, _ int) []string {
		return x.Strings()
	})
}

func (u Users) Random() *User {
	return util.RandomElement(u)
}

func (u Users) Clone() Users {
	return lo.Map(u, func(xx *User, _ int) *User {
		return xx.Clone()
	})
}
