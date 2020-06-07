package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

type saveProfileParams struct {
	Name    string `json:"name"`
	Choice  string `json:"choice"`
	Picture string `json:"picture"`
}

func saveProfile(s *Service, conn *connection, userID uuid.UUID, p *saveProfileParams) error {
	if len(p.Name) == 0 {
		p.Name = "Unnamed Member"
	}
	if p.Choice == "global" {
		err := s.UpdateMember(userID, p.Name, p.Picture)
		if err != nil {
			return err
		}
		p.Name = ""
		p.Picture = ""
	}
	dataSvc := dataFor(s, conn.Channel.Svc)

	current, err := dataSvc.Members.Get(conn.Channel.ID, userID)
	if err != nil {
		return err
	}

	if current.Name != p.Name || current.Picture != p.Picture {
		current, err = dataSvc.Members.Update(conn.Channel.ID, userID, p.Name, p.Picture)
		if err != nil {
			return err
		}
	}

	if conn.Channel == nil {
		return errors.New("no channel registered for [" + conn.ID.String() + "]")
	}
	return s.sendMemberUpdate(*conn.Channel, current)
}
