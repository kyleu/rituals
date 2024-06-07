// Package csprint - Content managed by Project Forge, see [projectforge.md] for details.
package csprint

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vsprint/vshistory"
)

func SprintHistoryList(w http.ResponseWriter, r *http.Request) {
	controller.Act("shistory.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("shistory", ps.Logger)
		ret, err := as.Services.SprintHistory.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Histories", ret)
		sprintIDsBySprintID := lo.Map(ret, func(x *shistory.SprintHistory, _ int) uuid.UUID {
			return x.SprintID
		})
		sprintsBySprintID, err := as.Services.Sprint.GetMultiple(ps.Context, nil, nil, ps.Logger, sprintIDsBySprintID...)
		if err != nil {
			return "", err
		}
		page := &vshistory.List{Models: ret, SprintsBySprintID: sprintsBySprintID, Params: ps.Params}
		return controller.Render(r, as, page, ps, "sprint", "shistory")
	})
}

func SprintHistoryDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("shistory.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := shistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (History)", ret)

		sprintBySprintID, _ := as.Services.Sprint.Get(ps.Context, nil, ret.SprintID, ps.Logger)

		return controller.Render(r, as, &vshistory.Detail{Model: ret, SprintBySprintID: sprintBySprintID}, ps, "sprint", "shistory", ret.TitleString()+"**history")
	})
}

func SprintHistoryCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("shistory.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &shistory.SprintHistory{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = shistory.Random()
			randomSprint, err := as.Services.Sprint.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomSprint != nil {
				ret.SprintID = randomSprint.ID
			}
		}
		ps.SetTitleAndData("Create [SprintHistory]", ret)
		ps.Data = ret
		return controller.Render(r, as, &vshistory.Edit{Model: ret, IsNew: true}, ps, "sprint", "shistory", "Create")
	})
}

func SprintHistoryRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("shistory.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.SprintHistory.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random SprintHistory")
		}
		return ret.WebPath(), nil
	})
}

func SprintHistoryCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("shistory.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := shistoryFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse SprintHistory from form")
		}
		err = as.Services.SprintHistory.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created SprintHistory")
		}
		msg := fmt.Sprintf("SprintHistory [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func SprintHistoryEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("shistory.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := shistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &vshistory.Edit{Model: ret}, ps, "sprint", "shistory", ret.String())
	})
}

func SprintHistoryEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("shistory.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := shistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := shistoryFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse SprintHistory from form")
		}
		frm.Slug = ret.Slug
		err = as.Services.SprintHistory.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update SprintHistory [%s]", frm.String())
		}
		msg := fmt.Sprintf("SprintHistory [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func SprintHistoryDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("shistory.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := shistoryFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.SprintHistory.Delete(ps.Context, nil, ret.Slug, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete history [%s]", ret.String())
		}
		msg := fmt.Sprintf("SprintHistory [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/sprint/history", ps)
	})
}

func shistoryFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*shistory.SprintHistory, error) {
	slugArg, err := cutil.PathString(r, "slug", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [slug] as a string argument")
	}
	return as.Services.SprintHistory.Get(ps.Context, nil, slugArg, ps.Logger)
}

func shistoryFromForm(r *http.Request, b []byte, setPK bool) (*shistory.SprintHistory, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := shistory.FromMap(frm, setPK)
	return ret, err
}
