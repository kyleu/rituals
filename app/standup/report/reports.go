// Content managed by Project Forge, see [projectforge.md] for details.
package report

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Reports []*Report

func (r Reports) Get(id uuid.UUID) *Report {
	for _, x := range r {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (r Reports) GetByIDs(ids ...uuid.UUID) Reports {
	var ret Reports
	for _, x := range r {
		if slices.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (r Reports) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(r)+1)
	for _, x := range r {
		ret = append(ret, x.ID)
	}
	return ret
}

func (r Reports) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range r {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (r Reports) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range r {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (r Reports) Clone() Reports {
	return slices.Clone(r)
}
