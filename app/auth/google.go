package auth

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"time"
)

type googleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
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

