package socket

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type connection struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	Channels []uuid.UUID
	socket   *websocket.Conn
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

func (m *Message) ParamJson() string {
	data, err := json.Marshal(m.Param)
	if err != nil {
		return "error: " + err.Error()
	}
	return string(data)
}
