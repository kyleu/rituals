package socket

import (
	"emperror.dev/errors"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/util"
)

type estimateSessionSaveParams struct {
	Title       string                 `json:"title"`
	Choices     string                 `json:"choices"`
	SprintID    string                 `json:"sprintID"`
	TeamID      string                 `json:"teamID"`
	Permissions permission.Permissions `json:"permissions"`
}

type addStoryParams struct {
	Title string `json:"title"`
}

type updateStoryParams struct {
	StoryID uuid.UUID `json:"storyID"`
	Title   string    `json:"title"`
}

type setStoryStatusParams struct {
	StoryID uuid.UUID `json:"storyID"`
	Status  string    `json:"status"`
}

type submitVoteParams struct {
	StoryID uuid.UUID `json:"storyID"`
	Choice  string    `json:"choice"`
}

func onEstimateMessage(s *Service, conn *Connection, cmd string, param json.RawMessage) error {
	var err error
	userID := conn.Profile.UserID
	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onEstimateConnect(s, conn, u)
	case ClientCmdUpdateSession:
		ess := estimateSessionSaveParams{}
		util.FromJSON(param, &ess, s.Logger)
		err = onEstimateSessionSave(s, *conn.Channel, userID, ess)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onRemoveMember(s, s.estimates.Data.Members, *conn.Channel, userID, u)
	case ClientCmdAddStory:
		asp := addStoryParams{}
		util.FromJSON(param, &asp, s.Logger)
		err = onAddStory(s, *conn.Channel, userID, asp)
	case ClientCmdUpdateStory:
		usp := updateStoryParams{}
		util.FromJSON(param, &usp, s.Logger)
		err = onUpdateStory(s, *conn.Channel, userID, usp)
	case ClientCmdRemoveStory:
		var u uuid.UUID
		util.FromJSON(param, &u, s.Logger)
		err = onRemoveStory(s, *conn.Channel, userID, u)
	case ClientCmdSetStoryStatus:
		sssp := setStoryStatusParams{}
		util.FromJSON(param, &sssp, s.Logger)
		err = onSetStoryStatus(s, *conn.Channel, userID, sssp)
	case ClientCmdSubmitVote:
		svp := submitVoteParams{}
		util.FromJSON(param, &svp, s.Logger)
		err = onSubmitVote(s, *conn.Channel, userID, svp)
	default:
		err = errors.New("unhandled estimate command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling estimate message")
}
