package cretro

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vretro/vrpermission"
)

func RetroPermissionList(w http.ResponseWriter, r *http.Request) {
	controller.Act("rpermission.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("rpermission", ps.Logger)
		ret, err := as.Services.RetroPermission.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Permissions", ret)
		retroIDsByRetroID := lo.Map(ret, func(x *rpermission.RetroPermission, _ int) uuid.UUID {
			return x.RetroID
		})
		retrosByRetroID, err := as.Services.Retro.GetMultiple(ps.Context, nil, nil, ps.Logger, retroIDsByRetroID...)
		if err != nil {
			return "", err
		}
		page := &vrpermission.List{Models: ret, RetrosByRetroID: retrosByRetroID, Params: ps.Params}
		return controller.Render(r, as, page, ps, "retro", "rpermission")
	})
}

//nolint:lll
func RetroPermissionDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("rpermission.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rpermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Permission)", ret)

		retroByRetroID, _ := as.Services.Retro.Get(ps.Context, nil, ret.RetroID, ps.Logger)

		return controller.Render(r, as, &vrpermission.Detail{Model: ret, RetroByRetroID: retroByRetroID}, ps, "retro", "rpermission", ret.TitleString()+"**permission")
	})
}

func RetroPermissionCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("rpermission.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &rpermission.RetroPermission{}
		if cutil.QueryStringString(r, "prototype") == util.KeyRandom {
			ret = rpermission.RandomRetroPermission()
			randomRetro, err := as.Services.Retro.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomRetro != nil {
				ret.RetroID = randomRetro.ID
			}
		}
		ps.SetTitleAndData("Create [RetroPermission]", ret)
		return controller.Render(r, as, &vrpermission.Edit{Model: ret, IsNew: true}, ps, "retro", "rpermission", "Create")
	})
}

func RetroPermissionRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("rpermission.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.RetroPermission.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random RetroPermission")
		}
		return ret.WebPath(), nil
	})
}

func RetroPermissionCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("rpermission.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rpermissionFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse RetroPermission from form")
		}
		err = as.Services.RetroPermission.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created RetroPermission")
		}
		msg := fmt.Sprintf("RetroPermission [%s] created", ret.TitleString())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func RetroPermissionEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("rpermission.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rpermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &vrpermission.Edit{Model: ret}, ps, "retro", "rpermission", ret.String())
	})
}

func RetroPermissionEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("rpermission.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rpermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := rpermissionFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse RetroPermission from form")
		}
		frm.RetroID = ret.RetroID
		frm.Key = ret.Key
		frm.Value = ret.Value
		err = as.Services.RetroPermission.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update RetroPermission [%s]", frm.String())
		}
		msg := fmt.Sprintf("RetroPermission [%s] updated", frm.TitleString())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func RetroPermissionDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("rpermission.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rpermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.RetroPermission.Delete(ps.Context, nil, ret.RetroID, ret.Key, ret.Value, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete permission [%s]", ret.String())
		}
		msg := fmt.Sprintf("RetroPermission [%s] deleted", ret.TitleString())
		return controller.FlashAndRedir(true, msg, "/admin/db/retro/permission", ps)
	})
}

func rpermissionFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*rpermission.RetroPermission, error) {
	retroIDArgStr, err := cutil.PathString(r, "retroID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [retroID] as an argument")
	}
	retroIDArgP := util.UUIDFromString(retroIDArgStr)
	if retroIDArgP == nil {
		return nil, errors.Errorf("argument [retroID] (%s) is not a valid UUID", retroIDArgStr)
	}
	retroIDArg := *retroIDArgP
	keyArg, err := cutil.PathString(r, "key", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [key] as a string argument")
	}
	valueArg, err := cutil.PathString(r, "value", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [value] as a string argument")
	}
	return as.Services.RetroPermission.Get(ps.Context, nil, retroIDArg, keyArg, valueArg, ps.Logger)
}

func rpermissionFromForm(r *http.Request, b []byte, setPK bool) (*rpermission.RetroPermission, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := rpermission.RetroPermissionFromMap(frm, setPK)
	return ret, err
}
