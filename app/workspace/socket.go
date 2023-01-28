package workspace

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/websocket"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) SocketOpen(sock *websocket.Service, conn *websocket.Connection, logger util.Logger) error {
	return nil
}

func (s *Service) SocketHandler(
	ctx context.Context, sock *websocket.Service, conn *websocket.Connection, svc string, cmd string, param json.RawMessage, logger util.Logger,
) error {
	logger.Infof("processing [%s] message of type [%s]", svc, cmd)
	svc, idStr := util.StringSplit(svc, ':', true)
	id := util.UUIDFromString(idStr)
	if id == nil {
		return errors.Errorf("invalid ID [%s]", idStr)
	}

	frm := util.ValueMap{}
	err := util.FromJSON(param, &frm)
	if err != nil {
		return err
	}

	p := NewParams(ctx, idStr, action.Act(cmd), frm, conn.Profile.ID, s, logger, conn.ID)

	switch enum.ModelService(svc) {
	case enum.ModelServiceTeam:
		_, _, _, err := s.ActionTeam(p)
		return err
	case enum.ModelServiceSprint:
		_, _, _, err := s.ActionSprint(p)
		return err
	case enum.ModelServiceEstimate:
		_, _, _, err := s.ActionEstimate(p)
		return err
	case enum.ModelServiceStandup:
		_, _, _, err := s.ActionStandup(p)
		return err
	case enum.ModelServiceRetro:
		_, _, _, err := s.ActionRetro(p)
		return err
	default:
		return errors.Errorf("invalid service [%s]", svc)
	}
}

func (s *Service) SocketClose(sock *websocket.Service, conn *websocket.Connection, logger util.Logger) error {
	param := util.ValueMap{"userID": conn.Profile.ID, "connected": false}
	for _, ch := range conn.Channels {
		svc, modelIDStr := util.StringSplit(ch, ':', true)
		if modelIDStr == "" {
			continue
		}
		modelID := util.UUIDFromString(modelIDStr)
		err := s.send(enum.ModelService(svc), *modelID, "online-update", param, &conn.Profile.ID, logger, conn.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
