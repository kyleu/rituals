// Content managed by Project Forge, see [projectforge.md] for details.
package csprint

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/views/vsprint/vshistory"
)

func SprintHistoryList(rc *fasthttp.RequestCtx) {
	controller.Act("shistory.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("shistory", nil, ps.Logger).Sanitize("shistory")
		ret, err := as.Services.SprintHistory.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Histories"
		ps.Data = ret
		sprintIDsBySprintID := lo.Map(ret, func(x *shistory.SprintHistory, _ int) uuid.UUID {
			return x.SprintID
		})
		sprintsBySprintID, err := as.Services.Sprint.GetMultiple(ps.Context, nil, ps.Logger, sprintIDsBySprintID...)
		if err != nil {
			return "", err
		}
		page := &vshistory.List{Models: ret, SprintsBySprintID: sprintsBySprintID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "sprint", "shistory")
	})
}

func SprintHistoryDetail(rc *fasthttp.RequestCtx) {
	controller.Act("shistory.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := shistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (History)"
		ps.Data = ret

		sprintBySprintID, _ := as.Services.Sprint.Get(ps.Context, nil, ret.SprintID, ps.Logger)

		return controller.Render(rc, as, &vshistory.Detail{Model: ret, SprintBySprintID: sprintBySprintID}, ps, "sprint", "shistory", ret.String())
	})
}

func SprintHistoryCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("shistory.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &shistory.SprintHistory{}
		ps.Title = "Create [SprintHistory]"
		ps.Data = ret
		return controller.Render(rc, as, &vshistory.Edit{Model: ret, IsNew: true}, ps, "sprint", "shistory", "Create")
	})
}

func SprintHistoryCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("shistory.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := shistory.Random()
		ps.Title = "Create Random SprintHistory"
		ps.Data = ret
		return controller.Render(rc, as, &vshistory.Edit{Model: ret, IsNew: true}, ps, "sprint", "shistory", "Create")
	})
}

func SprintHistoryCreate(rc *fasthttp.RequestCtx) {
	controller.Act("shistory.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := shistoryFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse SprintHistory from form")
		}
		err = as.Services.SprintHistory.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created SprintHistory")
		}
		msg := fmt.Sprintf("SprintHistory [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func SprintHistoryEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("shistory.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := shistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vshistory.Edit{Model: ret}, ps, "sprint", "shistory", ret.String())
	})
}

func SprintHistoryEdit(rc *fasthttp.RequestCtx) {
	controller.Act("shistory.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := shistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := shistoryFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse SprintHistory from form")
		}
		frm.Slug = ret.Slug
		err = as.Services.SprintHistory.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update SprintHistory [%s]", frm.String())
		}
		msg := fmt.Sprintf("SprintHistory [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func SprintHistoryDelete(rc *fasthttp.RequestCtx) {
	controller.Act("shistory.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := shistoryFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.SprintHistory.Delete(ps.Context, nil, ret.Slug, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete history [%s]", ret.String())
		}
		msg := fmt.Sprintf("SprintHistory [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/sprintHistory", rc, ps)
	})
}

func shistoryFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*shistory.SprintHistory, error) {
	slugArg, err := cutil.RCRequiredString(rc, "slug", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [slug] as a string argument")
	}
	return as.Services.SprintHistory.Get(ps.Context, nil, slugArg, ps.Logger)
}

func shistoryFromForm(rc *fasthttp.RequestCtx, setPK bool) (*shistory.SprintHistory, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return shistory.FromMap(frm, setPK)
}
