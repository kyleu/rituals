package invitation

import (
	"database/sql"
	"time"

	"github.com/kyleu/rituals.dev/app/database"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"
	"logur.dev/logur"
)

type Service struct {
	actions *action.Service
	db      *database.Service
	logger  logur.Logger
}

func NewService(service *action.Service, db *database.Service, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: util.KeyUser})

	return &Service{
		actions: service,
		db:      db,
		logger:  logger,
	}
}

func (s *Service) New(key string, k Type, v string, src *uuid.UUID, tgt *uuid.UUID, note string) (*Invitation, error) {
	s.logger.Info("creating invitation [" + key + "]")
	q := query.SQLInsert(util.KeyInvitation, []string{"key", "k", "v", "src", "tgt", "note", "status", "redeemed", "created"}, 1)
	dto := invitationDTO{
		Key:      key,
		K:        k.String(),
		V:        v,
		Src:      src,
		Tgt:      tgt,
		Note:     note,
		Status:   StatusPending.String(),
		Redeemed: nil,
		Created:  time.Now(),
	}
	err := s.db.Insert(q, nil, dto.Key, dto.K, dto.V, dto.Src, dto.Tgt, dto.Note, dto.Status, dto.Redeemed, dto.Created)
	if err != nil {
		return nil, err
	}
	return dto.ToInvitation(), nil
}

func (s *Service) List(params *query.Params) (Invitations, error) {
	params = query.ParamsWithDefaultOrdering(util.KeyInvitation, params, query.DefaultCreatedOrdering...)

	var dtos []invitationDTO
	q := query.SQLSelect("*", util.KeyInvitation, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		return nil, err
	}

	ret := make(Invitations, 0, len(dtos))

	for _, dto := range dtos {
		ret = append(ret, dto.ToInvitation())
	}

	return ret, nil
}

func (s *Service) GetByKey(key string) (*Invitation, error) {
	dto := &invitationDTO{}
	q := query.SQLSelect("*", util.KeyInvitation, "key = $1", "", 0, 0)
	err := s.db.Get(dto, q, nil, key)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return dto.ToInvitation(), nil
}
