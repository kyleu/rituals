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
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vretro/vfeedback"
)

func FeedbackList(rc *fasthttp.RequestCtx) {
	controller.Act("feedback.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("feedback", nil, ps.Logger).Sanitize("feedback")
		ret, err := as.Services.Feedback.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Feedbacks"
		ps.Data = ret
		retroIDsByRetroID := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			retroIDsByRetroID = append(retroIDsByRetroID, x.RetroID)
		}
		retrosByRetroID, err := as.Services.Retro.GetMultiple(ps.Context, nil, ps.Logger, retroIDsByRetroID...)
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
		page := &vfeedback.List{Models: ret, RetrosByRetroID: retrosByRetroID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "retro", "feedback")
	})
}

func FeedbackDetail(rc *fasthttp.RequestCtx) {
	controller.Act("feedback.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := feedbackFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Feedback)"
		ps.Data = ret
		return controller.Render(rc, as, &vfeedback.Detail{Model: ret}, ps, "retro", "feedback", ret.String())
	})
}

func FeedbackCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("feedback.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &feedback.Feedback{}
		ps.Title = "Create [Feedback]"
		ps.Data = ret
		return controller.Render(rc, as, &vfeedback.Edit{Model: ret, IsNew: true}, ps, "retro", "feedback", "Create")
	})
}

func FeedbackCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("feedback.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := feedback.Random()
		ps.Title = "Create Random Feedback"
		ps.Data = ret
		return controller.Render(rc, as, &vfeedback.Edit{Model: ret, IsNew: true}, ps, "retro", "feedback", "Create")
	})
}

func FeedbackCreate(rc *fasthttp.RequestCtx) {
	controller.Act("feedback.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := feedbackFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Feedback from form")
		}
		err = as.Services.Feedback.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Feedback")
		}
		msg := fmt.Sprintf("Feedback [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func FeedbackEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("feedback.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := feedbackFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vfeedback.Edit{Model: ret}, ps, "retro", "feedback", ret.String())
	})
}

func FeedbackEdit(rc *fasthttp.RequestCtx) {
	controller.Act("feedback.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := feedbackFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := feedbackFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Feedback from form")
		}
		frm.ID = ret.ID
		err = as.Services.Feedback.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Feedback [%s]", frm.String())
		}
		msg := fmt.Sprintf("Feedback [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func FeedbackDelete(rc *fasthttp.RequestCtx) {
	controller.Act("feedback.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := feedbackFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Feedback.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete feedback [%s]", ret.String())
		}
		msg := fmt.Sprintf("Feedback [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/feedback", rc, ps)
	})
}

func feedbackFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*feedback.Feedback, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
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

func feedbackFromForm(rc *fasthttp.RequestCtx, setPK bool) (*feedback.Feedback, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return feedback.FromMap(frm, setPK)
}
