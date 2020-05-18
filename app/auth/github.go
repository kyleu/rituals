package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/secrets"
	"github.com/kyleu/rituals.dev/app/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func githubConf(secure bool, host string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     secrets.GitHubClientID,
		ClientSecret: secrets.GitHubClientSecret,
		Endpoint:     github.Endpoint,
		RedirectURL:  callbackURL(secure, host, ProviderGitHub.Key),
		Scopes:       []string{"profile"},
	}
}

type githubUser struct {
	ID      string `json:"login"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"avatar_url"`
}

func githubAuth(tok *oauth2.Token) (*Record, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "token "+tok.AccessToken)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = response.Body.Close() }()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error reading GitHub response"))
	}
	var user = githubUser{}
	err = json.Unmarshal(contents, &user)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error marshalling GitHub user"))
	}

	ret := Record{
		ID:         util.UUID(),
		Provider:   &ProviderGitHub,
		ProviderID: user.ID,
		Expires:    &tok.Expiry,
		Name:       user.Name,
		Email:      user.Email,
		Picture:    user.Picture,
		Created:    time.Time{},
	}
	return &ret, nil
}
