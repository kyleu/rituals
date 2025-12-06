package feedback

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type Feedbacks []*Feedback

func (f Feedbacks) Get(id uuid.UUID) *Feedback {
	return lo.FindOrElse(f, nil, func(x *Feedback) bool {
		return x.ID == id
	})
}

func (f Feedbacks) IDs() []uuid.UUID {
	return lo.Map(f, func(xx *Feedback, _ int) uuid.UUID {
		return xx.ID
	})
}

func (f Feedbacks) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(f)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(f, func(x *Feedback, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (f Feedbacks) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(f)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(f, func(x *Feedback, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (f Feedbacks) GetByID(id uuid.UUID) Feedbacks {
	return lo.Filter(f, func(xx *Feedback, _ int) bool {
		return xx.ID == id
	})
}

func (f Feedbacks) GetByIDs(ids ...uuid.UUID) Feedbacks {
	return lo.Filter(f, func(xx *Feedback, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (f Feedbacks) RetroIDs() []uuid.UUID {
	return lo.Map(f, func(xx *Feedback, _ int) uuid.UUID {
		return xx.RetroID
	})
}

func (f Feedbacks) GetByRetroID(retroID uuid.UUID) Feedbacks {
	return lo.Filter(f, func(xx *Feedback, _ int) bool {
		return xx.RetroID == retroID
	})
}

func (f Feedbacks) GetByRetroIDs(retroIDs ...uuid.UUID) Feedbacks {
	return lo.Filter(f, func(xx *Feedback, _ int) bool {
		return lo.Contains(retroIDs, xx.RetroID)
	})
}

func (f Feedbacks) UserIDs() []uuid.UUID {
	return lo.Map(f, func(xx *Feedback, _ int) uuid.UUID {
		return xx.UserID
	})
}

func (f Feedbacks) GetByUserID(userID uuid.UUID) Feedbacks {
	return lo.Filter(f, func(xx *Feedback, _ int) bool {
		return xx.UserID == userID
	})
}

func (f Feedbacks) GetByUserIDs(userIDs ...uuid.UUID) Feedbacks {
	return lo.Filter(f, func(xx *Feedback, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (f Feedbacks) ToMap() map[uuid.UUID]*Feedback {
	return lo.SliceToMap(f, func(xx *Feedback) (uuid.UUID, *Feedback) {
		return xx.ID, xx
	})
}

func (f Feedbacks) ToMaps() []util.ValueMap {
	return lo.Map(f, func(xx *Feedback, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (f Feedbacks) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(f, func(x *Feedback, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (f Feedbacks) ToCSV() ([]string, [][]string) {
	return FeedbackFieldDescs.Keys(), lo.Map(f, func(x *Feedback, _ int) []string {
		return x.Strings()
	})
}

func (f Feedbacks) Random() *Feedback {
	return util.RandomElement(f)
}

func (f Feedbacks) Clone() Feedbacks {
	return lo.Map(f, func(xx *Feedback, _ int) *Feedback {
		return xx.Clone()
	})
}
