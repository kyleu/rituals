// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vaction"
)

func ActionList(rc *fasthttp.RequestCtx) {
	Act("action.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("action", nil, ps.Logger).Sanitize("action")
		ret, err := as.Services.Action.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Actions"
		ps.Data = ret
		userIDsByUserID := lo.Map(ret, func(x *action.Action, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vaction.List{Models: ret, UsersByUserID: usersByUserID, Params: ps.Params}
		return Render(rc, as, page, ps, "action")
	})
}

func ActionDetail(rc *fasthttp.RequestCtx) {
	Act("action.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := actionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Action)"
		ps.Data = ret

		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return Render(rc, as, &vaction.Detail{Model: ret, UserByUserID: userByUserID}, ps, "action", ret.String())
	})
}

func ActionCreateForm(rc *fasthttp.RequestCtx) {
	Act("action.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &action.Action{}
		ps.Title = "Create [Action]"
		ps.Data = ret
		return Render(rc, as, &vaction.Edit{Model: ret, IsNew: true}, ps, "action", "Create")
	})
}

func ActionCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("action.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := action.Random()
		ps.Title = "Create Random Action"
		ps.Data = ret
		return Render(rc, as, &vaction.Edit{Model: ret, IsNew: true}, ps, "action", "Create")
	})
}

func ActionCreate(rc *fasthttp.RequestCtx) {
	Act("action.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := actionFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Action from form")
		}
		err = as.Services.Action.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Action")
		}
		msg := fmt.Sprintf("Action [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func ActionEditForm(rc *fasthttp.RequestCtx) {
	Act("action.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := actionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vaction.Edit{Model: ret}, ps, "action", ret.String())
	})
}

func ActionEdit(rc *fasthttp.RequestCtx) {
	Act("action.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := actionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := actionFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Action from form")
		}
		frm.ID = ret.ID
		err = as.Services.Action.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Action [%s]", frm.String())
		}
		msg := fmt.Sprintf("Action [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func ActionDelete(rc *fasthttp.RequestCtx) {
	Act("action.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := actionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Action.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete action [%s]", ret.String())
		}
		msg := fmt.Sprintf("Action [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/admin/db/action", rc, ps)
	})
}

func actionFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*action.Action, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Action.Get(ps.Context, nil, idArg, ps.Logger)
}

func actionFromForm(rc *fasthttp.RequestCtx, setPK bool) (*action.Action, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return action.FromMap(frm, setPK)
}
