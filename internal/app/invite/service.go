package invite

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"logur.dev/logur"
	"time"
)

type Service struct {
	db      *sqlx.DB
	logger  logur.Logger
}

func NewInviteService(db *sqlx.DB, logger logur.LoggerFacade) Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "user"})

	return Service{
		db:      db,
		logger:  logger,
	}
}

func (s *Service) List() ([]Invitation, error) {
	dtos := []invitationDTO{}
	err := s.db.Select(&dtos, "select * from invitation")
	if err != nil {
		return nil, err
	}
	ret := make([]Invitation, 0)
	for _, dto := range dtos {
		ret = append(ret, dto.ToInvitation())
	}
	return ret, nil
}

func (s *Service) GetByKey(key string) (*Invitation, error) {
	dto := &invitationDTO{}
	err := s.db.Get(dto, "select * from invitation where key = $1", key)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	ret := dto.ToInvitation()
	return &ret, nil
}

func (s *Service) CreateNewInvitation(key string, k InvitationType, v string, src *uuid.UUID, tgt *uuid.UUID, note string) (*Invitation, error) {
	s.logger.Info("creating invitation [" + key + "]")
	q := "insert into invitation (key, k, v, src, tgt, note, status, redeemed, created) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	i := invitationDTO{
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
	_, err := s.db.Exec(q, i.Key, i.K, i.V, i.Src, i.Tgt, i.Note, i.Status, i.Redeemed, i.Created)
	if err != nil {
		return nil, err
	}
	ret := i.ToInvitation()
	return &ret, nil
}
