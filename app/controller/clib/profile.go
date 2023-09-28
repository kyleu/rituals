// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"net/url"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/csession"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/theme"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vprofile"
)

func Profile(rc *fasthttp.RequestCtx) {
	controller.Act("profile", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return profileAction(rc, as, ps)
	})
}

func ProfileSite(rc *fasthttp.RequestCtx) {
	controller.ActSite("profile", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return profileAction(rc, as, ps)
	})
}

func profileAction(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (string, error) {
	ps.Title = "Profile"
	ps.Data = ps.Profile
	thm := as.Themes.Get(ps.Profile.Theme, ps.Logger)

	prvs, err := as.Auth.Providers(ps.Logger)
	if err != nil {
		return "", errors.Wrap(err, "can't load providers")
	}

	redir := "/"
	ref := string(rc.Request.Header.Peek("Referer"))
	if ref != "" {
		u, err := url.Parse(ref)
		if err == nil && u != nil && u.Path != cutil.DefaultProfilePath {
			redir = u.Path
		}
	}

	page := &vprofile.Profile{Profile: ps.Profile, Theme: thm, Providers: prvs, Referrer: redir}
	return controller.Render(rc, as, page, ps, "Profile")
}

func ProfileSave(rc *fasthttp.RequestCtx) {
	controller.Act("profile.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}

		n := ps.Profile.Clone()

		referrerDefault := frm.GetStringOpt("referrer")
		if referrerDefault == "" {
			referrerDefault = cutil.DefaultProfilePath
		}

		n.Name = frm.GetStringOpt("name")
		n.Mode = frm.GetStringOpt("mode")
		n.Theme = frm.GetStringOpt("theme")
		if n.Theme == theme.Default.Key {
			n.Theme = ""
		}
		if ps.Profile.ID == util.UUIDDefault {
			n.ID = util.UUID()
		} else {
			n.ID = ps.Profile.ID
		}

		err = csession.SaveProfile(n, rc, ps.Session, ps.Logger)
		if err != nil {
			return "", err
		}

		curr, _ := as.Services.User.Get(ps.Context, nil, ps.Profile.ID, ps.Logger)
		if curr != nil {
			curr.Name = n.Name
			if curr.Picture == "" {
				curr.Picture = ps.Accounts.Image()
			}
			err = as.Services.User.Update(ps.Context, nil, curr, ps.Logger)
			if err != nil {
				return "", err
			}
		}

		return controller.ReturnToReferrer("Saved profile", referrerDefault, rc, ps)
	})
}
