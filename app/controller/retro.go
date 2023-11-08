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
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vretro"
)

func RetroList(rc *fasthttp.RequestCtx) {
	Act("retro.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("retro", nil, ps.Logger).Sanitize("retro")
		var ret retro.Retros
		var err error
		if q == "" {
			ret, err = as.Services.Retro.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Retro.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Retros", ret)
		teamIDsByTeamID := lo.Map(ret, func(x *retro.Retro, _ int) *uuid.UUID {
			return x.TeamID
		})
		teamsByTeamID, err := as.Services.Team.GetMultiple(ps.Context, nil, nil, ps.Logger, util.ArrayDereference(teamIDsByTeamID)...)
		if err != nil {
			return "", err
		}
		sprintIDsBySprintID := lo.Map(ret, func(x *retro.Retro, _ int) *uuid.UUID {
			return x.SprintID
		})
		sprintsBySprintID, err := as.Services.Sprint.GetMultiple(ps.Context, nil, nil, ps.Logger, util.ArrayDereference(sprintIDsBySprintID)...)
		if err != nil {
			return "", err
		}
		page := &vretro.List{Models: ret, TeamsByTeamID: teamsByTeamID, SprintsBySprintID: sprintsBySprintID, Params: ps.Params, SearchQuery: q}
		return Render(rc, as, page, ps, "retro")
	})
}

func RetroDetail(rc *fasthttp.RequestCtx) {
	Act("retro.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := retroFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Retro)", ret)

		var teamByTeamID *team.Team
		if ret.TeamID != nil {
			teamByTeamID, _ = as.Services.Team.Get(ps.Context, nil, *ret.TeamID, ps.Logger)
		}
		var sprintBySprintID *sprint.Sprint
		if ret.SprintID != nil {
			sprintBySprintID, _ = as.Services.Sprint.Get(ps.Context, nil, *ret.SprintID, ps.Logger)
		}

		relFeedbacksByRetroIDPrms := ps.Params.Get("feedback", nil, ps.Logger).Sanitize("feedback")
		relFeedbacksByRetroID, err := as.Services.Feedback.GetByRetroID(ps.Context, nil, ret.ID, relFeedbacksByRetroIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child feedbacks")
		}
		relRetroHistoriesByRetroIDPrms := ps.Params.Get("rhistory", nil, ps.Logger).Sanitize("rhistory")
		relRetroHistoriesByRetroID, err := as.Services.RetroHistory.GetByRetroID(ps.Context, nil, ret.ID, relRetroHistoriesByRetroIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		relRetroMembersByRetroIDPrms := ps.Params.Get("rmember", nil, ps.Logger).Sanitize("rmember")
		relRetroMembersByRetroID, err := as.Services.RetroMember.GetByRetroID(ps.Context, nil, ret.ID, relRetroMembersByRetroIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relRetroPermissionsByRetroIDPrms := ps.Params.Get("rpermission", nil, ps.Logger).Sanitize("rpermission")
		relRetroPermissionsByRetroID, err := as.Services.RetroPermission.GetByRetroID(ps.Context, nil, ret.ID, relRetroPermissionsByRetroIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		return Render(rc, as, &vretro.Detail{
			Model:            ret,
			TeamByTeamID:     teamByTeamID,
			SprintBySprintID: sprintBySprintID,
			Params:           ps.Params,

			RelFeedbacksByRetroID:        relFeedbacksByRetroID,
			RelRetroHistoriesByRetroID:   relRetroHistoriesByRetroID,
			RelRetroMembersByRetroID:     relRetroMembersByRetroID,
			RelRetroPermissionsByRetroID: relRetroPermissionsByRetroID,
		}, ps, "retro", ret.TitleString()+"**retro")
	})
}

func RetroCreateForm(rc *fasthttp.RequestCtx) {
	Act("retro.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &retro.Retro{}
		if string(rc.QueryArgs().Peek("prototype")) == util.KeyRandom {
			ret = retro.Random()
			randomTeam, err := as.Services.Team.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomTeam != nil {
				ret.TeamID = &randomTeam.ID
			}
			randomSprint, err := as.Services.Sprint.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomSprint != nil {
				ret.SprintID = &randomSprint.ID
			}
		}
		ps.SetTitleAndData("Create [Retro]", ret)
		ps.Data = ret
		return Render(rc, as, &vretro.Edit{Model: ret, IsNew: true}, ps, "retro", "Create")
	})
}

func RetroRandom(rc *fasthttp.RequestCtx) {
	Act("retro.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Retro.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Retro")
		}
		return ret.WebPath(), nil
	})
}

func RetroCreate(rc *fasthttp.RequestCtx) {
	Act("retro.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := retroFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Retro from form")
		}
		err = as.Services.Retro.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Retro")
		}
		msg := fmt.Sprintf("Retro [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func RetroEditForm(rc *fasthttp.RequestCtx) {
	Act("retro.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := retroFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(rc, as, &vretro.Edit{Model: ret}, ps, "retro", ret.String())
	})
}

func RetroEdit(rc *fasthttp.RequestCtx) {
	Act("retro.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := retroFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := retroFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Retro from form")
		}
		frm.ID = ret.ID
		err = as.Services.Retro.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Retro [%s]", frm.String())
		}
		msg := fmt.Sprintf("Retro [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func RetroDelete(rc *fasthttp.RequestCtx) {
	Act("retro.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := retroFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Retro.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete retro [%s]", ret.String())
		}
		msg := fmt.Sprintf("Retro [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/admin/db/retro", rc, ps)
	})
}

func retroFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*retro.Retro, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Retro.Get(ps.Context, nil, idArg, ps.Logger)
}

func retroFromForm(rc *fasthttp.RequestCtx, setPK bool) (*retro.Retro, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return retro.FromMap(frm, setPK)
}
