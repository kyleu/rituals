// Content managed by Project Forge, see [projectforge.md] for details.
package cestimate

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vestimate/vemember"
)

func EstimateMemberList(rc *fasthttp.RequestCtx) {
	controller.Act("emember.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("emember", nil, ps.Logger).Sanitize("emember")
		ret, err := as.Services.EstimateMember.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Members"
		ps.Data = ret
		estimateIDsByEstimateID := lo.Map(ret, func(x *emember.EstimateMember, _ int) uuid.UUID {
			return x.EstimateID
		})
		estimatesByEstimateID, err := as.Services.Estimate.GetMultiple(ps.Context, nil, ps.Logger, estimateIDsByEstimateID...)
		if err != nil {
			return "", err
		}
		userIDsByUserID := lo.Map(ret, func(x *emember.EstimateMember, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vemember.List{Models: ret, EstimatesByEstimateID: estimatesByEstimateID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "estimate", "emember")
	})
}

func EstimateMemberDetail(rc *fasthttp.RequestCtx) {
	controller.Act("emember.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ememberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Member)"
		ps.Data = ret

		estimateByEstimateID, _ := as.Services.Estimate.Get(ps.Context, nil, ret.EstimateID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(rc, as, &vemember.Detail{
			Model:                ret,
			EstimateByEstimateID: estimateByEstimateID,
			UserByUserID:         userByUserID,
		}, ps, "estimate", "emember", ret.String())
	})
}

func EstimateMemberCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("emember.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &emember.EstimateMember{}
		ps.Title = "Create [EstimateMember]"
		ps.Data = ret
		return controller.Render(rc, as, &vemember.Edit{Model: ret, IsNew: true}, ps, "estimate", "emember", "Create")
	})
}

func EstimateMemberCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("emember.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := emember.Random()
		ps.Title = "Create Random EstimateMember"
		ps.Data = ret
		return controller.Render(rc, as, &vemember.Edit{Model: ret, IsNew: true}, ps, "estimate", "emember", "Create")
	})
}

func EstimateMemberCreate(rc *fasthttp.RequestCtx) {
	controller.Act("emember.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ememberFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse EstimateMember from form")
		}
		err = as.Services.EstimateMember.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created EstimateMember")
		}
		msg := fmt.Sprintf("EstimateMember [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func EstimateMemberEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("emember.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ememberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vemember.Edit{Model: ret}, ps, "estimate", "emember", ret.String())
	})
}

func EstimateMemberEdit(rc *fasthttp.RequestCtx) {
	controller.Act("emember.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ememberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := ememberFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse EstimateMember from form")
		}
		frm.EstimateID = ret.EstimateID
		frm.UserID = ret.UserID
		err = as.Services.EstimateMember.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update EstimateMember [%s]", frm.String())
		}
		msg := fmt.Sprintf("EstimateMember [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func EstimateMemberDelete(rc *fasthttp.RequestCtx) {
	controller.Act("emember.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ememberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.EstimateMember.Delete(ps.Context, nil, ret.EstimateID, ret.UserID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete member [%s]", ret.String())
		}
		msg := fmt.Sprintf("EstimateMember [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/estimateMember", rc, ps)
	})
}

func ememberFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*emember.EstimateMember, error) {
	estimateIDArgStr, err := cutil.RCRequiredString(rc, "estimateID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [estimateID] as an argument")
	}
	estimateIDArgP := util.UUIDFromString(estimateIDArgStr)
	if estimateIDArgP == nil {
		return nil, errors.Errorf("argument [estimateID] (%s) is not a valid UUID", estimateIDArgStr)
	}
	estimateIDArg := *estimateIDArgP
	userIDArgStr, err := cutil.RCRequiredString(rc, "userID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [userID] as an argument")
	}
	userIDArgP := util.UUIDFromString(userIDArgStr)
	if userIDArgP == nil {
		return nil, errors.Errorf("argument [userID] (%s) is not a valid UUID", userIDArgStr)
	}
	userIDArg := *userIDArgP
	return as.Services.EstimateMember.Get(ps.Context, nil, estimateIDArg, userIDArg, ps.Logger)
}

func ememberFromForm(rc *fasthttp.RequestCtx, setPK bool) (*emember.EstimateMember, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return emember.FromMap(frm, setPK)
}
