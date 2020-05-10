package socket

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/util"
	"strings"
	"time"
)

func onAddReport(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	dString := strings.TrimSpace(param["d"].(string))
	if dString == "" {
		dString = time.Now().Format("2006-01-02")
	}
	d, err := time.Parse("2006-01-02", dString)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "invalid date [" + dString + "]"))
	}

	content := strings.TrimSpace(param["content"].(string))
	s.logger.Debug(fmt.Sprintf("adding [%s] report for [%s]", d, userID))

	report, err := s.standups.NewReport(ch.ID, d, content, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "cannot save new story"))
	}
	err = sendReportUpdate(s, ch, report)
	return errors.WithStack(errors.Wrap(err, "error sending stories"))
}

func sendReportUpdate(s *Service, ch channel, report *standup.Report) error {
	msg := Message{Svc: util.SvcStandup, Cmd: util.ServerCmdReportUpdate, Param: report}
	err := s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending report update"))
}
