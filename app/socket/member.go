package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
)

type updateMemberParams struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
}

func (s *Service) UpdateMember(id uuid.UUID, name string, picture string) error {
	return s.users.UpdateMember(id, name, picture)
}

func (s *Service) GetOnline(ch Channel) []uuid.UUID {
	connections, ok := s.channels[ch]
	if !ok {
		connections = make([]uuid.UUID, 0)
	}
	online := make([]uuid.UUID, 0)
	for _, cID := range connections {
		c, ok := s.connections[cID]
		if ok && c != nil && (!contains(online, c.Profile.UserID)) {
			online = append(online, c.Profile.UserID)
		}
	}

	return online
}

func (s *Service) sendOnlineUpdate(ch Channel, connID uuid.UUID, userID uuid.UUID, connected bool) error {
	p := OnlineUpdate{UserID: userID, Connected: connected}
	onlineMsg := NewMessage(util.SvcSystem, ServerCmdOnlineUpdate, p)
	return s.WriteChannel(ch, onlineMsg, connID)
}

func (s *Service) sendMemberUpdate(ch Channel, current *member.Entry, except ...uuid.UUID) error {
	if current == nil {
		return errors.New("no current member to send update for")
	}
	onlineMsg := NewMessage(util.SvcSystem, ServerCmdMemberUpdate, current)
	return s.WriteChannel(ch, onlineMsg, except...)
}

func (s *Service) sendMemberRemoved(ch Channel, member uuid.UUID, except ...uuid.UUID) error {
	onlineMsg := NewMessage(util.SvcSystem, ServerCmdMemberRemove, member)
	return s.WriteChannel(ch, onlineMsg, except...)
}

func onRemoveMember(s *Service, memberSvc *member.Service, ch Channel, userID uuid.UUID, target uuid.UUID) error {
	err := memberSvc.RemoveMember(ch.ID, userID, target)
	if err != nil {
		return errors.Wrap(err, "unable to remove member")
	}
	err = s.sendMemberRemoved(ch, target)
	return errors.Wrap(err, "unable to send member update")
}

func onUpdateMember(s *Service, memberSvc *member.Service, ch Channel, src uuid.UUID, params updateMemberParams) error {
	curr, err := memberSvc.UpdateMember(ch.ID, src, params.ID, params.Role)
	if err != nil {
		return errors.Wrap(err, "unable to remove member")
	}
	err = s.sendMemberUpdate(ch, curr)
	return errors.Wrap(err, "unable to send member update")
}
