package xls

import (
	"fmt"

	"github.com/kyleu/npn/npncore"

	"emperror.dev/errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/transcript"
)

func RenderResponse(rsp interface{}, f *excelize.File) (string, string, error) {
	switch r := rsp.(type) {
	case transcript.EmailResponse:
		return renderEmail(r, f)
	case team.Sessions:
		fn, title := renderTeamList(r, nil, f)
		return fn, title, nil
	case transcript.TeamResponse:
		return renderTeam(r, f)
	case sprint.Sessions:
		fn, title := renderSprintList(r, nil, f)
		return fn, title, nil
	case transcript.SprintResponse:
		return renderSprint(r, f)
	case estimate.Sessions:
		fn, title := renderEstimateList(r, nil, f)
		return fn, title, nil
	case transcript.EstimateResponse:
		return renderEstimate(r, f)
	case standup.Sessions:
		fn, title := renderStandupList(r, nil, f)
		return fn, title, nil
	case transcript.StandupResponse:
		return renderStandup(r, f)
	case retro.Sessions:
		fn, title := renderRetroList(r, nil, f)
		return fn, title, nil
	case transcript.RetroResponse:
		return renderRetro(r, f)
	default:
		return npncore.KeyError, npncore.Title(npncore.KeyError), errors.New(fmt.Sprintf("Invalid transcript type [%T]", rsp))
	}
}
