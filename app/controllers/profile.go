package controllers

import (
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npntemplate/gen/npntemplate"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnweb"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kyleu/npn/npnservice/auth"
)

type profileForm struct {
	Username  string `mapstructure:"username"`
	Theme     string `mapstructure:"theme"`
	LinkColor string `mapstructure:"linkColor"`
	NavColor  string `mapstructure:"navColor"`
	Ref       string `mapstructure:"ref"`
}

func Profile(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		params := npnweb.ParamSetFromRequest(r)
		auths := ctx.App.Auth().GetByUserID(ctx.Profile.UserID, params.Get(npncore.KeyAuth, ctx.Logger))
		ctx.Title = "User Profile"
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(npncore.KeyProfile)}
		ref := r.Header.Get("Referer")
		return npncontroller.T(npntemplate.Profile(auths, ref, ctx, w))
	})
}

func ProfileSave(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		prof := &profileForm{}
		err := npnweb.Decode(r, prof, ctx.Logger)
		if err != nil {
			return npncontroller.EResp(err)
		}

		if len(strings.TrimSpace(prof.Username)) == 0 {
			prof.Username = "Guest"
		}

		ctx.Profile.Name = strings.TrimSpace(prof.Username)
		ctx.Profile.Theme = npnuser.ThemeFromString(prof.Theme)
		ctx.Profile.NavColor = prof.NavColor
		ctx.Profile.LinkColor = prof.LinkColor

		_ = ctx.App.User().GetByID(ctx.Profile.UserID, true)
		_, err = ctx.App.User().SaveProfile(ctx.Profile)
		if err != nil {
			return npncontroller.EResp(err)
		}
		ref := strings.TrimSpace(prof.Ref)
		if len(ref) == 0 || strings.HasPrefix(ref, "http") {
			ref = "home"
		}
		return npncontroller.FlashAndRedir(true, "Profile saved", ref, w, r, ctx)
	})
}

func ProfilePic(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		if !ctx.App.Auth().Enabled {
			return "", auth.ErrorAuthDisabled
		}
		id, err := npnweb.IDFromParams(npncore.KeyID, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err, "invalid id")
		}
		a := ctx.App.Auth().GetByID(*id)
		ctx.Profile.Picture = a.Picture
		_, err = ctx.App.User().SaveProfile(ctx.Profile)
		if err != nil {
			return npncontroller.EResp(err, "can't save profile")
		}

		return ctx.Route(npncore.KeyProfile), nil
	})
}

func ProfileTheme(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		theme := npnuser.ThemeFromString(key)
		ctx.Profile.Theme = theme
		ctx.App.User().GetByID(ctx.Profile.UserID, true)
		_, err := ctx.App.User().SaveProfile(ctx.Profile)
		if err != nil {
			return npncontroller.EResp(err, "can't save profile")
		}

		_, err = w.Write([]byte(""))
		return npncontroller.EResp(err)
	})
}
