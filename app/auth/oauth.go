package auth

import (
	"context"
	"emperror.dev/errors"
	"golang.org/x/oauth2"
)

func (s *Service) getConfig(key string) *oauth2.Config {
	switch key {
	case "google":
		return &googleConf
	case "github":
		return &githubConf
	case "slack":
		return &slackConf
	default:
		return nil
	}
}

func (s *Service) UrlFor(key string) string {
	cfg := s.getConfig(key)
	if cfg == nil {
		return ""
	}
	return cfg.AuthCodeURL("state")
}

func (s *Service) getToken(key string, code string) (*oauth2.Token, error) {
	cfg := s.getConfig(key)
	if cfg == nil {
		return nil, errors.New("no auth config for [" + key + "]")
	}

	ctx := context.TODO()
	return cfg.Exchange(ctx, code)
}

func (s *Service) decodeRecord(key string, code string) (*Record, error) {
	tok, err := s.getToken(key, code)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error getting token"))
	}
	switch key {
	case "google":
		return googleAuth(tok)
	case "github":
		return githubAuth(tok)
	case "slack":
		return slackAuth(tok)
	default:
		return nil, nil
	}
}
