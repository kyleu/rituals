package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
)

type updateMemberParams struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
}

func sendOnlineUpdate(s *npnconnection.Service, ch npnconnection.Channel, connID uuid.UUID, userID uuid.UUID, connected bool) error {
	p := npnconnection.OnlineUpdate{UserID: userID, Connected: connected}
	onlineMsg := npnconnection.NewMessage(util.SvcSystem.Key, ServerCmdOnlineUpdate, p)
	return s.WriteChannel(ch, onlineMsg, connID)
}

func sendMemberUpdate(s *npnconnection.Service, ch npnconnection.Channel, current *member.Entry, except ...uuid.UUID) error {
	if current == nil {
		return errors.New("no current member to send update for")
	}
	onlineMsg := npnconnection.NewMessage(util.SvcSystem.Key, ServerCmdMemberUpdate, current)
	return s.WriteChannel(ch, onlineMsg, except...)
}

func sendMemberRemoved(s *npnconnection.Service, ch npnconnection.Channel, member uuid.UUID, except ...uuid.UUID) error {
	onlineMsg := npnconnection.NewMessage(util.SvcSystem.Key, ServerCmdMemberRemove, member)
	return s.WriteChannel(ch, onlineMsg, except...)
}

func onRemoveMember(s *npnconnection.Service, memberSvc *member.Service, ch npnconnection.Channel, userID uuid.UUID, target uuid.UUID) error {
	err := memberSvc.RemoveMember(ch.ID, userID, target)
	if err != nil {
		return errors.Wrap(err, "unable to remove member")
	}
	err = sendMemberRemoved(s, ch, target)
	return errors.Wrap(err, "unable to send member update")
}

func onUpdateMember(s *npnconnection.Service, memberSvc *member.Service, ch npnconnection.Channel, src uuid.UUID, params updateMemberParams) error {
	curr, err := memberSvc.UpdateMember(ch.ID, src, params.ID, params.Role)
	if err != nil {
		return errors.Wrap(err, "unable to remove member")
	}
	err = sendMemberUpdate(s, ch, curr)
	return errors.Wrap(err, "unable to send member update")
}
