// Package cstory - Content managed by Project Forge, see [projectforge.md] for details.
package cstory

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vestimate/vstory/vvote"
)

func VoteList(w http.ResponseWriter, r *http.Request) {
	controller.Act("vote.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Sanitized("vote", ps.Logger)
		ret, err := as.Services.Vote.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Votes", ret)
		storyIDsByStoryID := lo.Map(ret, func(x *vote.Vote, _ int) uuid.UUID {
			return x.StoryID
		})
		storiesByStoryID, err := as.Services.Story.GetMultiple(ps.Context, nil, nil, ps.Logger, storyIDsByStoryID...)
		if err != nil {
			return "", err
		}
		userIDsByUserID := lo.Map(ret, func(x *vote.Vote, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vvote.List{Models: ret, StoriesByStoryID: storiesByStoryID, UsersByUserID: usersByUserID, Params: ps.Params}
		return controller.Render(w, r, as, page, ps, "estimate", "story", "vote")
	})
}

func VoteDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("vote.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := voteFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Vote)", ret)

		storyByStoryID, _ := as.Services.Story.Get(ps.Context, nil, ret.StoryID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		return controller.Render(w, r, as, &vvote.Detail{
			Model:          ret,
			StoryByStoryID: storyByStoryID,
			UserByUserID:   userByUserID,
		}, ps, "estimate", "story", "vote", ret.TitleString()+"**vote-yea")
	})
}

func VoteCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("vote.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &vote.Vote{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = vote.Random()
			randomStory, err := as.Services.Story.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomStory != nil {
				ret.StoryID = randomStory.ID
			}
			randomUser, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomUser != nil {
				ret.UserID = randomUser.ID
			}
		}
		ps.SetTitleAndData("Create [Vote]", ret)
		ps.Data = ret
		return controller.Render(w, r, as, &vvote.Edit{Model: ret, IsNew: true}, ps, "estimate", "story", "vote", "Create")
	})
}

func VoteRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("vote.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Vote.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Vote")
		}
		return ret.WebPath(), nil
	})
}

func VoteCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("vote.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := voteFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Vote from form")
		}
		err = as.Services.Vote.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Vote")
		}
		msg := fmt.Sprintf("Vote [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), w, ps)
	})
}

func VoteEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("vote.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := voteFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(w, r, as, &vvote.Edit{Model: ret}, ps, "estimate", "story", "vote", ret.String())
	})
}

func VoteEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("vote.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := voteFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := voteFromForm(r, ps.RequestBody, false)
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
		return controller.FlashAndRedir(true, msg, frm.WebPath(), w, ps)
	})
}

func VoteDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("vote.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := voteFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Vote.Delete(ps.Context, nil, ret.StoryID, ret.UserID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete vote [%s]", ret.String())
		}
		msg := fmt.Sprintf("Vote [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/estimate/story/vote", w, ps)
	})
}

func voteFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*vote.Vote, error) {
	storyIDArgStr, err := cutil.PathString(r, "storyID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [storyID] as an argument")
	}
	storyIDArgP := util.UUIDFromString(storyIDArgStr)
	if storyIDArgP == nil {
		return nil, errors.Errorf("argument [storyID] (%s) is not a valid UUID", storyIDArgStr)
	}
	storyIDArg := *storyIDArgP
	userIDArgStr, err := cutil.PathString(r, "userID", false)
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

func voteFromForm(r *http.Request, b []byte, setPK bool) (*vote.Vote, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return vote.FromMap(frm, setPK)
}
