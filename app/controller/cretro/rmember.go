// Package cretro - Content managed by Project Forge, see [projectforge.md] for details.
package cretro

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
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
		ps.SetTitleAndData("Members", ret)
		retroIDsByRetroID := lo.Map(ret, func(x *rmember.RetroMember, _ int) uuid.UUID {
			return x.RetroID
		})
		retrosByRetroID, err := as.Services.Retro.GetMultiple(ps.Context, nil, nil, ps.Logger, retroIDsByRetroID...)
		if err != nil {
			return "", err
		}
		userIDsByUserID := lo.Map(ret, func(x *rmember.RetroMember, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vrmember.List{Models: ret, RetrosByRetroID: retrosByRetroID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "retro", "rmember")
	})
}

func RetroMemberDetail(rc *fasthttp.RequestCtx) {
	controller.Act("rmember.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rmemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Member)", ret)

		retroByRetroID, _ := as.Services.Retro.Get(ps.Context, nil, ret.RetroID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(rc, as, &vrmember.Detail{
			Model:          ret,
			RetroByRetroID: retroByRetroID,
			UserByUserID:   userByUserID,
		}, ps, "retro", "rmember", ret.TitleString()+"**users")
	})
}

func RetroMemberCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("rmember.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &rmember.RetroMember{}
		if string(rc.QueryArgs().Peek("prototype")) == util.KeyRandom {
			ret = rmember.Random()
			randomRetro, err := as.Services.Retro.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomRetro != nil {
				ret.RetroID = randomRetro.ID
			}
			randomUser, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomUser != nil {
				ret.UserID = randomUser.ID
			}
		}
		ps.SetTitleAndData("Create [RetroMember]", ret)
		ps.Data = ret
		return controller.Render(rc, as, &vrmember.Edit{Model: ret, IsNew: true}, ps, "retro", "rmember", "Create")
	})
}

func RetroMemberRandom(rc *fasthttp.RequestCtx) {
	controller.Act("rmember.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.RetroMember.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random RetroMember")
		}
		return ret.WebPath(), nil
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
		ps.SetTitleAndData("Edit "+ret.String(), ret)
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
		return controller.FlashAndRedir(true, msg, "/admin/db/retro/member", rc, ps)
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
