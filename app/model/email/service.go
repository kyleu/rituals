package email

import (
	"database/sql"
	"fmt"

	"github.com/kyleu/rituals.dev/app/database"
	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	logger  logur.Logger
	db      *database.Service
	Enabled bool
}

func NewService(db *database.Service, logger logur.Logger) *Service {
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

func (s *Service) List(params *query.Params) Emails {
	params = query.ParamsWithDefaultOrdering(util.KeyEmail, params, query.DefaultCreatedOrdering...)

	var dtos []emailDTO
	q := query.SQLSelect("*", util.KeyEmail, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)

	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving emails: %+v", err))
		return nil
	}

	return toEmails(dtos)
}

func (s *Service) GetByID(id string) *Email {
	var dto = &emailDTO{}
	q := query.SQLSelectSimple("*", util.KeyEmail, "id = $1")
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
	q := query.SQLInsert(util.KeyEmail, []string{util.KeyID, "recipients", "subject", "data", "plain", util.KeyHTML, util.WithDBID(util.KeyUser), util.KeyStatus}, 1)
	toS := database.ArrayToString(e.Recipients)
	json := util.ToJSON(e.Data, s.logger)
	return s.db.Insert(q, nil, e.ID, toS, e.Subject, json, e.Plain, e.HTML, e.UserID, e.Status)
}

func toEmails(dtos []emailDTO) Emails {
	ret := make(Emails, 0, len(dtos))

	for _, dto := range dtos {
		ret = append(ret, dto.toEmail())
	}

	return ret
}
