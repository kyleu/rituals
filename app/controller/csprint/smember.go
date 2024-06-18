package csprint

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vsprint/vsmember"
)

func SprintMemberList(w http.ResponseWriter, r *http.Request) {
	controller.Act("smember.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("smember", ps.Logger)
		ret, err := as.Services.SprintMember.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Members", ret)
		sprintIDsBySprintID := lo.Map(ret, func(x *smember.SprintMember, _ int) uuid.UUID {
			return x.SprintID
		})
		sprintsBySprintID, err := as.Services.Sprint.GetMultiple(ps.Context, nil, nil, ps.Logger, sprintIDsBySprintID...)
		if err != nil {
			return "", err
		}
		userIDsByUserID := lo.Map(ret, func(x *smember.SprintMember, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vsmember.List{Models: ret, SprintsBySprintID: sprintsBySprintID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(r, as, page, ps, "sprint", "smember")
	})
}

func SprintMemberDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("smember.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := smemberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Member)", ret)

		sprintBySprintID, _ := as.Services.Sprint.Get(ps.Context, nil, ret.SprintID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(r, as, &vsmember.Detail{
			Model:            ret,
			SprintBySprintID: sprintBySprintID,
			UserByUserID:     userByUserID,
		}, ps, "sprint", "smember", ret.TitleString()+"**users")
	})
}

func SprintMemberCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("smember.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &smember.SprintMember{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = smember.Random()
			randomSprint, err := as.Services.Sprint.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomSprint != nil {
				ret.SprintID = randomSprint.ID
			}
			randomUser, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomUser != nil {
				ret.UserID = randomUser.ID
			}
		}
		ps.SetTitleAndData("Create [SprintMember]", ret)
		ps.Data = ret
		return controller.Render(r, as, &vsmember.Edit{Model: ret, IsNew: true}, ps, "sprint", "smember", "Create")
	})
}

func SprintMemberRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("smember.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.SprintMember.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random SprintMember")
		}
		return ret.WebPath(), nil
	})
}

func SprintMemberCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("smember.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := smemberFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse SprintMember from form")
		}
		err = as.Services.SprintMember.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created SprintMember")
		}
		msg := fmt.Sprintf("SprintMember [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func SprintMemberEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("smember.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := smemberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &vsmember.Edit{Model: ret}, ps, "sprint", "smember", ret.String())
	})
}

func SprintMemberEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("smember.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := smemberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := smemberFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse SprintMember from form")
		}
		frm.SprintID = ret.SprintID
		frm.UserID = ret.UserID
		err = as.Services.SprintMember.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update SprintMember [%s]", frm.String())
		}
		msg := fmt.Sprintf("SprintMember [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func SprintMemberDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("smember.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := smemberFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.SprintMember.Delete(ps.Context, nil, ret.SprintID, ret.UserID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete member [%s]", ret.String())
		}
		msg := fmt.Sprintf("SprintMember [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/sprint/member", ps)
	})
}

func smemberFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*smember.SprintMember, error) {
	sprintIDArgStr, err := cutil.PathString(r, "sprintID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [sprintID] as an argument")
	}
	sprintIDArgP := util.UUIDFromString(sprintIDArgStr)
	if sprintIDArgP == nil {
		return nil, errors.Errorf("argument [sprintID] (%s) is not a valid UUID", sprintIDArgStr)
	}
	sprintIDArg := *sprintIDArgP
	userIDArgStr, err := cutil.PathString(r, "userID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [userID] as an argument")
	}
	userIDArgP := util.UUIDFromString(userIDArgStr)
	if userIDArgP == nil {
		return nil, errors.Errorf("argument [userID] (%s) is not a valid UUID", userIDArgStr)
	}
	userIDArg := *userIDArgP
	return as.Services.SprintMember.Get(ps.Context, nil, sprintIDArg, userIDArg, ps.Logger)
}

func smemberFromForm(r *http.Request, b []byte, setPK bool) (*smember.SprintMember, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := smember.FromMap(frm, setPK)
	return ret, err
}
