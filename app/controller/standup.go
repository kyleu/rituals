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
	"github.com/kyleu/rituals/app/standup"
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
		teamIDs := make([]*uuid.UUID, 0, len(ret))
		for _, x := range ret {
			teamIDs = append(teamIDs, x.TeamID)
		}
		teams, err := as.Services.Team.GetMultiple(ps.Context, nil, ps.Logger, util.ArrayDefererence(teamIDs)...)
		if err != nil {
			return "", err
		}
		sprintIDs := make([]*uuid.UUID, 0, len(ret))
		for _, x := range ret {
			sprintIDs = append(sprintIDs, x.SprintID)
		}
		sprints, err := as.Services.Sprint.GetMultiple(ps.Context, nil, ps.Logger, util.ArrayDefererence(sprintIDs)...)
		if err != nil {
			return "", err
		}
		page := &vstandup.List{Models: ret, Teams: teams, Sprints: sprints, Params: ps.Params, SearchQuery: q}
		return Render(rc, as, page, ps, "standup")
	})
}

func StandupDetail(rc *fasthttp.RequestCtx) {
	Act("standup.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := standupFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Standup)"
		ps.Data = ret
		reportPrms := ps.Params.Get("report", nil, ps.Logger).Sanitize("report")
		reportsByStandupID, err := as.Services.Report.GetByStandupID(ps.Context, nil, ret.ID, reportPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child reports")
		}
		standupHistoryPrms := ps.Params.Get("uhistory", nil, ps.Logger).Sanitize("uhistory")
		standupHistoriesByStandupID, err := as.Services.StandupHistory.GetByStandupID(ps.Context, nil, ret.ID, standupHistoryPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		standupMemberPrms := ps.Params.Get("umember", nil, ps.Logger).Sanitize("umember")
		standupMembersByStandupID, err := as.Services.StandupMember.GetByStandupID(ps.Context, nil, ret.ID, standupMemberPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		standupPermissionPrms := ps.Params.Get("upermission", nil, ps.Logger).Sanitize("upermission")
		standupPermissionsByStandupID, err := as.Services.StandupPermission.GetByStandupID(ps.Context, nil, ret.ID, standupPermissionPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		return Render(rc, as, &vstandup.Detail{
			Model:                         ret,
			Params:                        ps.Params,
			ReportsByStandupID:            reportsByStandupID,
			StandupHistoriesByStandupID:   standupHistoriesByStandupID,
			StandupMembersByStandupID:     standupMembersByStandupID,
			StandupPermissionsByStandupID: standupPermissionsByStandupID,
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
		return FlashAndRedir(true, msg, "/standup", rc, ps)
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
