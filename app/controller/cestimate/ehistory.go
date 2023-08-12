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
	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/views/vestimate/vehistory"
)

func EstimateHistoryList(rc *fasthttp.RequestCtx) {
	controller.Act("ehistory.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("ehistory", nil, ps.Logger).Sanitize("ehistory")
		ret, err := as.Services.EstimateHistory.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Histories"
		ps.Data = ret
		estimateIDsByEstimateID := lo.Map(ret, func(x *ehistory.EstimateHistory, _ int) uuid.UUID {
			return x.EstimateID
		})
		estimatesByEstimateID, err := as.Services.Estimate.GetMultiple(ps.Context, nil, ps.Logger, estimateIDsByEstimateID...)
		if err != nil {
			return "", err
		}
		page := &vehistory.List{Models: ret, EstimatesByEstimateID: estimatesByEstimateID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "estimate", "ehistory")
	})
}

func EstimateHistoryDetail(rc *fasthttp.RequestCtx) {
	controller.Act("ehistory.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ehistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (History)"
		ps.Data = ret

		estimateByEstimateID, _ := as.Services.Estimate.Get(ps.Context, nil, ret.EstimateID, ps.Logger)

		return controller.Render(rc, as, &vehistory.Detail{Model: ret, EstimateByEstimateID: estimateByEstimateID}, ps, "estimate", "ehistory", ret.String())
	})
}

func EstimateHistoryCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("ehistory.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &ehistory.EstimateHistory{}
		ps.Title = "Create [EstimateHistory]"
		ps.Data = ret
		return controller.Render(rc, as, &vehistory.Edit{Model: ret, IsNew: true}, ps, "estimate", "ehistory", "Create")
	})
}

func EstimateHistoryCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("ehistory.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := ehistory.Random()
		ps.Title = "Create Random EstimateHistory"
		ps.Data = ret
		return controller.Render(rc, as, &vehistory.Edit{Model: ret, IsNew: true}, ps, "estimate", "ehistory", "Create")
	})
}

func EstimateHistoryCreate(rc *fasthttp.RequestCtx) {
	controller.Act("ehistory.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ehistoryFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse EstimateHistory from form")
		}
		err = as.Services.EstimateHistory.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created EstimateHistory")
		}
		msg := fmt.Sprintf("EstimateHistory [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func EstimateHistoryEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("ehistory.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ehistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vehistory.Edit{Model: ret}, ps, "estimate", "ehistory", ret.String())
	})
}

func EstimateHistoryEdit(rc *fasthttp.RequestCtx) {
	controller.Act("ehistory.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ehistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := ehistoryFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse EstimateHistory from form")
		}
		frm.Slug = ret.Slug
		err = as.Services.EstimateHistory.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update EstimateHistory [%s]", frm.String())
		}
		msg := fmt.Sprintf("EstimateHistory [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func EstimateHistoryDelete(rc *fasthttp.RequestCtx) {
	controller.Act("ehistory.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ehistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.EstimateHistory.Delete(ps.Context, nil, ret.Slug, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete history [%s]", ret.String())
		}
		msg := fmt.Sprintf("EstimateHistory [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/estimate/history", rc, ps)
	})
}

func ehistoryFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*ehistory.EstimateHistory, error) {
	slugArg, err := cutil.RCRequiredString(rc, "slug", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [slug] as a string argument")
	}
	return as.Services.EstimateHistory.Get(ps.Context, nil, slugArg, ps.Logger)
}

func ehistoryFromForm(rc *fasthttp.RequestCtx, setPK bool) (*ehistory.EstimateHistory, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return ehistory.FromMap(frm, setPK)
}
