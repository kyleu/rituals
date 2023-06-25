// Content managed by Project Forge, see [projectforge.md] for details.
package cstory

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vestimate/vstory/vvote"
)

func VoteList(rc *fasthttp.RequestCtx) {
	controller.Act("vote.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("vote", nil, ps.Logger).Sanitize("vote")
		ret, err := as.Services.Vote.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Votes"
		ps.Data = ret
		storyIDsByStoryID := lo.Map(ret, func(x *vote.Vote, _ int) uuid.UUID {
			return x.StoryID
		})
		storiesByStoryID, err := as.Services.Story.GetMultiple(ps.Context, nil, ps.Logger, storyIDsByStoryID...)
		if err != nil {
			return "", err
		}
		userIDsByUserID := lo.Map(ret, func(x *vote.Vote, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vvote.List{Models: ret, StoriesByStoryID: storiesByStoryID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "estimate", "story", "vote")
	})
}

func VoteDetail(rc *fasthttp.RequestCtx) {
	controller.Act("vote.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := voteFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Vote)"
		ps.Data = ret

		storyByStoryID, _ := as.Services.Story.Get(ps.Context, nil, ret.StoryID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(rc, as, &vvote.Detail{
			Model:          ret,
			StoryByStoryID: storyByStoryID,
			UserByUserID:   userByUserID,
		}, ps, "estimate", "story", "vote", ret.String())
	})
}

func VoteCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("vote.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &vote.Vote{}
		ps.Title = "Create [Vote]"
		ps.Data = ret
		return controller.Render(rc, as, &vvote.Edit{Model: ret, IsNew: true}, ps, "estimate", "story", "vote", "Create")
	})
}

func VoteCreateFormRandom(rc *fasthttp.RequestCtx) {
	controller.Act("vote.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := vote.Random()
		ps.Title = "Create Random Vote"
		ps.Data = ret
		return controller.Render(rc, as, &vvote.Edit{Model: ret, IsNew: true}, ps, "estimate", "story", "vote", "Create")
	})
}

func VoteCreate(rc *fasthttp.RequestCtx) {
	controller.Act("vote.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := voteFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Vote from form")
		}
		err = as.Services.Vote.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Vote")
		}
		msg := fmt.Sprintf("Vote [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func VoteEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("vote.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := voteFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return controller.Render(rc, as, &vvote.Edit{Model: ret}, ps, "estimate", "story", "vote", ret.String())
	})
}

func VoteEdit(rc *fasthttp.RequestCtx) {
	controller.Act("vote.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := voteFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := voteFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Vote from form")
		}
		frm.StoryID = ret.StoryID
		frm.UserID = ret.UserID
		err = as.Services.Vote.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Vote [%s]", frm.String())
		}
		msg := fmt.Sprintf("Vote [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func VoteDelete(rc *fasthttp.RequestCtx) {
	controller.Act("vote.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := voteFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Vote.Delete(ps.Context, nil, ret.StoryID, ret.UserID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete vote [%s]", ret.String())
		}
		msg := fmt.Sprintf("Vote [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/vote", rc, ps)
	})
}

func voteFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*vote.Vote, error) {
	storyIDArgStr, err := cutil.RCRequiredString(rc, "storyID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [storyID] as an argument")
	}
	storyIDArgP := util.UUIDFromString(storyIDArgStr)
	if storyIDArgP == nil {
		return nil, errors.Errorf("argument [storyID] (%s) is not a valid UUID", storyIDArgStr)
	}
	storyIDArg := *storyIDArgP
	userIDArgStr, err := cutil.RCRequiredString(rc, "userID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [userID] as an argument")
	}
	userIDArgP := util.UUIDFromString(userIDArgStr)
	if userIDArgP == nil {
		return nil, errors.Errorf("argument [userID] (%s) is not a valid UUID", userIDArgStr)
	}
	userIDArg := *userIDArgP
	return as.Services.Vote.Get(ps.Context, nil, storyIDArg, userIDArg, ps.Logger)
}

func voteFromForm(rc *fasthttp.RequestCtx, setPK bool) (*vote.Vote, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return vote.FromMap(frm, setPK)
}
