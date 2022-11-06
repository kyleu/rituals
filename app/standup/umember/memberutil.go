package umember

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (t StandupMembers) Split(userID uuid.UUID) (*StandupMember, StandupMembers, error) {
	var match *StandupMember
	others := make(StandupMembers, 0, len(t))
	for _, x := range t {
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