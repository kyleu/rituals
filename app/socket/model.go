package socket

import (
	"encoding/json"
	"fmt"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/kyleu/rituals.dev/app/util"
)

type connection struct {
	ID      uuid.UUID
	Profile npnuser.Profile
	Svc     util.Service
	ModelID *uuid.UUID
	Channel *Channel
	socket  *websocket.Conn
	mu      sync.Mutex
}

func (c *connection) ToStatus() *Status {
	if c.Channel == nil {
		return &Status{ID: c.ID, UserID: c.Profile.UserID, Username: c.Profile.Name, ChannelSvc: util.SvcSystem, ChannelID: nil}
	}
	return &Status{ID: c.ID, UserID: c.Profile.UserID, Username: c.Profile.Name, ChannelSvc: c.Channel.Svc, ChannelID: &c.Channel.ID}
}

type Status struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"userID"`
	Username   string
	ChannelSvc util.Service
	ChannelID  *uuid.UUID
}

type Statuses = []*Status

type Message struct {
	Svc   string          `json:"svc"`
	Cmd   string          `json:"cmd"`
	Param json.RawMessage `json:"param"`
}

func NewMessage(svc util.Service, cmd string, param interface{}) *Message {
	return &Message{Svc: svc.Key, Cmd: cmd, Param: json.RawMessage(npncore.ToJSON(param, nil))}
}

func (m *Message) String() string {
	return fmt.Sprintf("%s/%s", m.Svc, m.Cmd)
}

type OnlineUpdate struct {
	UserID    uuid.UUID `json:"userID"`
	Connected bool      `json:"connected"`
}

func differentPointerValues(l *uuid.UUID, r *uuid.UUID) bool {
	switch {
	case l != nil && r != nil:
		return *l != *r
	case l == nil && r != nil:
		return true
	case l != nil && r == nil:
		return true
	default:
		return false
	}
}
