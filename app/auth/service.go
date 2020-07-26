package auth

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/database"
	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/user"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	Enabled          bool
	EnabledProviders Providers
	redir            string
	actions          *action.Service
	db               *database.Service
	logger           logur.Logger
	users            *user.Service
}

func NewService(enabled bool, redir string, actions *action.Service, db *database.Service, logger logur.Logger, users *user.Service) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{util.KeyService: util.KeyAuth})

	if !strings.HasPrefix(redir, "http") {
		redir = "https://" + redir
	}
	if !strings.HasSuffix(redir, "/") {
		redir += "/"
	}

	svc := Service{
		Enabled: enabled,
		redir:   redir,
		actions: actions,
		db:      db,
		logger:  logger,
		users:   users,
	}

	for _, p := range AllProviders {
		cfg := svc.getConfig(p)
		if cfg != nil {
			svc.EnabledProviders = append(svc.EnabledProviders, p)
		}
	}
	if len(svc.EnabledProviders) == 0 {
		svc.Enabled = false
	} else {
		logger.Info("auth service started for [" + strings.Join(svc.EnabledProviders.Names(), ", ") + "]")
	}

	return &svc
}

func (s *Service) GetDisplayByUserID(userID uuid.UUID, params *query.Params) (Records, Displays) {
	if !s.Enabled {
		return nil, nil
	}

	params = query.ParamsWithDefaultOrdering(util.KeyAuth, params, query.DefaultCreatedOrdering...)
	var dtos []recordDTO
	q := query.SQLSelect("*", util.KeyAuth, "user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving auth entries for user [%v]: %+v", userID, err))
		return nil, nil
	}
	rec := make(Records, 0, len(dtos))
	for _, dto := range dtos {
		rec = append(rec, dto.toRecord())
	}
	disp := make(Displays, 0, len(rec))
	for _, r := range rec {
		disp = append(disp, r.ToDisplay())
	}
	return rec, disp
}

func (s *Service) FullURL(path string) string {
	return s.redir + strings.TrimPrefix(path, "/")
}
