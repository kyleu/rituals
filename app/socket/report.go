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
		return errors.New(fmt.Sprintf("can't read string from [%v]", param["d"]))
	}
	d, err := parseDate(dString)
	if err != nil {
		return err
	}

	content, ok := getContent(param)
	if !ok {
		return errors.New(fmt.Sprintf("can't read content from [%v]", param["content"]))
	}

	s.logger.Debug(fmt.Sprintf("adding [%s] report for [%s]", util.ToYMD(d), userID))
	report, err := s.standups.NewReport(ch.ID, *d, content, userID)
	if err != nil {
		return errors.Wrap(err, "cannot save new report")
	}
	err = sendReportUpdate(s, ch, report)
	return errors.Wrap(err, "error sending report")
}

func onEditReport(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	id := getUUIDPointer(param, util.KeyID)
	if id == nil {
		return util.IDError(util.KeyReport, "")
	}

	d, err := parseDate(param["d"].(string))
	if err != nil {
		return err
	}

	content, ok := getContent(param)
	if !ok {
		return errors.New(fmt.Sprintf("can't read content from [%v]", param["content"]))
	}

	s.logger.Debug(fmt.Sprintf("updating [%s] report for [%s]", util.ToYMD(d), userID))
	report, err := s.standups.UpdateReport(*id, *d, content, userID)
	if err != nil {
		return errors.Wrap(err, "cannot update report")
	}
	err = sendReportUpdate(s, ch, report)
	return err
}

func onRemoveReport(s *Service, ch channel, userID uuid.UUID, param string) error {
	reportID, err := uuid.FromString(param)
	if err != nil {
		return util.IDError(util.KeyReport, param)
	}

	s.logger.Debug(fmt.Sprintf("removing report [%s]", reportID))
	err = s.standups.RemoveReport(reportID, userID)
	if err != nil {
		return errors.Wrap(err, "cannot remove report")
	}
	msg := Message{Svc: util.SvcStandup.Key, Cmd: ServerCmdReportRemove, Param: reportID}
	err = s.WriteChannel(ch, &msg)
	return errors.Wrap(err, "error sending report removal notification")
}

func sendReportUpdate(s *Service, ch channel, report *standup.Report) error {
	msg := Message{Svc: util.SvcStandup.Key, Cmd: ServerCmdReportUpdate, Param: report}
	err := s.WriteChannel(ch, &msg)
	return errors.Wrap(err, "error sending report update")
}

func parseDate(s string) (*time.Time, error) {
	dString := strings.TrimSpace(s)
	if dString == "" {
		t := time.Now()
		dString = util.ToYMD(&t)
	}
	return util.FromYMD(dString)
}
