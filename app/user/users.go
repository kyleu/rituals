// Content managed by Project Forge, see [projectforge.md] for details.
package user

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Users []*User

func (u Users) Get(id uuid.UUID) *User {
	for _, x := range u {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (u Users) GetByIDs(ids ...uuid.UUID) Users {
	var ret Users
	for _, x := range u {
		if lo.Contains(ids, x.ID) {
			ret = append(ret, x)
		}
	}
	return ret
}

func (u Users) IDs() []uuid.UUID {
	ret := make([]uuid.UUID, 0, len(u)+1)
	for _, x := range u {
		ret = append(ret, x.ID)
	}
	return ret
}

func (u Users) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(u)+1)
	if includeNil {
		ret = append(ret, "")
	}
	for _, x := range u {
		ret = append(ret, x.ID.String())
	}
	return ret
}

func (u Users) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(u)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	for _, x := range u {
		ret = append(ret, x.TitleString())
	}
	return ret
}

func (u Users) Clone() Users {
	return slices.Clone(u)
}
