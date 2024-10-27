package csprint

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint/spermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vsprint/vspermission"
)

func SprintPermissionList(w http.ResponseWriter, r *http.Request) {
	controller.Act("spermission.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("spermission", ps.Logger)
		ret, err := as.Services.SprintPermission.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Permissions", ret)
		sprintIDsBySprintID := lo.Map(ret, func(x *spermission.SprintPermission, _ int) uuid.UUID {
			return x.SprintID
		})
		sprintsBySprintID, err := as.Services.Sprint.GetMultiple(ps.Context, nil, nil, ps.Logger, sprintIDsBySprintID...)
		if err != nil {
			return "", err
		}
		page := &vspermission.List{Models: ret, SprintsBySprintID: sprintsBySprintID, Params: ps.Params}
		return controller.Render(r, as, page, ps, "sprint", "spermission")
	})
}

//nolint:lll
func SprintPermissionDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("spermission.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := spermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Permission)", ret)

		sprintBySprintID, _ := as.Services.Sprint.Get(ps.Context, nil, ret.SprintID, ps.Logger)

		return controller.Render(r, as, &vspermission.Detail{Model: ret, SprintBySprintID: sprintBySprintID}, ps, "sprint", "spermission", ret.TitleString()+"**permission")
	})
}

func SprintPermissionCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("spermission.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &spermission.SprintPermission{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = spermission.RandomSprintPermission()
			randomSprint, err := as.Services.Sprint.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomSprint != nil {
				ret.SprintID = randomSprint.ID
			}
		}
		ps.SetTitleAndData("Create [SprintPermission]", ret)
		ps.Data = ret
		return controller.Render(r, as, &vspermission.Edit{Model: ret, IsNew: true}, ps, "sprint", "spermission", "Create")
	})
}

func SprintPermissionRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("spermission.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.SprintPermission.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random SprintPermission")
		}
		return ret.WebPath(), nil
	})
}

func SprintPermissionCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("spermission.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := spermissionFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse SprintPermission from form")
		}
		err = as.Services.SprintPermission.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created SprintPermission")
		}
		msg := fmt.Sprintf("SprintPermission [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func SprintPermissionEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("spermission.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := spermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &vspermission.Edit{Model: ret}, ps, "sprint", "spermission", ret.String())
	})
}

func SprintPermissionEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("spermission.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := spermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := spermissionFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse SprintPermission from form")
		}
		frm.SprintID = ret.SprintID
		frm.Key = ret.Key
		frm.Value = ret.Value
		err = as.Services.SprintPermission.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update SprintPermission [%s]", frm.String())
		}
		msg := fmt.Sprintf("SprintPermission [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func SprintPermissionDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("spermission.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := spermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.SprintPermission.Delete(ps.Context, nil, ret.SprintID, ret.Key, ret.Value, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete permission [%s]", ret.String())
		}
		msg := fmt.Sprintf("SprintPermission [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/sprint/permission", ps)
	})
}

func spermissionFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*spermission.SprintPermission, error) {
	sprintIDArgStr, err := cutil.PathString(r, "sprintID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [sprintID] as an argument")
	}
	sprintIDArgP := util.UUIDFromString(sprintIDArgStr)
	if sprintIDArgP == nil {
		return nil, errors.Errorf("argument [sprintID] (%s) is not a valid UUID", sprintIDArgStr)
	}
	sprintIDArg := *sprintIDArgP
	keyArg, err := cutil.PathString(r, "key", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [key] as a string argument")
	}
	valueArg, err := cutil.PathString(r, "value", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [value] as a string argument")
	}
	return as.Services.SprintPermission.Get(ps.Context, nil, sprintIDArg, keyArg, valueArg, ps.Logger)
}

func spermissionFromForm(r *http.Request, b []byte, setPK bool) (*spermission.SprintPermission, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := spermission.SprintPermissionFromMap(frm, setPK)
	return ret, err
}
