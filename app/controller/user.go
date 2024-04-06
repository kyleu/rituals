// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vuser"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	Act("user.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("user", ps.Logger)
		var ret user.Users
		var err error
		if q == "" {
			ret, err = as.Services.User.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.User.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 {
				return FlashAndRedir(true, "single result found", ret[0].WebPath(), w, ps)
			}
		}
		ps.SetTitleAndData("Users", ret)
		page := &vuser.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(w, r, as, page, ps, "user")
	})
}

func UserDetail(w http.ResponseWriter, r *http.Request) {
	Act("user.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := userFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (User)", ret)

		relActionsByUserIDPrms := ps.Params.Sanitized("action", ps.Logger)
		relActionsByUserID, err := as.Services.Action.GetByUserID(ps.Context, nil, ret.ID, relActionsByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child actions")
		}
		relCommentsByUserIDPrms := ps.Params.Sanitized("comment", ps.Logger)
		relCommentsByUserID, err := as.Services.Comment.GetByUserID(ps.Context, nil, ret.ID, relCommentsByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child comments")
		}
		relEmailsByUserIDPrms := ps.Params.Sanitized("email", ps.Logger)
		relEmailsByUserID, err := as.Services.Email.GetByUserID(ps.Context, nil, ret.ID, relEmailsByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child emails")
		}
		relEstimateMembersByUserIDPrms := ps.Params.Sanitized("emember", ps.Logger)
		relEstimateMembersByUserID, err := as.Services.EstimateMember.GetByUserID(ps.Context, nil, ret.ID, relEstimateMembersByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relFeedbacksByUserIDPrms := ps.Params.Sanitized("feedback", ps.Logger)
		relFeedbacksByUserID, err := as.Services.Feedback.GetByUserID(ps.Context, nil, ret.ID, relFeedbacksByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child feedbacks")
		}
		relReportsByUserIDPrms := ps.Params.Sanitized("report", ps.Logger)
		relReportsByUserID, err := as.Services.Report.GetByUserID(ps.Context, nil, ret.ID, relReportsByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child reports")
		}
		relRetroMembersByUserIDPrms := ps.Params.Sanitized("rmember", ps.Logger)
		relRetroMembersByUserID, err := as.Services.RetroMember.GetByUserID(ps.Context, nil, ret.ID, relRetroMembersByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relSprintMembersByUserIDPrms := ps.Params.Sanitized("smember", ps.Logger)
		relSprintMembersByUserID, err := as.Services.SprintMember.GetByUserID(ps.Context, nil, ret.ID, relSprintMembersByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relStandupMembersByUserIDPrms := ps.Params.Sanitized("umember", ps.Logger)
		relStandupMembersByUserID, err := as.Services.StandupMember.GetByUserID(ps.Context, nil, ret.ID, relStandupMembersByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relStoriesByUserIDPrms := ps.Params.Sanitized("story", ps.Logger)
		relStoriesByUserID, err := as.Services.Story.GetByUserID(ps.Context, nil, ret.ID, relStoriesByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child stories")
		}
		relTeamMembersByUserIDPrms := ps.Params.Sanitized("tmember", ps.Logger)
		relTeamMembersByUserID, err := as.Services.TeamMember.GetByUserID(ps.Context, nil, ret.ID, relTeamMembersByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		relVotesByUserIDPrms := ps.Params.Sanitized("vote", ps.Logger)
		relVotesByUserID, err := as.Services.Vote.GetByUserID(ps.Context, nil, ret.ID, relVotesByUserIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child votes")
		}
		return Render(w, r, as, &vuser.Detail{
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
		}, ps, "user", ret.TitleString()+"**profile")
	})
}

func UserCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("user.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &user.User{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = user.Random()
		}
		ps.SetTitleAndData("Create [User]", ret)
		ps.Data = ret
		return Render(w, r, as, &vuser.Edit{Model: ret, IsNew: true}, ps, "user", "Create")
	})
}

func UserRandom(w http.ResponseWriter, r *http.Request) {
	Act("user.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random User")
		}
		return ret.WebPath(), nil
	})
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	Act("user.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := userFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse User from form")
		}
		err = as.Services.User.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created User")
		}
		msg := fmt.Sprintf("User [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), w, ps)
	})
}

func UserEditForm(w http.ResponseWriter, r *http.Request) {
	Act("user.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := userFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(w, r, as, &vuser.Edit{Model: ret}, ps, "user", ret.String())
	})
}

func UserEdit(w http.ResponseWriter, r *http.Request) {
	Act("user.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := userFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := userFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse User from form")
		}
		frm.ID = ret.ID
		err = as.Services.User.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update User [%s]", frm.String())
		}
		msg := fmt.Sprintf("User [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), w, ps)
	})
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	Act("user.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := userFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.User.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete user [%s]", ret.String())
		}
		msg := fmt.Sprintf("User [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/admin/db/user", w, ps)
	})
}

func userFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*user.User, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
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

func userFromForm(r *http.Request, b []byte, setPK bool) (*user.User, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return user.FromMap(frm, setPK)
}
