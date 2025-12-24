package controller

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vaction"
)

func ActionList(w http.ResponseWriter, r *http.Request) {
	Act("action.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("action", ps.Logger)
		ret, err := as.Services.Action.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Actions", ret)
		userIDsByUserID := lo.Map(ret, func(x *action.Action, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vaction.List{Models: ret, UsersByUserID: usersByUserID, Params: ps.Params}
		return Render(r, as, page, ps, "action")
	})
}

func ActionDetail(w http.ResponseWriter, r *http.Request) {
	Act("action.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := actionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Action)", ret)

		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return Render(r, as, &vaction.Detail{Model: ret, UserByUserID: userByUserID}, ps, "action", ret.TitleString()+"**action")
	})
}

func ActionCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("action.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &action.Action{}
		if cutil.QueryStringString(r, "prototype") == util.KeyRandom {
			ret = action.RandomAction()
			randomUser, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomUser != nil {
				ret.UserID = randomUser.ID
			}
		}
		ps.SetTitleAndData("Create [Action]", ret)
		return Render(r, as, &vaction.Edit{Model: ret, IsNew: true}, ps, "action", "Create")
	})
}

func ActionRandom(w http.ResponseWriter, r *http.Request) {
	Act("action.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Action.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Action")
		}
		return ret.WebPath(), nil
	})
}

func ActionCreate(w http.ResponseWriter, r *http.Request) {
	Act("action.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := actionFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Action from form")
		}
		err = as.Services.Action.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Action")
		}
		msg := fmt.Sprintf("Action [%s] created", ret.TitleString())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func ActionEditForm(w http.ResponseWriter, r *http.Request) {
	Act("action.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := actionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vaction.Edit{Model: ret}, ps, "action", ret.String())
	})
}

func ActionEdit(w http.ResponseWriter, r *http.Request) {
	Act("action.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := actionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := actionFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Action from form")
		}
		frm.ID = ret.ID
		err = as.Services.Action.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Action [%s]", frm.String())
		}
		msg := fmt.Sprintf("Action [%s] updated", frm.TitleString())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func ActionDelete(w http.ResponseWriter, r *http.Request) {
	Act("action.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := actionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Action.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete action [%s]", ret.String())
		}
		msg := fmt.Sprintf("Action [%s] deleted", ret.TitleString())
		return FlashAndRedir(true, msg, "/admin/db/action", ps)
	})
}

func actionFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*action.Action, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
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

func actionFromForm(r *http.Request, b []byte, setPK bool) (*action.Action, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := action.ActionFromMap(frm, setPK)
	return ret, err
}
