package history

import (
	"database/sql"
	"fmt"
	"github.com/kyleu/rituals.dev/app/database"

	"emperror.dev/errors"

	"logur.dev/logur"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/database/query"
)

type Service struct {
	db        *database.Service
	logger    logur.Logger
	svc       util.Service
	tableName string
	colName   string
}

func NewService(db *database.Service, logger logur.Logger, svc util.Service) *Service {
	return &Service{
		db:        db,
		logger:    logger,
		svc:       svc,
		tableName: svc.Key + "_history",
		colName:   util.WithDBID(svc.Key),
	}
}

func (s *Service) Get(slug string) *Entry {
	dto := entryDTO{}
	q := query.SQLSelectSimple("*", s.tableName, "slug = $1")
	err := s.db.Get(&dto, q, nil, slug)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		s.logger.Warn(fmt.Sprintf("error getting %v history: %+v", s.svc.Key, err))
		return nil
	}
	return dto.ToEntry()
}

func (s *Service) RemoveHistory(modelID uuid.UUID) error {
	q := query.SQLDelete(s.tableName, "model_id = $1")
	err := s.db.DeleteOne(q, nil, modelID)
	return errors.Wrap(err, "unable to remove history for ["+modelID.String()+"]")
}
