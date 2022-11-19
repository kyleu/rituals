// Content managed by Project Forge, see [projectforge.md] for details.
package csprint

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint/spermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vsprint/vspermission"
)

func SprintPermissionList(rc *fasthttp.RequestCtx) {
	controller.Act("spermission.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("spermission", nil, ps.Logger).Sanitize("spermission")
		ret, err := as.Services.SprintPermission.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Permissions"
		ps.Data = ret
		sprintIDs := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			sprintIDs = append(sprintIDs, x.SprintID)
		}
		sprints, err := as.Services.Sprint.GetMultiple(ps.Context, nil, ps.Logger, sprintIDs...)
		if err != nil {
			return "", err
		}
		return controller.Render(rc, as, &vspermission.List{Models: ret, Sprints: sprints, Params: ps.Params}, ps, "sprint", "spermission")
	})
}

func SprintPermissionDetail(rc *fasthttp.RequestCtx) {
	controller.Act("spermission.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := spermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Permission)"
		ps.Data = ret
		return controller.Render(rc, as, &vspermission.Detail{Model: ret}, ps, "sprint", "spermission", ret.String())
	})
}

func SprintPermissionCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("spermission.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &spermission.SprintPermission{}
		ps.Title = "Create [SprintPermission]"
		ps.Data = ret
		return controller.Render(rc, as, &vspermission.Edit{Model: ret, IsNew: true}, ps, "sprint", "spermission", "Create")
	})
}

func SprintPermissionCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("spermission.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := spermission.Random()
		ps.Title = "Create Random SprintPermission"
		ps.Data = ret
		return controller.Render(rc, as, &vspermission.Edit{Model: ret, IsNew: true}, ps, "sprint", "spermission", "Create")
	})
}

func SprintPermissionCreate(rc *fasthttp.RequestCtx) {
	controller.Act("spermission.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := spermissionFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse SprintPermission from form")
		}
		err = as.Services.SprintPermission.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created SprintPermission")
		}
		msg := fmt.Sprintf("SprintPermission [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func SprintPermissionEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("spermission.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := spermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vspermission.Edit{Model: ret}, ps, "sprint", "spermission", ret.String())
	})
}

func SprintPermissionEdit(rc *fasthttp.RequestCtx) {
	controller.Act("spermission.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := spermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := spermissionFromForm(rc, false)
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
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func SprintPermissionDelete(rc *fasthttp.RequestCtx) {
	controller.Act("spermission.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := spermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.SprintPermission.Delete(ps.Context, nil, ret.SprintID, ret.Key, ret.Value, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete permission [%s]", ret.String())
		}
		msg := fmt.Sprintf("SprintPermission [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/sprintPermission", rc, ps)
	})
}

func spermissionFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*spermission.SprintPermission, error) {
	sprintIDArgStr, err := cutil.RCRequiredString(rc, "sprintID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [sprintID] as an argument")
	}
	sprintIDArgP := util.UUIDFromString(sprintIDArgStr)
	if sprintIDArgP == nil {
		return nil, errors.Errorf("argument [sprintID] (%s) is not a valid UUID", sprintIDArgStr)
	}
	sprintIDArg := *sprintIDArgP
	keyArg, err := cutil.RCRequiredString(rc, "key", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [key] as an argument")
	}
	valueArg, err := cutil.RCRequiredString(rc, "value", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [value] as an argument")
	}
	return as.Services.SprintPermission.Get(ps.Context, nil, sprintIDArg, keyArg, valueArg, ps.Logger)
}

func spermissionFromForm(rc *fasthttp.RequestCtx, setPK bool) (*spermission.SprintPermission, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return spermission.FromMap(frm, setPK)
}
