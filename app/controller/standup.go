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
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vstandup"
)

func StandupList(w http.ResponseWriter, r *http.Request) {
	Act("standup.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("standup", ps.Logger)
		var ret standup.Standups
		var err error
		if q == "" {
			ret, err = as.Services.Standup.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Standup.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 && !cutil.IsContentTypeJSON(cutil.GetContentType(r)) {
				return FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
		}
		ps.SetTitleAndData("Standups", ret)
		teamIDsByTeamID := lo.Map(ret, func(x *standup.Standup, _ int) *uuid.UUID {
			return x.TeamID
		})
		teamsByTeamID, err := as.Services.Team.GetMultiple(ps.Context, nil, nil, ps.Logger, util.ArrayDereference(teamIDsByTeamID)...)
		if err != nil {
			return "", err
		}
		sprintIDsBySprintID := lo.Map(ret, func(x *standup.Standup, _ int) *uuid.UUID {
			return x.SprintID
		})
		sprintsBySprintID, err := as.Services.Sprint.GetMultiple(ps.Context, nil, nil, ps.Logger, util.ArrayDereference(sprintIDsBySprintID)...)
		if err != nil {
			return "", err
		}
		page := &vstandup.List{Models: ret, TeamsByTeamID: teamsByTeamID, SprintsBySprintID: sprintsBySprintID, Params: ps.Params, SearchQuery: q}
		return Render(r, as, page, ps, "standup")
	})
}

//nolint:lll
func StandupDetail(w http.ResponseWriter, r *http.Request) {
	Act("standup.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := standupFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Standup)", ret)

		var teamByTeamID *team.Team
		if ret.TeamID != nil {
			teamByTeamID, _ = as.Services.Team.Get(ps.Context, nil, *ret.TeamID, ps.Logger)
		}
		var sprintBySprintID *sprint.Sprint
		if ret.SprintID != nil {
			sprintBySprintID, _ = as.Services.Sprint.Get(ps.Context, nil, *ret.SprintID, ps.Logger)
		}

		relReportsByStandupIDPrms := ps.Params.Sanitized("report", ps.Logger)
		relReportsByStandupID, err := as.Services.Report.GetByStandupID(ps.Context, nil, ret.ID, relReportsByStandupIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child reports")
		}
		relStandupHistoriesByStandupIDPrms := ps.Params.Sanitized("uhistory", ps.Logger)
		relStandupHistoriesByStandupID, err := as.Services.StandupHistory.GetByStandupID(ps.Context, nil, ret.ID, relStandupHistoriesByStandupIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		relStandupMembersByStandupIDPrms := ps.Params.Sanitized("umember", ps.Logger)
		relStandupMembersByStandupID, err := as.Services.StandupMember.GetByStandupID(ps.Context, nil, ret.ID, relStandupMembersByStandupIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relStandupPermissionsByStandupIDPrms := ps.Params.Sanitized("upermission", ps.Logger)
		relStandupPermissionsByStandupID, err := as.Services.StandupPermission.GetByStandupID(ps.Context, nil, ret.ID, relStandupPermissionsByStandupIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		return Render(r, as, &vstandup.Detail{
			Model:            ret,
			TeamByTeamID:     teamByTeamID,
			SprintBySprintID: sprintBySprintID,
			Params:           ps.Params,

			RelReportsByStandupID:            relReportsByStandupID,
			RelStandupHistoriesByStandupID:   relStandupHistoriesByStandupID,
			RelStandupMembersByStandupID:     relStandupMembersByStandupID,
			RelStandupPermissionsByStandupID: relStandupPermissionsByStandupID,
		}, ps, "standup", ret.TitleString()+"**standup")
	})
}

func StandupCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("standup.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &standup.Standup{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = standup.RandomStandup()
			randomTeam, err := as.Services.Team.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomTeam != nil {
				ret.TeamID = &randomTeam.ID
			}
			randomSprint, err := as.Services.Sprint.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomSprint != nil {
				ret.SprintID = &randomSprint.ID
			}
		}
		ps.SetTitleAndData("Create [Standup]", ret)
		return Render(r, as, &vstandup.Edit{Model: ret, IsNew: true}, ps, "standup", "Create")
	})
}

func StandupRandom(w http.ResponseWriter, r *http.Request) {
	Act("standup.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Standup.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Standup")
		}
		return ret.WebPath(), nil
	})
}

func StandupCreate(w http.ResponseWriter, r *http.Request) {
	Act("standup.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := standupFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Standup from form")
		}
		err = as.Services.Standup.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Standup")
		}
		msg := fmt.Sprintf("Standup [%s] created", ret.TitleString())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func StandupEditForm(w http.ResponseWriter, r *http.Request) {
	Act("standup.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := standupFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vstandup.Edit{Model: ret}, ps, "standup", ret.String())
	})
}

func StandupEdit(w http.ResponseWriter, r *http.Request) {
	Act("standup.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := standupFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := standupFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Standup from form")
		}
		frm.ID = ret.ID
		err = as.Services.Standup.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Standup [%s]", frm.String())
		}
		msg := fmt.Sprintf("Standup [%s] updated", frm.TitleString())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func StandupDelete(w http.ResponseWriter, r *http.Request) {
	Act("standup.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := standupFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Standup.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete standup [%s]", ret.String())
		}
		msg := fmt.Sprintf("Standup [%s] deleted", ret.TitleString())
		return FlashAndRedir(true, msg, "/admin/db/standup", ps)
	})
}

func standupFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*standup.Standup, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Standup.Get(ps.Context, nil, idArg, ps.Logger)
}

func standupFromForm(r *http.Request, b []byte, setPK bool) (*standup.Standup, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := standup.StandupFromMap(frm, setPK)
	return ret, err
}
