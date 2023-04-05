// Content managed by Project Forge, see [projectforge.md] for details.
package cteam

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vteam/vtmember"
)

func TeamMemberList(rc *fasthttp.RequestCtx) {
	controller.Act("tmember.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("tmember", nil, ps.Logger).Sanitize("tmember")
		ret, err := as.Services.TeamMember.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Members"
		ps.Data = ret
		teamIDsByTeamID := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			teamIDsByTeamID = append(teamIDsByTeamID, x.TeamID)
		}
		teamsByTeamID, err := as.Services.Team.GetMultiple(ps.Context, nil, ps.Logger, teamIDsByTeamID...)
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
		page := &vtmember.List{Models: ret, TeamsByTeamID: teamsByTeamID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "team", "tmember")
	})
}

func TeamMemberDetail(rc *fasthttp.RequestCtx) {
	controller.Act("tmember.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tmemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Member)"
		ps.Data = ret

		teamByTeamID, _ := as.Services.Team.Get(ps.Context, nil, ret.TeamID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(rc, as, &vtmember.Detail{
			Model:        ret,
			TeamByTeamID: teamByTeamID,
			UserByUserID: userByUserID,
		}, ps, "team", "tmember", ret.String())
	})
}

func TeamMemberCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("tmember.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &tmember.TeamMember{}
		ps.Title = "Create [TeamMember]"
		ps.Data = ret
		return controller.Render(rc, as, &vtmember.Edit{Model: ret, IsNew: true}, ps, "team", "tmember", "Create")
	})
}

func TeamMemberCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("tmember.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := tmember.Random()
		ps.Title = "Create Random TeamMember"
		ps.Data = ret
		return controller.Render(rc, as, &vtmember.Edit{Model: ret, IsNew: true}, ps, "team", "tmember", "Create")
	})
}

func TeamMemberCreate(rc *fasthttp.RequestCtx) {
	controller.Act("tmember.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tmemberFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamMember from form")
		}
		err = as.Services.TeamMember.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created TeamMember")
		}
		msg := fmt.Sprintf("TeamMember [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func TeamMemberEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("tmember.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tmemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vtmember.Edit{Model: ret}, ps, "team", "tmember", ret.String())
	})
}

func TeamMemberEdit(rc *fasthttp.RequestCtx) {
	controller.Act("tmember.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tmemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := tmemberFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamMember from form")
		}
		frm.TeamID = ret.TeamID
		frm.UserID = ret.UserID
		err = as.Services.TeamMember.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update TeamMember [%s]", frm.String())
		}
		msg := fmt.Sprintf("TeamMember [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func TeamMemberDelete(rc *fasthttp.RequestCtx) {
	controller.Act("tmember.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tmemberFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.TeamMember.Delete(ps.Context, nil, ret.TeamID, ret.UserID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete member [%s]", ret.String())
		}
		msg := fmt.Sprintf("TeamMember [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/teamMember", rc, ps)
	})
}

func tmemberFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*tmember.TeamMember, error) {
	teamIDArgStr, err := cutil.RCRequiredString(rc, "teamID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [teamID] as an argument")
	}
	teamIDArgP := util.UUIDFromString(teamIDArgStr)
	if teamIDArgP == nil {
		return nil, errors.Errorf("argument [teamID] (%s) is not a valid UUID", teamIDArgStr)
	}
	teamIDArg := *teamIDArgP
	userIDArgStr, err := cutil.RCRequiredString(rc, "userID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [userID] as an argument")
	}
	userIDArgP := util.UUIDFromString(userIDArgStr)
	if userIDArgP == nil {
		return nil, errors.Errorf("argument [userID] (%s) is not a valid UUID", userIDArgStr)
	}
	userIDArg := *userIDArgP
	return as.Services.TeamMember.Get(ps.Context, nil, teamIDArg, userIDArg, ps.Logger)
}

func tmemberFromForm(rc *fasthttp.RequestCtx, setPK bool) (*tmember.TeamMember, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return tmember.FromMap(frm, setPK)
}
