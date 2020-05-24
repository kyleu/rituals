package auth

import (
	"context"

	"emperror.dev/errors"
	"golang.org/x/oauth2"
)

func callbackURL(secure bool, host string, k string) string {
	if secure {
		return "https://" + host + "/auth/callback/" + k
	}
	return "http://" + host + "/auth/callback/" + k
}

func (s *Service) getConfig(secure bool, host string, key string) *oauth2.Config {
	if len(host) == 0 {
		host = "localhost:6660"
	}

	switch key {
	case ProviderGoogle.Key:
		return googleConf(secure, host)
	case ProviderGitHub.Key:
		return githubConf(secure, host)
	case ProviderSlack.Key:
		return slackConf(secure, host)
	case ProviderAmazon.Key:
		return amazonConf(secure, host)
	default:
		return nil
	}
}

func (s *Service) URLFor(state string, secure bool, host string, key string) string {
	cfg := s.getConfig(secure, host, key)
	if cfg == nil {
		return ""
	}
	return cfg.AuthCodeURL(state)
}

func (s *Service) getToken(key string, code string) (*oauth2.Token, error) {
	cfg := s.getConfig(false, "", key)
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
	case ProviderGoogle.Key:
		return googleAuth(tok)
	case ProviderGitHub.Key:
		return githubAuth(tok)
	case ProviderSlack.Key:
		return slackAuth(tok)
	case ProviderAmazon.Key:
		return amazonAuth(tok)
	default:
		return nil, nil
	}
}
