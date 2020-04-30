package socket

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type connection struct {
	ID     uuid.UUID
	UserID uuid.UUID
	socket *websocket.Conn
}

type Message struct {
	Svc   string      `json:"svc"`
	Cmd   string      `json:"cmd"`
	Param interface{} `json:"param"`
}
