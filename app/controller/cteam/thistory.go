// Package cteam - Content managed by Project Forge, see [projectforge.md] for details.
package cteam

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vteam/vthistory"
)

func TeamHistoryList(w http.ResponseWriter, r *http.Request) {
	controller.Act("thistory.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("thistory", ps.Logger)
		ret, err := as.Services.TeamHistory.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Histories", ret)
		teamIDsByTeamID := lo.Map(ret, func(x *thistory.TeamHistory, _ int) uuid.UUID {
			return x.TeamID
		})
		teamsByTeamID, err := as.Services.Team.GetMultiple(ps.Context, nil, nil, ps.Logger, teamIDsByTeamID...)
		if err != nil {
			return "", err
		}
		page := &vthistory.List{Models: ret, TeamsByTeamID: teamsByTeamID, Params: ps.Params}
		return controller.Render(r, as, page, ps, "team", "thistory")
	})
}

func TeamHistoryDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("thistory.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := thistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (History)", ret)

		teamByTeamID, _ := as.Services.Team.Get(ps.Context, nil, ret.TeamID, ps.Logger)

		return controller.Render(r, as, &vthistory.Detail{Model: ret, TeamByTeamID: teamByTeamID}, ps, "team", "thistory", ret.TitleString()+"**history")
	})
}

func TeamHistoryCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("thistory.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &thistory.TeamHistory{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = thistory.Random()
			randomTeam, err := as.Services.Team.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomTeam != nil {
				ret.TeamID = randomTeam.ID
			}
		}
		ps.SetTitleAndData("Create [TeamHistory]", ret)
		ps.Data = ret
		return controller.Render(r, as, &vthistory.Edit{Model: ret, IsNew: true}, ps, "team", "thistory", "Create")
	})
}

func TeamHistoryRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("thistory.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.TeamHistory.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random TeamHistory")
		}
		return ret.WebPath(), nil
	})
}

func TeamHistoryCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("thistory.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := thistoryFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamHistory from form")
		}
		err = as.Services.TeamHistory.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created TeamHistory")
		}
		msg := fmt.Sprintf("TeamHistory [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func TeamHistoryEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("thistory.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := thistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &vthistory.Edit{Model: ret}, ps, "team", "thistory", ret.String())
	})
}

func TeamHistoryEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("thistory.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := thistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := thistoryFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamHistory from form")
		}
		frm.Slug = ret.Slug
		err = as.Services.TeamHistory.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update TeamHistory [%s]", frm.String())
		}
		msg := fmt.Sprintf("TeamHistory [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func TeamHistoryDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("thistory.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := thistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.TeamHistory.Delete(ps.Context, nil, ret.Slug, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete history [%s]", ret.String())
		}
		msg := fmt.Sprintf("TeamHistory [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/team/history", ps)
	})
}

func thistoryFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*thistory.TeamHistory, error) {
	slugArg, err := cutil.PathString(r, "slug", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [slug] as a string argument")
	}
	return as.Services.TeamHistory.Get(ps.Context, nil, slugArg, ps.Logger)
}

func thistoryFromForm(r *http.Request, b []byte, setPK bool) (*thistory.TeamHistory, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return thistory.FromMap(frm, setPK)
}
