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
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vestimate"
)

func EstimateList(w http.ResponseWriter, r *http.Request) {
	Act("estimate.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(cutil.QueryStringString(ps.URI, "q"))
		prms := ps.Params.Sanitized("estimate", ps.Logger)
		var ret estimate.Estimates
		var err error
		if q == "" {
			ret, err = as.Services.Estimate.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Estimate.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 && !cutil.IsContentTypeJSON(cutil.GetContentType(r)) {
				return FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
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
		return Render(r, as, page, ps, "estimate")
	})
}

//nolint:lll
func EstimateDetail(w http.ResponseWriter, r *http.Request) {
	Act("estimate.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromPath(r, as, ps)
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

		relEstimateHistoriesByEstimateIDPrms := ps.Params.Sanitized("ehistory", ps.Logger)
		relEstimateHistoriesByEstimateID, err := as.Services.EstimateHistory.GetByEstimateID(ps.Context, nil, ret.ID, relEstimateHistoriesByEstimateIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		relEstimateMembersByEstimateIDPrms := ps.Params.Sanitized("emember", ps.Logger)
		relEstimateMembersByEstimateID, err := as.Services.EstimateMember.GetByEstimateID(ps.Context, nil, ret.ID, relEstimateMembersByEstimateIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relEstimatePermissionsByEstimateIDPrms := ps.Params.Sanitized("epermission", ps.Logger)
		relEstimatePermissionsByEstimateID, err := as.Services.EstimatePermission.GetByEstimateID(ps.Context, nil, ret.ID, relEstimatePermissionsByEstimateIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		relStoriesByEstimateIDPrms := ps.Params.Sanitized("story", ps.Logger)
		relStoriesByEstimateID, err := as.Services.Story.GetByEstimateID(ps.Context, nil, ret.ID, relStoriesByEstimateIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child stories")
		}
		return Render(r, as, &vestimate.Detail{
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

func EstimateCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("estimate.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &estimate.Estimate{}
		if cutil.QueryStringString(ps.URI, "prototype") == util.KeyRandom {
			ret = estimate.RandomEstimate()
			randomTeam, err := as.Services.Team.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomTeam != nil {
				ret.TeamID = &randomTeam.ID
			}
			randomSprint, err := as.Services.Sprint.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomSprint != nil {
				ret.SprintID = &randomSprint.ID
			}
		}
		ps.SetTitleAndData("Create [Estimate]", ret)
		return Render(r, as, &vestimate.Edit{Model: ret, IsNew: true}, ps, "estimate", "Create")
	})
}

func EstimateRandom(w http.ResponseWriter, r *http.Request) {
	Act("estimate.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Estimate.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Estimate")
		}
		return ret.WebPath(), nil
	})
}

func EstimateCreate(w http.ResponseWriter, r *http.Request) {
	Act("estimate.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Estimate from form")
		}
		err = as.Services.Estimate.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Estimate")
		}
		msg := fmt.Sprintf("Estimate [%s] created", ret.TitleString())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func EstimateEditForm(w http.ResponseWriter, r *http.Request) {
	Act("estimate.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vestimate.Edit{Model: ret}, ps, "estimate", ret.String())
	})
}

func EstimateEdit(w http.ResponseWriter, r *http.Request) {
	Act("estimate.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := estimateFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Estimate from form")
		}
		frm.ID = ret.ID
		err = as.Services.Estimate.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Estimate [%s]", frm.String())
		}
		msg := fmt.Sprintf("Estimate [%s] updated", frm.TitleString())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func EstimateDelete(w http.ResponseWriter, r *http.Request) {
	Act("estimate.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Estimate.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete estimate [%s]", ret.String())
		}
		msg := fmt.Sprintf("Estimate [%s] deleted", ret.TitleString())
		return FlashAndRedir(true, msg, "/admin/db/estimate", ps)
	})
}

func estimateFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*estimate.Estimate, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
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

func estimateFromForm(r *http.Request, b []byte, setPK bool) (*estimate.Estimate, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := estimate.EstimateFromMap(frm, setPK)
	return ret, err
}
