package workspace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionStandup(
	ctx context.Context, slug string, act action.Act, frm util.ValueMap, userID uuid.UUID, logger util.Logger,
) (*FullStandup, string, string, error) {
	fu, err := s.LoadStandup(ctx, slug, userID, "", nil, nil, logger)
	if err != nil {
		return nil, "", "", err
	}
	switch act {
	case action.ActUpdate:
		return standupUpdate(ctx, fu, userID, frm, slug, s, logger)
	case action.ActReportAdd:
		return standupReportAdd(ctx, fu, userID, frm, s, logger)
	case action.ActReportUpdate:
		return standupReportUpdate(ctx, fu, userID, frm, s, logger)
	case action.ActReportRemove:
		return standupReportRemove(ctx, fu, userID, frm, s, logger)
	case action.ActMemberUpdate:
		return standupMemberUpdate(ctx, fu, frm, s, logger)
	case action.ActMemberRemove:
		return standupMemberRemove(ctx, fu, frm, s, logger)
	case action.ActMemberSelf:
		return standupUpdateSelf(ctx, fu, frm, s, logger)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}

func standupUpdate(
	ctx context.Context, fu *FullStandup, userID uuid.UUID, frm util.ValueMap, slug string, s *Service, logger util.Logger,
) (*FullStandup, string, string, error) {
	tgt := fu.Standup.Clone()
	tgt.Title = frm.GetStringOpt("title")
	tgt.Slug = frm.GetStringOpt("slug")
	if tgt.Slug == "" {
		tgt.Slug = util.Slugify(tgt.Title)
	}
	tgt.Slug = s.u.Slugify(ctx, tgt.ID, tgt.Slug, slug, s.uh, nil, logger)
	tgt.Icon = frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.TeamID, _ = frm.GetUUID(util.KeyTeam, true)
	tgt.SprintID, _ = frm.GetUUID(util.KeySprint, true)
	model, err := s.SaveStandup(ctx, tgt, userID, nil, logger)
	if err != nil {
		return nil, "", "", err
	}
	fu.Standup = model
	return fu, "Standup saved", model.PublicWebPath(), nil
}

func standupReportAdd(
	ctx context.Context, fu *FullStandup, userID uuid.UUID, frm util.ValueMap, s *Service, logger util.Logger,
) (*FullStandup, string, string, error) {
	day, _ := frm.GetTime("day", false)
	if day == nil {
		return nil, "", "", errors.New("must provide [day]")
	}
	day = util.TimeTruncate(day)
	content := frm.GetStringOpt("content")
	html := util.ToHTML(content)
	rpt := &report.Report{
		ID: util.UUID(), StandupID: fu.Standup.ID, Day: *day, UserID: userID, Content: content, HTML: html, Created: time.Now(),
	}
	err := s.rt.Create(ctx, nil, logger, rpt)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited report")
	}
	return fu, "Report added", fu.Standup.PublicWebPath(), nil
}

func standupReportUpdate(
	ctx context.Context, fu *FullStandup, userID uuid.UUID, frm util.ValueMap, s *Service, logger util.Logger,
) (*FullStandup, string, string, error) {
	id, _ := frm.GetUUID("reportID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [id]")
	}
	curr := fu.Reports.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no report found with id [%s]", id.String())
	}
	day, _ := frm.GetTime("day", false)
	if day == nil {
		return nil, "", "", errors.New("must provide [day]")
	}
	day = util.TimeTruncate(day)
	content := frm.GetStringOpt("content")
	html := util.ToHTML(content)
	rpt := &report.Report{
		ID: *id, StandupID: fu.Standup.ID, Day: *day, UserID: userID, Content: content, HTML: html, Created: curr.Created, Updated: util.TimeToday(),
	}
	err := s.rt.Update(ctx, nil, rpt, logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited report")
	}
	return fu, "Report saved", fu.Standup.PublicWebPath(), nil
}

func standupReportRemove(
	ctx context.Context, fu *FullStandup, userID uuid.UUID, frm util.ValueMap, s *Service, logger util.Logger,
) (*FullStandup, string, string, error) {
	id, _ := frm.GetUUID("reportID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [id]")
	}
	curr := fu.Reports.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no report found with id [%s]", id.String())
	}
	err := s.rt.Delete(ctx, nil, *id, logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to delete report")
	}
	return fu, "Report deleted", fu.Standup.PublicWebPath(), nil
}

func standupMemberUpdate(ctx context.Context, fu *FullStandup, frm util.ValueMap, s *Service, logger util.Logger) (*FullStandup, string, string, error) {
	if fu.Self == nil {
		return nil, "", "", errors.New("you are not a member of this standup")
	}
	if fu.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this standup")
	}
	userID, _ := frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	role := frm.GetStringOpt("role")
	if role == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	curr := fu.Members.Get(fu.Standup.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this standup", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := s.um.Update(ctx, nil, curr, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fu, "Member updated", fu.Standup.PublicWebPath(), nil
}

func standupMemberRemove(ctx context.Context, fu *FullStandup, frm util.ValueMap, s *Service, logger util.Logger) (*FullStandup, string, string, error) {
	if fu.Self == nil {
		return nil, "", "", errors.New("you are not a member of this standup")
	}
	if fu.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this standup")
	}
	userID, _ := frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	curr := fu.Members.Get(fu.Standup.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this standup", userID.String())
	}
	err := s.um.Delete(ctx, nil, curr.StandupID, curr.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fu, "Member removed", fu.Standup.PublicWebPath(), nil
}

func standupUpdateSelf(ctx context.Context, fu *FullStandup, frm util.ValueMap, s *Service, logger util.Logger) (*FullStandup, string, string, error) {
	if fu.Self == nil {
		return nil, "", "", errors.New("you are not a member of this standup")
	}
	choice := frm.GetStringOpt("choice")
	name := frm.GetStringOpt("name")
	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	fu.Self.Name = name
	err := s.um.Update(ctx, nil, fu.Self, logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == "global" {
		return nil, "", "", errors.New("can't change global name yet")
	}
	return fu, "Profile edited", fu.Standup.PublicWebPath(), nil
}
