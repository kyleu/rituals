package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2/amazon"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/secrets"
	"github.com/kyleu/rituals.dev/app/util"
	"golang.org/x/oauth2"
)

func amazonConf(secure bool, host string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     secrets.AmazonClientID,
		ClientSecret: secrets.AmazonClientSecret,
		Endpoint:     amazon.Endpoint,
		RedirectURL:  callbackURL(secure, host, ProviderAmazon.Key),
		Scopes:       []string{"profile"},
	}
}

type amazonUser struct {
	ID      string `json:"login"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"avatar_url"`
}

func amazonAuth(tok *oauth2.Token) (*Record, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.amazon.com/user/profile", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+tok.AccessToken)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = response.Body.Close() }()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error reading Amazon response"))
	}

	var user = amazonUser{}
	err = json.Unmarshal(contents, &user)

	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error marshalling Amazon user"))
	}

	ret := Record{
		ID:         util.UUID(),
		Provider:   &ProviderAmazon,
		ProviderID: user.ID,
		Expires:    &tok.Expiry,
		Name:       user.Name,
		Email:      user.Email,
		Picture:    user.Picture,
		Created:    time.Time{},
	}
	return &ret, nil
}
