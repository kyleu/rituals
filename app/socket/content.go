package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

func (s *Service) SendContentUpdate(svc string, id *uuid.UUID) error {
	if id != nil {
		err := s.WriteChannel(channel{Svc: svc, ID: *id}, &Message{Svc: svc, Cmd: ServerCmdContentUpdate, Param: nil})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error writing "+svc+" content update message"))
		}
	}
	return nil
}
