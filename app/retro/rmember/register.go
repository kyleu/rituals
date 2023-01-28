package rmember

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
	ctx context.Context, r uuid.UUID, userID uuid.UUID, name string, role enum.MemberStatus, tx *sqlx.Tx,
	actSvc *action.Service, send action.SendFn, logger util.Logger,
) (*RetroMember, error) {
	m := &RetroMember{RetroID: r, UserID: userID, Name: name, Role: role, Created: time.Now()}
	err := s.Save(ctx, tx, logger, m)
	if err != nil {
		return nil, err
	}
	err = actSvc.Post(ctx, enum.ModelServiceRetro, r, userID, action.ActMemberAdd, util.ValueMap{"payload": m}, nil, logger)
	if err != nil {
		return nil, err
	}
	if send != nil {
		err = send(enum.ModelServiceRetro, r, action.ActMemberAdd, m, nil, logger)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}