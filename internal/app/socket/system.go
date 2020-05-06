package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/internal/app/member"
	"github.com/kyleu/rituals.dev/internal/app/util"
)

func onSystemMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface {}) error {
	if conn.UserID != userID {
		return errors.WithStack(errors.New("received name change for wrong user [" + userID.String() + "]"))
	}
	var err error = nil
	switch cmd {
	case "member-name-save":
		err = saveName(s, conn, userID, param.(map[string]interface {}))
	default:
		err = errors.New("unhandled system command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling system message"))
}

func saveName(s *Service, conn *connection, userID uuid.UUID, o map[string]interface {}) error {
	name := o["name"].(string)
	choice := o["choice"].(string)
	if choice == "global" {
		err := s.UpdateName(userID, name)
		if err != nil {
			return err
		}
	}
	memberSvc, err := memberSvcFor(s, conn.Channel.Svc)
	if err != nil {
		return err
	}

	current, err := memberSvc.Get(conn.Channel.ID, userID)
	if err != nil {
		return err
	}

	if current.Name != name {
		err = memberSvc.UpdateName(conn.Channel.ID, userID, name)
		if err != nil {
			return err
		}
		if conn.Channel == nil {
			return errors.New("no channel registered for [" + conn.ID.String() + "]")
		}
		return s.SendMembers(memberSvc, *conn.Channel, nil)
	}
	return nil
}

func memberSvcFor(s *Service, svc string) (*member.Service, error) {
	var ret *member.Service = nil
	switch svc {
	case util.SvcEstimate:
		ret = &s.estimates.Members
	case util.SvcStandup:
		ret = &s.standups.Members
	case util.SvcRetro:
		ret = &s.retros.Members
	default:
		return nil, errors.New("invalid service [" + svc + "]")
	}
	return ret, nil
}
