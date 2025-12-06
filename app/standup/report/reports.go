package report

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type Reports []*Report

func (r Reports) Get(id uuid.UUID) *Report {
	return lo.FindOrElse(r, nil, func(x *Report) bool {
		return x.ID == id
	})
}

func (r Reports) IDs() []uuid.UUID {
	return lo.Map(r, func(xx *Report, _ int) uuid.UUID {
		return xx.ID
	})
}

func (r Reports) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *Report, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (r Reports) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(r, func(x *Report, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (r Reports) GetByID(id uuid.UUID) Reports {
	return lo.Filter(r, func(xx *Report, _ int) bool {
		return xx.ID == id
	})
}

func (r Reports) GetByIDs(ids ...uuid.UUID) Reports {
	return lo.Filter(r, func(xx *Report, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (r Reports) StandupIDs() []uuid.UUID {
	return lo.Map(r, func(xx *Report, _ int) uuid.UUID {
		return xx.StandupID
	})
}

func (r Reports) GetByStandupID(standupID uuid.UUID) Reports {
	return lo.Filter(r, func(xx *Report, _ int) bool {
		return xx.StandupID == standupID
	})
}

func (r Reports) GetByStandupIDs(standupIDs ...uuid.UUID) Reports {
	return lo.Filter(r, func(xx *Report, _ int) bool {
		return lo.Contains(standupIDs, xx.StandupID)
	})
}

func (r Reports) UserIDs() []uuid.UUID {
	return lo.Map(r, func(xx *Report, _ int) uuid.UUID {
		return xx.UserID
	})
}

func (r Reports) GetByUserID(userID uuid.UUID) Reports {
	return lo.Filter(r, func(xx *Report, _ int) bool {
		return xx.UserID == userID
	})
}

func (r Reports) GetByUserIDs(userIDs ...uuid.UUID) Reports {
	return lo.Filter(r, func(xx *Report, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (r Reports) ToMap() map[uuid.UUID]*Report {
	return lo.SliceToMap(r, func(xx *Report) (uuid.UUID, *Report) {
		return xx.ID, xx
	})
}

func (r Reports) ToMaps() []util.ValueMap {
	return lo.Map(r, func(xx *Report, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (r Reports) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(r, func(x *Report, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (r Reports) ToCSV() ([]string, [][]string) {
	return ReportFieldDescs.Keys(), lo.Map(r, func(x *Report, _ int) []string {
		return x.Strings()
	})
}

func (r Reports) Random() *Report {
	return util.RandomElement(r)
}

func (r Reports) Clone() Reports {
	return lo.Map(r, func(xx *Report, _ int) *Report {
		return xx.Clone()
	})
}
