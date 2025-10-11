package websocket

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/user"
	dbuser "github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/app/util"
)

type Connection struct {
	ID       uuid.UUID     `json:"id"`
	User     *dbuser.User  `json:"user,omitzero"`
	Profile  *user.Profile `json:"profile,omitzero"`
	Accounts user.Accounts `json:"accounts,omitempty"`
	Svc      string        `json:"svc,omitzero"`
	ModelID  *uuid.UUID    `json:"modelID,omitzero"`
	Channels []string      `json:"channels,omitempty"`
	Started  time.Time     `json:"started,omitzero"`
	Stats    *Stats        `json:"stats,omitempty"`
	handler  Handler
	socket   *websocket.Conn
	mu       sync.Mutex
}

func NewConnection(svc string, usr *dbuser.User, profile *user.Profile, accounts user.Accounts, socket *websocket.Conn, handler Handler) *Connection {
	return &Connection{
		ID:       util.UUID(),
		User:     usr,
		Profile:  profile,
		Accounts: accounts,
		Svc:      svc,
		Started:  util.TimeCurrent(),
		Stats:    NewStats(),
		handler:  handler,
		socket:   socket,
	}
}

func (c *Connection) ToStatus() *Status {
	if c.Channels == nil {
		return &Status{ID: c.ID, Username: c.Profile.Name, Channels: nil}
	}
	return &Status{ID: c.ID, Username: c.Profile.Name, Channels: c.Channels}
}

func (c *Connection) Username() string {
	if c.User != nil {
		return c.User.Name
	}
	return c.Profile.Name
}

func (c *Connection) Write(b []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	err := c.socket.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		return errors.Wrap(err, "unable to write to websocket")
	}
	return nil
}

func (c *Connection) Read() ([]byte, error) {
	_, message, err := c.socket.ReadMessage()
	return message, errors.Wrap(err, "unable to write to websocket")
}

func (c *Connection) Close(logger util.Logger) error {
	if logger != nil {
		logger.Infof("closing connection [%s]: %s", c.ID, c.Stats.String())
	}
	return c.socket.Close()
}

func (c *Connection) String() string {
	return fmt.Sprintf("[%s][%s::%s][%s]", c.ID, c.Svc, c.ModelID, c.Profile.String())
}
