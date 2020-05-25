package auth

import (
	"context"
	"emperror.dev/errors"
	"golang.org/x/oauth2"
	"os"
)

func (s *Service) callbackURL(secure bool, k string) string {
	if secure {
		return s.Redir + "/auth/callback/" + k
	}
	return s.Redir + "/auth/callback/" + k
}

func (s *Service) getConfig(secure bool, prv *Provider) *oauth2.Config {
	idKey, secretKey := envsFor(prv)
	id := os.Getenv(idKey)
	secret := os.Getenv(secretKey)
	if id == "" || secret == "" {
		return nil
	}

	ret := oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		Endpoint:     prv.Endpoint,
		RedirectURL:  s.callbackURL(secure, prv.Key),
		Scopes:       prv.Scopes,
	}

	return &ret
}

func (s *Service) URLFor(state string, secure bool, prv *Provider) string {
	cfg := s.getConfig(secure, prv)
	if cfg == nil {
		return ""
	}
	return cfg.AuthCodeURL(state)
}

func (s *Service) getToken(prv *Provider, code string) (*oauth2.Token, error) {
	cfg := s.getConfig(false, prv)
	if cfg == nil {
		return nil, errors.New("no auth config for [" + prv.Key + "]")
	}

	ctx := context.TODO()
	return cfg.Exchange(ctx, code)
}

func (s *Service) decodeRecord(prv *Provider, code string) (*Record, error) {
	tok, err := s.getToken(prv, code)
	if err != nil {
		return nil, errors.Wrap(err, "error getting token")
	}

	switch prv {
	case &ProviderGoogle:
		return googleAuth(tok)
	case &ProviderGitHub:
		return githubAuth(tok)
	case &ProviderSlack:
		return slackAuth(tok)
	case &ProviderAmazon:
		return amazonAuth(tok)
	case &ProviderMicrosoft:
		return microsoftAuth(tok)
	default:
		return nil, nil
	}
}

func envsFor(prv *Provider) (string, string) {
	var id = "rituals_client_id_" + prv.Key
	var secret = "rituals_client_secret_" + prv.Key
	return id, secret
}
