package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/kyleu/rituals/app/util"
)

func (s *Service) CreateIfNeeded(ctx context.Context, userID uuid.UUID, username string, tx *sqlx.Tx, logger util.Logger) error {
	if curr, _ := s.Get(ctx, tx, userID, logger); curr == nil {
		err := s.Create(ctx, tx, logger, &User{ID: userID, Name: username, Created: time.Now()})
		if err != nil {
			return err
		}
	}
	return nil
}
