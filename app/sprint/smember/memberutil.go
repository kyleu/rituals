package smember

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/util"
)

func (s SprintMembers) ToMembers(online []uuid.UUID) util.Members {
	ret := make(util.Members, 0, len(s))
	for _, x := range s {
		o := slices.Contains(online, x.UserID)
		ret = append(ret, &util.Member{UserID: x.UserID, Name: x.Name, Picture: x.Picture, Role: x.Role, Online: o})
	}
	return ret
}
