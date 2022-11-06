// Content managed by Project Forge, see [projectforge.md] for details.
package cretro

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vretro/vrpermission"
)

func RetroPermissionList(rc *fasthttp.RequestCtx) {
	controller.Act("rpermission.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("rpermission", nil, ps.Logger).Sanitize("rpermission")
		ret, err := as.Services.RetroPermission.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Permissions"
		ps.Data = ret
		retroIDs := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			retroIDs = append(retroIDs, x.RetroID)
		}
		retros, err := as.Services.Retro.GetMultiple(ps.Context, nil, ps.Logger, retroIDs...)
		if err != nil {
			return "", err
		}
		return controller.Render(rc, as, &vrpermission.List{Models: ret, Retros: retros, Params: ps.Params}, ps, "retro", "rpermission")
	})
}

func RetroPermissionDetail(rc *fasthttp.RequestCtx) {
	controller.Act("rpermission.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rpermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Permission)"
		ps.Data = ret
		return controller.Render(rc, as, &vrpermission.Detail{Model: ret}, ps, "retro", "rpermission", ret.String())
	})
}

func RetroPermissionCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("rpermission.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &rpermission.RetroPermission{}
		ps.Title = "Create [RetroPermission]"
		ps.Data = ret
		return controller.Render(rc, as, &vrpermission.Edit{Model: ret, IsNew: true}, ps, "retro", "rpermission", "Create")
	})
}

func RetroPermissionCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("rpermission.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := rpermission.Random()
		ps.Title = "Create Random RetroPermission"
		ps.Data = ret
		return controller.Render(rc, as, &vrpermission.Edit{Model: ret, IsNew: true}, ps, "retro", "rpermission", "Create")
	})
}

func RetroPermissionCreate(rc *fasthttp.RequestCtx) {
	controller.Act("rpermission.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rpermissionFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse RetroPermission from form")
		}
		err = as.Services.RetroPermission.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created RetroPermission")
		}
		msg := fmt.Sprintf("RetroPermission [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func RetroPermissionEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("rpermission.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rpermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vrpermission.Edit{Model: ret}, ps, "retro", "rpermission", ret.String())
	})
}

func RetroPermissionEdit(rc *fasthttp.RequestCtx) {
	controller.Act("rpermission.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rpermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := rpermissionFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse RetroPermission from form")
		}
		frm.RetroID = ret.RetroID
		frm.K = ret.K
		frm.V = ret.V
		err = as.Services.RetroPermission.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update RetroPermission [%s]", frm.String())
		}
		msg := fmt.Sprintf("RetroPermission [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func RetroPermissionDelete(rc *fasthttp.RequestCtx) {
	controller.Act("rpermission.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rpermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.RetroPermission.Delete(ps.Context, nil, ret.RetroID, ret.K, ret.V, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete permission [%s]", ret.String())
		}
		msg := fmt.Sprintf("RetroPermission [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/retroPermission", rc, ps)
	})
}

func rpermissionFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*rpermission.RetroPermission, error) {
	retroIDArgStr, err := cutil.RCRequiredString(rc, "retroID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [retroID] as an argument")
	}
	retroIDArgP := util.UUIDFromString(retroIDArgStr)
	if retroIDArgP == nil {
		return nil, errors.Errorf("argument [retroID] (%s) is not a valid UUID", retroIDArgStr)
	}
	retroIDArg := *retroIDArgP
	kArg, err := cutil.RCRequiredString(rc, "k", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [k] as an argument")
	}
	vArg, err := cutil.RCRequiredString(rc, "v", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [v] as an argument")
	}
	return as.Services.RetroPermission.Get(ps.Context, nil, retroIDArg, kArg, vArg, ps.Logger)
}

func rpermissionFromForm(rc *fasthttp.RequestCtx, setPK bool) (*rpermission.RetroPermission, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return rpermission.FromMap(frm, setPK)
}