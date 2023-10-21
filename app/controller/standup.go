// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vstandup"
)

func StandupList(rc *fasthttp.RequestCtx) {
	Act("standup.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("standup", nil, ps.Logger).Sanitize("standup")
		var ret standup.Standups
		var err error
		if q == "" {
			ret, err = as.Services.Standup.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Standup.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.Title = "Standups"
		ps.Data = ret
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
		return Render(rc, as, page, ps, "standup")
	})
}

//nolint:lll
func StandupDetail(rc *fasthttp.RequestCtx) {
	Act("standup.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := standupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Standup)"
		ps.Data = ret

		var teamByTeamID *team.Team
		if ret.TeamID != nil {
			teamByTeamID, _ = as.Services.Team.Get(ps.Context, nil, *ret.TeamID, ps.Logger)
		}
		var sprintBySprintID *sprint.Sprint
		if ret.SprintID != nil {
			sprintBySprintID, _ = as.Services.Sprint.Get(ps.Context, nil, *ret.SprintID, ps.Logger)
		}

		relReportsByStandupIDPrms := ps.Params.Get("report", nil, ps.Logger).Sanitize("report")
		relReportsByStandupID, err := as.Services.Report.GetByStandupID(ps.Context, nil, ret.ID, relReportsByStandupIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child reports")
		}
		relStandupHistoriesByStandupIDPrms := ps.Params.Get("uhistory", nil, ps.Logger).Sanitize("uhistory")
		relStandupHistoriesByStandupID, err := as.Services.StandupHistory.GetByStandupID(ps.Context, nil, ret.ID, relStandupHistoriesByStandupIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		relStandupMembersByStandupIDPrms := ps.Params.Get("umember", nil, ps.Logger).Sanitize("umember")
		relStandupMembersByStandupID, err := as.Services.StandupMember.GetByStandupID(ps.Context, nil, ret.ID, relStandupMembersByStandupIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relStandupPermissionsByStandupIDPrms := ps.Params.Get("upermission", nil, ps.Logger).Sanitize("upermission")
		relStandupPermissionsByStandupID, err := as.Services.StandupPermission.GetByStandupID(ps.Context, nil, ret.ID, relStandupPermissionsByStandupIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		return Render(rc, as, &vstandup.Detail{
			Model:            ret,
			TeamByTeamID:     teamByTeamID,
			SprintBySprintID: sprintBySprintID,
			Params:           ps.Params,

			RelReportsByStandupID:            relReportsByStandupID,
			RelStandupHistoriesByStandupID:   relStandupHistoriesByStandupID,
			RelStandupMembersByStandupID:     relStandupMembersByStandupID,
			RelStandupPermissionsByStandupID: relStandupPermissionsByStandupID,
		}, ps, "standup", ret.String())
	})
}

func StandupCreateForm(rc *fasthttp.RequestCtx) {
	Act("standup.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &standup.Standup{}
		ps.Title = "Create [Standup]"
		ps.Data = ret
		return Render(rc, as, &vstandup.Edit{Model: ret, IsNew: true}, ps, "standup", "Create")
	})
}

func StandupCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("standup.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := standup.Random()
		ps.Title = "Create Random Standup"
		ps.Data = ret
		return Render(rc, as, &vstandup.Edit{Model: ret, IsNew: true}, ps, "standup", "Create")
	})
}

func StandupCreate(rc *fasthttp.RequestCtx) {
	Act("standup.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := standupFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Standup from form")
		}
		err = as.Services.Standup.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Standup")
		}
		msg := fmt.Sprintf("Standup [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func StandupEditForm(rc *fasthttp.RequestCtx) {
	Act("standup.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := standupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vstandup.Edit{Model: ret}, ps, "standup", ret.String())
	})
}

func StandupEdit(rc *fasthttp.RequestCtx) {
	Act("standup.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := standupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := standupFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Standup from form")
		}
		frm.ID = ret.ID
		err = as.Services.Standup.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Standup [%s]", frm.String())
		}
		msg := fmt.Sprintf("Standup [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func StandupDelete(rc *fasthttp.RequestCtx) {
	Act("standup.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := standupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Standup.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete standup [%s]", ret.String())
		}
		msg := fmt.Sprintf("Standup [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/admin/db/standup", rc, ps)
	})
}

func standupFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*standup.Standup, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
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

func standupFromForm(rc *fasthttp.RequestCtx, setPK bool) (*standup.Standup, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return standup.FromMap(frm, setPK)
}
