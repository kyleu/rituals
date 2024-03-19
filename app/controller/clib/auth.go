// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/auth"
	"github.com/kyleu/rituals/app/util"
)

const signinMsg = "signed in using %s as [%s]"

func AuthDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("auth.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prv, err := getProvider(as, r, ps.Logger)
		if err != nil {
			return "", err
		}
		u, _, err := auth.CompleteUserAuth(prv, w, r, ps.Session, ps.Logger)
		if err == nil {
			msg := fmt.Sprintf(signinMsg, auth.AvailableProviderNames[prv.ID], u.Email)
			return controller.ReturnToReferrer(msg, cutil.DefaultProfilePath, w, ps)
		}
		return auth.BeginAuthHandler(prv, w, r, ps.Session, ps.Logger)
	})
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {
	controller.Act("auth.callback", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prv, err := getProvider(as, r, ps.Logger)
		if err != nil {
			return "", err
		}
		u, _, err := auth.CompleteUserAuth(prv, w, r, ps.Session, ps.Logger)
		if err != nil {
			return "", err
		}
		msg := fmt.Sprintf(signinMsg, auth.AvailableProviderNames[prv.ID], u.Email)
		return controller.ReturnToReferrer(msg, cutil.DefaultProfilePath, w, ps)
	})
}

func AuthLogout(w http.ResponseWriter, r *http.Request) {
	controller.Act("auth.logout", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(r, "key", false)
		if err != nil {
			return "", err
		}
		err = auth.Logout(w, r, ps.Session, ps.Logger, key)
		if err != nil {
			return "", err
		}

		return ps.ProfilePath, nil
	})
}

func getProvider(as *app.State, r *http.Request, logger util.Logger) (*auth.Provider, error) {
	key, err := cutil.RCRequiredString(r, "key", false)
	if err != nil {
		return nil, err
	}
	prvs, err := as.Auth.Providers(logger)
	if err != nil {
		return nil, errors.Wrap(err, "can't load providers")
	}
	prv := prvs.Get(key)
	if prv == nil {
		return nil, errors.Errorf("no provider available with id [%s]", key)
	}
	return prv, nil
}
