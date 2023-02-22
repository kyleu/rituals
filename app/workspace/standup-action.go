package workspace

import (
	"time"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionStandup(p *Params) (*FullStandup, string, string, error) {
	lp := NewLoadParams(p.Ctx, p.Slug, p.Profile, p.Accounts, nil, nil, p.Logger)
	fu, err := p.Svc.LoadStandup(lp, func() (team.Teams, error) {
		return p.Svc.t.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	}, func() (sprint.Sprints, error) {
		return p.Svc.s.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	})
	if err != nil {
		return nil, "", "", err
	}
	switch p.Act {
	case action.ActUpdate:
		return standupUpdate(p, fu)
	case action.ActChildAdd:
		return standupReportAdd(p, fu)
	case action.ActChildUpdate:
		return standupReportUpdate(p, fu)
	case action.ActChildRemove:
		return standupReportRemove(p, fu)
	case action.ActMemberUpdate:
		return standupMemberUpdate(p, fu)
	case action.ActMemberRemove:
		return standupMemberRemove(p, fu)
	case action.ActMemberSelf:
		return standupUpdateSelf(p, fu)
	case action.ActComment:
		return standupComment(p, fu)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", p.Act)
	}
}

func standupUpdate(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	tgt := fu.Standup.Clone()
	tgt.Title = p.Frm.GetStringOpt("title")
	if tgt.Title == "" {
		return nil, "", "", errors.New("title may not be empty")
	}
	tgt.Slug = p.Frm.GetStringOpt("slug")
	if tgt.Slug == "" {
		tgt.Slug = util.Slugify(tgt.Title)
	}
	tx, err := p.Svc.db.StartTransaction(p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	defer func() {
		_ = tx.Rollback()
	}()
	tgt.Slug = p.Svc.u.Slugify(p.Ctx, tgt.ID, tgt.Slug, p.Slug, p.Svc.uh, tx, p.Logger)
	tgt.Icon = p.Frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.TeamID, _ = p.Frm.GetUUID(util.KeyTeam, true)
	tgt.SprintID, _ = p.Frm.GetUUID(util.KeySprint, true)
	perms := loadPermissionsForm(p.Frm)
	modelChanged := len(fu.Standup.Diff(tgt)) > 0
	permsChanged := len(fu.Permissions.ToPermissions().Diff(perms)) > 0
	if !modelChanged && !permsChanged {
		return fu, MsgNoChangesNeeded, fu.Standup.PublicWebPath(), nil
	}
	if modelChanged {
		model, err := p.Svc.SaveStandup(p.Ctx, tgt, fu.Self.UserID, tx, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
		err = updateTeam(
			util.KeyStandup, fu.Standup.TeamID, model.TeamID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fu.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		err = updateSprint(
			util.KeyStandup, fu.Standup.SprintID, model.SprintID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fu.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		fu.Standup = model
		err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActUpdate, model, &fu.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
		err = sendTeamSprintUpdates(util.KeyStandup, model.TeamID, model.SprintID, model, &fu.Self.UserID, p.Svc, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	if permsChanged {
		if err := p.Svc.up.DeleteWhere(p.Ctx, tx, "standup_id = $1", len(fu.Permissions), p.Logger, tgt.ID); err != nil {
			return nil, "", "", err
		}
		newPerms := make(upermission.StandupPermissions, 0, len(perms))
		for _, x := range perms {
			newPerms = append(newPerms, &upermission.StandupPermission{StandupID: tgt.ID, Key: x.Key, Value: x.Value})
		}
		if err = p.Svc.up.Save(p.Ctx, tx, p.Logger, newPerms...); err != nil {
			return nil, "", "", err
		}
		err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActPermissions, perms, &fu.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, "", "", err
	}
	return fu, "Standup updated", fu.Standup.PublicWebPath(), nil
}

func standupReportAdd(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	day, _ := p.Frm.GetTime("day", false)
	if day == nil {
		return nil, "", "", errors.New("must provide [day]")
	}
	day = util.TimeTruncate(day)
	content := p.Frm.GetStringOpt("content")
	if content == "" {
		return nil, "", "", errors.New("must provide [content]")
	}
	html := util.ToHTML(content, true)
	rpt := &report.Report{
		ID: util.UUID(), StandupID: fu.Standup.ID, Day: *day, UserID: fu.Self.UserID, Content: content, HTML: html, Created: time.Now(),
	}
	err := p.Svc.rt.Create(p.Ctx, nil, p.Logger, rpt)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited report")
	}
	err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActChildAdd, rpt, &fu.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fu, "Report added", fu.Standup.PublicWebPath(), nil
}

func standupReportUpdate(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	id, _ := p.Frm.GetUUID("reportID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [id]")
	}
	curr := fu.Reports.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no report found with id [%s]", id.String())
	}
	rpt := curr.Clone()
	day, _ := p.Frm.GetTime("day", false)
	if day == nil {
		return nil, "", "", errors.New("must provide [day]")
	}
	day = util.TimeTruncate(day)
	rpt.Day = *day
	rpt.Content = p.Frm.GetStringOpt("content")
	rpt.HTML = util.ToHTML(rpt.Content, true)
	if len(curr.Diff(rpt)) == 0 {
		return fu, MsgNoChangesNeeded, fu.Standup.PublicWebPath(), nil
	}
	err := p.Svc.rt.Update(p.Ctx, nil, rpt, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited report")
	}
	err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActChildUpdate, rpt, &fu.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fu, "Report saved", fu.Standup.PublicWebPath(), nil
}

func standupReportRemove(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	id, _ := p.Frm.GetUUID("reportID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [id]")
	}
	curr := fu.Reports.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no report found with id [%s]", id.String())
	}
	err := p.Svc.rt.Delete(p.Ctx, nil, *id, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to delete report")
	}
	err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActChildRemove, id, &fu.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fu, "Report deleted", fu.Standup.PublicWebPath(), nil
}

func standupMemberUpdate(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	if fu.Self == nil {
		return nil, "", "", errors.New("you are not a member of this standup")
	}
	if fu.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this standup")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	role := p.Frm.GetStringOpt("role")
	if role == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	curr := fu.Members.Get(fu.Standup.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this standup", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := p.Svc.um.Update(p.Ctx, nil, curr, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActMemberUpdate, curr, &fu.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fu, MsgMemberUpdated, fu.Standup.PublicWebPath(), nil
}

func standupMemberRemove(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	if fu.Self == nil {
		return nil, "", "", errors.New("you are not a member of this standup")
	}
	if fu.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this standup")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	curr := fu.Members.Get(fu.Standup.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this standup", userID.String())
	}
	err := p.Svc.um.Delete(p.Ctx, nil, curr.StandupID, curr.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActMemberRemove, userID, &fu.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fu, MsgMemberRemoved, fu.Standup.PublicWebPath(), nil
}

func standupUpdateSelf(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	if fu.Self == nil {
		return nil, "", "", errors.New("you are not a member of this standup")
	}
	name := p.Frm.GetStringOpt("name")
	choice := p.Frm.GetStringOpt("choice")
	picture := p.Frm.GetStringOpt("picture")

	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	if name == fu.Self.Name && picture == fu.Self.Picture {
		return fu, MsgNoChangesNeeded, fu.Standup.PublicWebPath(), nil
	}

	fu.Self.Picture = picture
	fu.Self.Name = name
	err := p.Svc.um.Update(p.Ctx, nil, fu.Self, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == KeyGlobal {
		err = p.Svc.setName(p.Ctx, p.Profile.ID, name, picture, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	arg := util.ValueMap{"userID": fu.Self.UserID, "name": name, "role": fu.Self.Role}
	if picture != "" {
		arg["picture"] = picture
	}
	err = p.Svc.send(enum.ModelServiceStandup, fu.Standup.ID, action.ActMemberUpdate, arg, &fu.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fu, MsgProfileEdited, fu.Standup.PublicWebPath(), nil
}

func standupComment(p *Params, fu *FullStandup) (*FullStandup, string, string, error) {
	if fu.Self == nil {
		return nil, "", "", errors.New("you are not a member of this standup")
	}
	c, u, err := commentFromForm(p.Frm, fu.Self.UserID)
	if err != nil {
		return nil, "", "", err
	}
	switch c.Svc {
	case enum.ModelServiceStandup:
		if c.ModelID != fu.Standup.ID {
			return nil, "", "", errors.New("this comment refers to a different standup")
		}
	case enum.ModelServiceReport:
		if curr := fu.Reports.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to a report that isn't part of this standup")
		}
	default:
		return nil, "", "", errors.Errorf("can't comment on object of type [%s]", c.Svc)
	}
	err = p.Svc.c.Save(p.Ctx, nil, p.Logger, c)
	if err != nil {
		return nil, "", "", err
	}
	err = sendComment(enum.ModelServiceStandup, fu.Standup.ID, c, &fu.Self.UserID, fu.Standup.TeamID, fu.Standup.SprintID, p.Svc.send, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fu, MsgCommentAdded, fu.Standup.PublicWebPath() + u, nil
}
