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
	T string `json:"t"`
	K string `json:"k"`
	V string `json:"v"`
}
