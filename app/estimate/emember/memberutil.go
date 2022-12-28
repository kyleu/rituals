package emember

import (
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (e EstimateMembers) Split(userID uuid.UUID) (*EstimateMember, EstimateMembers, error) {
	var match *EstimateMember
	others := make(EstimateMembers, 0, len(e))
	for _, x := range e {
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

func (e EstimateMembers) ToMembers() util.Members {
	ret := make(util.Members, 0, len(e))
	for _, x := range e {
		ret = append(ret, &util.Member{UserID: x.UserID, Name: x.Name, Picture: x.Picture, Role: x.Role})
	}
	return ret
}
