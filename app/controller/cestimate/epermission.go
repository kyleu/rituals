// Package cestimate - Content managed by Project Forge, see [projectforge.md] for details.
package cestimate

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
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
		ps.SetTitleAndData("Permissions", ret)
		estimateIDsByEstimateID := lo.Map(ret, func(x *epermission.EstimatePermission, _ int) uuid.UUID {
			return x.EstimateID
		})
		estimatesByEstimateID, err := as.Services.Estimate.GetMultiple(ps.Context, nil, nil, ps.Logger, estimateIDsByEstimateID...)
		if err != nil {
			return "", err
		}
		page := &vepermission.List{Models: ret, EstimatesByEstimateID: estimatesByEstimateID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "estimate", "epermission")
	})
}

//nolint:lll
func EstimatePermissionDetail(rc *fasthttp.RequestCtx) {
	controller.Act("epermission.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := epermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Permission)", ret)

		estimateByEstimateID, _ := as.Services.Estimate.Get(ps.Context, nil, ret.EstimateID, ps.Logger)

		return controller.Render(rc, as, &vepermission.Detail{Model: ret, EstimateByEstimateID: estimateByEstimateID}, ps, "estimate", "epermission", ret.TitleString()+"**permission")
	})
}

func EstimatePermissionCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("epermission.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &epermission.EstimatePermission{}
		if string(rc.QueryArgs().Peek("prototype")) == "random" {
			ret = epermission.Random()
		}
		ps.SetTitleAndData("Create [EstimatePermission]", ret)
		ps.Data = ret
		return controller.Render(rc, as, &vepermission.Edit{Model: ret, IsNew: true}, ps, "estimate", "epermission", "Create")
	})
}

func EstimatePermissionRandom(rc *fasthttp.RequestCtx) {
	controller.Act("epermission.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.EstimatePermission.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random EstimatePermission")
		}
		return ret.WebPath(), nil
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
		ps.SetTitleAndData("Edit "+ret.String(), ret)
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
		return controller.FlashAndRedir(true, msg, "/admin/db/estimate/permission", rc, ps)
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
		return nil, errors.Wrap(err, "must provide [key] as a string argument")
	}
	valueArg, err := cutil.RCRequiredString(rc, "value", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [value] as a string argument")
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
