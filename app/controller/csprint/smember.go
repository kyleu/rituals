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
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vsprint/vsmember"
)

func SprintMemberList(rc *fasthttp.RequestCtx) {
	controller.Act("smember.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("smember", nil, ps.Logger).Sanitize("smember")
		ret, err := as.Services.SprintMember.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Members"
		ps.Data = ret
		sprintIDsBySprintID := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			sprintIDsBySprintID = append(sprintIDsBySprintID, x.SprintID)
		}
		sprintsBySprintID, err := as.Services.Sprint.GetMultiple(ps.Context, nil, ps.Logger, sprintIDsBySprintID...)
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
		page := &vsmember.List{Models: ret, SprintsBySprintID: sprintsBySprintID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "sprint", "smember")
	})
}

func SprintMemberDetail(rc *fasthttp.RequestCtx) {
	controller.Act("smember.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := smemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Member)"
		ps.Data = ret

		sprintBySprintID, _ := as.Services.Sprint.Get(ps.Context, nil, ret.SprintID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(rc, as, &vsmember.Detail{
			Model:            ret,
			SprintBySprintID: sprintBySprintID,
			UserByUserID:     userByUserID,
		}, ps, "sprint", "smember", ret.String())
	})
}

func SprintMemberCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("smember.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &smember.SprintMember{}
		ps.Title = "Create [SprintMember]"
		ps.Data = ret
		return controller.Render(rc, as, &vsmember.Edit{Model: ret, IsNew: true}, ps, "sprint", "smember", "Create")
	})
}

func SprintMemberCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("smember.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := smember.Random()
		ps.Title = "Create Random SprintMember"
		ps.Data = ret
		return controller.Render(rc, as, &vsmember.Edit{Model: ret, IsNew: true}, ps, "sprint", "smember", "Create")
	})
}

func SprintMemberCreate(rc *fasthttp.RequestCtx) {
	controller.Act("smember.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := smemberFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse SprintMember from form")
		}
		err = as.Services.SprintMember.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created SprintMember")
		}
		msg := fmt.Sprintf("SprintMember [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func SprintMemberEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("smember.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := smemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vsmember.Edit{Model: ret}, ps, "sprint", "smember", ret.String())
	})
}

func SprintMemberEdit(rc *fasthttp.RequestCtx) {
	controller.Act("smember.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := smemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := smemberFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse SprintMember from form")
		}
		frm.SprintID = ret.SprintID
		frm.UserID = ret.UserID
		err = as.Services.SprintMember.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update SprintMember [%s]", frm.String())
		}
		msg := fmt.Sprintf("SprintMember [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func SprintMemberDelete(rc *fasthttp.RequestCtx) {
	controller.Act("smember.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := smemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.SprintMember.Delete(ps.Context, nil, ret.SprintID, ret.UserID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete member [%s]", ret.String())
		}
		msg := fmt.Sprintf("SprintMember [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/sprintMember", rc, ps)
	})
}

func smemberFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*smember.SprintMember, error) {
	sprintIDArgStr, err := cutil.RCRequiredString(rc, "sprintID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [sprintID] as an argument")
	}
	sprintIDArgP := util.UUIDFromString(sprintIDArgStr)
	if sprintIDArgP == nil {
		return nil, errors.Errorf("argument [sprintID] (%s) is not a valid UUID", sprintIDArgStr)
	}
	sprintIDArg := *sprintIDArgP
	userIDArgStr, err := cutil.RCRequiredString(rc, "userID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [userID] as an argument")
	}
	userIDArgP := util.UUIDFromString(userIDArgStr)
	if userIDArgP == nil {
		return nil, errors.Errorf("argument [userID] (%s) is not a valid UUID", userIDArgStr)
	}
	userIDArg := *userIDArgP
	return as.Services.SprintMember.Get(ps.Context, nil, sprintIDArg, userIDArg, ps.Logger)
}

func smemberFromForm(rc *fasthttp.RequestCtx, setPK bool) (*smember.SprintMember, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return smember.FromMap(frm, setPK)
}
