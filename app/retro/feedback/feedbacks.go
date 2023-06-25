// Content managed by Project Forge, see [projectforge.md] for details.
package feedback

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Feedbacks []*Feedback

func (f Feedbacks) Get(id uuid.UUID) *Feedback {
	return lo.FindOrElse(f, nil, func(x *Feedback) bool {
		return x.ID == id
	})
}

func (f Feedbacks) GetByIDs(ids ...uuid.UUID) Feedbacks {
	return lo.Filter(f, func(x *Feedback, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (f Feedbacks) IDs() []uuid.UUID {
	return lo.Map(f, func(x *Feedback, _ int) uuid.UUID {
		return x.ID
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

func (f Feedbacks) Clone() Feedbacks {
	return slices.Clone(f)
}
