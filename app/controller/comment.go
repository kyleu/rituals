package controller

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vcomment"
)

func CommentList(w http.ResponseWriter, r *http.Request) {
	Act("comment.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("comment", ps.Logger)
		ret, err := as.Services.Comment.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Comments", ret)
		userIDsByUserID := lo.Map(ret, func(x *comment.Comment, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vcomment.List{Models: ret, UsersByUserID: usersByUserID, Params: ps.Params}
		return Render(r, as, page, ps, "comment")
	})
}

func CommentDetail(w http.ResponseWriter, r *http.Request) {
	Act("comment.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := commentFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Comment)", ret)

		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return Render(r, as, &vcomment.Detail{Model: ret, UserByUserID: userByUserID}, ps, "comment", ret.TitleString()+"**comments")
	})
}

func CommentCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("comment.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &comment.Comment{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = comment.RandomComment()
			randomUser, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomUser != nil {
				ret.UserID = randomUser.ID
			}
		}
		ps.SetTitleAndData("Create [Comment]", ret)
		ps.Data = ret
		return Render(r, as, &vcomment.Edit{Model: ret, IsNew: true}, ps, "comment", "Create")
	})
}

func CommentRandom(w http.ResponseWriter, r *http.Request) {
	Act("comment.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Comment.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Comment")
		}
		return ret.WebPath(), nil
	})
}

func CommentCreate(w http.ResponseWriter, r *http.Request) {
	Act("comment.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := commentFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Comment from form")
		}
		err = as.Services.Comment.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Comment")
		}
		msg := fmt.Sprintf("Comment [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func CommentEditForm(w http.ResponseWriter, r *http.Request) {
	Act("comment.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := commentFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vcomment.Edit{Model: ret}, ps, "comment", ret.String())
	})
}

func CommentEdit(w http.ResponseWriter, r *http.Request) {
	Act("comment.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := commentFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := commentFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Comment from form")
		}
		frm.ID = ret.ID
		err = as.Services.Comment.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Comment [%s]", frm.String())
		}
		msg := fmt.Sprintf("Comment [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func CommentDelete(w http.ResponseWriter, r *http.Request) {
	Act("comment.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := commentFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Comment.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete comment [%s]", ret.String())
		}
		msg := fmt.Sprintf("Comment [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/admin/db/comment", ps)
	})
}

func commentFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*comment.Comment, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Comment.Get(ps.Context, nil, idArg, ps.Logger)
}

func commentFromForm(r *http.Request, b []byte, setPK bool) (*comment.Comment, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := comment.CommentFromMap(frm, setPK)
	return ret, err
}
