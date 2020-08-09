package email

import (
	"bytes"
	"fmt"
	"time"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/gen/transcripttemplates"
)

var nightlyRecipients = []string{"kyle@kyleu.com"}

func (s *Service) GetNightlyEmail(ymd string, tx *transcript.Context) (string, *transcript.EmailResponse, error) {
	rsp, err := transcript.Email.Resolve(tx.App, tx.UserID, ymd)
	if err != nil {
		return "", nil, errors.Wrap(err, "error running email transcript")
	}
	er := rsp.(transcript.EmailResponse)

	b := &bytes.Buffer{}
	transcripttemplates.PrintEmail(er, tx, b)

	return b.String(), &er, nil
}

func (s *Service) SendNightlyEmail(d *time.Time, force bool, tx *transcript.Context) error {
	ymd, skip := shouldSkip(s, d)
	if !force && skip {
		return nil
	}

	html, er, err := s.GetNightlyEmail(ymd, tx)
	if err != nil {
		return err
	}

	err = s.sendSMTP(nightlyRecipients, er.Subject(), html)
	if err != nil {
		return err
	}

	err = s.New(Email{
		ID:         "nightly-" + ymd,
		Recipients: nightlyRecipients,
		Subject:    er.Subject(),
		Data:       er,
		Plain:      "",
		HTML:       html,
		UserID:     tx.UserID,
		Status:     "ok",
		Created:    time.Time{},
	})

	return errors.Wrap(err, "can't save email")
}

func shouldSkip(s *Service, d *time.Time) (string, bool) {
	if d == nil {
		n := time.Now()
		d = &n
	}
	ymd := npncore.ToYMD(d)
	curr := s.GetByID("nightly-" + ymd)
	if curr != nil {
		s.logger.Info(fmt.Sprintf("skipping nightly email for [%v], it was sent at [%v]", ymd, curr.Created))
		return ymd, true
	}
	return ymd, false
}
