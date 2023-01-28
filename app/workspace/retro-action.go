package workspace

import (
	"time"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionRetro(p *Params) (*FullRetro, string, string, error) {
	lp := NewLoadParams(p.Ctx, p.Slug, p.UserID, "", nil, nil, p.Logger)
	fr, err := p.Svc.LoadRetro(lp)
	if err != nil {
		return nil, "", "", err
	}

	switch p.Act {
	case action.ActUpdate:
		return retroUpdate(p, fr)
	case action.ActFeedbackAdd:
		return retroFeedbackAdd(p, fr)
	case action.ActFeedbackUpdate:
		return retroFeedbackUpdate(p, fr)
	case action.ActFeedbackRemove:
		return retroFeedbackRemove(p, fr)
	case action.ActMemberUpdate:
		return retroMemberUpdate(p, fr)
	case action.ActMemberRemove:
		return retroMemberRemove(p, fr)
	case action.ActMemberSelf:
		return retroUpdateSelf(p, fr)
	case action.ActComment:
		return retroComment(p, fr)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", p.Act)
	}
}

func retroUpdate(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	tgt := fr.Retro.Clone()
	tgt.Title = p.Frm.GetStringOpt("title")
	tgt.Slug = p.Frm.GetStringOpt("slug")
	if tgt.Slug == "" {
		tgt.Slug = util.Slugify(tgt.Title)
	}
	tgt.Slug = p.Svc.r.Slugify(p.Ctx, tgt.ID, tgt.Slug, p.Slug, p.Svc.rh, nil, p.Logger)
	tgt.Icon = p.Frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.Categories = util.StringSplitAndTrim(p.Frm.GetStringOpt("categories"), ",")
	tgt.TeamID, _ = p.Frm.GetUUID(util.KeyTeam, true)
	tgt.SprintID, _ = p.Frm.GetUUID(util.KeySprint, true)
	model, err := p.Svc.SaveRetro(p.Ctx, tgt, fr.Self.UserID, nil, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	fr.Retro = model
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActUpdate, model, &fr.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Retro saved", model.PublicWebPath(), nil
}

func retroFeedbackAdd(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	category := p.Frm.GetStringOpt("category")
	content := p.Frm.GetStringOpt("content")
	html := util.ToHTML(content, true)
	f := &feedback.Feedback{
		ID: util.UUID(), RetroID: fr.Retro.ID, Category: category, UserID: fr.Self.UserID, Content: content, HTML: html, Created: time.Now(),
	}
	err := p.Svc.f.Create(p.Ctx, nil, p.Logger, f)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited feedback")
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActFeedbackAdd, f, &fr.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Feedback added", fr.Retro.PublicWebPath(), nil
}

func retroFeedbackUpdate(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	id, _ := p.Frm.GetUUID("feedbackID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [id]")
	}
	curr := fr.Feedbacks.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no feedback found with id [%s]", id.String())
	}
	category := p.Frm.GetStringOpt("category")
	content := p.Frm.GetStringOpt("content")
	html := util.ToHTML(content, true)
	f := &feedback.Feedback{
		ID: *id, RetroID: fr.Retro.ID, Category: category, UserID: fr.Self.UserID, Content: content, HTML: html, Created: curr.Created, Updated: util.TimeToday(),
	}
	err := p.Svc.f.Update(p.Ctx, nil, f, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited feedback")
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActFeedbackUpdate, f, &fr.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Feedback saved", fr.Retro.PublicWebPath(), nil
}

func retroFeedbackRemove(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	id, _ := p.Frm.GetUUID("feedbackID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [id]")
	}
	curr := fr.Feedbacks.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no feedback found with id [%s]", id.String())
	}
	err := p.Svc.f.Delete(p.Ctx, nil, *id, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to delete feedback")
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActFeedbackRemove, id, &fr.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Feedback deleted", fr.Retro.PublicWebPath(), nil
}

func retroMemberUpdate(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	if fr.Self == nil {
		return nil, "", "", errors.New("you are not a member of this retro")
	}
	if fr.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this retro")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	role := p.Frm.GetStringOpt("role")
	if role == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	curr := fr.Members.Get(fr.Retro.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this retro", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := p.Svc.rm.Update(p.Ctx, nil, curr, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActMemberUpdate, curr, &fr.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Member updated", fr.Retro.PublicWebPath(), nil
}

func retroMemberRemove(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	if fr.Self == nil {
		return nil, "", "", errors.New("you are not a member of this retro")
	}
	if fr.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this retro")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	curr := fr.Members.Get(fr.Retro.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this retro", userID.String())
	}
	err := p.Svc.rm.Delete(p.Ctx, nil, curr.RetroID, curr.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActMemberRemove, userID, &fr.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Member removed", fr.Retro.PublicWebPath(), nil
}

func retroUpdateSelf(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	if fr.Self == nil {
		return nil, "", "", errors.New("you are not a member of this retro")
	}
	choice := p.Frm.GetStringOpt("choice")
	name := p.Frm.GetStringOpt("name")
	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	fr.Self.Name = name
	err := p.Svc.rm.Update(p.Ctx, nil, fr.Self, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == "global" {
		return nil, "", "", errors.New("can't change global name yet")
	}
	arg := util.ValueMap{"userID": fr.Self.UserID, "name": name}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActMemberUpdate, arg, &fr.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Profile edited", fr.Retro.PublicWebPath(), nil
}

func retroComment(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	if fr.Self == nil {
		return nil, "", "", errors.New("you are not a member of this retro")
	}
	c, u, err := commentFromForm(p.Frm, fr.Self.UserID)
	if err != nil {
		return nil, "", "", err
	}
	switch c.Svc {
	case enum.ModelServiceRetro:
		if c.ModelID != fr.Retro.ID {
			return nil, "", "", errors.New("this comment refers to a different retro")
		}
	case enum.ModelServiceFeedback:
		if curr := fr.Feedbacks.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to a feedback that isn't part of this retro")
		}
	default:
		return nil, "", "", errors.Errorf("can't comment on object of type [%s]", c.Svc)
	}
	err = p.Svc.c.Save(p.Ctx, nil, p.Logger, c)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActComment, c, &fr.Self.UserID, p.Logger, p.Except...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Comment added", fr.Retro.PublicWebPath() + u, nil
}
