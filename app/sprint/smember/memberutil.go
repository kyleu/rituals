package smember

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/member"
)

func (s SprintMembers) ToMembers(online []uuid.UUID) member.Members {
	return lo.Map(s, func(x *SprintMember, _ int) *member.Member {
		return &member.Member{UserID: x.UserID, Name: x.Name, Picture: x.Picture, Role: x.Role, Online: slices.Contains(online, x.UserID)}
	})
}
