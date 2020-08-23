package socket

import (
	"strings"

	"github.com/kyleu/npn/npnconnection"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

func SendContentUpdate(s *npnconnection.Service, svc string, id *uuid.UUID) error {
	if id != nil {
		msg := npnconnection.NewMessage(svc, ServerCmdContentUpdate, nil)
		err := s.WriteChannel(npnconnection.Channel{Svc: svc, ID: *id}, msg)
		if err != nil {
			return errors.Wrap(err, "error writing "+svc+" content update message")
		}
	}
	return nil
}

func getContent(c string) string {
	c = strings.TrimSpace(c)
	if len(c) == 0 {
		c = "-no text-"
	}
	return c
}
