package member

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"github.com/kyleu/npn/npnservice/user"

	"emperror.dev/errors"

	"logur.dev/logur"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
)

type Service struct {
	actions   *action.Service
	users     user.Service
	db        *npndatabase.Service
	logger    logur.Logger
	svc       util.Service
	tableName string
	colName   string
}

func NewService(actions *action.Service, users user.Service, db *npndatabase.Service, logger logur.Logger, svc util.Service) *Service {
	return &Service{
		actions:   actions,
		users:     users,
		db:        db,
		logger:    logger,
		svc:       svc,
		tableName: svc.Key + "_member",
		colName:   npncore.WithDBID(svc.Key),
	}
}

const nameClause = "(case when name = '' then (select name from system_user su where su.id = user_id) else name end) as name"
const pictureClause = "(case when picture = '' then (select picture from system_user su where su.id = user_id) else picture end) as picture"

func (s *Service) GetByModelID(id uuid.UUID, params *npncore.Params) Entries {
	var defaultOrdering = npncore.Orderings{{Column: npncore.KeyName, Asc: true}}
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyMember, params, defaultOrdering...)
	var dtos []entryDTO

	where := fmt.Sprintf("%s = $1", s.colName)
	cols := fmt.Sprintf("user_id, %s, %s, role, created", nameClause, pictureClause)
	q := npndatabase.SQLSelect(cols, s.tableName, where, params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, id)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving member entries for model [%v]: %+v", id, err))
		return nil
	}

	ret := make(Entries, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.toEntry())
	}
	return ret
}

func (s *Service) Get(modelID uuid.UUID, userID uuid.UUID) (*Entry, error) {
	dto := entryDTO{}
	cols := fmt.Sprintf("user_id, %s, %s, role, created", nameClause, pictureClause)
	where := fmt.Sprintf("%s = $1 and user_id = $2", s.colName)
	q := npndatabase.SQLSelectSimple(cols, s.tableName, where)
	err := s.db.Get(&dto, q, nil, modelID, userID)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return dto.toEntry(), nil
}

func (s *Service) Update(modelID uuid.UUID, userID uuid.UUID, name string, picture string) (*Entry, error) {
	cols := []string{"name", "picture"}
	q := npndatabase.SQLUpdate(s.tableName, cols, fmt.Sprintf("%v = $%v and user_id = $%v", s.colName, len(cols)+1, len(cols)+1+1))
	err := s.db.Insert(q, nil, name, picture, modelID, userID)
	if err != nil {
		return nil, err
	}
	return s.Get(modelID, userID)
}

func (s *Service) RemoveMember(modelID uuid.UUID, src uuid.UUID, target uuid.UUID) error {
	q := npndatabase.SQLDelete(s.tableName, fmt.Sprintf("%v = $1 and user_id = $2", s.colName))
	err := s.db.DeleteOne(q, nil, modelID, target)
	return errors.Wrap(err, "unable to remove member ["+target.String()+"] as ["+src.String()+"]")
}

func (s *Service) UpdateMember(modelID uuid.UUID, src uuid.UUID, target uuid.UUID, role string) (*Entry, error) {
	q := npndatabase.SQLUpdate(s.tableName, []string{"role"}, fmt.Sprintf("%v = $2 and user_id = $3", s.colName))
	err := s.db.DeleteOne(q, nil, role, modelID, target)
	if err != nil {
		return nil, errors.Wrap(err, "unable to update member ["+target.String()+"] as ["+src.String()+"]")
	}
	return s.Get(modelID, target)
}

func (s *Service) UpdateStatus(sessionID uuid.UUID, status string, userID uuid.UUID) error {
	cols := []string{npncore.KeyStatus}
	q := npndatabase.SQLUpdate(s.svc.Key, cols, fmt.Sprintf("%v = $%v", npncore.KeyID, len(cols)+1))
	err := s.db.UpdateOne(q, nil, status, sessionID)
	s.actions.Post(s.svc.Key, sessionID, userID, action.ActUpdate, nil)
	return errors.Wrap(err, "error updating retro session")
}
