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
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vestimate"
)

func EstimateList(rc *fasthttp.RequestCtx) {
	Act("estimate.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("estimate", nil, ps.Logger).Sanitize("estimate")
		var ret estimate.Estimates
		var err error
		if q == "" {
			ret, err = as.Services.Estimate.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Estimate.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Estimates", ret)
		teamIDsByTeamID := lo.Map(ret, func(x *estimate.Estimate, _ int) *uuid.UUID {
			return x.TeamID
		})
		teamsByTeamID, err := as.Services.Team.GetMultiple(ps.Context, nil, nil, ps.Logger, util.ArrayDereference(teamIDsByTeamID)...)
		if err != nil {
			return "", err
		}
		sprintIDsBySprintID := lo.Map(ret, func(x *estimate.Estimate, _ int) *uuid.UUID {
			return x.SprintID
		})
		sprintsBySprintID, err := as.Services.Sprint.GetMultiple(ps.Context, nil, nil, ps.Logger, util.ArrayDereference(sprintIDsBySprintID)...)
		if err != nil {
			return "", err
		}
		page := &vestimate.List{Models: ret, TeamsByTeamID: teamsByTeamID, SprintsBySprintID: sprintsBySprintID, Params: ps.Params, SearchQuery: q}
		return Render(rc, as, page, ps, "estimate")
	})
}

//nolint:lll
func EstimateDetail(rc *fasthttp.RequestCtx) {
	Act("estimate.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Estimate)", ret)

		var teamByTeamID *team.Team
		if ret.TeamID != nil {
			teamByTeamID, _ = as.Services.Team.Get(ps.Context, nil, *ret.TeamID, ps.Logger)
		}
		var sprintBySprintID *sprint.Sprint
		if ret.SprintID != nil {
			sprintBySprintID, _ = as.Services.Sprint.Get(ps.Context, nil, *ret.SprintID, ps.Logger)
		}

		relEstimateHistoriesByEstimateIDPrms := ps.Params.Get("ehistory", nil, ps.Logger).Sanitize("ehistory")
		relEstimateHistoriesByEstimateID, err := as.Services.EstimateHistory.GetByEstimateID(ps.Context, nil, ret.ID, relEstimateHistoriesByEstimateIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		relEstimateMembersByEstimateIDPrms := ps.Params.Get("emember", nil, ps.Logger).Sanitize("emember")
		relEstimateMembersByEstimateID, err := as.Services.EstimateMember.GetByEstimateID(ps.Context, nil, ret.ID, relEstimateMembersByEstimateIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relEstimatePermissionsByEstimateIDPrms := ps.Params.Get("epermission", nil, ps.Logger).Sanitize("epermission")
		relEstimatePermissionsByEstimateID, err := as.Services.EstimatePermission.GetByEstimateID(ps.Context, nil, ret.ID, relEstimatePermissionsByEstimateIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		relStoriesByEstimateIDPrms := ps.Params.Get("story", nil, ps.Logger).Sanitize("story")
		relStoriesByEstimateID, err := as.Services.Story.GetByEstimateID(ps.Context, nil, ret.ID, relStoriesByEstimateIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child stories")
		}
		return Render(rc, as, &vestimate.Detail{
			Model:            ret,
			TeamByTeamID:     teamByTeamID,
			SprintBySprintID: sprintBySprintID,
			Params:           ps.Params,

			RelEstimateHistoriesByEstimateID:   relEstimateHistoriesByEstimateID,
			RelEstimateMembersByEstimateID:     relEstimateMembersByEstimateID,
			RelEstimatePermissionsByEstimateID: relEstimatePermissionsByEstimateID,
			RelStoriesByEstimateID:             relStoriesByEstimateID,
		}, ps, "estimate", ret.TitleString()+"**estimate")
	})
}

func EstimateCreateForm(rc *fasthttp.RequestCtx) {
	Act("estimate.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &estimate.Estimate{}
		if string(rc.QueryArgs().Peek("prototype")) == util.KeyRandom {
			ret = estimate.Random()
		}
		ps.SetTitleAndData("Create [Estimate]", ret)
		ps.Data = ret
		return Render(rc, as, &vestimate.Edit{Model: ret, IsNew: true}, ps, "estimate", "Create")
	})
}

func EstimateRandom(rc *fasthttp.RequestCtx) {
	Act("estimate.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Estimate.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Estimate")
		}
		return ret.WebPath(), nil
	})
}

func EstimateCreate(rc *fasthttp.RequestCtx) {
	Act("estimate.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Estimate from form")
		}
		err = as.Services.Estimate.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Estimate")
		}
		msg := fmt.Sprintf("Estimate [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func EstimateEditForm(rc *fasthttp.RequestCtx) {
	Act("estimate.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(rc, as, &vestimate.Edit{Model: ret}, ps, "estimate", ret.String())
	})
}

func EstimateEdit(rc *fasthttp.RequestCtx) {
	Act("estimate.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := estimateFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Estimate from form")
		}
		frm.ID = ret.ID
		err = as.Services.Estimate.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Estimate [%s]", frm.String())
		}
		msg := fmt.Sprintf("Estimate [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func EstimateDelete(rc *fasthttp.RequestCtx) {
	Act("estimate.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Estimate.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete estimate [%s]", ret.String())
		}
		msg := fmt.Sprintf("Estimate [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/admin/db/estimate", rc, ps)
	})
}

func estimateFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*estimate.Estimate, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Estimate.Get(ps.Context, nil, idArg, ps.Logger)
}

func estimateFromForm(rc *fasthttp.RequestCtx, setPK bool) (*estimate.Estimate, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return estimate.FromMap(frm, setPK)
}
