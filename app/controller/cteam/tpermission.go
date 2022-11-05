// Content managed by Project Forge, see [projectforge.md] for details.
package cteam

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vteam/vtpermission"
)

func TeamPermissionList(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("tpermission", nil, ps.Logger).Sanitize("tpermission")
		ret, err := as.Services.TeamPermission.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Permissions"
		ps.Data = ret
		teamIDs := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			teamIDs = append(teamIDs, x.TeamID)
		}
		teams, err := as.Services.Team.GetMultiple(ps.Context, nil, ps.Logger, teamIDs...)
		if err != nil {
			return "", err
		}
		return controller.Render(rc, as, &vtpermission.List{Models: ret, Teams: teams, Params: ps.Params}, ps, "team", "tpermission")
	})
}

func TeamPermissionDetail(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Permission)"
		ps.Data = ret
		return controller.Render(rc, as, &vtpermission.Detail{Model: ret}, ps, "team", "tpermission", ret.String())
	})
}

func TeamPermissionCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &tpermission.TeamPermission{}
		ps.Title = "Create [TeamPermission]"
		ps.Data = ret
		return controller.Render(rc, as, &vtpermission.Edit{Model: ret, IsNew: true}, ps, "team", "tpermission", "Create")
	})
}

func TeamPermissionCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := tpermission.Random()
		ps.Title = "Create Random TeamPermission"
		ps.Data = ret
		return controller.Render(rc, as, &vtpermission.Edit{Model: ret, IsNew: true}, ps, "team", "tpermission", "Create")
	})
}

func TeamPermissionCreate(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamPermission from form")
		}
		err = as.Services.TeamPermission.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created TeamPermission")
		}
		msg := fmt.Sprintf("TeamPermission [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func TeamPermissionEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vtpermission.Edit{Model: ret}, ps, "team", "tpermission", ret.String())
	})
}

func TeamPermissionEdit(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := tpermissionFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamPermission from form")
		}
		frm.TeamID = ret.TeamID
		frm.K = ret.K
		frm.V = ret.V
		err = as.Services.TeamPermission.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update TeamPermission [%s]", frm.String())
		}
		msg := fmt.Sprintf("TeamPermission [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func TeamPermissionDelete(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.TeamPermission.Delete(ps.Context, nil, ret.TeamID, ret.K, ret.V, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete permission [%s]", ret.String())
		}
		msg := fmt.Sprintf("TeamPermission [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/teamPermission", rc, ps)
	})
}

func tpermissionFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*tpermission.TeamPermission, error) {
	teamIDArgStr, err := cutil.RCRequiredString(rc, "teamID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [teamID] as an argument")
	}
	teamIDArgP := util.UUIDFromString(teamIDArgStr)
	if teamIDArgP == nil {
		return nil, errors.Errorf("argument [teamID] (%s) is not a valid UUID", teamIDArgStr)
	}
	teamIDArg := *teamIDArgP
	kArg, err := cutil.RCRequiredString(rc, "k", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [k] as an argument")
	}
	vArg, err := cutil.RCRequiredString(rc, "v", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [v] as an argument")
	}
	return as.Services.TeamPermission.Get(ps.Context, nil, teamIDArg, kArg, vArg, ps.Logger)
}

func tpermissionFromForm(rc *fasthttp.RequestCtx, setPK bool) (*tpermission.TeamPermission, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return tpermission.FromMap(frm, setPK)
}
