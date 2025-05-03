package cstandup

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vstandup/vreport"
)

func ReportList(w http.ResponseWriter, r *http.Request) {
	controller.Act("report.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("report", ps.Logger)
		ret, err := as.Services.Report.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Reports", ret)
		standupIDsByStandupID := lo.Map(ret, func(x *report.Report, _ int) uuid.UUID {
			return x.StandupID
		})
		standupsByStandupID, err := as.Services.Standup.GetMultiple(ps.Context, nil, nil, ps.Logger, standupIDsByStandupID...)
		if err != nil {
			return "", err
		}
		userIDsByUserID := lo.Map(ret, func(x *report.Report, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vreport.List{Models: ret, StandupsByStandupID: standupsByStandupID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(r, as, page, ps, "standup", "report")
	})
}

func ReportDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("report.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := reportFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Report)", ret)

		standupByStandupID, _ := as.Services.Standup.Get(ps.Context, nil, ret.StandupID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(r, as, &vreport.Detail{
			Model:              ret,
			StandupByStandupID: standupByStandupID,
			UserByUserID:       userByUserID,
		}, ps, "standup", "report", ret.TitleString()+"**file-alt")
	})
}

func ReportCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("report.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &report.Report{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = report.RandomReport()
			randomStandup, err := as.Services.Standup.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomStandup != nil {
				ret.StandupID = randomStandup.ID
			}
			randomUser, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomUser != nil {
				ret.UserID = randomUser.ID
			}
		}
		ps.SetTitleAndData("Create [Report]", ret)
		return controller.Render(r, as, &vreport.Edit{Model: ret, IsNew: true}, ps, "standup", "report", "Create")
	})
}

func ReportRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("report.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Report.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Report")
		}
		return ret.WebPath(), nil
	})
}

func ReportCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("report.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := reportFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Report from form")
		}
		err = as.Services.Report.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Report")
		}
		msg := fmt.Sprintf("Report [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func ReportEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("report.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := reportFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &vreport.Edit{Model: ret}, ps, "standup", "report", ret.String())
	})
}

func ReportEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("report.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := reportFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := reportFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Report from form")
		}
		frm.ID = ret.ID
		err = as.Services.Report.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Report [%s]", frm.String())
		}
		msg := fmt.Sprintf("Report [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func ReportDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("report.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := reportFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Report.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete report [%s]", ret.String())
		}
		msg := fmt.Sprintf("Report [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/standup/report", ps)
	})
}

func reportFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*report.Report, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
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

func reportFromForm(r *http.Request, b []byte, setPK bool) (*report.Report, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := report.ReportFromMap(frm, setPK)
	return ret, err
}
