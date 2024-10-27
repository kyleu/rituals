package cestimate

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vestimate/vstory"
)

func StoryList(w http.ResponseWriter, r *http.Request) {
	controller.Act("story.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("story", ps.Logger)
		var ret story.Stories
		var err error
		if q == "" {
			ret, err = as.Services.Story.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Story.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 {
				return controller.FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
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
		return controller.Render(r, as, page, ps, "estimate", "story")
	})
}

func StoryDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("story.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := storyFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Story)", ret)

		estimateByEstimateID, _ := as.Services.Estimate.Get(ps.Context, nil, ret.EstimateID, ps.Logger)
		userByUserID, _ := as.Services.User.Get(ps.Context, nil, ret.UserID, ps.Logger)

		relVotesByStoryIDPrms := ps.Params.Sanitized("vote", ps.Logger)
		relVotesByStoryID, err := as.Services.Vote.GetByStoryID(ps.Context, nil, ret.ID, relVotesByStoryIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child votes")
		}
		return controller.Render(r, as, &vstory.Detail{
			Model:                ret,
			EstimateByEstimateID: estimateByEstimateID,
			UserByUserID:         userByUserID,
			Params:               ps.Params,

			RelVotesByStoryID: relVotesByStoryID,
		}, ps, "estimate", "story", ret.TitleString()+"**story")
	})
}

func StoryCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("story.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &story.Story{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = story.RandomStory()
			randomEstimate, err := as.Services.Estimate.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomEstimate != nil {
				ret.EstimateID = randomEstimate.ID
			}
			randomUser, err := as.Services.User.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomUser != nil {
				ret.UserID = randomUser.ID
			}
		}
		ps.SetTitleAndData("Create [Story]", ret)
		ps.Data = ret
		return controller.Render(r, as, &vstory.Edit{Model: ret, IsNew: true}, ps, "estimate", "story", "Create")
	})
}

func StoryRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("story.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Story.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Story")
		}
		return ret.WebPath(), nil
	})
}

func StoryCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("story.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := storyFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Story from form")
		}
		err = as.Services.Story.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Story")
		}
		msg := fmt.Sprintf("Story [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func StoryEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("story.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := storyFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &vstory.Edit{Model: ret}, ps, "estimate", "story", ret.String())
	})
}

func StoryEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("story.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := storyFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := storyFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Story from form")
		}
		frm.ID = ret.ID
		err = as.Services.Story.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Story [%s]", frm.String())
		}
		msg := fmt.Sprintf("Story [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func StoryDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("story.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := storyFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Story.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete story [%s]", ret.String())
		}
		msg := fmt.Sprintf("Story [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/estimate/story", ps)
	})
}

func storyFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*story.Story, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
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

func storyFromForm(r *http.Request, b []byte, setPK bool) (*story.Story, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := story.StoryFromMap(frm, setPK)
	return ret, err
}
