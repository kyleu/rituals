// Package cstandup - Content managed by Project Forge, see [projectforge.md] for details.
package cstandup

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/views/vstandup/vuhistory"
)

func StandupHistoryList(rc *fasthttp.RequestCtx) {
	controller.Act("uhistory.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("uhistory", nil, ps.Logger).Sanitize("uhistory")
		ret, err := as.Services.StandupHistory.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Histories", ret)
		standupIDsByStandupID := lo.Map(ret, func(x *uhistory.StandupHistory, _ int) uuid.UUID {
			return x.StandupID
		})
		standupsByStandupID, err := as.Services.Standup.GetMultiple(ps.Context, nil, nil, ps.Logger, standupIDsByStandupID...)
		if err != nil {
			return "", err
		}
		page := &vuhistory.List{Models: ret, StandupsByStandupID: standupsByStandupID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "standup", "uhistory")
	})
}

//nolint:lll
func StandupHistoryDetail(rc *fasthttp.RequestCtx) {
	controller.Act("uhistory.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := uhistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (History)", ret)

		standupByStandupID, _ := as.Services.Standup.Get(ps.Context, nil, ret.StandupID, ps.Logger)

		return controller.Render(rc, as, &vuhistory.Detail{Model: ret, StandupByStandupID: standupByStandupID}, ps, "standup", "uhistory", ret.TitleString()+"**history")
	})
}

func StandupHistoryCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("uhistory.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &uhistory.StandupHistory{}
		if string(rc.QueryArgs().Peek("prototype")) == "random" {
			ret = uhistory.Random()
		}
		ps.SetTitleAndData("Create [StandupHistory]", ret)
		ps.Data = ret
		return controller.Render(rc, as, &vuhistory.Edit{Model: ret, IsNew: true}, ps, "standup", "uhistory", "Create")
	})
}

func StandupHistoryRandom(rc *fasthttp.RequestCtx) {
	controller.Act("uhistory.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.StandupHistory.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random StandupHistory")
		}
		return ret.WebPath(), nil
	})
}

func StandupHistoryCreate(rc *fasthttp.RequestCtx) {
	controller.Act("uhistory.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := uhistoryFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse StandupHistory from form")
		}
		err = as.Services.StandupHistory.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created StandupHistory")
		}
		msg := fmt.Sprintf("StandupHistory [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func StandupHistoryEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("uhistory.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := uhistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(rc, as, &vuhistory.Edit{Model: ret}, ps, "standup", "uhistory", ret.String())
	})
}

func StandupHistoryEdit(rc *fasthttp.RequestCtx) {
	controller.Act("uhistory.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := uhistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := uhistoryFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse StandupHistory from form")
		}
		frm.Slug = ret.Slug
		err = as.Services.StandupHistory.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update StandupHistory [%s]", frm.String())
		}
		msg := fmt.Sprintf("StandupHistory [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func StandupHistoryDelete(rc *fasthttp.RequestCtx) {
	controller.Act("uhistory.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := uhistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.StandupHistory.Delete(ps.Context, nil, ret.Slug, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete history [%s]", ret.String())
		}
		msg := fmt.Sprintf("StandupHistory [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/standup/history", rc, ps)
	})
}

func uhistoryFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*uhistory.StandupHistory, error) {
	slugArg, err := cutil.RCRequiredString(rc, "slug", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [slug] as a string argument")
	}
	return as.Services.StandupHistory.Get(ps.Context, nil, slugArg, ps.Logger)
}

func uhistoryFromForm(rc *fasthttp.RequestCtx, setPK bool) (*uhistory.StandupHistory, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return uhistory.FromMap(frm, setPK)
}
