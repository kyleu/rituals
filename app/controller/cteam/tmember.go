package cteam

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vteam/vtmember"
)

func TeamMemberList(w http.ResponseWriter, r *http.Request) {
	controller.Act("tmember.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("tmember", ps.Logger)
		ret, err := as.Services.TeamMember.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Members", ret)
		teamIDsByTeamID := lo.Map(ret, func(x *tmember.TeamMember, _ int) uuid.UUID {
			return x.TeamID
		})
		teamsByTeamID, err := as.Services.Team.GetMultiple(ps.Context, nil, nil, ps.Logger, teamIDsByTeamID...)
		if err != nil {
			return "", err
		}
		userIDsByUserID := lo.Map(ret, func(x *tmember.TeamMember, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vtmember.List{Models: ret, TeamsByTeamID: teamsByTeamID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(r, as, page, ps, "team", "tmember")
	})
}

func TeamMemberDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("tmember.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tmemberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Member)", ret)

		teamByTeamID, _ := as.Services.Team.Get(ps.Context, nil, ret.TeamID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(r, as, &vtmember.Detail{
			Model:        ret,
			TeamByTeamID: teamByTeamID,
			UserByUserID: userByUserID,
		}, ps, "team", "tmember", ret.TitleString()+"**users")
	})
}

func TeamMemberCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("tmember.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &tmember.TeamMember{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = tmember.RandomTeamMember()
			randomTeam, err := as.Services.Team.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomTeam != nil {
				ret.TeamID = randomTeam.ID
			}
			randomUser, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomUser != nil {
				ret.UserID = randomUser.ID
			}
		}
		ps.SetTitleAndData("Create [TeamMember]", ret)
		return controller.Render(r, as, &vtmember.Edit{Model: ret, IsNew: true}, ps, "team", "tmember", "Create")
	})
}

func TeamMemberRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("tmember.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.TeamMember.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random TeamMember")
		}
		return ret.WebPath(), nil
	})
}

func TeamMemberCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("tmember.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tmemberFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamMember from form")
		}
		err = as.Services.TeamMember.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created TeamMember")
		}
		msg := fmt.Sprintf("TeamMember [%s] created", ret.TitleString())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func TeamMemberEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("tmember.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tmemberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &vtmember.Edit{Model: ret}, ps, "team", "tmember", ret.String())
	})
}

func TeamMemberEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("tmember.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tmemberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := tmemberFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamMember from form")
		}
		frm.TeamID = ret.TeamID
		frm.UserID = ret.UserID
		err = as.Services.TeamMember.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update TeamMember [%s]", frm.String())
		}
		msg := fmt.Sprintf("TeamMember [%s] updated", frm.TitleString())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func TeamMemberDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("tmember.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tmemberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.TeamMember.Delete(ps.Context, nil, ret.TeamID, ret.UserID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete member [%s]", ret.String())
		}
		msg := fmt.Sprintf("TeamMember [%s] deleted", ret.TitleString())
		return controller.FlashAndRedir(true, msg, "/admin/db/team/member", ps)
	})
}

func tmemberFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*tmember.TeamMember, error) {
	teamIDArgStr, err := cutil.PathString(r, "teamID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [teamID] as an argument")
	}
	teamIDArgP := util.UUIDFromString(teamIDArgStr)
	if teamIDArgP == nil {
		return nil, errors.Errorf("argument [teamID] (%s) is not a valid UUID", teamIDArgStr)
	}
	teamIDArg := *teamIDArgP
	userIDArgStr, err := cutil.PathString(r, "userID", false)
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

func tmemberFromForm(r *http.Request, b []byte, setPK bool) (*tmember.TeamMember, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := tmember.TeamMemberFromMap(frm, setPK)
	return ret, err
}
