package workspace

import (
	"encoding/json"

	"github.com/kyleu/rituals/app/lib/websocket"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) SocketOpen(sock *websocket.Service, conn *websocket.Connection) error {
	sock.Logger.Infof("OPEN: %s", util.ToJSON(conn))
	return nil
}

func (s *Service) SocketHandler(sock *websocket.Service, conn *websocket.Connection, svc string, cmd string, param json.RawMessage) error {
	sock.Logger.Infof("HANDLE: %s", util.ToJSON(conn))
	return nil
}

func (s *Service) SocketClose(sock *websocket.Service, conn *websocket.Connection) error {
	sock.Logger.Infof("CLOSE: %s", util.ToJSON(conn))
	return nil
}
