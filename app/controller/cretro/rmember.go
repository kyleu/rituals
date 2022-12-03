// Content managed by Project Forge, see [projectforge.md] for details.
package cretro

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/retro/rmember"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vretro/vrmember"
)

func RetroMemberList(rc *fasthttp.RequestCtx) {
	controller.Act("rmember.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("rmember", nil, ps.Logger).Sanitize("rmember")
		ret, err := as.Services.RetroMember.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Members"
		ps.Data = ret
		retroIDs := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			retroIDs = append(retroIDs, x.RetroID)
		}
		retros, err := as.Services.Retro.GetMultiple(ps.Context, nil, ps.Logger, retroIDs...)
		if err != nil {
			return "", err
		}
		userIDs := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			userIDs = append(userIDs, x.UserID)
		}
		users, err := as.Services.User.GetMultiple(ps.Context, nil, ps.Logger, userIDs...)
		if err != nil {
			return "", err
		}
		page := &vrmember.List{Models: ret, Retros: retros, Users: users, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "retro", "rmember")
	})
}

func RetroMemberDetail(rc *fasthttp.RequestCtx) {
	controller.Act("rmember.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rmemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Member)"
		ps.Data = ret
		return controller.Render(rc, as, &vrmember.Detail{Model: ret}, ps, "retro", "rmember", ret.String())
	})
}

func RetroMemberCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("rmember.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &rmember.RetroMember{}
		ps.Title = "Create [RetroMember]"
		ps.Data = ret
		return controller.Render(rc, as, &vrmember.Edit{Model: ret, IsNew: true}, ps, "retro", "rmember", "Create")
	})
}

func RetroMemberCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("rmember.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := rmember.Random()
		ps.Title = "Create Random RetroMember"
		ps.Data = ret
		return controller.Render(rc, as, &vrmember.Edit{Model: ret, IsNew: true}, ps, "retro", "rmember", "Create")
	})
}

func RetroMemberCreate(rc *fasthttp.RequestCtx) {
	controller.Act("rmember.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rmemberFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse RetroMember from form")
		}
		err = as.Services.RetroMember.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created RetroMember")
		}
		msg := fmt.Sprintf("RetroMember [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func RetroMemberEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("rmember.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rmemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vrmember.Edit{Model: ret}, ps, "retro", "rmember", ret.String())
	})
}

func RetroMemberEdit(rc *fasthttp.RequestCtx) {
	controller.Act("rmember.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rmemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := rmemberFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse RetroMember from form")
		}
		frm.RetroID = ret.RetroID
		frm.UserID = ret.UserID
		err = as.Services.RetroMember.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update RetroMember [%s]", frm.String())
		}
		msg := fmt.Sprintf("RetroMember [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func RetroMemberDelete(rc *fasthttp.RequestCtx) {
	controller.Act("rmember.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rmemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.RetroMember.Delete(ps.Context, nil, ret.RetroID, ret.UserID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete member [%s]", ret.String())
		}
		msg := fmt.Sprintf("RetroMember [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/retroMember", rc, ps)
	})
}

func rmemberFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*rmember.RetroMember, error) {
	retroIDArgStr, err := cutil.RCRequiredString(rc, "retroID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [retroID] as an argument")
	}
	retroIDArgP := util.UUIDFromString(retroIDArgStr)
	if retroIDArgP == nil {
		return nil, errors.Errorf("argument [retroID] (%s) is not a valid UUID", retroIDArgStr)
	}
	retroIDArg := *retroIDArgP
	userIDArgStr, err := cutil.RCRequiredString(rc, "userID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [userID] as an argument")
	}
	userIDArgP := util.UUIDFromString(userIDArgStr)
	if userIDArgP == nil {
		return nil, errors.Errorf("argument [userID] (%s) is not a valid UUID", userIDArgStr)
	}
	userIDArg := *userIDArgP
	return as.Services.RetroMember.Get(ps.Context, nil, retroIDArg, userIDArg, ps.Logger)
}

func rmemberFromForm(rc *fasthttp.RequestCtx, setPK bool) (*rmember.RetroMember, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return rmember.FromMap(frm, setPK)
}
