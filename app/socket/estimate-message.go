package socket

import (
	"encoding/json"
	"github.com/kyleu/npn/npnconnection"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/permission"
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

func onEstimateMessage(s *npnconnection.Service, conn *npnconnection.Connection, cmd string, param json.RawMessage) error {
	dataSvc := estimates(s)
	var err error
	userID := conn.Profile.UserID
	switch cmd {
	case ClientCmdConnect:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onEstimateConnect(s, conn, u)
	case ClientCmdUpdateSession:
		ess := estimateSessionSaveParams{}
		_ = npncore.FromJSON(param, &ess)
		err = onEstimateSessionSave(s, conn, userID, ess)
	case ClientCmdUpdateMember:
		u := updateMemberParams{}
		_ = npncore.FromJSON(param, &u)
		err = onUpdateMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	case ClientCmdRemoveMember:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onRemoveMember(s, dataSvc.Data.Members, *conn.Channel, userID, u)
	case ClientCmdAddStory:
		asp := addStoryParams{}
		_ = npncore.FromJSON(param, &asp)
		err = onAddStory(s, *conn.Channel, userID, asp)
	case ClientCmdUpdateStory:
		usp := updateStoryParams{}
		_ = npncore.FromJSON(param, &usp)
		err = onUpdateStory(s, *conn.Channel, userID, usp)
	case ClientCmdRemoveStory:
		var u uuid.UUID
		_ = npncore.FromJSON(param, &u)
		err = onRemoveStory(s, *conn.Channel, userID, u)
	case ClientCmdSetStoryStatus:
		sssp := setStoryStatusParams{}
		_ = npncore.FromJSON(param, &sssp)
		err = onSetStoryStatus(s, *conn.Channel, userID, sssp)
	case ClientCmdSubmitVote:
		svp := submitVoteParams{}
		_ = npncore.FromJSON(param, &svp)
		err = onSubmitVote(s, *conn.Channel, userID, svp)
	default:
		err = errors.New("unhandled estimate command [" + cmd + "]")
	}
	return errors.Wrap(err, "error handling estimate message")
}
