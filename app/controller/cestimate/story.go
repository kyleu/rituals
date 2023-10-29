// Package cestimate - Content managed by Project Forge, see [projectforge.md] for details.
package cestimate

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vestimate/vstory"
)

func StoryList(rc *fasthttp.RequestCtx) {
	controller.Act("story.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("story", nil, ps.Logger).Sanitize("story")
		var ret story.Stories
		var err error
		if q == "" {
			ret, err = as.Services.Story.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Story.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Stories", ret)
		estimateIDsByEstimateID := lo.Map(ret, func(x *story.Story, _ int) uuid.UUID {
			return x.EstimateID
		})
		estimatesByEstimateID, err := as.Services.Estimate.GetMultiple(ps.Context, nil, nil, ps.Logger, estimateIDsByEstimateID...)
		if err != nil {
			return "", err
		}
		userIDsByUserID := lo.Map(ret, func(x *story.Story, _ int) uuid.UUID {
			return x.UserID
		})
		usersByUserID, err := as.Services.User.GetMultiple(ps.Context, nil, nil, ps.Logger, userIDsByUserID...)
		if err != nil {
			return "", err
		}
		page := &vstory.List{Models: ret, EstimatesByEstimateID: estimatesByEstimateID, UsersByUserID: usersByUserID, Params: ps.Params, SearchQuery: q}
		return controller.Render(rc, as, page, ps, "estimate", "story")
	})
}

func StoryDetail(rc *fasthttp.RequestCtx) {
	controller.Act("story.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := storyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Story)", ret)

		estimateByEstimateID, _ := as.Services.Estimate.Get(ps.Context, nil, ret.EstimateID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		relVotesByStoryIDPrms := ps.Params.Get("vote", nil, ps.Logger).Sanitize("vote")
		relVotesByStoryID, err := as.Services.Vote.GetByStoryID(ps.Context, nil, ret.ID, relVotesByStoryIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child votes")
		}
		return controller.Render(rc, as, &vstory.Detail{
			Model:                ret,
			EstimateByEstimateID: estimateByEstimateID,
			UserByUserID:         userByUserID,
			Params:               ps.Params,

			RelVotesByStoryID: relVotesByStoryID,
		}, ps, "estimate", "story", ret.TitleString()+"**story")
	})
}

func StoryCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("story.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &story.Story{}
		if string(rc.QueryArgs().Peek("prototype")) == util.KeyRandom {
			ret = story.Random()
		}
		ps.SetTitleAndData("Create [Story]", ret)
		ps.Data = ret
		return controller.Render(rc, as, &vstory.Edit{Model: ret, IsNew: true}, ps, "estimate", "story", "Create")
	})
}

func StoryRandom(rc *fasthttp.RequestCtx) {
	controller.Act("story.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Story.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Story")
		}
		return ret.WebPath(), nil
	})
}

func StoryCreate(rc *fasthttp.RequestCtx) {
	controller.Act("story.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := storyFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Story from form")
		}
		err = as.Services.Story.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Story")
		}
		msg := fmt.Sprintf("Story [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func StoryEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("story.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := storyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(rc, as, &vstory.Edit{Model: ret}, ps, "estimate", "story", ret.String())
	})
}

func StoryEdit(rc *fasthttp.RequestCtx) {
	controller.Act("story.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := storyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := storyFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Story from form")
		}
		frm.ID = ret.ID
		err = as.Services.Story.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Story [%s]", frm.String())
		}
		msg := fmt.Sprintf("Story [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func StoryDelete(rc *fasthttp.RequestCtx) {
	controller.Act("story.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := storyFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Story.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete story [%s]", ret.String())
		}
		msg := fmt.Sprintf("Story [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/estimate/story", rc, ps)
	})
}

func storyFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*story.Story, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Story.Get(ps.Context, nil, idArg, ps.Logger)
}

func storyFromForm(rc *fasthttp.RequestCtx, setPK bool) (*story.Story, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return story.FromMap(frm, setPK)
}
