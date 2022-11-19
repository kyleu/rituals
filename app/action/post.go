package action

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Post(
	ctx context.Context, svc enum.ModelService, id uuid.UUID, user uuid.UUID, a Act, t util.ValueMap, tx *sqlx.Tx, logger util.Logger,
) error {
	action := &Action{ID: util.UUID(), Svc: svc, ModelID: id, UserID: user, Act: string(a), Content: t, Created: time.Now()}
	return s.Create(ctx, tx, logger, action)
}
