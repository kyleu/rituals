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
	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vestimate/vehistory"
)

func EstimateHistoryList(w http.ResponseWriter, r *http.Request) {
	controller.Act("ehistory.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("ehistory", ps.Logger)
		ret, err := as.Services.EstimateHistory.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Histories", ret)
		estimateIDsByEstimateID := lo.Map(ret, func(x *ehistory.EstimateHistory, _ int) uuid.UUID {
			return x.EstimateID
		})
		estimatesByEstimateID, err := as.Services.Estimate.GetMultiple(ps.Context, nil, nil, ps.Logger, estimateIDsByEstimateID...)
		if err != nil {
			return "", err
		}
		page := &vehistory.List{Models: ret, EstimatesByEstimateID: estimatesByEstimateID, Params: ps.Params}
		return controller.Render(w, r, as, page, ps, "estimate", "ehistory")
	})
}

//nolint:lll
func EstimateHistoryDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("ehistory.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ehistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (History)", ret)

		estimateByEstimateID, _ := as.Services.Estimate.Get(ps.Context, nil, ret.EstimateID, ps.Logger)

		return controller.Render(w, r, as, &vehistory.Detail{Model: ret, EstimateByEstimateID: estimateByEstimateID}, ps, "estimate", "ehistory", ret.TitleString()+"**history")
	})
}

func EstimateHistoryCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("ehistory.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &ehistory.EstimateHistory{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = ehistory.Random()
			randomEstimate, err := as.Services.Estimate.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomEstimate != nil {
				ret.EstimateID = randomEstimate.ID
			}
		}
		ps.SetTitleAndData("Create [EstimateHistory]", ret)
		ps.Data = ret
		return controller.Render(w, r, as, &vehistory.Edit{Model: ret, IsNew: true}, ps, "estimate", "ehistory", "Create")
	})
}

func EstimateHistoryRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("ehistory.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.EstimateHistory.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random EstimateHistory")
		}
		return ret.WebPath(), nil
	})
}

func EstimateHistoryCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("ehistory.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ehistoryFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse EstimateHistory from form")
		}
		err = as.Services.EstimateHistory.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created EstimateHistory")
		}
		msg := fmt.Sprintf("EstimateHistory [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), w, ps)
	})
}

func EstimateHistoryEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("ehistory.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ehistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(w, r, as, &vehistory.Edit{Model: ret}, ps, "estimate", "ehistory", ret.String())
	})
}

func EstimateHistoryEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("ehistory.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ehistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := ehistoryFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse EstimateHistory from form")
		}
		frm.Slug = ret.Slug
		err = as.Services.EstimateHistory.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update EstimateHistory [%s]", frm.String())
		}
		msg := fmt.Sprintf("EstimateHistory [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), w, ps)
	})
}

func EstimateHistoryDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("ehistory.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := ehistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.EstimateHistory.Delete(ps.Context, nil, ret.Slug, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete history [%s]", ret.String())
		}
		msg := fmt.Sprintf("EstimateHistory [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/estimate/history", w, ps)
	})
}

func ehistoryFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*ehistory.EstimateHistory, error) {
	slugArg, err := cutil.PathString(r, "slug", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [slug] as a string argument")
	}
	return as.Services.EstimateHistory.Get(ps.Context, nil, slugArg, ps.Logger)
}

func ehistoryFromForm(r *http.Request, b []byte, setPK bool) (*ehistory.EstimateHistory, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return ehistory.FromMap(frm, setPK)
}
