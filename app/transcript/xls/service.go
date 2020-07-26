package xls

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/app/util"
)

var defSheet = "Sheet1"

func Render(rsp interface{}, url string) (string, []byte, error) {
	f := newDoc()
	filename, title, err := renderResponse(rsp, f)
	if err != nil {
		return filename, nil, err
	}
	writeAboutSheet(url, title, f)
	setFirstSheetTitle(filename, f)
	return response(filename, f)
}

func renderResponse(rsp interface{}, f *excelize.File) (string, string, error) {
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
		return util.KeyError, util.Title(util.KeyError), errors.New(fmt.Sprintf("Invalid transcript type [%T]", rsp))
	}
}

func newDoc() *excelize.File {
	f := excelize.NewFile()
	f.SetActiveSheet(1)
	return f
}

func writeAboutSheet(url string, title string, f *excelize.File) {
	key := util.AppName
	f.NewSheet(key)
	f.SetCellValue(key, "A1", util.AppName)
	f.SetCellValue(key, "A2", url)
	f.SetCellValue(key, "A3", title)
	setColumnWidths(key, []int{64}, f)
}

func response(fn string, f *excelize.File) (string, []byte, error) {
	buff, err := f.WriteToBuffer()
	if err != nil {
		return fn, nil, errors.Wrap(err, "error writing PDF output")
	}

	if len(fn) == 0 {
		fn = util.KeyExport
	}
	return fn, buff.Bytes(), nil
}
