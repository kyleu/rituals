// Content managed by Project Forge, see [projectforge.md] for details.
package cestimate

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vestimate/vepermission"
)

func EstimatePermissionList(rc *fasthttp.RequestCtx) {
	controller.Act("epermission.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("epermission", nil, ps.Logger).Sanitize("epermission")
		ret, err := as.Services.EstimatePermission.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Permissions"
		ps.Data = ret
		estimateIDs := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			estimateIDs = append(estimateIDs, x.EstimateID)
		}
		estimates, err := as.Services.Estimate.GetMultiple(ps.Context, nil, ps.Logger, estimateIDs...)
		if err != nil {
			return "", err
		}
		return controller.Render(rc, as, &vepermission.List{Models: ret, Estimates: estimates, Params: ps.Params}, ps, "estimate", "epermission")
	})
}

func EstimatePermissionDetail(rc *fasthttp.RequestCtx) {
	controller.Act("epermission.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := epermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Permission)"
		ps.Data = ret
		return controller.Render(rc, as, &vepermission.Detail{Model: ret}, ps, "estimate", "epermission", ret.String())
	})
}

func EstimatePermissionCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("epermission.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &epermission.EstimatePermission{}
		ps.Title = "Create [EstimatePermission]"
		ps.Data = ret
		return controller.Render(rc, as, &vepermission.Edit{Model: ret, IsNew: true}, ps, "estimate", "epermission", "Create")
	})
}

func EstimatePermissionCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("epermission.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := epermission.Random()
		ps.Title = "Create Random EstimatePermission"
		ps.Data = ret
		return controller.Render(rc, as, &vepermission.Edit{Model: ret, IsNew: true}, ps, "estimate", "epermission", "Create")
	})
}

func EstimatePermissionCreate(rc *fasthttp.RequestCtx) {
	controller.Act("epermission.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := epermissionFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse EstimatePermission from form")
		}
		err = as.Services.EstimatePermission.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created EstimatePermission")
		}
		msg := fmt.Sprintf("EstimatePermission [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func EstimatePermissionEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("epermission.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := epermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vepermission.Edit{Model: ret}, ps, "estimate", "epermission", ret.String())
	})
}

func EstimatePermissionEdit(rc *fasthttp.RequestCtx) {
	controller.Act("epermission.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := epermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := epermissionFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse EstimatePermission from form")
		}
		frm.EstimateID = ret.EstimateID
		frm.Key = ret.Key
		frm.Value = ret.Value
		err = as.Services.EstimatePermission.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update EstimatePermission [%s]", frm.String())
		}
		msg := fmt.Sprintf("EstimatePermission [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func EstimatePermissionDelete(rc *fasthttp.RequestCtx) {
	controller.Act("epermission.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := epermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.EstimatePermission.Delete(ps.Context, nil, ret.EstimateID, ret.Key, ret.Value, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete permission [%s]", ret.String())
		}
		msg := fmt.Sprintf("EstimatePermission [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/estimatePermission", rc, ps)
	})
}

func epermissionFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*epermission.EstimatePermission, error) {
	estimateIDArgStr, err := cutil.RCRequiredString(rc, "estimateID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [estimateID] as an argument")
	}
	estimateIDArgP := util.UUIDFromString(estimateIDArgStr)
	if estimateIDArgP == nil {
		return nil, errors.Errorf("argument [estimateID] (%s) is not a valid UUID", estimateIDArgStr)
	}
	estimateIDArg := *estimateIDArgP
	keyArg, err := cutil.RCRequiredString(rc, "key", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [key] as an argument")
	}
	valueArg, err := cutil.RCRequiredString(rc, "value", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [value] as an argument")
	}
	return as.Services.EstimatePermission.Get(ps.Context, nil, estimateIDArg, keyArg, valueArg, ps.Logger)
}

func epermissionFromForm(rc *fasthttp.RequestCtx, setPK bool) (*epermission.EstimatePermission, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return epermission.FromMap(frm, setPK)
}
