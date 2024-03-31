// Package cestimate - Content managed by Project Forge, see [projectforge.md] for details.
package cestimate

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vestimate/vepermission"
)

func EstimatePermissionList(w http.ResponseWriter, r *http.Request) {
	controller.Act("epermission.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("epermission", ps.Logger)
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
		return controller.Render(w, r, as, page, ps, "estimate", "epermission")
	})
}

//nolint:lll
func EstimatePermissionDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("epermission.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := epermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Permission)", ret)

		estimateByEstimateID, _ := as.Services.Estimate.Get(ps.Context, nil, ret.EstimateID, ps.Logger)

		return controller.Render(w, r, as, &vepermission.Detail{Model: ret, EstimateByEstimateID: estimateByEstimateID}, ps, "estimate", "epermission", ret.TitleString()+"**permission")
	})
}

func EstimatePermissionCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("epermission.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &epermission.EstimatePermission{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = epermission.Random()
			randomEstimate, err := as.Services.Estimate.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomEstimate != nil {
				ret.EstimateID = randomEstimate.ID
			}
		}
		ps.SetTitleAndData("Create [EstimatePermission]", ret)
		ps.Data = ret
		return controller.Render(w, r, as, &vepermission.Edit{Model: ret, IsNew: true}, ps, "estimate", "epermission", "Create")
	})
}

func EstimatePermissionRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("epermission.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.EstimatePermission.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random EstimatePermission")
		}
		return ret.WebPath(), nil
	})
}

func EstimatePermissionCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("epermission.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := epermissionFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse EstimatePermission from form")
		}
		err = as.Services.EstimatePermission.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created EstimatePermission")
		}
		msg := fmt.Sprintf("EstimatePermission [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), w, ps)
	})
}

func EstimatePermissionEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("epermission.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := epermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(w, r, as, &vepermission.Edit{Model: ret}, ps, "estimate", "epermission", ret.String())
	})
}

func EstimatePermissionEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("epermission.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := epermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := epermissionFromForm(r, ps.RequestBody, false)
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
		return controller.FlashAndRedir(true, msg, frm.WebPath(), w, ps)
	})
}

func EstimatePermissionDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("epermission.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := epermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.EstimatePermission.Delete(ps.Context, nil, ret.EstimateID, ret.Key, ret.Value, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete permission [%s]", ret.String())
		}
		msg := fmt.Sprintf("EstimatePermission [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/estimate/permission", w, ps)
	})
}

func epermissionFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*epermission.EstimatePermission, error) {
	estimateIDArgStr, err := cutil.RCRequiredString(r, "estimateID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [estimateID] as an argument")
	}
	estimateIDArgP := util.UUIDFromString(estimateIDArgStr)
	if estimateIDArgP == nil {
		return nil, errors.Errorf("argument [estimateID] (%s) is not a valid UUID", estimateIDArgStr)
	}
	estimateIDArg := *estimateIDArgP
	keyArg, err := cutil.RCRequiredString(r, "key", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [key] as a string argument")
	}
	valueArg, err := cutil.RCRequiredString(r, "value", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [value] as a string argument")
	}
	return as.Services.EstimatePermission.Get(ps.Context, nil, estimateIDArg, keyArg, valueArg, ps.Logger)
}

func epermissionFromForm(r *http.Request, b []byte, setPK bool) (*epermission.EstimatePermission, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return epermission.FromMap(frm, setPK)
}
