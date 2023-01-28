package emember

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/util"
)

func (e EstimateMembers) ToMembers(online []uuid.UUID) util.Members {
	ret := make(util.Members, 0, len(e))
	for _, x := range e {
		o := slices.Contains(online, x.UserID)
		ret = append(ret, &util.Member{UserID: x.UserID, Name: x.Name, Picture: x.Picture, Role: x.Role, Online: o})
	}
	return ret
}
