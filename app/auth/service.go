package auth

import (
	"emperror.dev/errors"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/actions"
	"github.com/kyleu/rituals.dev/app/user"
	"github.com/kyleu/rituals.dev/app/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"logur.dev/logur"
)

type Service struct {
	actions    *actions.Service
	db         *sqlx.DB
	logger     logur.Logger
	users      *user.Service
	googleConf *oauth2.Config
	githubConf *oauth2.Config
}

func NewService(actions *actions.Service, db *sqlx.DB, logger logur.Logger, users *user.Service) *Service {
	googleConf := oauth2.Config{
		ClientID:     "37858401718-arnqoomkhbbm2r5494f763b06b6d43h5.apps.googleusercontent.com",
		ClientSecret: "e7CypFkwa7B_kvi9zOuTpeXu",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:6660/auth/callback/google",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
	}

	githubConf := oauth2.Config{
		ClientID:     "566d7b94991e79edfa71",
		ClientSecret: "d67a6c671eb850fa98a855c2c6b5946cdcea11c3",
		Endpoint:     github.Endpoint,
		RedirectURL:  "http://localhost:6660/auth/callback/github",
		Scopes: []string{
			"profile",
		},
	}

	logger = logur.WithFields(logger, map[string]interface{}{"service": "auth"})
	svc := Service{
		actions:    actions,
		db:         db,
		logger:     logger,
		users:      users,
		googleConf: &googleConf,
		githubConf: &githubConf,
	}
	return &svc
}

func (s *Service) Handle(profile *util.UserProfile, key string, code string) (*Record, error) {
	if profile == nil {
		return nil, errors.New("no user profile for auth")
	}

	cfg := s.getConfig(key)
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

	curr, err := s.GetByKV(record.K, record.V)
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
			s.logger.Warn("TODO insert auth record with matching users")
			return s.mergeProfile(profile, record)
		} else {
			s.logger.Warn("TODO insert auth record with conflicting users")
			return s.mergeProfile(profile, record)
		}
	}
}

func (s *Service) mergeProfile(p *util.UserProfile, record *Record) (*Record, error) {
	p.Name = record.Name
	p.Picture = record.Picture

	p, err := s.users.SaveProfile(p)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving user profile"))
	}
	return record, nil
}
