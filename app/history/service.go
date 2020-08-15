package history

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/action"

	"logur.dev/logur"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
)

type Service struct {
	actions   *action.Service
	db        *npndatabase.Service
	logger    logur.Logger
	svc       util.Service
	tableName string
	colName   string
}

func NewService(actions *action.Service, db *npndatabase.Service, logger logur.Logger, svc util.Service) *Service {
	return &Service{
		actions:   actions,
		db:        db,
		logger:    logger,
		svc:       svc,
		tableName: svc.Key + "_history",
		colName:   npncore.WithDBID(svc.Key),
	}
}

func (s *Service) Get(slug string) *Entry {
	dto := entryDTO{}
	q := npndatabase.SQLSelectSimple("*", s.tableName, "slug = $1")
	err := s.db.Get(&dto, q, nil, slug)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		s.logger.Warn(fmt.Sprintf("error getting %v history: %+v", s.svc.Key, err))
		return nil
	}
	return dto.toEntry()
}

func (s *Service) UpdateSlug(sessID uuid.UUID, oSlug string, oTitle string, title string, userID uuid.UUID) (string, error) {
	nSlug := slugify(title)
	if oSlug == nSlug {
		return oSlug, nil
	}

	tgt, err := s.NewSlugFor(&sessID, title)
	if err != nil {
		return "", errors.Wrap(err, "error getting new slug slug for ["+title+"]")
	}

	if tgt == oSlug {
		return oSlug, nil
	}

	cols := []string{npncore.KeySlug, npncore.WithDBID(npncore.KeyModel), "model_name"}
	q := npndatabase.SQLInsert(s.tableName, cols, 1)
	err = s.db.Insert(q, nil, oSlug, sessID, oTitle)
	if err != nil {
		return "", errors.Wrap(err, "error getting new slug for ["+title+"]")
	}

	q = npndatabase.SQLUpdate(s.svc.Key, []string{npncore.KeySlug}, npncore.KeyID+" = $2 and slug = $3")
	err = s.db.UpdateOne(q, nil, tgt, sessID, oSlug)
	if err != nil {
		return "", errors.Wrap(err, "error getting new slug for ["+title+"]")
	}

	actionContent := map[string]interface{}{"src": oSlug, "tgt": tgt}
	s.actions.Post(s.svc.Key, sessID, userID, action.ActUpdate, actionContent)
	return tgt, nil
}

func (s *Service) Remove(slug string) error {
	q := npndatabase.SQLDelete(s.tableName, "slug = $1")
	err := s.db.DeleteOne(q, nil, slug)
	return errors.Wrap(err, "unable to remove "+s.svc.Key+" history for ["+slug+"]")
}

func (s *Service) GetByModelID(id uuid.UUID, params *npncore.Params) Entries {
	var dtos []entryDTO

	q := npndatabase.SQLSelect("*", s.tableName, "model_id = $1", params.OrderByString(), params.Limit, params.Offset)
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
		ret = append(ret, dto.toEntry())
	}
	return ret
}
