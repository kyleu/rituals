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

//nolint:funlen,gocognit
func UserDetail(rc *fasthttp.RequestCtx) {
	Act("user.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := userFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (User)"
		ps.Data = ret
		actionPrms := ps.Params.Get("action", nil, ps.Logger).Sanitize("action")
		actionsByUserID, err := as.Services.Action.GetByUserID(ps.Context, nil, ret.ID, actionPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child actions")
		}
		commentPrms := ps.Params.Get("comment", nil, ps.Logger).Sanitize("comment")
		commentsByUserID, err := as.Services.Comment.GetByUserID(ps.Context, nil, ret.ID, commentPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child comments")
		}
		emailPrms := ps.Params.Get("email", nil, ps.Logger).Sanitize("email")
		emailsByUserID, err := as.Services.Email.GetByUserID(ps.Context, nil, ret.ID, emailPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child emails")
		}
		estimatePrms := ps.Params.Get("estimate", nil, ps.Logger).Sanitize("estimate")
		estimatesByOwner, err := as.Services.Estimate.GetByOwner(ps.Context, nil, ret.ID, estimatePrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child estimates")
		}
		estimateMemberPrms := ps.Params.Get("emember", nil, ps.Logger).Sanitize("emember")
		estimateMembersByUserID, err := as.Services.EstimateMember.GetByUserID(ps.Context, nil, ret.ID, estimateMemberPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		feedbackPrms := ps.Params.Get("feedback", nil, ps.Logger).Sanitize("feedback")
		feedbacksByUserID, err := as.Services.Feedback.GetByUserID(ps.Context, nil, ret.ID, feedbackPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child feedbacks")
		}
		reportPrms := ps.Params.Get("report", nil, ps.Logger).Sanitize("report")
		reportsByUserID, err := as.Services.Report.GetByUserID(ps.Context, nil, ret.ID, reportPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child reports")
		}
		retroPrms := ps.Params.Get("retro", nil, ps.Logger).Sanitize("retro")
		retrosByOwner, err := as.Services.Retro.GetByOwner(ps.Context, nil, ret.ID, retroPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child retros")
		}
		retroMemberPrms := ps.Params.Get("rmember", nil, ps.Logger).Sanitize("rmember")
		retroMembersByUserID, err := as.Services.RetroMember.GetByUserID(ps.Context, nil, ret.ID, retroMemberPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		sprintPrms := ps.Params.Get("sprint", nil, ps.Logger).Sanitize("sprint")
		sprintsByOwner, err := as.Services.Sprint.GetByOwner(ps.Context, nil, ret.ID, sprintPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child sprints")
		}
		sprintMemberPrms := ps.Params.Get("smember", nil, ps.Logger).Sanitize("smember")
		sprintMembersByUserID, err := as.Services.SprintMember.GetByUserID(ps.Context, nil, ret.ID, sprintMemberPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		standupPrms := ps.Params.Get("standup", nil, ps.Logger).Sanitize("standup")
		standupsByOwner, err := as.Services.Standup.GetByOwner(ps.Context, nil, ret.ID, standupPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child standups")
		}
		standupMemberPrms := ps.Params.Get("umember", nil, ps.Logger).Sanitize("umember")
		standupMembersByUserID, err := as.Services.StandupMember.GetByUserID(ps.Context, nil, ret.ID, standupMemberPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		storyPrms := ps.Params.Get("story", nil, ps.Logger).Sanitize("story")
		storiesByUserID, err := as.Services.Story.GetByUserID(ps.Context, nil, ret.ID, storyPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child stories")
		}
		teamPrms := ps.Params.Get("team", nil, ps.Logger).Sanitize("team")
		teamsByOwner, err := as.Services.Team.GetByOwner(ps.Context, nil, ret.ID, teamPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child teams")
		}
		teamMemberPrms := ps.Params.Get("tmember", nil, ps.Logger).Sanitize("tmember")
		teamMembersByUserID, err := as.Services.TeamMember.GetByUserID(ps.Context, nil, ret.ID, teamMemberPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		votePrms := ps.Params.Get("vote", nil, ps.Logger).Sanitize("vote")
		votesByUserID, err := as.Services.Vote.GetByUserID(ps.Context, nil, ret.ID, votePrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child votes")
		}
		return Render(rc, as, &vuser.Detail{
			Model:                   ret,
			Params:                  ps.Params,
			ActionsByUserID:         actionsByUserID,
			CommentsByUserID:        commentsByUserID,
			EmailsByUserID:          emailsByUserID,
			EstimatesByOwner:        estimatesByOwner,
			EstimateMembersByUserID: estimateMembersByUserID,
			FeedbacksByUserID:       feedbacksByUserID,
			ReportsByUserID:         reportsByUserID,
			RetrosByOwner:           retrosByOwner,
			RetroMembersByUserID:    retroMembersByUserID,
			SprintsByOwner:          sprintsByOwner,
			SprintMembersByUserID:   sprintMembersByUserID,
			StandupsByOwner:         standupsByOwner,
			StandupMembersByUserID:  standupMembersByUserID,
			StoriesByUserID:         storiesByUserID,
			TeamsByOwner:            teamsByOwner,
			TeamMembersByUserID:     teamMembersByUserID,
			VotesByUserID:           votesByUserID,
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
