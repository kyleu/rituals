package tmember

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Register(
	ctx context.Context, t uuid.UUID, userID uuid.UUID, name string, role enum.MemberStatus, tx *sqlx.Tx, actSvc *action.Service, logger util.Logger,
) (*TeamMember, error) {
	m := &TeamMember{TeamID: t, UserID: userID, Name: name, Role: role, Created: time.Now()}
	err := s.Save(ctx, tx, logger, m)
	if err != nil {
		return nil, err
	}
	err = actSvc.Post(ctx, enum.ModelServiceTeam, t, userID, action.ActMemberAdd, nil, nil, logger)
	if err != nil {
		return nil, err
	}
	return m, nil
}
