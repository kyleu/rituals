package socket

import (
	"fmt"
	"strings"
	"time"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/util"
)

func onAddReport(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	dString, ok := param["d"].(string)
	if !ok {
		return errors.WithStack(errors.New(fmt.Sprintf("can't read string from [%v]", param["d"])))
	}
	d, err := parseDate(dString)
	if err != nil {
		return errors.WithStack(err)
	}

	c, ok := param["content"].(string)
	if !ok {
		return errors.WithStack(errors.New("cannot read content"))
	}
	content := strings.TrimSpace(c)
	if len(content) == 0 {
		content = util.KeyNoText
	}

	s.logger.Debug(fmt.Sprintf("adding [%s] report for [%s]", d.Format("2006-01-02"), userID))
	report, err := s.standups.NewReport(ch.ID, *d, content, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot save new report"))
	}
	err = sendReportUpdate(s, ch, report)
	return errors.WithStack(errors.Wrap(err, "error sending report"))
}

func onEditReport(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	id := getUUIDPointer(param, util.KeyID)
	if id == nil {
		return errors.WithStack(errors.New("invalid id"))
	}

	d, err := parseDate(param["d"].(string))
	if err != nil {
		return errors.WithStack(err)
	}

	c, ok := param["content"].(string)
	if !ok {
		return errors.WithStack(errors.Wrap(err, "cannot read report content"))
	}
	content := strings.TrimSpace(c)
	if len(content) == 0 {
		content = util.KeyNoText
	}

	s.logger.Debug(fmt.Sprintf("updating [%s] report for [%s]", d.Format("2006-01-02"), userID))
	report, err := s.standups.UpdateReport(*id, *d, content, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot update report"))
	}
	err = sendReportUpdate(s, ch, report)
	return errors.WithStack(err)
}

func onRemoveReport(s *Service, ch channel, userID uuid.UUID, param string) error {
	reportID, err := uuid.FromString(param)

	if err != nil {
		return errors.New("invalid report id [" + param + "]")
	}

	s.logger.Debug(fmt.Sprintf("removing report [%s]", reportID))
	err = s.standups.RemoveReport(reportID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot remove report"))
	}
	msg := Message{Svc: util.SvcStandup.Key, Cmd: ServerCmdReportRemove, Param: reportID}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending report removal notification"))
}

func sendReportUpdate(s *Service, ch channel, report *standup.Report) error {
	msg := Message{Svc: util.SvcStandup.Key, Cmd: ServerCmdReportUpdate, Param: report}
	err := s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending report update"))
}

func parseDate(s string) (*time.Time, error) {
	dString := strings.TrimSpace(s)

	if dString == "" {
		dString = time.Now().Format("2006-01-02")
	}

	t, err := time.Parse("2006-01-02", dString)
	if err != nil {
		return nil, errors.WithStack(errors.New("invalid date [" + dString + "] (expected 2020-01-15)"))
	}
	return &t, nil
}
