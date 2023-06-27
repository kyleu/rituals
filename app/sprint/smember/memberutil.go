package smember

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/util"
)

func (s SprintMembers) ToMembers(online []uuid.UUID) util.Members {
	return lo.Map(s, func(x *SprintMember, _ int) *util.Member {
		return &util.Member{UserID: x.UserID, Name: x.Name, Picture: x.Picture, Role: x.Role, Online: slices.Contains(online, x.UserID)}
	})
}
