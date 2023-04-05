// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vteam"
)

func TeamList(rc *fasthttp.RequestCtx) {
	Act("team.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("team", nil, ps.Logger).Sanitize("team")
		var ret team.Teams
		var err error
		if q == "" {
			ret, err = as.Services.Team.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Team.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.Title = "Teams"
		ps.Data = ret
		page := &vteam.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(rc, as, page, ps, "team")
	})
}

func TeamDetail(rc *fasthttp.RequestCtx) {
	Act("team.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := teamFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Team)"
		ps.Data = ret

		relEstimatesByTeamIDPrms := ps.Params.Get("estimate", nil, ps.Logger).Sanitize("estimate")
		relEstimatesByTeamID, err := as.Services.Estimate.GetByTeamID(ps.Context, nil, &ret.ID, relEstimatesByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child estimates")
		}
		relRetrosByTeamIDPrms := ps.Params.Get("retro", nil, ps.Logger).Sanitize("retro")
		relRetrosByTeamID, err := as.Services.Retro.GetByTeamID(ps.Context, nil, &ret.ID, relRetrosByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child retros")
		}
		relSprintsByTeamIDPrms := ps.Params.Get("sprint", nil, ps.Logger).Sanitize("sprint")
		relSprintsByTeamID, err := as.Services.Sprint.GetByTeamID(ps.Context, nil, &ret.ID, relSprintsByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child sprints")
		}
		relStandupsByTeamIDPrms := ps.Params.Get("standup", nil, ps.Logger).Sanitize("standup")
		relStandupsByTeamID, err := as.Services.Standup.GetByTeamID(ps.Context, nil, &ret.ID, relStandupsByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child standups")
		}
		relTeamHistoriesByTeamIDPrms := ps.Params.Get("thistory", nil, ps.Logger).Sanitize("thistory")
		relTeamHistoriesByTeamID, err := as.Services.TeamHistory.GetByTeamID(ps.Context, nil, ret.ID, relTeamHistoriesByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		relTeamMembersByTeamIDPrms := ps.Params.Get("tmember", nil, ps.Logger).Sanitize("tmember")
		relTeamMembersByTeamID, err := as.Services.TeamMember.GetByTeamID(ps.Context, nil, ret.ID, relTeamMembersByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relTeamPermissionsByTeamIDPrms := ps.Params.Get("tpermission", nil, ps.Logger).Sanitize("tpermission")
		relTeamPermissionsByTeamID, err := as.Services.TeamPermission.GetByTeamID(ps.Context, nil, ret.ID, relTeamPermissionsByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		return Render(rc, as, &vteam.Detail{
			Model:                      ret,
			Params:                     ps.Params,
			RelEstimatesByTeamID:       relEstimatesByTeamID,
			RelRetrosByTeamID:          relRetrosByTeamID,
			RelSprintsByTeamID:         relSprintsByTeamID,
			RelStandupsByTeamID:        relStandupsByTeamID,
			RelTeamHistoriesByTeamID:   relTeamHistoriesByTeamID,
			RelTeamMembersByTeamID:     relTeamMembersByTeamID,
			RelTeamPermissionsByTeamID: relTeamPermissionsByTeamID,
		}, ps, "team", ret.String())
	})
}

func TeamCreateForm(rc *fasthttp.RequestCtx) {
	Act("team.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &team.Team{}
		ps.Title = "Create [Team]"
		ps.Data = ret
		return Render(rc, as, &vteam.Edit{Model: ret, IsNew: true}, ps, "team", "Create")
	})
}

func TeamCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("team.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := team.Random()
		ps.Title = "Create Random Team"
		ps.Data = ret
		return Render(rc, as, &vteam.Edit{Model: ret, IsNew: true}, ps, "team", "Create")
	})
}

func TeamCreate(rc *fasthttp.RequestCtx) {
	Act("team.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := teamFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Team from form")
		}
		err = as.Services.Team.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Team")
		}
		msg := fmt.Sprintf("Team [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func TeamEditForm(rc *fasthttp.RequestCtx) {
	Act("team.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := teamFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vteam.Edit{Model: ret}, ps, "team", ret.String())
	})
}

func TeamEdit(rc *fasthttp.RequestCtx) {
	Act("team.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := teamFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := teamFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Team from form")
		}
		frm.ID = ret.ID
		err = as.Services.Team.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Team [%s]", frm.String())
		}
		msg := fmt.Sprintf("Team [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func TeamDelete(rc *fasthttp.RequestCtx) {
	Act("team.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := teamFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Team.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete team [%s]", ret.String())
		}
		msg := fmt.Sprintf("Team [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/team", rc, ps)
	})
}

func teamFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*team.Team, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Team.Get(ps.Context, nil, idArg, ps.Logger)
}

func teamFromForm(rc *fasthttp.RequestCtx, setPK bool) (*team.Team, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return team.FromMap(frm, setPK)
}
