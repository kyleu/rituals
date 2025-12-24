package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vsprint"
)

func SprintList(w http.ResponseWriter, r *http.Request) {
	Act("sprint.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(cutil.QueryStringString(r, "q"))
		prms := ps.Params.Sanitized("sprint", ps.Logger)
		var ret sprint.Sprints
		var err error
		if q == "" {
			ret, err = as.Services.Sprint.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Sprint.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 && !cutil.IsContentTypeJSON(cutil.GetContentType(r)) {
				return FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
		}
		ps.SetTitleAndData("Sprints", ret)
		teamIDsByTeamID := lo.Map(ret, func(x *sprint.Sprint, _ int) *uuid.UUID {
			return x.TeamID
		})
		teamsByTeamID, err := as.Services.Team.GetMultiple(ps.Context, nil, nil, ps.Logger, util.ArrayDereference(teamIDsByTeamID)...)
		if err != nil {
			return "", err
		}
		page := &vsprint.List{Models: ret, TeamsByTeamID: teamsByTeamID, Params: ps.Params, SearchQuery: q}
		return Render(r, as, page, ps, "sprint")
	})
}

func SprintDetail(w http.ResponseWriter, r *http.Request) {
	Act("sprint.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := sprintFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Sprint)", ret)

		var teamByTeamID *team.Team
		if ret.TeamID != nil {
			teamByTeamID, _ = as.Services.Team.Get(ps.Context, nil, *ret.TeamID, ps.Logger)
		}

		relEstimatesBySprintIDPrms := ps.Params.Sanitized("estimate", ps.Logger)
		relEstimatesBySprintID, err := as.Services.Estimate.GetBySprintID(ps.Context, nil, &ret.ID, relEstimatesBySprintIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child estimates")
		}
		relRetrosBySprintIDPrms := ps.Params.Sanitized("retro", ps.Logger)
		relRetrosBySprintID, err := as.Services.Retro.GetBySprintID(ps.Context, nil, &ret.ID, relRetrosBySprintIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child retros")
		}
		relSprintHistoriesBySprintIDPrms := ps.Params.Sanitized("shistory", ps.Logger)
		relSprintHistoriesBySprintID, err := as.Services.SprintHistory.GetBySprintID(ps.Context, nil, ret.ID, relSprintHistoriesBySprintIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		relSprintMembersBySprintIDPrms := ps.Params.Sanitized("smember", ps.Logger)
		relSprintMembersBySprintID, err := as.Services.SprintMember.GetBySprintID(ps.Context, nil, ret.ID, relSprintMembersBySprintIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relSprintPermissionsBySprintIDPrms := ps.Params.Sanitized("spermission", ps.Logger)
		relSprintPermissionsBySprintID, err := as.Services.SprintPermission.GetBySprintID(ps.Context, nil, ret.ID, relSprintPermissionsBySprintIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		relStandupsBySprintIDPrms := ps.Params.Sanitized("standup", ps.Logger)
		relStandupsBySprintID, err := as.Services.Standup.GetBySprintID(ps.Context, nil, &ret.ID, relStandupsBySprintIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child standups")
		}
		return Render(r, as, &vsprint.Detail{
			Model:        ret,
			TeamByTeamID: teamByTeamID,
			Params:       ps.Params,

			RelEstimatesBySprintID:         relEstimatesBySprintID,
			RelRetrosBySprintID:            relRetrosBySprintID,
			RelSprintHistoriesBySprintID:   relSprintHistoriesBySprintID,
			RelSprintMembersBySprintID:     relSprintMembersBySprintID,
			RelSprintPermissionsBySprintID: relSprintPermissionsBySprintID,
			RelStandupsBySprintID:          relStandupsBySprintID,
		}, ps, "sprint", ret.TitleString()+"**sprint")
	})
}

func SprintCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("sprint.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &sprint.Sprint{}
		if cutil.QueryStringString(r, "prototype") == util.KeyRandom {
			ret = sprint.RandomSprint()
			randomTeam, err := as.Services.Team.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomTeam != nil {
				ret.TeamID = &randomTeam.ID
			}
		}
		ps.SetTitleAndData("Create [Sprint]", ret)
		return Render(r, as, &vsprint.Edit{Model: ret, IsNew: true}, ps, "sprint", "Create")
	})
}

func SprintRandom(w http.ResponseWriter, r *http.Request) {
	Act("sprint.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Sprint.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Sprint")
		}
		return ret.WebPath(), nil
	})
}

func SprintCreate(w http.ResponseWriter, r *http.Request) {
	Act("sprint.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := sprintFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Sprint from form")
		}
		err = as.Services.Sprint.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Sprint")
		}
		msg := fmt.Sprintf("Sprint [%s] created", ret.TitleString())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func SprintEditForm(w http.ResponseWriter, r *http.Request) {
	Act("sprint.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := sprintFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vsprint.Edit{Model: ret}, ps, "sprint", ret.String())
	})
}

func SprintEdit(w http.ResponseWriter, r *http.Request) {
	Act("sprint.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := sprintFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := sprintFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Sprint from form")
		}
		frm.ID = ret.ID
		err = as.Services.Sprint.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Sprint [%s]", frm.String())
		}
		msg := fmt.Sprintf("Sprint [%s] updated", frm.TitleString())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func SprintDelete(w http.ResponseWriter, r *http.Request) {
	Act("sprint.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := sprintFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Sprint.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete sprint [%s]", ret.String())
		}
		msg := fmt.Sprintf("Sprint [%s] deleted", ret.TitleString())
		return FlashAndRedir(true, msg, "/admin/db/sprint", ps)
	})
}

func sprintFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*sprint.Sprint, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Sprint.Get(ps.Context, nil, idArg, ps.Logger)
}

func sprintFromForm(r *http.Request, b []byte, setPK bool) (*sprint.Sprint, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := sprint.SprintFromMap(frm, setPK)
	return ret, err
}
