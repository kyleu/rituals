package xls

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/standup"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/model/transcript"
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
	switch rsp.(type) {
	case transcript.EmailResponse:
		return renderEmail(rsp.(transcript.EmailResponse), f)
	case team.Sessions:
		return renderTeamList(rsp.(team.Sessions), nil, f)
	case transcript.TeamResponse:
		return renderTeam(rsp.(transcript.TeamResponse), f)
	case sprint.Sessions:
		return renderSprintList(rsp.(sprint.Sessions), nil, f)
	case transcript.SprintResponse:
		return renderSprint(rsp.(transcript.SprintResponse), f)
	case estimate.Sessions:
		return renderEstimateList(rsp.(estimate.Sessions), nil, f)
	case transcript.EstimateResponse:
		return renderEstimate(rsp.(transcript.EstimateResponse), f)
	case standup.Sessions:
		return renderStandupList(rsp.(standup.Sessions), nil, f)
	case transcript.StandupResponse:
		return renderStandup(rsp.(transcript.StandupResponse), f)
	case retro.Sessions:
		return renderRetroList(rsp.(retro.Sessions), nil, f)
	case transcript.RetroResponse:
		return renderRetro(rsp.(transcript.RetroResponse), f)
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
