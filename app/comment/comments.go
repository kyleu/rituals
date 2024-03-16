// Package comment - Content managed by Project Forge, see [projectforge.md] for details.
package comment

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

type Comments []*Comment

func (c Comments) Get(id uuid.UUID) *Comment {
	return lo.FindOrElse(c, nil, func(x *Comment) bool {
		return x.ID == id
	})
}

func (c Comments) IDs() []uuid.UUID {
	return lo.Map(c, func(xx *Comment, _ int) uuid.UUID {
		return xx.ID
	})
}

func (c Comments) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(c)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(c, func(x *Comment, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (c Comments) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(c)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(c, func(x *Comment, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (c Comments) GetByID(id uuid.UUID) Comments {
	return lo.Filter(c, func(xx *Comment, _ int) bool {
		return xx.ID == id
	})
}

func (c Comments) GetByIDs(ids ...uuid.UUID) Comments {
	return lo.Filter(c, func(xx *Comment, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (c Comments) UserIDs() []uuid.UUID {
	return lo.Map(c, func(xx *Comment, _ int) uuid.UUID {
		return xx.UserID
	})
}

func (c Comments) GetByUserID(userID uuid.UUID) Comments {
	return lo.Filter(c, func(xx *Comment, _ int) bool {
		return xx.UserID == userID
	})
}

func (c Comments) GetByUserIDs(userIDs ...uuid.UUID) Comments {
	return lo.Filter(c, func(xx *Comment, _ int) bool {
		return lo.Contains(userIDs, xx.UserID)
	})
}

func (c Comments) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), lo.Map(c, func(x *Comment, _ int) []string {
		return x.Strings()
	})
}

func (c Comments) Random() *Comment {
	if len(c) == 0 {
		return nil
	}
	return c[util.RandomInt(len(c))]
}

func (c Comments) Clone() Comments {
	return slices.Clone(c)
}
