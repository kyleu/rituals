// Content managed by Project Forge, see [projectforge.md] for details.
package cstandup

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vstandup/vupermission"
)

func StandupPermissionList(rc *fasthttp.RequestCtx) {
	controller.Act("upermission.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("upermission", nil, ps.Logger).Sanitize("upermission")
		ret, err := as.Services.StandupPermission.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Permissions"
		ps.Data = ret
		standupIDs := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			standupIDs = append(standupIDs, x.StandupID)
		}
		standupsByStandupID, err := as.Services.Standup.GetMultiple(ps.Context, nil, ps.Logger, standupIDs...)
		if err != nil {
			return "", err
		}
		page := &vupermission.List{Models: ret, StandupsByStandupID: standupsByStandupID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "standup", "upermission")
	})
}

func StandupPermissionDetail(rc *fasthttp.RequestCtx) {
	controller.Act("upermission.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := upermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Permission)"
		ps.Data = ret
		return controller.Render(rc, as, &vupermission.Detail{Model: ret}, ps, "standup", "upermission", ret.String())
	})
}

func StandupPermissionCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("upermission.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &upermission.StandupPermission{}
		ps.Title = "Create [StandupPermission]"
		ps.Data = ret
		return controller.Render(rc, as, &vupermission.Edit{Model: ret, IsNew: true}, ps, "standup", "upermission", "Create")
	})
}

func StandupPermissionCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("upermission.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := upermission.Random()
		ps.Title = "Create Random StandupPermission"
		ps.Data = ret
		return controller.Render(rc, as, &vupermission.Edit{Model: ret, IsNew: true}, ps, "standup", "upermission", "Create")
	})
}

func StandupPermissionCreate(rc *fasthttp.RequestCtx) {
	controller.Act("upermission.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := upermissionFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse StandupPermission from form")
		}
		err = as.Services.StandupPermission.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created StandupPermission")
		}
		msg := fmt.Sprintf("StandupPermission [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func StandupPermissionEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("upermission.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := upermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vupermission.Edit{Model: ret}, ps, "standup", "upermission", ret.String())
	})
}

func StandupPermissionEdit(rc *fasthttp.RequestCtx) {
	controller.Act("upermission.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := upermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := upermissionFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse StandupPermission from form")
		}
		frm.StandupID = ret.StandupID
		frm.Key = ret.Key
		frm.Value = ret.Value
		err = as.Services.StandupPermission.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update StandupPermission [%s]", frm.String())
		}
		msg := fmt.Sprintf("StandupPermission [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func StandupPermissionDelete(rc *fasthttp.RequestCtx) {
	controller.Act("upermission.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := upermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.StandupPermission.Delete(ps.Context, nil, ret.StandupID, ret.Key, ret.Value, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete permission [%s]", ret.String())
		}
		msg := fmt.Sprintf("StandupPermission [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/standupPermission", rc, ps)
	})
}

func upermissionFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*upermission.StandupPermission, error) {
	standupIDArgStr, err := cutil.RCRequiredString(rc, "standupID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [standupID] as an argument")
	}
	standupIDArgP := util.UUIDFromString(standupIDArgStr)
	if standupIDArgP == nil {
		return nil, errors.Errorf("argument [standupID] (%s) is not a valid UUID", standupIDArgStr)
	}
	standupIDArg := *standupIDArgP
	keyArg, err := cutil.RCRequiredString(rc, "key", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [key] as an argument")
	}
	valueArg, err := cutil.RCRequiredString(rc, "value", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [value] as an argument")
	}
	return as.Services.StandupPermission.Get(ps.Context, nil, standupIDArg, keyArg, valueArg, ps.Logger)
}

func upermissionFromForm(rc *fasthttp.RequestCtx, setPK bool) (*upermission.StandupPermission, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return upermission.FromMap(frm, setPK)
}
