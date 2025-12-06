package email

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type Emails []*Email

func (e Emails) Get(id uuid.UUID) *Email {
	return lo.FindOrElse(e, nil, func(x *Email) bool {
		return x.ID == id
	})
}

func (e Emails) IDs() []uuid.UUID {
	return lo.Map(e, func(xx *Email, _ int) uuid.UUID {
		return xx.ID
	})
}

func (e Emails) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(e, func(x *Email, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (e Emails) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(e)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(e, func(x *Email, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (e Emails) GetByID(id uuid.UUID) Emails {
	return lo.Filter(e, func(xx *Email, _ int) bool {
		return xx.ID == id
	})
}

func (e Emails) GetByIDs(ids ...uuid.UUID) Emails {
	return lo.Filter(e, func(xx *Email, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (e Emails) UserIDs() []uuid.UUID {
	return lo.Map(e, func(xx *Email, _ int) uuid.UUID {
		return xx.UserID
	})
}

func (e Emails) GetByUserID(userID uuid.UUID) Emails {
	return lo.Filter(e, func(xx *Email, _ int) bool {
		return xx.UserID == userID
	})
}

func (e Emails) GetByUserIDs(userIDs ...uuid.UUID) Emails {
	return lo.Filter(e, func(xx *Email, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (e Emails) ToMap() map[uuid.UUID]*Email {
	return lo.SliceToMap(e, func(xx *Email) (uuid.UUID, *Email) {
		return xx.ID, xx
	})
}

func (e Emails) ToMaps() []util.ValueMap {
	return lo.Map(e, func(xx *Email, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (e Emails) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(e, func(x *Email, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (e Emails) ToCSV() ([]string, [][]string) {
	return EmailFieldDescs.Keys(), lo.Map(e, func(x *Email, _ int) []string {
		return x.Strings()
	})
}

func (e Emails) Random() *Email {
	return util.RandomElement(e)
}

func (e Emails) Clone() Emails {
	return lo.Map(e, func(xx *Email, _ int) *Email {
		return xx.Clone()
	})
}
