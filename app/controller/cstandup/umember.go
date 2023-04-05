// Content managed by Project Forge, see [projectforge.md] for details.
package cstandup

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vstandup/vumember"
)

func StandupMemberList(rc *fasthttp.RequestCtx) {
	controller.Act("umember.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("umember", nil, ps.Logger).Sanitize("umember")
		ret, err := as.Services.StandupMember.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Members"
		ps.Data = ret
		standupIDsByStandupID := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			standupIDsByStandupID = append(standupIDsByStandupID, x.StandupID)
		}
		standupsByStandupID, err := as.Services.Standup.GetMultiple(ps.Context, nil, ps.Logger, standupIDsByStandupID...)
		if err != nil {
			return "", err
		}
		userIDsByUserID := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			userIDsByUserID = append(userIDsByUserID, x.UserID)
		}
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vumember.List{Models: ret, StandupsByStandupID: standupsByStandupID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "standup", "umember")
	})
}

func StandupMemberDetail(rc *fasthttp.RequestCtx) {
	controller.Act("umember.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := umemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Member)"
		ps.Data = ret

		standupByStandupID, _ := as.Services.Standup.Get(ps.Context, nil, ret.StandupID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(rc, as, &vumember.Detail{
			Model:              ret,
			StandupByStandupID: standupByStandupID,
			UserByUserID:       userByUserID,
		}, ps, "standup", "umember", ret.String())
	})
}

func StandupMemberCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("umember.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &umember.StandupMember{}
		ps.Title = "Create [StandupMember]"
		ps.Data = ret
		return controller.Render(rc, as, &vumember.Edit{Model: ret, IsNew: true}, ps, "standup", "umember", "Create")
	})
}

func StandupMemberCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("umember.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := umember.Random()
		ps.Title = "Create Random StandupMember"
		ps.Data = ret
		return controller.Render(rc, as, &vumember.Edit{Model: ret, IsNew: true}, ps, "standup", "umember", "Create")
	})
}

func StandupMemberCreate(rc *fasthttp.RequestCtx) {
	controller.Act("umember.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := umemberFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse StandupMember from form")
		}
		err = as.Services.StandupMember.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created StandupMember")
		}
		msg := fmt.Sprintf("StandupMember [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func StandupMemberEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("umember.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := umemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vumember.Edit{Model: ret}, ps, "standup", "umember", ret.String())
	})
}

func StandupMemberEdit(rc *fasthttp.RequestCtx) {
	controller.Act("umember.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := umemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := umemberFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse StandupMember from form")
		}
		frm.StandupID = ret.StandupID
		frm.UserID = ret.UserID
		err = as.Services.StandupMember.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update StandupMember [%s]", frm.String())
		}
		msg := fmt.Sprintf("StandupMember [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func StandupMemberDelete(rc *fasthttp.RequestCtx) {
	controller.Act("umember.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := umemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.StandupMember.Delete(ps.Context, nil, ret.StandupID, ret.UserID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete member [%s]", ret.String())
		}
		msg := fmt.Sprintf("StandupMember [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/standupMember", rc, ps)
	})
}

func umemberFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*umember.StandupMember, error) {
	standupIDArgStr, err := cutil.RCRequiredString(rc, "standupID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [standupID] as an argument")
	}
	standupIDArgP := util.UUIDFromString(standupIDArgStr)
	if standupIDArgP == nil {
		return nil, errors.Errorf("argument [standupID] (%s) is not a valid UUID", standupIDArgStr)
	}
	standupIDArg := *standupIDArgP
	userIDArgStr, err := cutil.RCRequiredString(rc, "userID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [userID] as an argument")
	}
	userIDArgP := util.UUIDFromString(userIDArgStr)
	if userIDArgP == nil {
		return nil, errors.Errorf("argument [userID] (%s) is not a valid UUID", userIDArgStr)
	}
	userIDArg := *userIDArgP
	return as.Services.StandupMember.Get(ps.Context, nil, standupIDArg, userIDArg, ps.Logger)
}

func umemberFromForm(rc *fasthttp.RequestCtx, setPK bool) (*umember.StandupMember, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return umember.FromMap(frm, setPK)
}
