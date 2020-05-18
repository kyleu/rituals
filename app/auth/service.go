package auth

import (
	"emperror.dev/errors"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/user"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	actions *action.Service
	db      *sqlx.DB
	logger  logur.Logger
	users   *user.Service
}

func NewService(actions *action.Service, db *sqlx.DB, logger logur.Logger, users *user.Service) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": "auth"})
	svc := Service{
		actions: actions,
		db:      db,
		logger:  logger,
		users:   users,
	}
	return &svc
}

func (s *Service) Handle(profile *util.UserProfile, key string, code string) (*Record, error) {
	if profile == nil {
		return nil, errors.New("no user profile for auth")
	}

	cfg := s.getConfig(false, "", key)
	if cfg == nil {
		return nil, errors.New("no auth config for [" + key + "]")
	}

	record, err := s.decodeRecord(key, code)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error retrieving auth profile"))
	}
	if record == nil {
		return nil, errors.WithStack(errors.New("cannot retrieve auth profile"))
	}
	record.UserID = profile.UserID

	curr, err := s.GetByProviderID(record.Provider.Key, record.ProviderID)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error retrieving auth record"))
	}
	if curr == nil {
		record, err = s.NewRecord(record)
		if err != nil {
			return nil, errors.WithStack(errors.Wrap(err, "error saving new auth record"))
		}
		return s.mergeProfile(profile, record)
	} else {
		if curr.UserID == profile.UserID {
			record.ID = curr.ID
			err = s.UpdateRecord(record)
			return s.mergeProfile(profile, record)
		} else {
			s.logger.Warn("TODO insert auth record with conflicting users")
			record, err = s.NewRecord(record)
			if err != nil {
				return nil, errors.WithStack(errors.Wrap(err, "error saving new auth record"))
			}
			return s.mergeProfile(profile, record)
		}
	}
}

func (s *Service) mergeProfile(p *util.UserProfile, record *Record) (*Record, error) {
	p.Name = record.Name
	if p.Name == "" {
		p.Name = record.Provider.Title + " User"
	}
	p.Picture = record.Picture

	_, err := s.users.SaveProfile(p)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving user profile"))
	}
	return record, nil
}
