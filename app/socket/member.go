package socket

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) UpdateName(id uuid.UUID, name string) error {
	return s.users.UpdateUserName(id, name)
}

func (s *Service) GetOnline(ch channel) ([]uuid.UUID, error) {
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

	return online, nil
}

func (s *Service) sendOnlineUpdate(ch channel, connID uuid.UUID, userID uuid.UUID, connected bool) error {
	p := OnlineUpdate{UserID: userID, Connected: connected}
	onlineMsg := Message{Svc: util.SvcSystem.Key, Cmd: ServerCmdOnlineUpdate, Param: p}
	return s.WriteChannel(ch, &onlineMsg, connID)
}

func (s *Service) sendMemberUpdate(ch channel, current *member.Entry, except ...uuid.UUID) error {
	onlineMsg := Message{Svc: util.SvcSystem.Key, Cmd: ServerCmdMemberUpdate, Param: current}
	return s.WriteChannel(ch, &onlineMsg, except...)
}
