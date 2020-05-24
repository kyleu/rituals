package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
	"time"
)

func (s *Service) UpdateName(id uuid.UUID, name string) error {
	return s.users.UpdateUserName(id, name)
}

func (s *Service) GetOnline(ch channel) []uuid.UUID {
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

func (s *Service) sendOnlineUpdate(ch channel, connID uuid.UUID, userID uuid.UUID, connected bool) error {
	p := OnlineUpdate{UserID: userID, Connected: connected}
	onlineMsg := Message{Svc: util.SvcSystem.Key, Cmd: ServerCmdOnlineUpdate, Param: p}
	return s.WriteChannel(ch, &onlineMsg, connID)
}

func (s *Service) sendMemberUpdate(ch channel, current *member.Entry, except ...uuid.UUID) error {
	onlineMsg := Message{Svc: util.SvcSystem.Key, Cmd: ServerCmdMemberUpdate, Param: current}
	return s.WriteChannel(ch, &onlineMsg, except...)
}

func onRemoveMember(s *Service, memberSvc *member.Service, ch channel, userID uuid.UUID, targetString string) error {
	target, err := uuid.FromString(targetString)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "invalid target id [" + targetString + "]"))
	}

	err = memberSvc.RemoveMember(ch.ID, target)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "unable to remove member from team"))
	}

	err = s.sendMemberUpdate(ch, &member.Entry{
		UserID:  target,
		Name:    "::delete",
		Role:    member.Role{},
		Created: time.Time{},
	})

	return errors.WithStack(errors.Wrap(err, "unable to send member update from team"))
}
