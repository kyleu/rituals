package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

func onRetroMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error
	switch cmd {
	default:
		err = errors.New("unhandled retro command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling estimate message"))
}
