package action

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

type (
	SendFn     func(svc enum.ModelService, modelID uuid.UUID, act Act, param any, userID *uuid.UUID, logger util.Logger, except ...uuid.UUID) error
	SendUserFn func(connID uuid.UUID, svc enum.ModelService, modelID uuid.UUID, act Act, param any, userID *uuid.UUID, logger util.Logger) error
)

func (s *Service) Post(
	ctx context.Context, svc enum.ModelService, id uuid.UUID, userID uuid.UUID,
	a Act, t util.ValueMap, tx *sqlx.Tx, logger util.Logger, sends ...SendFn,
) error {
	action := &Action{ID: util.UUID(), Svc: svc, ModelID: id, UserID: userID, Act: string(a), Content: t, Created: time.Now()}
	err := s.Create(ctx, tx, logger, action)
	if err != nil {
		return err
	}
	for _, x := range sends {
		var p any
		if t != nil {
			p = t["payload"]
		}
		err = x(svc, id, a, p, &userID, logger, userID)
		if err != nil {
			return err
		}
	}
	return nil
}
