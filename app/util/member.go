package util

import (
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/enum"
)

type Member struct {
	UserID  uuid.UUID         `json:"userID"`
	Name    string            `json:"name"`
	Picture string            `json:"picture"`
	Role    enum.MemberStatus `json:"role"`
	Online  bool              `json:"online,omitempty"`
}

func (m *Member) NameSafe() string {
	if m.Name == "" {
		return m.UserID.String()[0:8]
	}
	return m.Name
}

type Members []*Member

func (m Members) Get(userID uuid.UUID) *Member {
	for _, x := range m {
		if x.UserID == userID {
			return x
		}
	}
	return nil
}

func (m Members) Name(userID *uuid.UUID) string {
	if userID == nil {
		return "-"
	}
	if x := m.Get(*userID); x != nil {
		return x.NameSafe()
	}
	return userID.String()[0:8]
}

func (m Members) Sort() {
	slices.SortFunc(m, func(l, r *Member) bool {
		return strings.ToLower(l.Name) < strings.ToLower(r.Name)
	})
}

func (m Members) Split(userID uuid.UUID) (*Member, Members, error) {
	var match *Member
	others := make(Members, 0, len(m))
	for _, x := range m {
		if x.UserID == userID {
			if match != nil {
				return nil, nil, errors.Errorf("multiple members found with user ID [%s]", x.UserID.String())
			}
			match = x
		} else {
			others = append(others, x)
		}
	}
	if match == nil {
		return nil, nil, errors.Errorf("user [%s] is not a member", userID.String())
	}
	others.Sort()
	return match, others, nil
}
