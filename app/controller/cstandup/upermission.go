// Package cstandup - Content managed by Project Forge, see [projectforge.md] for details.
package cstandup

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vstandup/vupermission"
)

func StandupPermissionList(w http.ResponseWriter, r *http.Request) {
	controller.Act("upermission.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("upermission", ps.Logger)
		ret, err := as.Services.StandupPermission.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Permissions", ret)
		standupIDsByStandupID := lo.Map(ret, func(x *upermission.StandupPermission, _ int) uuid.UUID {
			return x.StandupID
		})
		standupsByStandupID, err := as.Services.Standup.GetMultiple(ps.Context, nil, nil, ps.Logger, standupIDsByStandupID...)
		if err != nil {
			return "", err
		}
		page := &vupermission.List{Models: ret, StandupsByStandupID: standupsByStandupID, Params: ps.Params}
		return controller.Render(w, r, as, page, ps, "standup", "upermission")
	})
}

//nolint:lll
func StandupPermissionDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("upermission.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := upermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Permission)", ret)

		standupByStandupID, _ := as.Services.Standup.Get(ps.Context, nil, ret.StandupID, ps.Logger)

		return controller.Render(w, r, as, &vupermission.Detail{Model: ret, StandupByStandupID: standupByStandupID}, ps, "standup", "upermission", ret.TitleString()+"**permission")
	})
}

func StandupPermissionCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("upermission.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &upermission.StandupPermission{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = upermission.Random()
			randomStandup, err := as.Services.Standup.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomStandup != nil {
				ret.StandupID = randomStandup.ID
			}
		}
		ps.SetTitleAndData("Create [StandupPermission]", ret)
		ps.Data = ret
		return controller.Render(w, r, as, &vupermission.Edit{Model: ret, IsNew: true}, ps, "standup", "upermission", "Create")
	})
}

func StandupPermissionRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("upermission.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.StandupPermission.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random StandupPermission")
		}
		return ret.WebPath(), nil
	})
}

func StandupPermissionCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("upermission.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := upermissionFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse StandupPermission from form")
		}
		err = as.Services.StandupPermission.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created StandupPermission")
		}
		msg := fmt.Sprintf("StandupPermission [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), w, ps)
	})
}

func StandupPermissionEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("upermission.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := upermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(w, r, as, &vupermission.Edit{Model: ret}, ps, "standup", "upermission", ret.String())
	})
}

func StandupPermissionEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("upermission.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := upermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := upermissionFromForm(r, ps.RequestBody, false)
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
		return controller.FlashAndRedir(true, msg, frm.WebPath(), w, ps)
	})
}

func StandupPermissionDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("upermission.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := upermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.StandupPermission.Delete(ps.Context, nil, ret.StandupID, ret.Key, ret.Value, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete permission [%s]", ret.String())
		}
		msg := fmt.Sprintf("StandupPermission [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/standup/permission", w, ps)
	})
}

func upermissionFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*upermission.StandupPermission, error) {
	standupIDArgStr, err := cutil.RCRequiredString(r, "standupID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [standupID] as an argument")
	}
	standupIDArgP := util.UUIDFromString(standupIDArgStr)
	if standupIDArgP == nil {
		return nil, errors.Errorf("argument [standupID] (%s) is not a valid UUID", standupIDArgStr)
	}
	standupIDArg := *standupIDArgP
	keyArg, err := cutil.RCRequiredString(r, "key", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [key] as a string argument")
	}
	valueArg, err := cutil.RCRequiredString(r, "value", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [value] as a string argument")
	}
	return as.Services.StandupPermission.Get(ps.Context, nil, standupIDArg, keyArg, valueArg, ps.Logger)
}

func upermissionFromForm(r *http.Request, b []byte, setPK bool) (*upermission.StandupPermission, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return upermission.FromMap(frm, setPK)
}
