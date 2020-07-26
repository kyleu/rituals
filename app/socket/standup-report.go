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

func onAddReport(s *Service, ch Channel, userID uuid.UUID, param addReportParams) error {
	d, err := parseDate(param.D)
	if err != nil {
		return err
	}
	content := getContent(param.Content)
	s.Logger.Debug(fmt.Sprintf("adding [%s] report for [%s]", util.ToYMD(d), userID))
	report, err := s.standups.NewReport(ch.ID, *d, content, userID)
	if err != nil {
		return errors.Wrap(err, "cannot save new report")
	}
	err = sendReportUpdate(s, ch, report)
	return errors.Wrap(err, "error sending report")
}

func onEditReport(s *Service, ch Channel, userID uuid.UUID, param editReportParams) error {
	d, err := parseDate(param.D)
	if err != nil {
		return err
	}
	content := getContent(param.Content)
	s.Logger.Debug(fmt.Sprintf("updating [%s] report for [%s]", util.ToYMD(d), userID))
	report, err := s.standups.UpdateReport(param.ID, *d, content, userID)
	if err != nil {
		return errors.Wrap(err, "cannot update report")
	}
	err = sendReportUpdate(s, ch, report)
	return err
}

func onRemoveReport(s *Service, ch Channel, userID uuid.UUID, reportID uuid.UUID) error {
	s.Logger.Debug(fmt.Sprintf("removing report [%s]", reportID))
	err := s.standups.RemoveReport(reportID, userID)
	if err != nil {
		return errors.Wrap(err, "cannot remove report")
	}
	err = s.WriteChannel(ch, NewMessage(util.SvcStandup, ServerCmdReportRemove, reportID))
	return errors.Wrap(err, "error sending report removal notification")
}

func sendReportUpdate(s *Service, ch Channel, report *standup.Report) error {
	err := s.WriteChannel(ch, NewMessage(util.SvcStandup, ServerCmdReportUpdate, report))
	return errors.Wrap(err, "error sending report update")
}

func parseDate(s string) (*time.Time, error) {
	dString := strings.TrimSpace(s)
	if len(dString) == 0 {
		t := time.Now()
		dString = util.ToYMD(&t)
	}
	return util.FromYMD(dString)
}
