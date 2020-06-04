package email

import (
	"bytes"
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/model/transcript"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/gen/transcripttemplates"
	"time"
)

var nightlyRecipients = []string{"kyle@kyleu.com"}

func (s *Service) SendNightlyEmail(app *config.AppInfo, userID uuid.UUID, d *time.Time, force bool) error {
	ymd, skip := shouldSkip(s, d)
	if !force && skip {
		return nil
	}

	rsp, err := transcript.Email.Resolve(app, userID, ymd, "html")
	if err != nil {
		return errors.Wrap(err, "error running email transcript")
	}
	er := rsp.(transcript.EmailResponse)

	b := &bytes.Buffer{}
	transcripttemplates.TranscriptEmailHTML(er, b)

	html := b.String()

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
		UserID:     userID,
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
	ymd := util.ToYMD(d)
	curr := s.GetByID("nightly-" + ymd)
	if curr != nil {
		s.logger.Info(fmt.Sprintf("skipping nightly email for [%v], it was sent at [%v]", ymd, curr.Created))
		return ymd, true
	}
	return ymd, false
}
