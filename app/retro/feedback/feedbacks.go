// Content managed by Project Forge, see [projectforge.md] for details.
package feedback

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Feedbacks []*Feedback

func (f Feedbacks) Get(id uuid.UUID) *Feedback {
	for _, x := range f {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (f Feedbacks) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(f)+1)
	for _, x := range f {
		ret = append(ret, x.ID)
	}
	return ret
}

func (f Feedbacks) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(f)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range f {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (f Feedbacks) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(f)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range f {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (f Feedbacks) Clone() Feedbacks {
	return slices.Clone(f)
}
