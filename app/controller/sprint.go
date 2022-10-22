// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vsprint"
)

func SprintList(rc *fasthttp.RequestCtx) {
	Act("sprint.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		params := cutil.ParamSetFromRequest(rc)
		prms := params.Get("sprint", nil, ps.Logger).Sanitize("sprint")
		ret, err := as.Services.Sprint.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Sprints"
		ps.Data = ret
		userIDs := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			userIDs = append(userIDs, x.Owner)
		}
		users, err := as.Services.User.GetMultiple(ps.Context, nil, ps.Logger, userIDs...)
		if err != nil {
			return "", err
		}
		teamIDs := make([]*uuid.UUID, 0, len(ret))
		for _, x := range ret {
			teamIDs = append(teamIDs, x.TeamID)
		}
		teams, err := as.Services.Team.GetMultiple(ps.Context, nil, ps.Logger, util.ArrayDefererence(teamIDs)...)
		if err != nil {
			return "", err
		}
		return Render(rc, as, &vsprint.List{Models: ret, Users: users, Teams: teams, Params: params}, ps, "sprint")
	})
}

func SprintDetail(rc *fasthttp.RequestCtx) {
	Act("sprint.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		params := cutil.ParamSetFromRequest(rc)
		ret, err := sprintFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Sprint)"
		ps.Data = ret
		estimatePrms := params.Get("estimate", nil, ps.Logger).Sanitize("estimate")
		estimatesBySprintID, err := as.Services.Estimate.GetBySprintID(ps.Context, nil, &ret.ID, estimatePrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child estimates")
		}
		retroPrms := params.Get("retro", nil, ps.Logger).Sanitize("retro")
		retrosBySprintID, err := as.Services.Retro.GetBySprintID(ps.Context, nil, &ret.ID, retroPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child retros")
		}
		sprintHistoryPrms := params.Get("sprintHistory", nil, ps.Logger).Sanitize("sprintHistory")
		sprintHistoriesBySprintID, err := as.Services.SprintHistory.GetBySprintID(ps.Context, nil, ret.ID, sprintHistoryPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		sprintMemberPrms := params.Get("sprintMember", nil, ps.Logger).Sanitize("sprintMember")
		sprintMembersBySprintID, err := as.Services.SprintMember.GetBySprintID(ps.Context, nil, ret.ID, sprintMemberPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		sprintPermissionPrms := params.Get("sprintPermission", nil, ps.Logger).Sanitize("sprintPermission")
		sprintPermissionsBySprintID, err := as.Services.SprintPermission.GetBySprintID(ps.Context, nil, ret.ID, sprintPermissionPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		standupPrms := params.Get("standup", nil, ps.Logger).Sanitize("standup")
		standupsBySprintID, err := as.Services.Standup.GetBySprintID(ps.Context, nil, &ret.ID, standupPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child standups")
		}
		return Render(rc, as, &vsprint.Detail{
			Model:                       ret,
			Params:                      params,
			EstimatesBySprintID:         estimatesBySprintID,
			RetrosBySprintID:            retrosBySprintID,
			SprintHistoriesBySprintID:   sprintHistoriesBySprintID,
			SprintMembersBySprintID:     sprintMembersBySprintID,
			SprintPermissionsBySprintID: sprintPermissionsBySprintID,
			StandupsBySprintID:          standupsBySprintID,
		}, ps, "sprint", ret.String())
	})
}

func SprintCreateForm(rc *fasthttp.RequestCtx) {
	Act("sprint.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &sprint.Sprint{}
		ps.Title = "Create [Sprint]"
		ps.Data = ret
		return Render(rc, as, &vsprint.Edit{Model: ret, IsNew: true}, ps, "sprint", "Create")
	})
}

func SprintCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("sprint.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := sprint.Random()
		ps.Title = "Create Random Sprint"
		ps.Data = ret
		return Render(rc, as, &vsprint.Edit{Model: ret, IsNew: true}, ps, "sprint", "Create")
	})
}

func SprintCreate(rc *fasthttp.RequestCtx) {
	Act("sprint.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := sprintFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Sprint from form")
		}
		err = as.Services.Sprint.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Sprint")
		}
		msg := fmt.Sprintf("Sprint [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func SprintEditForm(rc *fasthttp.RequestCtx) {
	Act("sprint.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := sprintFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vsprint.Edit{Model: ret}, ps, "sprint", ret.String())
	})
}

func SprintEdit(rc *fasthttp.RequestCtx) {
	Act("sprint.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := sprintFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := sprintFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Sprint from form")
		}
		frm.ID = ret.ID
		err = as.Services.Sprint.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Sprint [%s]", frm.String())
		}
		msg := fmt.Sprintf("Sprint [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func SprintDelete(rc *fasthttp.RequestCtx) {
	Act("sprint.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := sprintFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Sprint.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete sprint [%s]", ret.String())
		}
		msg := fmt.Sprintf("Sprint [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/sprint", rc, ps)
	})
}

func sprintFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*sprint.Sprint, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
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

func sprintFromForm(rc *fasthttp.RequestCtx, setPK bool) (*sprint.Sprint, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return sprint.FromMap(frm, setPK)
}
