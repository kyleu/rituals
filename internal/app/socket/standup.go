package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

func onStandupMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error = nil
	switch cmd {
	default:
		err = errors.New("unhandled standup command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling estimate message"))
}
