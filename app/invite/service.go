package invite

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"
	"logur.dev/logur"
)

type Service struct {
	actions *action.Service
	db      *sqlx.DB
	logger  logur.Logger
}

func NewService(service *action.Service, db *sqlx.DB, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "user"})

	return &Service{
		actions: service,
		db:      db,
		logger:  logger,
	}
}

func (s *Service) New(key string, k InvitationType, v string, src *uuid.UUID, tgt *uuid.UUID, note string) (*Invitation, error) {
	s.logger.Info("creating invitation [" + key + "]")
	q := "insert into invitation (key, k, v, src, tgt, note, status, redeemed, created) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	dto := invitationDTO{
		Key:      key,
		K:        k.String(),
		V:        v,
		Src:      src,
		Tgt:      tgt,
		Note:     note,
		Status:   InvitationStatusPending.String(),
		Redeemed: nil,
		Created:  time.Now(),
	}
	_, err := s.db.Exec(q, dto.Key, dto.K, dto.V, dto.Src, dto.Tgt, dto.Note, dto.Status, dto.Redeemed, dto.Created)
	if err != nil {
		return nil, err
	}
	return dto.ToInvitation(), nil
}

func (s *Service) List(params *query.Params) ([]*Invitation, error) {
	params = query.ParamsWithDefaultOrdering("invite", params, &query.Ordering{Column: "created", Asc: false})
	var dtos []invitationDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", "invitation", "", params.OrderByString(), params.Limit, params.Offset))
	if err != nil {
		return nil, err
	}
	ret := make([]*Invitation, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToInvitation())
	}
	return ret, nil
}

func (s *Service) GetByKey(key string) (*Invitation, error) {
	dto := &invitationDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", "invitation", "key = $1", "", 0, 0), key)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return dto.ToInvitation(), nil
}
