package workspace

import (
	"time"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/util"
)

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
		return nil, "", "", errors.New("must provide [reportID]")
	}
	curr := fu.Reports.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no report found with id [%s]", id.String())
	}
	if curr.UserID != fu.Self.UserID && (!fu.Admin()) {
		return nil, "", "", errors.New("you do not have permission to update this report")
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
		return nil, "", "", errors.New("must provide [reportID]")
	}
	curr := fu.Reports.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no report found with id [%s]", id.String())
	}
	if curr.UserID != fu.Self.UserID && (!fu.Admin()) {
		return nil, "", "", errors.New("you do not have permission to remove this report")
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
