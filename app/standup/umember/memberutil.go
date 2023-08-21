package umember

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/util"
)

func (s StandupMembers) ToMembers(online []uuid.UUID) util.Members {
	return lo.Map(s, func(x *StandupMember, _ int) *util.Member {
		return &util.Member{UserID: x.UserID, Name: x.Name, Picture: x.Picture, Role: x.Role, Online: slices.Contains(online, x.UserID)}
	})
}
