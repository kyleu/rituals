package cworkspace

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
)

type requestForm struct {
	Form   util.ValueMap `json:"form"`
	ID     uuid.UUID     `json:"id"`
	Title  string        `json:"title"`
	Name   string        `json:"name"`
	Team   *uuid.UUID    `json:"team"`
	Sprint *uuid.UUID    `json:"sprint"`
}

func parseRequestForm(r *http.Request, b []byte, defaultName string) (*requestForm, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "no form provided")
	}
	id, _ := frm.GetUUID("id", false)
	if id == nil {
		id = util.UUIDP()
	}
	title := strings.TrimSpace(frm.GetStringOpt("title"))
	if title == "" {
		return nil, errors.New("field [title] is required")
	}
	name := strings.TrimSpace(frm.GetStringOpt("name"))
	if name == "" {
		name = defaultName
	}
	teamID, _ := frm.GetUUID(util.KeyTeam, false)
	sprintID, _ := frm.GetUUID(util.KeySprint, false)
	return &requestForm{Form: frm, ID: *id, Title: title, Name: name, Team: teamID, Sprint: sprintID}, nil
}
