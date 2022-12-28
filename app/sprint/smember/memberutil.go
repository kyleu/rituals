package smember

import (
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (s SprintMembers) Split(userID uuid.UUID) (*SprintMember, SprintMembers, error) {
	var match *SprintMember
	others := make(SprintMembers, 0, len(s))
	for _, x := range s {
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
	return match, others, nil
}

func (s SprintMembers) ToMembers() util.Members {
	ret := make(util.Members, 0, len(s))
	for _, x := range s {
		ret = append(ret, &util.Member{UserID: x.UserID, Name: x.Name, Picture: x.Picture, Role: x.Role})
	}
	return ret
}
