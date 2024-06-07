// Package cestimate - Content managed by Project Forge, see [projectforge.md] for details.
package cestimate

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vestimate/vemember"
)

func EstimateMemberList(w http.ResponseWriter, r *http.Request) {
	controller.Act("emember.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("emember", ps.Logger)
		ret, err := as.Services.EstimateMember.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Members", ret)
		estimateIDsByEstimateID := lo.Map(ret, func(x *emember.EstimateMember, _ int) uuid.UUID {
			return x.EstimateID
		})
		estimatesByEstimateID, err := as.Services.Estimate.GetMultiple(ps.Context, nil, nil, ps.Logger, estimateIDsByEstimateID...)
		if err != nil {
			return "", err
		}
		userIDsByUserID := lo.Map(ret, func(x *emember.EstimateMember, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vemember.List{Models: ret, EstimatesByEstimateID: estimatesByEstimateID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(r, as, page, ps, "estimate", "emember")
	})
}

func EstimateMemberDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("emember.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ememberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Member)", ret)

		estimateByEstimateID, _ := as.Services.Estimate.Get(ps.Context, nil, ret.EstimateID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(r, as, &vemember.Detail{
			Model:                ret,
			EstimateByEstimateID: estimateByEstimateID,
			UserByUserID:         userByUserID,
		}, ps, "estimate", "emember", ret.TitleString()+"**users")
	})
}

func EstimateMemberCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("emember.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &emember.EstimateMember{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = emember.Random()
			randomEstimate, err := as.Services.Estimate.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomEstimate != nil {
				ret.EstimateID = randomEstimate.ID
			}
			randomUser, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomUser != nil {
				ret.UserID = randomUser.ID
			}
		}
		ps.SetTitleAndData("Create [EstimateMember]", ret)
		ps.Data = ret
		return controller.Render(r, as, &vemember.Edit{Model: ret, IsNew: true}, ps, "estimate", "emember", "Create")
	})
}

func EstimateMemberRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("emember.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.EstimateMember.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random EstimateMember")
		}
		return ret.WebPath(), nil
	})
}

func EstimateMemberCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("emember.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ememberFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse EstimateMember from form")
		}
		err = as.Services.EstimateMember.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created EstimateMember")
		}
		msg := fmt.Sprintf("EstimateMember [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func EstimateMemberEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("emember.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ememberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &vemember.Edit{Model: ret}, ps, "estimate", "emember", ret.String())
	})
}

func EstimateMemberEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("emember.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ememberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := ememberFromForm(r, ps.RequestBody, false)
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
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func EstimateMemberDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("emember.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ememberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.EstimateMember.Delete(ps.Context, nil, ret.EstimateID, ret.UserID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete member [%s]", ret.String())
		}
		msg := fmt.Sprintf("EstimateMember [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/estimate/member", ps)
	})
}

func ememberFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*emember.EstimateMember, error) {
	estimateIDArgStr, err := cutil.PathString(r, "estimateID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [estimateID] as an argument")
	}
	estimateIDArgP := util.UUIDFromString(estimateIDArgStr)
	if estimateIDArgP == nil {
		return nil, errors.Errorf("argument [estimateID] (%s) is not a valid UUID", estimateIDArgStr)
	}
	estimateIDArg := *estimateIDArgP
	userIDArgStr, err := cutil.PathString(r, "userID", false)
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

func ememberFromForm(r *http.Request, b []byte, setPK bool) (*emember.EstimateMember, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := emember.FromMap(frm, setPK)
	return ret, err
}
