package cworkspace

import (
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
)

type requestForm struct {
	Form   util.ValueMap `json:"form"`
	ID     uuid.UUID     `json:"id"`
	Title  string        `json:"title"`
	Slug   string        `json:"slug"`
	Name   string        `json:"name"`
	Team   *uuid.UUID    `json:"team"`
	Sprint *uuid.UUID    `json:"sprint"`
}

func parseRequestForm(rc *fasthttp.RequestCtx, defaultName string) (*requestForm, error) {
	frm, err := cutil.ParseForm(rc)
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
	slug := frm.GetStringOpt("slug")
	if slug == "" {
		slug = slugify(title)
	} else {
		slug = slugify(slug)
	}
	name := strings.TrimSpace(frm.GetStringOpt("name"))
	if name == "" {
		name = defaultName
	}
	teamID, _ := frm.GetUUID("team", false)
	sprintID, _ := frm.GetUUID("sprint", false)
	return &requestForm{Form: frm, ID: *id, Title: title, Slug: slug, Name: name, Team: teamID, Sprint: sprintID}, nil
}
