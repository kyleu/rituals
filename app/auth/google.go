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
	"golang.org/x/oauth2/google"
)

func googleConf(secure bool, host string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     secrets.GoogleClientID,
		ClientSecret: secrets.GoogleClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  callbackURL(secure, host, ProviderGoogle.Key),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
	}
}

type googleUser struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func googleAuth(tok *oauth2.Token) (*Record, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken)
	if err != nil {
		return nil, err
	}

	defer func() { _ = response.Body.Close() }()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error reading Google response"))
	}
	var user = googleUser{}
	err = json.Unmarshal(contents, &user)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error marshalling Google user"))
	}

	ret := Record{
		ID:         util.UUID(),
		Provider:   &ProviderGoogle,
		ProviderID: user.ID,
		Expires:    &tok.Expiry,
		Name:       user.Name,
		Email:      user.Email,
		Picture:    user.Picture,
		Created:    time.Time{},
	}
	return &ret, nil
}
