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
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/views/vretro/vrhistory"
)

func RetroHistoryList(rc *fasthttp.RequestCtx) {
	controller.Act("rhistory.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("rhistory", nil, ps.Logger).Sanitize("rhistory")
		ret, err := as.Services.RetroHistory.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Histories"
		ps.Data = ret
		retroIDsByRetroID := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			retroIDsByRetroID = append(retroIDsByRetroID, x.RetroID)
		}
		retrosByRetroID, err := as.Services.Retro.GetMultiple(ps.Context, nil, ps.Logger, retroIDsByRetroID...)
		if err != nil {
			return "", err
		}
		page := &vrhistory.List{Models: ret, RetrosByRetroID: retrosByRetroID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "retro", "rhistory")
	})
}

func RetroHistoryDetail(rc *fasthttp.RequestCtx) {
	controller.Act("rhistory.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rhistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (History)"
		ps.Data = ret
		return controller.Render(rc, as, &vrhistory.Detail{Model: ret}, ps, "retro", "rhistory", ret.String())
	})
}

func RetroHistoryCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("rhistory.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &rhistory.RetroHistory{}
		ps.Title = "Create [RetroHistory]"
		ps.Data = ret
		return controller.Render(rc, as, &vrhistory.Edit{Model: ret, IsNew: true}, ps, "retro", "rhistory", "Create")
	})
}

func RetroHistoryCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("rhistory.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := rhistory.Random()
		ps.Title = "Create Random RetroHistory"
		ps.Data = ret
		return controller.Render(rc, as, &vrhistory.Edit{Model: ret, IsNew: true}, ps, "retro", "rhistory", "Create")
	})
}

func RetroHistoryCreate(rc *fasthttp.RequestCtx) {
	controller.Act("rhistory.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rhistoryFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse RetroHistory from form")
		}
		err = as.Services.RetroHistory.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created RetroHistory")
		}
		msg := fmt.Sprintf("RetroHistory [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func RetroHistoryEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("rhistory.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rhistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vrhistory.Edit{Model: ret}, ps, "retro", "rhistory", ret.String())
	})
}

func RetroHistoryEdit(rc *fasthttp.RequestCtx) {
	controller.Act("rhistory.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rhistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := rhistoryFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse RetroHistory from form")
		}
		frm.Slug = ret.Slug
		err = as.Services.RetroHistory.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update RetroHistory [%s]", frm.String())
		}
		msg := fmt.Sprintf("RetroHistory [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func RetroHistoryDelete(rc *fasthttp.RequestCtx) {
	controller.Act("rhistory.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := rhistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.RetroHistory.Delete(ps.Context, nil, ret.Slug, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete history [%s]", ret.String())
		}
		msg := fmt.Sprintf("RetroHistory [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/retroHistory", rc, ps)
	})
}

func rhistoryFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*rhistory.RetroHistory, error) {
	slugArg, err := cutil.RCRequiredString(rc, "slug", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [slug] as an argument")
	}
	return as.Services.RetroHistory.Get(ps.Context, nil, slugArg, ps.Logger)
}

func rhistoryFromForm(rc *fasthttp.RequestCtx, setPK bool) (*rhistory.RetroHistory, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return rhistory.FromMap(frm, setPK)
}
