package workspace

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/websocket"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) SocketOpen(_ *websocket.Service, _ *websocket.Connection, _ util.Logger) error {
	return nil
}

func (s *Service) SocketHandler(
	ctx context.Context, _ *websocket.Service, conn *websocket.Connection, svc string, cmd string, param []byte, logger util.Logger,
) error {
	logger.Infof("processing [%s] message of type [%s]", svc, cmd)
	svc, idStr := util.StringCut(svc, ':', true)
	id := util.UUIDFromString(idStr)
	if id == nil {
		return errors.Errorf("invalid ID [%s]", idStr)
	}

	frm := util.ValueMap{}
	err := util.FromJSON(param, &frm)
	if err != nil {
		return err
	}
	var msg string
	p := NewParams(ctx, idStr, action.Act(cmd), frm, conn.Profile, conn.Accounts, s, logger, conn.ID)

	service := enum.AllModelServices.Get(svc, nil)
	switch service {
	case enum.ModelServiceTeam:
		_, msg, _, err = s.ActionTeam(p)
	case enum.ModelServiceSprint:
		_, msg, _, err = s.ActionSprint(p)
	case enum.ModelServiceEstimate:
		_, msg, _, err = s.ActionEstimate(p)
	case enum.ModelServiceStandup:
		_, msg, _, err = s.ActionStandup(p)
	case enum.ModelServiceRetro:
		_, msg, _, err = s.ActionRetro(p)
	default:
		err = errors.Errorf("invalid service [%s]", svc)
	}
	if err != nil {
		prm := map[string]any{"level": "error", "message": err.Error()}
		_ = s.sendUser(conn.ID, service, *id, action.ActMessage, prm, &conn.Profile.ID, logger)
		return err
	}
	if msg != "" {
		prm := map[string]any{"level": "success", "message": msg}
		err = s.sendUser(conn.ID, service, *id, action.ActMessage, prm, &conn.Profile.ID, logger)
		return err
	}
	return nil
}

func (s *Service) SocketClose(_ *websocket.Service, conn *websocket.Connection, logger util.Logger) error {
	param := util.ValueMap{"userID": conn.Profile.ID, "connected": false}
	for _, ch := range conn.Channels {
		svc, modelIDStr := util.StringCut(ch, ':', true)
		if modelIDStr == "" {
			continue
		}
		modelID := util.UUIDFromString(modelIDStr)
		err := s.send(enum.AllModelServices.Get(svc, nil), *modelID, "online-update", param, &conn.Profile.ID, logger, conn.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
