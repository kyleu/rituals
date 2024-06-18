package controller

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vemail"
)

func EmailList(w http.ResponseWriter, r *http.Request) {
	Act("email.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("email", ps.Logger)
		ret, err := as.Services.Email.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Emails", ret)
		userIDsByUserID := lo.Map(ret, func(x *email.Email, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vemail.List{Models: ret, UsersByUserID: usersByUserID, Params: ps.Params}
		return Render(r, as, page, ps, "email")
	})
}

func EmailDetail(w http.ResponseWriter, r *http.Request) {
	Act("email.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := emailFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Email)", ret)

		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return Render(r, as, &vemail.Detail{Model: ret, UserByUserID: userByUserID}, ps, "email", ret.TitleString()+"**email")
	})
}

func EmailCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("email.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &email.Email{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = email.Random()
			randomUser, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomUser != nil {
				ret.UserID = randomUser.ID
			}
		}
		ps.SetTitleAndData("Create [Email]", ret)
		ps.Data = ret
		return Render(r, as, &vemail.Edit{Model: ret, IsNew: true}, ps, "email", "Create")
	})
}

func EmailRandom(w http.ResponseWriter, r *http.Request) {
	Act("email.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Email.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Email")
		}
		return ret.WebPath(), nil
	})
}

func EmailCreate(w http.ResponseWriter, r *http.Request) {
	Act("email.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := emailFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Email from form")
		}
		err = as.Services.Email.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Email")
		}
		msg := fmt.Sprintf("Email [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func EmailEditForm(w http.ResponseWriter, r *http.Request) {
	Act("email.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := emailFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vemail.Edit{Model: ret}, ps, "email", ret.String())
	})
}

func EmailEdit(w http.ResponseWriter, r *http.Request) {
	Act("email.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := emailFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := emailFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Email from form")
		}
		frm.ID = ret.ID
		err = as.Services.Email.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Email [%s]", frm.String())
		}
		msg := fmt.Sprintf("Email [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func EmailDelete(w http.ResponseWriter, r *http.Request) {
	Act("email.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := emailFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Email.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete email [%s]", ret.String())
		}
		msg := fmt.Sprintf("Email [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/admin/db/email", ps)
	})
}

func emailFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*email.Email, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Email.Get(ps.Context, nil, idArg, ps.Logger)
}

func emailFromForm(r *http.Request, b []byte, setPK bool) (*email.Email, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := email.FromMap(frm, setPK)
	return ret, err
}
