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
	"golang.org/x/oauth2/slack"
)

func slackConf(secure bool, host string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     secrets.SlackClientID,
		ClientSecret: secrets.SlackClientSecret,
		Endpoint:     slack.Endpoint,
		RedirectURL:  callbackURL(secure, host, ProviderSlack.Key),
		Scopes:       []string{"users.profile:read"},
	}
}

type slackResponse struct {
	Ok      bool          `json:"ok"`
	Profile *slackProfile `json:"profile"`
}

type slackProfile struct {
	Email   string `json:"email"`
	Name    string `json:"real_name"`
	Picture string `json:"image_192"`
}

func slackAuth(tok *oauth2.Token) (*Record, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://slack.com/api/users.profile.get", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+tok.AccessToken)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error reading response from Slack"))
	}
	var rsp = slackResponse{}
	err = json.Unmarshal(contents, &rsp)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error marshalling slack user"))
	}

	ret := Record{
		ID:         util.UUID(),
		Provider:   &ProviderSlack,
		ProviderID: rsp.Profile.Email,
		Expires:    &tok.Expiry,
		Name:       rsp.Profile.Name,
		Email:      rsp.Profile.Email,
		Picture:    rsp.Profile.Picture,
		Created:    time.Time{},
	}
	return &ret, nil
}
