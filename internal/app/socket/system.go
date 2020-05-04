package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

func onSystemMessage(s *Service, connID uuid.UUID, userID uuid.UUID, cmd string, param interface {}) error {
	var err error = nil
	switch cmd {
	case "member-name-save":
		err = saveName(s, userID, param.(map[string]interface {}))
	default:
		err = errors.New("Unhandled estimate command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling estimate message"))
}

func saveName(s *Service, userID uuid.UUID, o map[string]interface {}) error {
	svc := o["svc"].(string)
	id := o["id"].(string)
	name := o["name"].(string)
	choice := o["choice"].(string)
	println(svc + ": " + id + " / " + name + ": " + choice)
	return nil
}
