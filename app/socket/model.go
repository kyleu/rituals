package socket

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/kyleu/rituals.dev/app/util"
)

type connection struct {
	ID      uuid.UUID
	Profile util.Profile
	Svc     string
	ModelID *uuid.UUID
	Channel *channel
	socket  *websocket.Conn
	mu      sync.Mutex
}

func (c *connection) ToStatus() *Status {
	if c.Channel == nil {
		return &Status{ID: c.ID, UserID: c.Profile.UserID, Username: c.Profile.Name, ChannelSvc: "", ChannelID: nil}
	}
	return &Status{ID: c.ID, UserID: c.Profile.UserID, Username: c.Profile.Name, ChannelSvc: c.Channel.Svc, ChannelID: &c.Channel.ID}
}

type Status struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"userID"`
	Username   string
	ChannelSvc string
	ChannelID  *uuid.UUID
}

type Message struct {
	Svc   string      `json:"svc"`
	Cmd   string      `json:"cmd"`
	Param interface{} `json:"param"`
}

func (m Message) String() string {
	return fmt.Sprintf("%s/%s", m.Svc, m.Cmd)
}

func (m *Message) ParamJSON() string {
	data, err := json.Marshal(m.Param)
	if err != nil {
		return "error: " + err.Error()
	}
	return string(data)
}

type OnlineUpdate struct {
	UserID    uuid.UUID `json:"userID"`
	Connected bool      `json:"connected"`
}
