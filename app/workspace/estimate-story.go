package workspace

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/util"
)

func estimateStoryAdd(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	title := strings.TrimSpace(p.Frm.GetStringOpt("title"))
	if title == "" {
		return nil, "", "", errors.New("must provide [title]")
	}
	st := &story.Story{
		ID: util.UUID(), EstimateID: fe.Estimate.ID, Idx: len(fe.Stories), UserID: fe.Self.UserID, Title: title, Status: enum.SessionStatusNew, Created: time.Now(),
	}
	err := p.Svc.st.Create(p.Ctx, nil, p.Logger, st)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited story")
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActChildAdd, st, &fe.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Story added", st.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryUpdate(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	id, _ := p.Frm.GetUUID("storyID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [storyID]")
	}
	curr := fe.Stories.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
	}
	if curr.UserID != fe.Self.UserID && (!fe.Admin()) {
		return nil, "", "", errors.New("you do not have permission to edit this story")
	}
	st := curr.Clone()
	st.Title = strings.TrimSpace(p.Frm.GetStringOpt("title"))
	if st.Title == "" {
		return nil, "", "", errors.New("must provide [title]")
	}
	if len(curr.Diff(st)) == 0 {
		return fe, MsgNoChangesNeeded, st.PublicWebPath(fe.Estimate.Slug), nil
	}
	err := p.Svc.st.Update(p.Ctx, nil, st, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited story")
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActChildUpdate, st, &fe.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Story saved", st.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryStatus(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	id, _ := p.Frm.GetUUID("storyID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [storyID]")
	}
	curr := fe.Stories.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
	}
	statusStr := strings.TrimSpace(p.Frm.GetStringOpt("status"))
	if statusStr == "" {
		return nil, "", "", errors.New("must provide [status]")
	}
	status := enum.SessionStatus(statusStr)
	st := &story.Story{
		ID: *id, EstimateID: fe.Estimate.ID, Idx: curr.Idx, UserID: fe.Self.UserID,
		Title: curr.Title, Status: status, FinalVote: curr.FinalVote, Created: curr.Created,
	}
	fe.Stories.Replace(st)
	err := p.Svc.st.Update(p.Ctx, nil, st, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save new status for story")
	}
	param := map[string]any{"story": st}
	if statusStr == "complete" {
		if v := fe.Votes.GetByStoryIDs(st.ID); len(v) > 0 {
			param["votes"] = v
			param["results"] = v.Results()
		}
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActChildStatus, param, &fe.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Story status updated", st.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryVote(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	id, _ := p.Frm.GetUUID("storyID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [storyID]")
	}
	currStory := fe.Stories.Get(*id)
	if currStory == nil {
		return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
	}
	if typ := strings.TrimSpace(p.Frm.GetStringOpt("typ")); typ == "" {
		return estimateStoryUserVote(p, fe, currStory)
	} else {
		return estimateStoryFinalVote(p, fe, currStory)
	}
}

func estimateStoryUserVote(p *Params, fe *FullEstimate, s *story.Story) (*FullEstimate, string, string, error) {
	voteStr := strings.TrimSpace(p.Frm.GetStringOpt("vote"))
	if voteStr == "" {
		return nil, "", "", errors.New("must provide [vote]")
	}
	if !slices.Contains(fe.Estimate.Choices, voteStr) {
		return nil, "", "", errors.Errorf("vote choice [%s] is not one of the valid choices [%s]", voteStr, strings.Join(fe.Estimate.Choices, ", "))
	}
	v := fe.Votes.Get(s.ID, fe.Self.UserID)
	if v == nil {
		v = &vote.Vote{StoryID: s.ID, UserID: fe.Self.UserID, Created: time.Now()}
	}
	v.Choice = voteStr
	err := p.Svc.v.Save(p.Ctx, nil, p.Logger, v)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save vote for story")
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActVote, v, &fe.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Vote recorded", s.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryFinalVote(p *Params, fe *FullEstimate, s *story.Story) (*FullEstimate, string, string, error) {
	valueStr := strings.TrimSpace(p.Frm.GetStringOpt("value"))
	if valueStr == "" {
		return nil, "", "", errors.New("must provide [value]")
	}
	s.FinalVote = valueStr
	err := p.Svc.st.Save(p.Ctx, nil, p.Logger, s)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save final vote for story")
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActChildUpdate, s, &fe.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Final vote recorded", s.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryRemove(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	id, _ := p.Frm.GetUUID("storyID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [storyID]")
	}
	curr := fe.Stories.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
	}
	if curr.UserID != fe.Self.UserID && (!fe.Admin()) {
		return nil, "", "", errors.New("you do not have permission to remove this story")
	}
	for _, v := range fe.Votes.GetByStoryIDs(*id) {
		err := p.Svc.v.Delete(p.Ctx, nil, v.StoryID, v.UserID, p.Logger)
		if err != nil {
			return nil, "", "", errors.Wrap(err, "unable to delete vote")
		}
	}
	err := p.Svc.st.Delete(p.Ctx, nil, *id, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to delete story")
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActChildRemove, id, &fe.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Story deleted", fe.Estimate.PublicWebPath(), nil
}
