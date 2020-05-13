package auth

import (
	"context"
	"emperror.dev/errors"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"time"
)

func (s *Service) getConfig(key string) *oauth2.Config {
	switch key {
	case "google":
		return s.googleConf
	case "github":
		return s.githubConf
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
	default:
		return nil, nil
	}
}

func googleAuth(tok *oauth2.Token) (*Record, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken)
	if err != nil {
		return nil, err
	}
	defer func() { _ = response.Body.Close() }()

	contents, err := ioutil.ReadAll(response.Body)
	var user = googleUser{}
	err = json.Unmarshal(contents, &user)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error marshalling google user"))
	}

	ret := Record{
		ID:      util.UUID(),
		UserID:  uuid.UUID{},
		K:       "google",
		V:       user.ID,
		Expires: &tok.Expiry,
		Name:    user.Name,
		Email:   user.Email,
		Picture: user.Picture,
		Created: time.Time{},
	}
	return &ret, nil
}

func githubAuth(tok *oauth2.Token) (*Record, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "token " + tok.AccessToken)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = response.Body.Close() }()

	contents, err := ioutil.ReadAll(response.Body)
	var user = githubUser{}
	err = json.Unmarshal(contents, &user)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error marshalling google user"))
	}

	ret := Record{
		ID:      util.UUID(),
		K:       "github",
		V:       user.ID,
		Expires: &tok.Expiry,
		Name:    user.Name,
		Email:   user.Email,
		Picture: user.Picture,
		Created: time.Time{},
	}
	return &ret, nil
}

