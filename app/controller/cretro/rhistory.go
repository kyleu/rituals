// Package cretro - Content managed by Project Forge, see [projectforge.md] for details.
package cretro

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vretro/vrhistory"
)

func RetroHistoryList(w http.ResponseWriter, r *http.Request) {
	controller.Act("rhistory.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("rhistory", nil, ps.Logger).Sanitize("rhistory")
		ret, err := as.Services.RetroHistory.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Histories", ret)
		retroIDsByRetroID := lo.Map(ret, func(x *rhistory.RetroHistory, _ int) uuid.UUID {
			return x.RetroID
		})
		retrosByRetroID, err := as.Services.Retro.GetMultiple(ps.Context, nil, nil, ps.Logger, retroIDsByRetroID...)
		if err != nil {
			return "", err
		}
		page := &vrhistory.List{Models: ret, RetrosByRetroID: retrosByRetroID, Params: ps.Params}
		return controller.Render(w, r, as, page, ps, "retro", "rhistory")
	})
}

func RetroHistoryDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("rhistory.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rhistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (History)", ret)

		retroByRetroID, _ := as.Services.Retro.Get(ps.Context, nil, ret.RetroID, ps.Logger)

		return controller.Render(w, r, as, &vrhistory.Detail{Model: ret, RetroByRetroID: retroByRetroID}, ps, "retro", "rhistory", ret.TitleString()+"**history")
	})
}

func RetroHistoryCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("rhistory.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &rhistory.RetroHistory{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = rhistory.Random()
			randomRetro, err := as.Services.Retro.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomRetro != nil {
				ret.RetroID = randomRetro.ID
			}
		}
		ps.SetTitleAndData("Create [RetroHistory]", ret)
		ps.Data = ret
		return controller.Render(w, r, as, &vrhistory.Edit{Model: ret, IsNew: true}, ps, "retro", "rhistory", "Create")
	})
}

func RetroHistoryRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("rhistory.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.RetroHistory.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random RetroHistory")
		}
		return ret.WebPath(), nil
	})
}

func RetroHistoryCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("rhistory.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rhistoryFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse RetroHistory from form")
		}
		err = as.Services.RetroHistory.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created RetroHistory")
		}
		msg := fmt.Sprintf("RetroHistory [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), w, ps)
	})
}

func RetroHistoryEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("rhistory.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rhistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(w, r, as, &vrhistory.Edit{Model: ret}, ps, "retro", "rhistory", ret.String())
	})
}

func RetroHistoryEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("rhistory.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rhistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := rhistoryFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse RetroHistory from form")
		}
		frm.Slug = ret.Slug
		err = as.Services.RetroHistory.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update RetroHistory [%s]", frm.String())
		}
		msg := fmt.Sprintf("RetroHistory [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), w, ps)
	})
}

func RetroHistoryDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("rhistory.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rhistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.RetroHistory.Delete(ps.Context, nil, ret.Slug, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete history [%s]", ret.String())
		}
		msg := fmt.Sprintf("RetroHistory [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/retro/history", w, ps)
	})
}

func rhistoryFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*rhistory.RetroHistory, error) {
	slugArg, err := cutil.RCRequiredString(r, "slug", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [slug] as a string argument")
	}
	return as.Services.RetroHistory.Get(ps.Context, nil, slugArg, ps.Logger)
}

func rhistoryFromForm(r *http.Request, b []byte, setPK bool) (*rhistory.RetroHistory, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return rhistory.FromMap(frm, setPK)
}
