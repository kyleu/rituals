package socket

import (
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) SendContentUpdate(svc util.Service, id *uuid.UUID) error {
	if id != nil {
		msg := NewMessage(svc, ServerCmdContentUpdate, nil)
		err := s.WriteChannel(channel{Svc: svc, ID: *id}, msg)
		if err != nil {
			return errors.Wrap(err, "error writing "+svc.Key+" content update message")
		}
	}
	return nil
}

func getContent(c string) string {
	c = strings.TrimSpace(c)
	if len(c) == 0 {
		c = util.KeyNoText
	}
	return c
}
