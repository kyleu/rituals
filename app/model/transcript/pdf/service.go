package pdf

import (
	"fmt"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/standup"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/model/transcript"
)

func Render(rsp interface{}) ([]byte, error) {
	switch rsp.(type) {
	case transcript.EmailResponse:
    return renderEmail(rsp.(transcript.EmailResponse))
	case team.Sessions:
		return renderTeamList(rsp.(team.Sessions))
	case transcript.TeamResponse:
		return renderTeam(rsp.(transcript.TeamResponse))
	case sprint.Sessions:
		return renderSprintList(rsp.(sprint.Sessions))
	case transcript.SprintResponse:
		return renderSprint(rsp.(transcript.SprintResponse))
	case estimate.Sessions:
		return renderEstimateList(rsp.(estimate.Sessions))
	case transcript.EstimateResponse:
		return renderEstimate(rsp.(transcript.EstimateResponse))
	case standup.Sessions:
		return renderStandupList(rsp.(standup.Sessions))
	case transcript.StandupResponse:
		return renderStandup(rsp.(transcript.StandupResponse))
	case retro.Sessions:
		return renderRetroList(rsp.(retro.Sessions))
	case transcript.RetroResponse:
		return renderRetro(rsp.(transcript.RetroResponse))
	default:
    return []byte(fmt.Sprintf("Invalid transcript type [%T]", rsp)), nil
  }
}

func renderEmail(response transcript.EmailResponse) ([]byte, error) {
	return nil, nil
}

func renderTeam(rsp transcript.TeamResponse) ([]byte, error) {
	return nil, nil
}

func renderTeamList(sessions team.Sessions) ([]byte, error) {
	return nil, nil
}

func renderEstimate(response transcript.EstimateResponse) ([]byte, error) {
	return nil, nil
}

func renderEstimateList(sessions estimate.Sessions) ([]byte, error) {
	return nil, nil
}

func renderSprint(response transcript.SprintResponse) ([]byte, error) {
	return nil, nil
}

func renderSprintList(sessions sprint.Sessions) ([]byte, error) {
	return nil, nil
}

func renderStandupList(sessions standup.Sessions) ([]byte, error) {
	return nil, nil
}

func renderStandup(response transcript.StandupResponse) ([]byte, error) {
	return nil, nil
}

func renderRetroList(sessions retro.Sessions) ([]byte, error) {
	return nil, nil
}

func renderRetro(response transcript.RetroResponse) ([]byte, error) {
	return nil, nil
}

