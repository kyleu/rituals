// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vcomment"
)

func CommentList(rc *fasthttp.RequestCtx) {
	Act("comment.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("comment", nil, ps.Logger).Sanitize("comment")
		ret, err := as.Services.Comment.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Comments"
		ps.Data = ret
		userIDsByUserID := lo.Map(ret, func(x *comment.Comment, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vcomment.List{Models: ret, UsersByUserID: usersByUserID, Params: ps.Params}
		return Render(rc, as, page, ps, "comment")
	})
}

func CommentDetail(rc *fasthttp.RequestCtx) {
	Act("comment.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := commentFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Comment)"
		ps.Data = ret

		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return Render(rc, as, &vcomment.Detail{Model: ret, UserByUserID: userByUserID}, ps, "comment", ret.String())
	})
}

func CommentCreateForm(rc *fasthttp.RequestCtx) {
	Act("comment.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &comment.Comment{}
		ps.Title = "Create [Comment]"
		ps.Data = ret
		return Render(rc, as, &vcomment.Edit{Model: ret, IsNew: true}, ps, "comment", "Create")
	})
}

func CommentCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("comment.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := comment.Random()
		ps.Title = "Create Random Comment"
		ps.Data = ret
		return Render(rc, as, &vcomment.Edit{Model: ret, IsNew: true}, ps, "comment", "Create")
	})
}

func CommentCreate(rc *fasthttp.RequestCtx) {
	Act("comment.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := commentFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Comment from form")
		}
		err = as.Services.Comment.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Comment")
		}
		msg := fmt.Sprintf("Comment [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func CommentEditForm(rc *fasthttp.RequestCtx) {
	Act("comment.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := commentFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vcomment.Edit{Model: ret}, ps, "comment", ret.String())
	})
}

func CommentEdit(rc *fasthttp.RequestCtx) {
	Act("comment.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := commentFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := commentFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Comment from form")
		}
		frm.ID = ret.ID
		err = as.Services.Comment.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Comment [%s]", frm.String())
		}
		msg := fmt.Sprintf("Comment [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func CommentDelete(rc *fasthttp.RequestCtx) {
	Act("comment.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := commentFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Comment.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete comment [%s]", ret.String())
		}
		msg := fmt.Sprintf("Comment [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/admin/db/comment", rc, ps)
	})
}

func commentFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*comment.Comment, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
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

func commentFromForm(rc *fasthttp.RequestCtx, setPK bool) (*comment.Comment, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return comment.FromMap(frm, setPK)
}
