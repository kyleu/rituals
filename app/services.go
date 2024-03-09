package app

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/database/migrate"
	"github.com/kyleu/rituals/app/lib/websocket"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/queries/migrations"
)

type Services struct {
	GeneratedServices

	Workspace *workspace.Service
	Socket    *websocket.Service
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	migrations.LoadMigrations(st.Debug)
	err := migrate.Migrate(ctx, st.DB, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to run database migrations")
	}

	gen := initGeneratedServices(ctx, st.DB, logger)
	w := workspace.NewService(
		gen.Team, gen.TeamHistory, gen.TeamMember, gen.TeamPermission,
		gen.Sprint, gen.SprintHistory, gen.SprintMember, gen.SprintPermission,
		gen.Estimate, gen.EstimateHistory, gen.EstimateMember, gen.EstimatePermission, gen.Story, gen.Vote,
		gen.Standup, gen.StandupHistory, gen.StandupMember, gen.StandupPermission, gen.Report,
		gen.Retro, gen.RetroHistory, gen.RetroMember, gen.RetroPermission, gen.Feedback,
		gen.User, gen.Action, gen.Comment, gen.Email, st.DB,
	)
	ws := websocket.NewService(w.SocketOpen, w.SocketHandler, w.SocketClose)
	w.RegisterSend(func(svc enum.ModelService, id uuid.UUID, act action.Act, param any, userID *uuid.UUID, logger util.Logger, except ...uuid.UUID) error {
		ch := fmt.Sprintf("%s:%s", svc.Key, id.String())
		msg := websocket.NewMessage(userID, ch, string(act), param)
		return ws.WriteChannel(msg, logger, except...)
	}, func(connID uuid.UUID, svc enum.ModelService, id uuid.UUID, act action.Act, param any, userID *uuid.UUID, logger util.Logger) error {
		ch := fmt.Sprintf("%s:%s", svc.Key, id.String())
		msg := websocket.NewMessage(userID, ch, string(act), param)
		return ws.WriteMessage(connID, msg, logger)
	})
	w.RegisterOnline(ws.GetOnline)
	return &Services{GeneratedServices: gen, Workspace: w, Socket: ws}, nil
}

func (s *Services) Close(_ context.Context, _ util.Logger) error {
	return nil
}
