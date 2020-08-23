package socket

import (
	"fmt"
	"strings"
	"time"

	"github.com/kyleu/npn/npnconnection"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/util"
)

func onAddReport(s *npnconnection.Service, ch npnconnection.Channel, userID uuid.UUID, param addReportParams) error {
	d, err := parseDate(param.D)
	if err != nil {
		return err
	}
	content := getContent(param.Content)
	s.Logger.Debug(fmt.Sprintf("adding [%s] report for [%s]", npncore.ToYMD(d), userID))
	report, err := standups(s).NewReport(ch.ID, *d, content, userID)
	if err != nil {
		return errors.Wrap(err, "cannot save new report")
	}
	err = sendReportUpdate(s, ch, report)
	return errors.Wrap(err, "error sending report")
}

func onEditReport(s *npnconnection.Service, ch npnconnection.Channel, userID uuid.UUID, param editReportParams) error {
	d, err := parseDate(param.D)
	if err != nil {
		return err
	}
	content := getContent(param.Content)
	s.Logger.Debug(fmt.Sprintf("updating [%s] report for [%s]", npncore.ToYMD(d), userID))
	report, err := standups(s).UpdateReport(param.ID, *d, content, userID)
	if err != nil {
		return errors.Wrap(err, "cannot update report")
	}
	err = sendReportUpdate(s, ch, report)
	return err
}

func onRemoveReport(s *npnconnection.Service, ch npnconnection.Channel, userID uuid.UUID, reportID uuid.UUID) error {
	s.Logger.Debug(fmt.Sprintf("removing report [%s]", reportID))
	err := standups(s).RemoveReport(reportID, userID)
	if err != nil {
		return errors.Wrap(err, "cannot remove report")
	}
	err = s.WriteChannel(ch, npnconnection.NewMessage(util.SvcStandup.Key, ServerCmdReportRemove, reportID))
	return errors.Wrap(err, "error sending report removal notification")
}

func sendReportUpdate(s *npnconnection.Service, ch npnconnection.Channel, report *standup.Report) error {
	err := s.WriteChannel(ch, npnconnection.NewMessage(util.SvcStandup.Key, ServerCmdReportUpdate, report))
	return errors.Wrap(err, "error sending report update")
}

func parseDate(s string) (*time.Time, error) {
	dString := strings.TrimSpace(s)
	if len(dString) == 0 {
		t := time.Now()
		dString = npncore.ToYMD(&t)
	}
	return npncore.FromYMD(dString)
}
