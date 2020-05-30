package auth

import (
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/database"
	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/model/action"
	"github.com/kyleu/rituals.dev/app/model/user"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	Enabled          bool
	EnabledProviders Providers
	Redir            string
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

	svc := Service{
		Enabled: enabled,
		Redir:   redir,
		actions: actions,
		db:      db,
		logger:  logger,
		users:   users,
	}

	for _, p := range AllProviders {
		cfg := svc.getConfig(false, p)
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

	params = query.ParamsWithDefaultOrdering(util.KeyMember, params, query.DefaultCreatedOrdering...)
	var dtos []recordDTO
	q := query.SQLSelect("*", util.KeyAuth, "user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		util.LogError(s.logger, "error retrieving auth entries for user [%v]: %+v", userID, err)
		return nil, nil
	}
	rec := make(Records, 0, len(dtos))
	for _, dto := range dtos {
		rec = append(rec, dto.ToRecord())
	}
	disp := make(Displays, 0, len(rec))
	for _, r := range rec {
		disp = append(disp, r.ToDisplay())
	}
	return rec, disp
}

func (s *Service) Handle(profile *util.UserProfile, prv *Provider, code string) (*Record, error) {
	if !s.Enabled {
		return nil, ErrorAuthDisabled
	}

	if profile == nil {
		return nil, errors.New("no user profile for auth")
	}

	cfg := s.getConfig(false, prv)
	if cfg == nil {
		return nil, errors.New("no auth config for [" + prv.Key + "]")
	}

	record, err := s.decodeRecord(prv, code)
	if err != nil {
		return nil, errors.Wrap(err, "error retrieving auth profile")
	}
	if record == nil {
		return nil, errors.New("cannot retrieve auth profile")
	}
	record.UserID = profile.UserID

	curr := s.GetByProviderID(record.Provider.Key, record.ProviderID)
	if curr == nil {
		record, err = s.NewRecord(record)
		if err != nil {
			return nil, errors.Wrap(err, "error saving new auth record")
		}

		return s.mergeProfile(profile, record)
	}
	if curr.UserID == profile.UserID {
		record.ID = curr.ID

		err = s.UpdateRecord(record)
		if err != nil {
			return nil, errors.Wrap(err, "error updating auth record")
		}

		return s.mergeProfile(profile, record)
	}

	record, err = s.NewRecord(record)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new auth record")
	}

	return s.mergeProfile(profile, record)
}

func (s *Service) mergeProfile(p *util.UserProfile, record *Record) (*Record, error) {
	p.Name = record.Name
	if p.Name == "" {
		p.Name = record.Provider.Title + " User"
	}
	p.Picture = record.Picture

	_, err := s.users.SaveProfile(p)
	if err != nil {
		return nil, errors.Wrap(err, "error saving user profile")
	}

	return record, nil
}
