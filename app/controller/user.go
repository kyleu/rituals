// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vuser"
)

func UserList(rc *fasthttp.RequestCtx) {
	Act("user.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("user", nil, ps.Logger).Sanitize("user")
		var ret user.Users
		var err error
		if q == "" {
			ret, err = as.Services.User.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.User.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.Title = "Users"
		ps.Data = ret
		page := &vuser.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(rc, as, page, ps, "user")
	})
}

func UserDetail(rc *fasthttp.RequestCtx) {
	Act("user.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := userFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (User)"
		ps.Data = ret

		relActionsByUserIDPrms := ps.Params.Get("action", nil, ps.Logger).Sanitize("action")
		relActionsByUserID, err := as.Services.Action.GetByUserID(ps.Context, nil, ret.ID, relActionsByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child actions")
		}
		relCommentsByUserIDPrms := ps.Params.Get("comment", nil, ps.Logger).Sanitize("comment")
		relCommentsByUserID, err := as.Services.Comment.GetByUserID(ps.Context, nil, ret.ID, relCommentsByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child comments")
		}
		relEmailsByUserIDPrms := ps.Params.Get("email", nil, ps.Logger).Sanitize("email")
		relEmailsByUserID, err := as.Services.Email.GetByUserID(ps.Context, nil, ret.ID, relEmailsByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child emails")
		}
		relEstimateMembersByUserIDPrms := ps.Params.Get("emember", nil, ps.Logger).Sanitize("emember")
		relEstimateMembersByUserID, err := as.Services.EstimateMember.GetByUserID(ps.Context, nil, ret.ID, relEstimateMembersByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relFeedbacksByUserIDPrms := ps.Params.Get("feedback", nil, ps.Logger).Sanitize("feedback")
		relFeedbacksByUserID, err := as.Services.Feedback.GetByUserID(ps.Context, nil, ret.ID, relFeedbacksByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child feedbacks")
		}
		relReportsByUserIDPrms := ps.Params.Get("report", nil, ps.Logger).Sanitize("report")
		relReportsByUserID, err := as.Services.Report.GetByUserID(ps.Context, nil, ret.ID, relReportsByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child reports")
		}
		relRetroMembersByUserIDPrms := ps.Params.Get("rmember", nil, ps.Logger).Sanitize("rmember")
		relRetroMembersByUserID, err := as.Services.RetroMember.GetByUserID(ps.Context, nil, ret.ID, relRetroMembersByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relSprintMembersByUserIDPrms := ps.Params.Get("smember", nil, ps.Logger).Sanitize("smember")
		relSprintMembersByUserID, err := as.Services.SprintMember.GetByUserID(ps.Context, nil, ret.ID, relSprintMembersByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relStandupMembersByUserIDPrms := ps.Params.Get("umember", nil, ps.Logger).Sanitize("umember")
		relStandupMembersByUserID, err := as.Services.StandupMember.GetByUserID(ps.Context, nil, ret.ID, relStandupMembersByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relStoriesByUserIDPrms := ps.Params.Get("story", nil, ps.Logger).Sanitize("story")
		relStoriesByUserID, err := as.Services.Story.GetByUserID(ps.Context, nil, ret.ID, relStoriesByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child stories")
		}
		relTeamMembersByUserIDPrms := ps.Params.Get("tmember", nil, ps.Logger).Sanitize("tmember")
		relTeamMembersByUserID, err := as.Services.TeamMember.GetByUserID(ps.Context, nil, ret.ID, relTeamMembersByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relVotesByUserIDPrms := ps.Params.Get("vote", nil, ps.Logger).Sanitize("vote")
		relVotesByUserID, err := as.Services.Vote.GetByUserID(ps.Context, nil, ret.ID, relVotesByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child votes")
		}
		return Render(rc, as, &vuser.Detail{
			Model:  ret,
			Params: ps.Params,

			RelActionsByUserID:         relActionsByUserID,
			RelCommentsByUserID:        relCommentsByUserID,
			RelEmailsByUserID:          relEmailsByUserID,
			RelEstimateMembersByUserID: relEstimateMembersByUserID,
			RelFeedbacksByUserID:       relFeedbacksByUserID,
			RelReportsByUserID:         relReportsByUserID,
			RelRetroMembersByUserID:    relRetroMembersByUserID,
			RelSprintMembersByUserID:   relSprintMembersByUserID,
			RelStandupMembersByUserID:  relStandupMembersByUserID,
			RelStoriesByUserID:         relStoriesByUserID,
			RelTeamMembersByUserID:     relTeamMembersByUserID,
			RelVotesByUserID:           relVotesByUserID,
		}, ps, "user", ret.String())
	})
}

func UserCreateForm(rc *fasthttp.RequestCtx) {
	Act("user.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &user.User{}
		ps.Title = "Create [User]"
		ps.Data = ret
		return Render(rc, as, &vuser.Edit{Model: ret, IsNew: true}, ps, "user", "Create")
	})
}

func UserCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("user.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := user.Random()
		ps.Title = "Create Random User"
		ps.Data = ret
		return Render(rc, as, &vuser.Edit{Model: ret, IsNew: true}, ps, "user", "Create")
	})
}

func UserCreate(rc *fasthttp.RequestCtx) {
	Act("user.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := userFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse User from form")
		}
		err = as.Services.User.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created User")
		}
		msg := fmt.Sprintf("User [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func UserEditForm(rc *fasthttp.RequestCtx) {
	Act("user.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := userFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vuser.Edit{Model: ret}, ps, "user", ret.String())
	})
}

func UserEdit(rc *fasthttp.RequestCtx) {
	Act("user.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := userFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := userFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse User from form")
		}
		frm.ID = ret.ID
		err = as.Services.User.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update User [%s]", frm.String())
		}
		msg := fmt.Sprintf("User [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func UserDelete(rc *fasthttp.RequestCtx) {
	Act("user.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := userFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.User.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete user [%s]", ret.String())
		}
		msg := fmt.Sprintf("User [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/user", rc, ps)
	})
}

func userFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*user.User, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.User.Get(ps.Context, nil, idArg, ps.Logger)
}

func userFromForm(rc *fasthttp.RequestCtx, setPK bool) (*user.User, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return user.FromMap(frm, setPK)
}
