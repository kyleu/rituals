package util

import (
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/enum"
)

type Member struct {
	UserID  uuid.UUID         `json:"userID"`
	Name    string            `json:"name"`
	Picture string            `json:"picture"`
	Role    enum.MemberStatus `json:"role"`
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
