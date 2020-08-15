package pdf

import (
	"fmt"

	"emperror.dev/errors"
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/transcript"
)

func RenderCallback(rsp interface{}, m pdfgen.Maroto) (string, error) {
	switch r := rsp.(type) {
	case transcript.EmailResponse:
		return renderEmail(r, m)
	case team.Sessions:
		renderTeamList(r, nil, m)
	case transcript.TeamResponse:
		return renderTeam(r, m)
	case sprint.Sessions:
		renderSprintList(r, nil, m)
	case transcript.SprintResponse:
		return renderSprint(r, m)
	case estimate.Sessions:
		renderEstimateList(r, nil, m)
	case transcript.EstimateResponse:
		return renderEstimate(r, m), nil
	case standup.Sessions:
		renderStandupList(r, nil, m)
	case transcript.StandupResponse:
		return renderStandup(r, m), nil
	case retro.Sessions:
		renderRetroList(r, nil, m)
	case transcript.RetroResponse:
		return renderRetro(r, m), nil
	default:
		return "error", errors.New(fmt.Sprintf("Invalid transcript type [%T]", rsp))
	}
	return "", nil
}
