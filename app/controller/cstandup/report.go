// Content managed by Project Forge, see [projectforge.md] for details.
package cstandup

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vstandup/vreport"
)

func ReportList(rc *fasthttp.RequestCtx) {
	controller.Act("report.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("report", nil, ps.Logger).Sanitize("report")
		ret, err := as.Services.Report.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Reports"
		ps.Data = ret
		standupIDsByStandupID := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			standupIDsByStandupID = append(standupIDsByStandupID, x.StandupID)
		}
		standupsByStandupID, err := as.Services.Standup.GetMultiple(ps.Context, nil, ps.Logger, standupIDsByStandupID...)
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
		page := &vreport.List{Models: ret, StandupsByStandupID: standupsByStandupID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "standup", "report")
	})
}

func ReportDetail(rc *fasthttp.RequestCtx) {
	controller.Act("report.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := reportFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Report)"
		ps.Data = ret

		standupByStandupID, _ := as.Services.Standup.Get(ps.Context, nil, ret.StandupID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(rc, as, &vreport.Detail{
			Model:              ret,
			StandupByStandupID: standupByStandupID,
			UserByUserID:       userByUserID,
		}, ps, "standup", "report", ret.String())
	})
}

func ReportCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("report.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &report.Report{}
		ps.Title = "Create [Report]"
		ps.Data = ret
		return controller.Render(rc, as, &vreport.Edit{Model: ret, IsNew: true}, ps, "standup", "report", "Create")
	})
}

func ReportCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("report.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := report.Random()
		ps.Title = "Create Random Report"
		ps.Data = ret
		return controller.Render(rc, as, &vreport.Edit{Model: ret, IsNew: true}, ps, "standup", "report", "Create")
	})
}

func ReportCreate(rc *fasthttp.RequestCtx) {
	controller.Act("report.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := reportFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Report from form")
		}
		err = as.Services.Report.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Report")
		}
		msg := fmt.Sprintf("Report [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func ReportEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("report.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := reportFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vreport.Edit{Model: ret}, ps, "standup", "report", ret.String())
	})
}

func ReportEdit(rc *fasthttp.RequestCtx) {
	controller.Act("report.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := reportFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := reportFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Report from form")
		}
		frm.ID = ret.ID
		err = as.Services.Report.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Report [%s]", frm.String())
		}
		msg := fmt.Sprintf("Report [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func ReportDelete(rc *fasthttp.RequestCtx) {
	controller.Act("report.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := reportFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Report.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete report [%s]", ret.String())
		}
		msg := fmt.Sprintf("Report [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/report", rc, ps)
	})
}

func reportFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*report.Report, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Report.Get(ps.Context, nil, idArg, ps.Logger)
}

func reportFromForm(rc *fasthttp.RequestCtx, setPK bool) (*report.Report, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return report.FromMap(frm, setPK)
}
