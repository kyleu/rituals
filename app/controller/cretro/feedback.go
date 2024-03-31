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
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vretro/vfeedback"
)

func FeedbackList(w http.ResponseWriter, r *http.Request) {
	controller.Act("feedback.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("feedback", ps.Logger)
		ret, err := as.Services.Feedback.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Feedbacks", ret)
		retroIDsByRetroID := lo.Map(ret, func(x *feedback.Feedback, _ int) uuid.UUID {
			return x.RetroID
		})
		retrosByRetroID, err := as.Services.Retro.GetMultiple(ps.Context, nil, nil, ps.Logger, retroIDsByRetroID...)
		if err != nil {
			return "", err
		}
		userIDsByUserID := lo.Map(ret, func(x *feedback.Feedback, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vfeedback.List{Models: ret, RetrosByRetroID: retrosByRetroID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(w, r, as, page, ps, "retro", "feedback")
	})
}

func FeedbackDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("feedback.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := feedbackFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Feedback)", ret)

		retroByRetroID, _ := as.Services.Retro.Get(ps.Context, nil, ret.RetroID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(w, r, as, &vfeedback.Detail{
			Model:          ret,
			RetroByRetroID: retroByRetroID,
			UserByUserID:   userByUserID,
		}, ps, "retro", "feedback", ret.TitleString()+"**comment")
	})
}

func FeedbackCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("feedback.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &feedback.Feedback{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = feedback.Random()
			randomRetro, err := as.Services.Retro.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomRetro != nil {
				ret.RetroID = randomRetro.ID
			}
			randomUser, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomUser != nil {
				ret.UserID = randomUser.ID
			}
		}
		ps.SetTitleAndData("Create [Feedback]", ret)
		ps.Data = ret
		return controller.Render(w, r, as, &vfeedback.Edit{Model: ret, IsNew: true}, ps, "retro", "feedback", "Create")
	})
}

func FeedbackRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("feedback.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Feedback.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Feedback")
		}
		return ret.WebPath(), nil
	})
}

func FeedbackCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("feedback.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := feedbackFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Feedback from form")
		}
		err = as.Services.Feedback.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Feedback")
		}
		msg := fmt.Sprintf("Feedback [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), w, ps)
	})
}

func FeedbackEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("feedback.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := feedbackFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(w, r, as, &vfeedback.Edit{Model: ret}, ps, "retro", "feedback", ret.String())
	})
}

func FeedbackEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("feedback.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := feedbackFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := feedbackFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Feedback from form")
		}
		frm.ID = ret.ID
		err = as.Services.Feedback.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Feedback [%s]", frm.String())
		}
		msg := fmt.Sprintf("Feedback [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), w, ps)
	})
}

func FeedbackDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("feedback.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := feedbackFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Feedback.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete feedback [%s]", ret.String())
		}
		msg := fmt.Sprintf("Feedback [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/retro/feedback", w, ps)
	})
}

func feedbackFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*feedback.Feedback, error) {
	idArgStr, err := cutil.RCRequiredString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Feedback.Get(ps.Context, nil, idArg, ps.Logger)
}

func feedbackFromForm(r *http.Request, b []byte, setPK bool) (*feedback.Feedback, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return feedback.FromMap(frm, setPK)
}
