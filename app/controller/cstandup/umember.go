// Package cstandup - Content managed by Project Forge, see [projectforge.md] for details.
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
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vstandup/vumember"
)

func StandupMemberList(w http.ResponseWriter, r *http.Request) {
	controller.Act("umember.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("umember", ps.Logger)
		ret, err := as.Services.StandupMember.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Members", ret)
		standupIDsByStandupID := lo.Map(ret, func(x *umember.StandupMember, _ int) uuid.UUID {
			return x.StandupID
		})
		standupsByStandupID, err := as.Services.Standup.GetMultiple(ps.Context, nil, nil, ps.Logger, standupIDsByStandupID...)
		if err != nil {
			return "", err
		}
		userIDsByUserID := lo.Map(ret, func(x *umember.StandupMember, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vumember.List{Models: ret, StandupsByStandupID: standupsByStandupID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(w, r, as, page, ps, "standup", "umember")
	})
}

func StandupMemberDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("umember.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := umemberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Member)", ret)

		standupByStandupID, _ := as.Services.Standup.Get(ps.Context, nil, ret.StandupID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(w, r, as, &vumember.Detail{
			Model:              ret,
			StandupByStandupID: standupByStandupID,
			UserByUserID:       userByUserID,
		}, ps, "standup", "umember", ret.TitleString()+"**users")
	})
}

func StandupMemberCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("umember.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &umember.StandupMember{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = umember.Random()
			randomStandup, err := as.Services.Standup.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomStandup != nil {
				ret.StandupID = randomStandup.ID
			}
			randomUser, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomUser != nil {
				ret.UserID = randomUser.ID
			}
		}
		ps.SetTitleAndData("Create [StandupMember]", ret)
		ps.Data = ret
		return controller.Render(w, r, as, &vumember.Edit{Model: ret, IsNew: true}, ps, "standup", "umember", "Create")
	})
}

func StandupMemberRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("umember.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.StandupMember.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random StandupMember")
		}
		return ret.WebPath(), nil
	})
}

func StandupMemberCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("umember.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := umemberFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse StandupMember from form")
		}
		err = as.Services.StandupMember.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created StandupMember")
		}
		msg := fmt.Sprintf("StandupMember [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), w, ps)
	})
}

func StandupMemberEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("umember.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := umemberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(w, r, as, &vumember.Edit{Model: ret}, ps, "standup", "umember", ret.String())
	})
}

func StandupMemberEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("umember.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := umemberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := umemberFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse StandupMember from form")
		}
		frm.StandupID = ret.StandupID
		frm.UserID = ret.UserID
		err = as.Services.StandupMember.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update StandupMember [%s]", frm.String())
		}
		msg := fmt.Sprintf("StandupMember [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), w, ps)
	})
}

func StandupMemberDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("umember.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := umemberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.StandupMember.Delete(ps.Context, nil, ret.StandupID, ret.UserID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete member [%s]", ret.String())
		}
		msg := fmt.Sprintf("StandupMember [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/standup/member", w, ps)
	})
}

func umemberFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*umember.StandupMember, error) {
	standupIDArgStr, err := cutil.PathString(r, "standupID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [standupID] as an argument")
	}
	standupIDArgP := util.UUIDFromString(standupIDArgStr)
	if standupIDArgP == nil {
		return nil, errors.Errorf("argument [standupID] (%s) is not a valid UUID", standupIDArgStr)
	}
	standupIDArg := *standupIDArgP
	userIDArgStr, err := cutil.PathString(r, "userID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [userID] as an argument")
	}
	userIDArgP := util.UUIDFromString(userIDArgStr)
	if userIDArgP == nil {
		return nil, errors.Errorf("argument [userID] (%s) is not a valid UUID", userIDArgStr)
	}
	userIDArg := *userIDArgP
	return as.Services.StandupMember.Get(ps.Context, nil, standupIDArg, userIDArg, ps.Logger)
}

func umemberFromForm(r *http.Request, b []byte, setPK bool) (*umember.StandupMember, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return umember.FromMap(frm, setPK)
}
