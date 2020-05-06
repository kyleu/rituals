package socket

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"sync"
)

type connection struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	Svc     string
	ModelID *uuid.UUID
	Channel *channel
	socket  *websocket.Conn
	mu      sync.Mutex
}

type Status struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"userID"`
}

type Message struct {
	Svc   string      `json:"svc"`
	Cmd   string      `json:"cmd"`
	Param interface{} `json:"param"`
}

func (m Message) String() string {
	return fmt.Sprintf("%s/%s", m.Svc, m.Cmd)
}

func (m *Message) ParamJson() string {
	data, err := json.Marshal(m.Param)
	if err != nil {
		return "error: " + err.Error()
	}
	return string(data)
}
