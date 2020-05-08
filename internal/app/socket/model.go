package socket

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/kyleu/rituals.dev/internal/app/util"
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

type Status struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"userID"`
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
