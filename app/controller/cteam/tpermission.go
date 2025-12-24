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
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vteam/vtpermission"
)

func TeamPermissionList(w http.ResponseWriter, r *http.Request) {
	controller.Act("tpermission.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("tpermission", ps.Logger)
		ret, err := as.Services.TeamPermission.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Permissions", ret)
		teamIDsByTeamID := lo.Map(ret, func(x *tpermission.TeamPermission, _ int) uuid.UUID {
			return x.TeamID
		})
		teamsByTeamID, err := as.Services.Team.GetMultiple(ps.Context, nil, nil, ps.Logger, teamIDsByTeamID...)
		if err != nil {
			return "", err
		}
		page := &vtpermission.List{Models: ret, TeamsByTeamID: teamsByTeamID, Params: ps.Params}
		return controller.Render(r, as, page, ps, "team", "tpermission")
	})
}

func TeamPermissionDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("tpermission.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Permission)", ret)

		teamByTeamID, _ := as.Services.Team.Get(ps.Context, nil, ret.TeamID, ps.Logger)

		return controller.Render(r, as, &vtpermission.Detail{Model: ret, TeamByTeamID: teamByTeamID}, ps, "team", "tpermission", ret.TitleString()+"**permission")
	})
}

func TeamPermissionCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("tpermission.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &tpermission.TeamPermission{}
		if cutil.QueryStringString(r, "prototype") == util.KeyRandom {
			ret = tpermission.RandomTeamPermission()
			randomTeam, err := as.Services.Team.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomTeam != nil {
				ret.TeamID = randomTeam.ID
			}
		}
		ps.SetTitleAndData("Create [TeamPermission]", ret)
		return controller.Render(r, as, &vtpermission.Edit{Model: ret, IsNew: true}, ps, "team", "tpermission", "Create")
	})
}

func TeamPermissionRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("tpermission.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.TeamPermission.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random TeamPermission")
		}
		return ret.WebPath(), nil
	})
}

func TeamPermissionCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("tpermission.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamPermission from form")
		}
		err = as.Services.TeamPermission.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created TeamPermission")
		}
		msg := fmt.Sprintf("TeamPermission [%s] created", ret.TitleString())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func TeamPermissionEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("tpermission.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &vtpermission.Edit{Model: ret}, ps, "team", "tpermission", ret.String())
	})
}

func TeamPermissionEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("tpermission.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := tpermissionFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamPermission from form")
		}
		frm.TeamID = ret.TeamID
		frm.Key = ret.Key
		frm.Value = ret.Value
		err = as.Services.TeamPermission.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update TeamPermission [%s]", frm.String())
		}
		msg := fmt.Sprintf("TeamPermission [%s] updated", frm.TitleString())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func TeamPermissionDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("tpermission.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.TeamPermission.Delete(ps.Context, nil, ret.TeamID, ret.Key, ret.Value, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete permission [%s]", ret.String())
		}
		msg := fmt.Sprintf("TeamPermission [%s] deleted", ret.TitleString())
		return controller.FlashAndRedir(true, msg, "/admin/db/team/permission", ps)
	})
}

func tpermissionFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*tpermission.TeamPermission, error) {
	teamIDArgStr, err := cutil.PathString(r, "teamID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [teamID] as an argument")
	}
	teamIDArgP := util.UUIDFromString(teamIDArgStr)
	if teamIDArgP == nil {
		return nil, errors.Errorf("argument [teamID] (%s) is not a valid UUID", teamIDArgStr)
	}
	teamIDArg := *teamIDArgP
	keyArg, err := cutil.PathString(r, "key", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [key] as a string argument")
	}
	valueArg, err := cutil.PathString(r, "value", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [value] as a string argument")
	}
	return as.Services.TeamPermission.Get(ps.Context, nil, teamIDArg, keyArg, valueArg, ps.Logger)
}

func tpermissionFromForm(r *http.Request, b []byte, setPK bool) (*tpermission.TeamPermission, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := tpermission.TeamPermissionFromMap(frm, setPK)
	return ret, err
}
