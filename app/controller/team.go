package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vteam"
)

func TeamList(w http.ResponseWriter, r *http.Request) {
	Act("team.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("team", ps.Logger)
		var ret team.Teams
		var err error
		if q == "" {
			ret, err = as.Services.Team.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Team.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 && !cutil.IsContentTypeJSON(cutil.GetContentType(r)) {
				return FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
		}
		ps.SetTitleAndData("Teams", ret)
		page := &vteam.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(r, as, page, ps, "team")
	})
}

func TeamDetail(w http.ResponseWriter, r *http.Request) {
	Act("team.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := teamFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Team)", ret)

		relEstimatesByTeamIDPrms := ps.Params.Sanitized("estimate", ps.Logger)
		relEstimatesByTeamID, err := as.Services.Estimate.GetByTeamID(ps.Context, nil, &ret.ID, relEstimatesByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child estimates")
		}
		relRetrosByTeamIDPrms := ps.Params.Sanitized("retro", ps.Logger)
		relRetrosByTeamID, err := as.Services.Retro.GetByTeamID(ps.Context, nil, &ret.ID, relRetrosByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child retros")
		}
		relSprintsByTeamIDPrms := ps.Params.Sanitized("sprint", ps.Logger)
		relSprintsByTeamID, err := as.Services.Sprint.GetByTeamID(ps.Context, nil, &ret.ID, relSprintsByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child sprints")
		}
		relStandupsByTeamIDPrms := ps.Params.Sanitized("standup", ps.Logger)
		relStandupsByTeamID, err := as.Services.Standup.GetByTeamID(ps.Context, nil, &ret.ID, relStandupsByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child standups")
		}
		relTeamHistoriesByTeamIDPrms := ps.Params.Sanitized("thistory", ps.Logger)
		relTeamHistoriesByTeamID, err := as.Services.TeamHistory.GetByTeamID(ps.Context, nil, ret.ID, relTeamHistoriesByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		relTeamMembersByTeamIDPrms := ps.Params.Sanitized("tmember", ps.Logger)
		relTeamMembersByTeamID, err := as.Services.TeamMember.GetByTeamID(ps.Context, nil, ret.ID, relTeamMembersByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relTeamPermissionsByTeamIDPrms := ps.Params.Sanitized("tpermission", ps.Logger)
		relTeamPermissionsByTeamID, err := as.Services.TeamPermission.GetByTeamID(ps.Context, nil, ret.ID, relTeamPermissionsByTeamIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		return Render(r, as, &vteam.Detail{
			Model:  ret,
			Params: ps.Params,

			RelEstimatesByTeamID:       relEstimatesByTeamID,
			RelRetrosByTeamID:          relRetrosByTeamID,
			RelSprintsByTeamID:         relSprintsByTeamID,
			RelStandupsByTeamID:        relStandupsByTeamID,
			RelTeamHistoriesByTeamID:   relTeamHistoriesByTeamID,
			RelTeamMembersByTeamID:     relTeamMembersByTeamID,
			RelTeamPermissionsByTeamID: relTeamPermissionsByTeamID,
		}, ps, "team", ret.TitleString()+"**team")
	})
}

func TeamCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("team.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &team.Team{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = team.RandomTeam()
		}
		ps.SetTitleAndData("Create [Team]", ret)
		return Render(r, as, &vteam.Edit{Model: ret, IsNew: true}, ps, "team", "Create")
	})
}

func TeamRandom(w http.ResponseWriter, r *http.Request) {
	Act("team.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Team.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Team")
		}
		return ret.WebPath(), nil
	})
}

func TeamCreate(w http.ResponseWriter, r *http.Request) {
	Act("team.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := teamFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Team from form")
		}
		err = as.Services.Team.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Team")
		}
		msg := fmt.Sprintf("Team [%s] created", ret.TitleString())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func TeamEditForm(w http.ResponseWriter, r *http.Request) {
	Act("team.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := teamFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vteam.Edit{Model: ret}, ps, "team", ret.String())
	})
}

func TeamEdit(w http.ResponseWriter, r *http.Request) {
	Act("team.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := teamFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := teamFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Team from form")
		}
		frm.ID = ret.ID
		err = as.Services.Team.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Team [%s]", frm.String())
		}
		msg := fmt.Sprintf("Team [%s] updated", frm.TitleString())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func TeamDelete(w http.ResponseWriter, r *http.Request) {
	Act("team.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := teamFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Team.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete team [%s]", ret.String())
		}
		msg := fmt.Sprintf("Team [%s] deleted", ret.TitleString())
		return FlashAndRedir(true, msg, "/admin/db/team", ps)
	})
}

func teamFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*team.Team, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
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

func teamFromForm(r *http.Request, b []byte, setPK bool) (*team.Team, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := team.TeamFromMap(frm, setPK)
	return ret, err
}
