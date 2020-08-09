package email

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"logur.dev/logur"
)

type Service struct {
	logger  logur.Logger
	db      *npndatabase.Service
	Enabled bool
}

func NewService(db *npndatabase.Service, logger logur.Logger) *Service {
	cfg := getCfg()
	if !cfg.Enabled() {
		logger.Warn("email service in not enabled")
		logger.Warn("set the following environment variables to enable it:")
		for _, x := range []string{"host", "port", "username", "password", "from"} {
			logger.Warn(fmt.Sprintf("  - rituals_mail_%v", x))
		}
	}
	return &Service{logger: logger, db: db, Enabled: cfg.Enabled()}
}

func (s *Service) List(params *npncore.Params) Emails {
	params = npncore.ParamsWithDefaultOrdering(npncore.KeyEmail, params, npncore.DefaultCreatedOrdering...)

	var dtos []emailDTO
	q := npndatabase.SQLSelect("*", npncore.KeyEmail, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving emails: %+v", err))
		return nil
	}

	return toEmails(dtos)
}

func (s *Service) GetByID(id string) *Email {
	var dto = &emailDTO{}
	q := npndatabase.SQLSelectSimple("*", npncore.KeyEmail, "id = $1")
	err := s.db.Get(dto, q, nil, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		s.logger.Error(fmt.Sprintf("error getting email by id [%v]: %+v", id, err))
		return nil
	}
	return dto.toEmail()
}

func (s *Service) New(e Email) error {
	q := npndatabase.SQLInsert(npncore.KeyEmail, []string{npncore.KeyID, "recipients", "subject", "data", "plain", npncore.KeyHTML, npncore.WithDBID(npncore.KeyUser), npncore.KeyStatus}, 1)
	toS := npndatabase.ArrayToString(e.Recipients)
	json := npncore.ToJSON(e.Data, s.logger)
	return s.db.Insert(q, nil, e.ID, toS, e.Subject, json, e.Plain, e.HTML, e.UserID, e.Status)
}

func toEmails(dtos []emailDTO) Emails {
	ret := make(Emails, 0, len(dtos))

	for _, dto := range dtos {
		ret = append(ret, dto.toEmail())
	}

	return ret
}
