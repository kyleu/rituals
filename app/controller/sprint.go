// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vsprint"
)

func SprintList(rc *fasthttp.RequestCtx) {
	Act("sprint.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("sprint", nil, ps.Logger).Sanitize("sprint")
		var ret sprint.Sprints
		var err error
		if q == "" {
			ret, err = as.Services.Sprint.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Sprint.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.Title = "Sprints"
		ps.Data = ret
		teamIDsByTeamID := make([]*uuid.UUID, 0, len(ret))
		for _, x := range ret {
			teamIDsByTeamID = append(teamIDsByTeamID, x.TeamID)
		}
		teamsByTeamID, err := as.Services.Team.GetMultiple(ps.Context, nil, ps.Logger, util.ArrayDereference(teamIDsByTeamID)...)
		if err != nil {
			return "", err
		}
		page := &vsprint.List{Models: ret, TeamsByTeamID: teamsByTeamID, Params: ps.Params, SearchQuery: q}
		return Render(rc, as, page, ps, "sprint")
	})
}

func SprintDetail(rc *fasthttp.RequestCtx) {
	Act("sprint.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := sprintFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Sprint)"
		ps.Data = ret

		var teamByTeamID *team.Team
		if ret.TeamID != nil {
			teamByTeamID, _ = as.Services.Team.Get(ps.Context, nil, *ret.TeamID, ps.Logger)
		}

		relEstimatesBySprintIDPrms := ps.Params.Get("estimate", nil, ps.Logger).Sanitize("estimate")
		relEstimatesBySprintID, err := as.Services.Estimate.GetBySprintID(ps.Context, nil, &ret.ID, relEstimatesBySprintIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child estimates")
		}
		relRetrosBySprintIDPrms := ps.Params.Get("retro", nil, ps.Logger).Sanitize("retro")
		relRetrosBySprintID, err := as.Services.Retro.GetBySprintID(ps.Context, nil, &ret.ID, relRetrosBySprintIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child retros")
		}
		relSprintHistoriesBySprintIDPrms := ps.Params.Get("shistory", nil, ps.Logger).Sanitize("shistory")
		relSprintHistoriesBySprintID, err := as.Services.SprintHistory.GetBySprintID(ps.Context, nil, ret.ID, relSprintHistoriesBySprintIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		relSprintMembersBySprintIDPrms := ps.Params.Get("smember", nil, ps.Logger).Sanitize("smember")
		relSprintMembersBySprintID, err := as.Services.SprintMember.GetBySprintID(ps.Context, nil, ret.ID, relSprintMembersBySprintIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relSprintPermissionsBySprintIDPrms := ps.Params.Get("spermission", nil, ps.Logger).Sanitize("spermission")
		relSprintPermissionsBySprintID, err := as.Services.SprintPermission.GetBySprintID(ps.Context, nil, ret.ID, relSprintPermissionsBySprintIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		relStandupsBySprintIDPrms := ps.Params.Get("standup", nil, ps.Logger).Sanitize("standup")
		relStandupsBySprintID, err := as.Services.Standup.GetBySprintID(ps.Context, nil, &ret.ID, relStandupsBySprintIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child standups")
		}
		return Render(rc, as, &vsprint.Detail{
			Model:        ret,
			TeamByTeamID: teamByTeamID,
			Params:       ps.Params,

			RelEstimatesBySprintID:         relEstimatesBySprintID,
			RelRetrosBySprintID:            relRetrosBySprintID,
			RelSprintHistoriesBySprintID:   relSprintHistoriesBySprintID,
			RelSprintMembersBySprintID:     relSprintMembersBySprintID,
			RelSprintPermissionsBySprintID: relSprintPermissionsBySprintID,
			RelStandupsBySprintID:          relStandupsBySprintID,
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
