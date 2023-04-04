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
	"github.com/kyleu/rituals/app/retro"
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
		ps.Title = "Retros"
		ps.Data = ret
		teamIDsByTeamID := make([]*uuid.UUID, 0, len(ret))
		for _, x := range ret {
			teamIDsByTeamID = append(teamIDsByTeamID, x.TeamID)
		}
		teamsByTeamID, err := as.Services.Team.GetMultiple(ps.Context, nil, ps.Logger, util.ArrayDefererence(teamIDsByTeamID)...)
		if err != nil {
			return "", err
		}
		sprintIDsBySprintID := make([]*uuid.UUID, 0, len(ret))
		for _, x := range ret {
			sprintIDsBySprintID = append(sprintIDsBySprintID, x.SprintID)
		}
		sprintsBySprintID, err := as.Services.Sprint.GetMultiple(ps.Context, nil, ps.Logger, util.ArrayDefererence(sprintIDsBySprintID)...)
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
		ps.Title = ret.TitleString() + " (Retro)"
		ps.Data = ret
		feedbacksByRetroIDPrms := ps.Params.Get("feedback", nil, ps.Logger).Sanitize("feedback")
		feedbacksByRetroID, err := as.Services.Feedback.GetByRetroID(ps.Context, nil, ret.ID, feedbacksByRetroIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child feedbacks")
		}
		retroHistoriesByRetroIDPrms := ps.Params.Get("rhistory", nil, ps.Logger).Sanitize("rhistory")
		retroHistoriesByRetroID, err := as.Services.RetroHistory.GetByRetroID(ps.Context, nil, ret.ID, retroHistoriesByRetroIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		retroMembersByRetroIDPrms := ps.Params.Get("rmember", nil, ps.Logger).Sanitize("rmember")
		retroMembersByRetroID, err := as.Services.RetroMember.GetByRetroID(ps.Context, nil, ret.ID, retroMembersByRetroIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		retroPermissionsByRetroIDPrms := ps.Params.Get("rpermission", nil, ps.Logger).Sanitize("rpermission")
		retroPermissionsByRetroID, err := as.Services.RetroPermission.GetByRetroID(ps.Context, nil, ret.ID, retroPermissionsByRetroIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		return Render(rc, as, &vretro.Detail{
			Model:                     ret,
			Params:                    ps.Params,
			FeedbacksByRetroID:        feedbacksByRetroID,
			RetroHistoriesByRetroID:   retroHistoriesByRetroID,
			RetroMembersByRetroID:     retroMembersByRetroID,
			RetroPermissionsByRetroID: retroPermissionsByRetroID,
		}, ps, "retro", ret.String())
	})
}

func RetroCreateForm(rc *fasthttp.RequestCtx) {
	Act("retro.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &retro.Retro{}
		ps.Title = "Create [Retro]"
		ps.Data = ret
		return Render(rc, as, &vretro.Edit{Model: ret, IsNew: true}, ps, "retro", "Create")
	})
}

func RetroCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("retro.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := retro.Random()
		ps.Title = "Create Random Retro"
		ps.Data = ret
		return Render(rc, as, &vretro.Edit{Model: ret, IsNew: true}, ps, "retro", "Create")
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
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
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
		return FlashAndRedir(true, msg, "/retro", rc, ps)
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
