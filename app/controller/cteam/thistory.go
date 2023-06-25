// Content managed by Project Forge, see [projectforge.md] for details.
package cteam

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/views/vteam/vthistory"
)

func TeamHistoryList(rc *fasthttp.RequestCtx) {
	controller.Act("thistory.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("thistory", nil, ps.Logger).Sanitize("thistory")
		ret, err := as.Services.TeamHistory.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Histories"
		ps.Data = ret
		teamIDsByTeamID := lo.Map(ret, func(x *thistory.TeamHistory, _ int) uuid.UUID {
			return x.TeamID
		})
		teamsByTeamID, err := as.Services.Team.GetMultiple(ps.Context, nil, ps.Logger, teamIDsByTeamID...)
		if err != nil {
			return "", err
		}
		page := &vthistory.List{Models: ret, TeamsByTeamID: teamsByTeamID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "team", "thistory")
	})
}

func TeamHistoryDetail(rc *fasthttp.RequestCtx) {
	controller.Act("thistory.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := thistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (History)"
		ps.Data = ret

		teamByTeamID, _ := as.Services.Team.Get(ps.Context, nil, ret.TeamID, ps.Logger)

		return controller.Render(rc, as, &vthistory.Detail{Model: ret, TeamByTeamID: teamByTeamID}, ps, "team", "thistory", ret.String())
	})
}

func TeamHistoryCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("thistory.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &thistory.TeamHistory{}
		ps.Title = "Create [TeamHistory]"
		ps.Data = ret
		return controller.Render(rc, as, &vthistory.Edit{Model: ret, IsNew: true}, ps, "team", "thistory", "Create")
	})
}

func TeamHistoryCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("thistory.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := thistory.Random()
		ps.Title = "Create Random TeamHistory"
		ps.Data = ret
		return controller.Render(rc, as, &vthistory.Edit{Model: ret, IsNew: true}, ps, "team", "thistory", "Create")
	})
}

func TeamHistoryCreate(rc *fasthttp.RequestCtx) {
	controller.Act("thistory.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := thistoryFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamHistory from form")
		}
		err = as.Services.TeamHistory.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created TeamHistory")
		}
		msg := fmt.Sprintf("TeamHistory [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func TeamHistoryEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("thistory.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := thistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vthistory.Edit{Model: ret}, ps, "team", "thistory", ret.String())
	})
}

func TeamHistoryEdit(rc *fasthttp.RequestCtx) {
	controller.Act("thistory.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := thistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := thistoryFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamHistory from form")
		}
		frm.Slug = ret.Slug
		err = as.Services.TeamHistory.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update TeamHistory [%s]", frm.String())
		}
		msg := fmt.Sprintf("TeamHistory [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func TeamHistoryDelete(rc *fasthttp.RequestCtx) {
	controller.Act("thistory.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := thistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.TeamHistory.Delete(ps.Context, nil, ret.Slug, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete history [%s]", ret.String())
		}
		msg := fmt.Sprintf("TeamHistory [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/teamHistory", rc, ps)
	})
}

func thistoryFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*thistory.TeamHistory, error) {
	slugArg, err := cutil.RCRequiredString(rc, "slug", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [slug] as a string argument")
	}
	return as.Services.TeamHistory.Get(ps.Context, nil, slugArg, ps.Logger)
}

func thistoryFromForm(rc *fasthttp.RequestCtx, setPK bool) (*thistory.TeamHistory, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return thistory.FromMap(frm, setPK)
}
