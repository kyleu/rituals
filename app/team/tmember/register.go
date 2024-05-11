package tmember

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Register(
	ctx context.Context, t uuid.UUID, userID uuid.UUID, name string, picture string, role enum.MemberStatus, tx *sqlx.Tx,
	actSvc *action.Service, send action.SendFn, us *user.Service, logger util.Logger,
) (*TeamMember, error) {
	err := us.CreateIfNeeded(ctx, userID, name, nil, logger)
	if err != nil {
		return nil, err
	}
	if name == "" {
		name = "Guest"
	}
	m := &TeamMember{TeamID: t, UserID: userID, Name: name, Picture: picture, Role: role, Created: time.Now()}
	err = s.Save(ctx, tx, logger, m)
	if err != nil {
		return nil, err
	}
	err = actSvc.Post(ctx, enum.ModelServiceTeam, t, userID, action.ActMemberAdd, util.ValueMap{"payload": m}, nil, logger)
	if err != nil {
		return nil, err
	}
	if send != nil {
		err = send(enum.ModelServiceTeam, t, action.ActMemberAdd, m, nil, logger)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}
