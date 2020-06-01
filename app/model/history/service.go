package history

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/rituals.dev/app/database"
	"github.com/kyleu/rituals.dev/app/model/action"

	"emperror.dev/errors"

	"logur.dev/logur"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/database/query"
)

type Service struct {
	actions   *action.Service
	db        *database.Service
	logger    logur.Logger
	svc       util.Service
	tableName string
	colName   string
}

func NewService(actions *action.Service, db *database.Service, logger logur.Logger, svc util.Service) *Service {
	return &Service{
		actions:   actions,
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

func (s *Service) UpdateSlug(sessID uuid.UUID, oSlug string, oTitle string, title string, userID uuid.UUID) (string, error) {
	tgt, err := s.NewSlugFor(&sessID, title)
	if err != nil {
		return "", errors.Wrap(err, "error getting new slug slug for ["+title+"]")
	}

	if tgt == oSlug {
		return oSlug, nil
	}

	cols := []string{util.KeySlug, util.WithDBID(util.KeyModel), "model_name"}
	q := query.SQLInsert(s.tableName, cols, 1)
	err = s.db.Insert(q, nil, oSlug, sessID, oTitle)
	if err != nil {
		return "", errors.Wrap(err, "error getting new slug for ["+title+"]")
	}

	q = query.SQLUpdate(s.svc.Key, []string{util.KeySlug}, util.KeyID+" = $2 and slug = $3")
	err = s.db.UpdateOne(q, nil, tgt, sessID, oSlug)
	if err != nil {
		return "", errors.Wrap(err, "error getting new slug for ["+title+"]")
	}

	actionContent := map[string]interface{}{"src": oSlug, "tgt": tgt}
	s.actions.Post(s.svc, sessID, userID, action.ActUpdate, actionContent, "")
	return tgt, nil
}

func (s *Service) Remove(slug string) error {
	q := query.SQLDelete(s.tableName, "slug = $1")
	err := s.db.DeleteOne(q, nil, slug)
	return errors.Wrap(err, "unable to remove " + s.svc.Key + " history for ["+slug+"]")
}

func (s *Service) GetByModelID(id uuid.UUID, params *query.Params) Entries {
	var dtos []entryDTO

	q := query.SQLSelect("*", s.tableName, "model_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving history for model [%v:%v]: %+v", s.svc.Key, id, err))
		return nil
	}
	return toEntries(dtos)
}

func toEntries(dtos []entryDTO) Entries {
	ret := make(Entries, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToEntry())
	}
	return ret
}
