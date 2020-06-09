package pdf

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/standup"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/model/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

func Render(rsp interface{}, url string) (string, []byte, error) {
	m := newDoc()
	writeDocHeader(url, m)
	filename, err := renderResponse(rsp, m)
	if err != nil {
		return filename, nil, err
	}
	return response(filename, m)
}

func renderResponse(rsp interface{}, m pdfgen.Maroto) (string, error) {
	switch rsp.(type) {
	case transcript.EmailResponse:
		return renderEmail(rsp.(transcript.EmailResponse), m)
	case team.Sessions:
		return renderTeamList(rsp.(team.Sessions), nil, m)
	case transcript.TeamResponse:
		return renderTeam(rsp.(transcript.TeamResponse), m)
	case sprint.Sessions:
		return renderSprintList(rsp.(sprint.Sessions), nil, m)
	case transcript.SprintResponse:
		return renderSprint(rsp.(transcript.SprintResponse), m)
	case estimate.Sessions:
		return renderEstimateList(rsp.(estimate.Sessions), nil, m)
	case transcript.EstimateResponse:
		return renderEstimate(rsp.(transcript.EstimateResponse), m)
	case standup.Sessions:
		return renderStandupList(rsp.(standup.Sessions), nil, m)
	case transcript.StandupResponse:
		return renderStandup(rsp.(transcript.StandupResponse), m)
	case retro.Sessions:
		return renderRetroList(rsp.(retro.Sessions), nil, m)
	case transcript.RetroResponse:
		return renderRetro(rsp.(transcript.RetroResponse), m)
	default:
		return "error", errors.New(fmt.Sprintf("Invalid transcript type [%T]", rsp))
	}
}

func newDoc() pdfgen.Maroto {
	m := pdfgen.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 10, 10)
	return m
}

func writeDocHeader(url string, m pdfgen.Maroto) {
	m.RegisterHeader(func() {
		tr(func() { col(func() {
			m.Text(util.AppName, props.Text{Size: 16, Align: consts.Left})
			m.Text(url, props.Text{Size: 8, Align: consts.Right})
		}, 12, m) }, 10, m)
	})
}

func response(fn string, m pdfgen.Maroto) (string, []byte, error) {
	buff, err := m.Output()
	if err != nil {
		return fn, nil, errors.Wrap(err, "error writing PDF output")
	}

	if len(fn) == 0 {
		fn = util.KeyExport
	}
	return fn, buff.Bytes(), nil
}
