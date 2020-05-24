package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
)

func (s *Service) updatePerms(ch channel, userID uuid.UUID, permSvc *permission.Service, param map[string]interface{}) error {
	println("1")
	permParam, ok := param["permissions"]
	if !ok {
		return nil
	}
	permArray := permParam.([]interface{})
	var perms permission.Permissions
	for _, a := range permArray {
		p := a.(map[string]interface{})
		k := p["k"].(string)
		v := p["v"].(string)
		access := p["access"].(string)
		perms = append(perms, &permission.Permission{K: k, V: v, Access: member.RoleFromString(access)})
	}

	final, err := permSvc.SetAll(ch.ID, perms, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "unable to set permissions"))
	}

	err = sendPermissionsUpdate(s, ch, final)
	return errors.WithStack(errors.Wrap(err, "unable to send permissions update"))
}

func sendPermissionsUpdate(s *Service, ch channel, perms permission.Permissions) error {
	err := s.WriteChannel(ch, &Message{Svc: ch.Svc, Cmd: ServerCmdPermissionsUpdate, Param: perms})
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error writing permission update message"))
	}

	return err
}
