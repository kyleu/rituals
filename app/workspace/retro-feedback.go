package workspace

import (
	"time"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/util"
)

func retroFeedbackAdd(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	category := p.Frm.GetStringOpt("category")
	content := p.Frm.GetStringOpt("content")
	if content == "" {
		return nil, "", "", errors.New("must provide [content]")
	}
	html := util.ToHTML(content, true)
	f := &feedback.Feedback{
		ID: util.UUID(), RetroID: fr.Retro.ID, Category: category, UserID: fr.Self.UserID, Content: content, HTML: html, Created: time.Now(),
	}
	err := p.Svc.f.Create(p.Ctx, nil, p.Logger, f)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited feedback")
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActChildAdd, f, &fr.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Feedback added", fr.Retro.PublicWebPath(), nil
}

func retroFeedbackUpdate(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	id, _ := p.Frm.GetUUID("feedbackID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [feedbackID]")
	}
	curr := fr.Feedbacks.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no feedback found with id [%s]", id.String())
	}
	if curr.UserID != fr.Self.UserID && (!fr.Admin()) {
		return nil, "", "", errors.New("you do not have permission to update this feedback")
	}
	f := curr.Clone()
	f.Category = p.Frm.GetStringOpt("category")
	f.Content = p.Frm.GetStringOpt("content")
	f.HTML = util.ToHTML(f.Content, true)
	if len(curr.Diff(f)) == 0 {
		return fr, MsgNoChangesNeeded, fr.Retro.PublicWebPath(), nil
	}
	err := p.Svc.f.Update(p.Ctx, nil, f, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited feedback")
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActChildUpdate, f, &fr.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Feedback saved", fr.Retro.PublicWebPath(), nil
}

func retroFeedbackRemove(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	id, _ := p.Frm.GetUUID("feedbackID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [feedbackID]")
	}
	curr := fr.Feedbacks.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no feedback found with id [%s]", id.String())
	}
	if curr.UserID != fr.Self.UserID && (!fr.Admin()) {
		return nil, "", "", errors.New("you do not have permission to remove this feedback")
	}
	err := p.Svc.f.Delete(p.Ctx, nil, *id, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to delete feedback")
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActChildRemove, id, &fr.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Feedback deleted", fr.Retro.PublicWebPath(), nil
}
