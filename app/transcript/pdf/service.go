package pdf

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/johnfercher/maroto/pkg/consts"
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/transcript"
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

func newDoc() pdfgen.Maroto {
	m := pdfgen.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 10, 10)
	return m
}

func writeDocHeader(url string, m pdfgen.Maroto) {
	m.RegisterHeader(func() {
		tr(func() {
			col(func() {
				m.Text(util.AppName, props.Text{Size: 16, Align: consts.Left})
				m.Text(url, props.Text{Size: 8, Align: consts.Right})
			}, 12, m)
		}, 10, m)
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
