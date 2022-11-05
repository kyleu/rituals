// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vemail"
)

func EmailList(rc *fasthttp.RequestCtx) {
	Act("email.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("email", nil, ps.Logger).Sanitize("email")
		ret, err := as.Services.Email.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Emails"
		ps.Data = ret
		userIDs := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			userIDs = append(userIDs, x.UserID)
		}
		users, err := as.Services.User.GetMultiple(ps.Context, nil, ps.Logger, userIDs...)
		if err != nil {
			return "", err
		}
		return Render(rc, as, &vemail.List{Models: ret, Users: users, Params: ps.Params}, ps, "email")
	})
}

func EmailDetail(rc *fasthttp.RequestCtx) {
	Act("email.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := emailFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Email)"
		ps.Data = ret
		return Render(rc, as, &vemail.Detail{Model: ret}, ps, "email", ret.String())
	})
}

func EmailCreateForm(rc *fasthttp.RequestCtx) {
	Act("email.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &email.Email{}
		ps.Title = "Create [Email]"
		ps.Data = ret
		return Render(rc, as, &vemail.Edit{Model: ret, IsNew: true}, ps, "email", "Create")
	})
}

func EmailCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("email.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := email.Random()
		ps.Title = "Create Random Email"
		ps.Data = ret
		return Render(rc, as, &vemail.Edit{Model: ret, IsNew: true}, ps, "email", "Create")
	})
}

func EmailCreate(rc *fasthttp.RequestCtx) {
	Act("email.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := emailFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Email from form")
		}
		err = as.Services.Email.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Email")
		}
		msg := fmt.Sprintf("Email [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func EmailEditForm(rc *fasthttp.RequestCtx) {
	Act("email.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := emailFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vemail.Edit{Model: ret}, ps, "email", ret.String())
	})
}

func EmailEdit(rc *fasthttp.RequestCtx) {
	Act("email.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := emailFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := emailFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Email from form")
		}
		frm.ID = ret.ID
		err = as.Services.Email.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Email [%s]", frm.String())
		}
		msg := fmt.Sprintf("Email [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func EmailDelete(rc *fasthttp.RequestCtx) {
	Act("email.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := emailFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Email.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete email [%s]", ret.String())
		}
		msg := fmt.Sprintf("Email [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/email", rc, ps)
	})
}

func emailFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*email.Email, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
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

func emailFromForm(rc *fasthttp.RequestCtx, setPK bool) (*email.Email, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return email.FromMap(frm, setPK)
}
