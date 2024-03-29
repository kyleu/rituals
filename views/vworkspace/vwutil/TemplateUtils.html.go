// Code generated by qtc from "TemplateUtils.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwutil/TemplateUtils.html:1
package vwutil

//line views/vworkspace/vwutil/TemplateUtils.html:1
import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

//line views/vworkspace/vwutil/TemplateUtils.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwutil/TemplateUtils.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwutil/TemplateUtils.html:9
func StreamBanner(qw422016 *qt422016.Writer, t *team.Team, s *sprint.Sprint, mdl string) {
//line views/vworkspace/vwutil/TemplateUtils.html:9
	qw422016.N().S(`<em id="model-banner">`)
//line views/vworkspace/vwutil/TemplateUtils.html:11
	if s != nil {
//line views/vworkspace/vwutil/TemplateUtils.html:11
		qw422016.N().S(`<a href="`)
//line views/vworkspace/vwutil/TemplateUtils.html:11
		qw422016.E().S(s.PublicWebPath())
//line views/vworkspace/vwutil/TemplateUtils.html:11
		qw422016.N().S(`">`)
//line views/vworkspace/vwutil/TemplateUtils.html:11
		qw422016.E().S(s.TitleString())
//line views/vworkspace/vwutil/TemplateUtils.html:11
		qw422016.N().S(`</a>`)
//line views/vworkspace/vwutil/TemplateUtils.html:11
		qw422016.N().S(` `)
//line views/vworkspace/vwutil/TemplateUtils.html:11
	}
//line views/vworkspace/vwutil/TemplateUtils.html:12
	qw422016.E().S(mdl)
//line views/vworkspace/vwutil/TemplateUtils.html:13
	if t != nil {
//line views/vworkspace/vwutil/TemplateUtils.html:13
		qw422016.N().S(` `)
//line views/vworkspace/vwutil/TemplateUtils.html:13
		qw422016.N().S(`in <a href="`)
//line views/vworkspace/vwutil/TemplateUtils.html:13
		qw422016.E().S(t.PublicWebPath())
//line views/vworkspace/vwutil/TemplateUtils.html:13
		qw422016.N().S(`">`)
//line views/vworkspace/vwutil/TemplateUtils.html:13
		qw422016.E().S(t.TitleString())
//line views/vworkspace/vwutil/TemplateUtils.html:13
		qw422016.N().S(`</a>`)
//line views/vworkspace/vwutil/TemplateUtils.html:13
	}
//line views/vworkspace/vwutil/TemplateUtils.html:13
	qw422016.N().S(`</em>`)
//line views/vworkspace/vwutil/TemplateUtils.html:15
}

//line views/vworkspace/vwutil/TemplateUtils.html:15
func WriteBanner(qq422016 qtio422016.Writer, t *team.Team, s *sprint.Sprint, mdl string) {
//line views/vworkspace/vwutil/TemplateUtils.html:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/TemplateUtils.html:15
	StreamBanner(qw422016, t, s, mdl)
//line views/vworkspace/vwutil/TemplateUtils.html:15
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/TemplateUtils.html:15
}

//line views/vworkspace/vwutil/TemplateUtils.html:15
func Banner(t *team.Team, s *sprint.Sprint, mdl string) string {
//line views/vworkspace/vwutil/TemplateUtils.html:15
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/TemplateUtils.html:15
	WriteBanner(qb422016, t, s, mdl)
//line views/vworkspace/vwutil/TemplateUtils.html:15
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/TemplateUtils.html:15
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/TemplateUtils.html:15
	return qs422016
//line views/vworkspace/vwutil/TemplateUtils.html:15
}

//line views/vworkspace/vwutil/TemplateUtils.html:17
func StreamEditWorkspaceForm(qw422016 *qt422016.Writer, svc string, teamID *uuid.UUID, sprintID *uuid.UUID, placeholder string) {
//line views/vworkspace/vwutil/TemplateUtils.html:17
	qw422016.N().S(`<form action="/`)
//line views/vworkspace/vwutil/TemplateUtils.html:18
	qw422016.E().S(svc)
//line views/vworkspace/vwutil/TemplateUtils.html:18
	qw422016.N().S(`" method="post">`)
//line views/vworkspace/vwutil/TemplateUtils.html:19
	if teamID != nil {
//line views/vworkspace/vwutil/TemplateUtils.html:19
		qw422016.N().S(`<input type="hidden" name="`)
//line views/vworkspace/vwutil/TemplateUtils.html:20
		qw422016.E().S(util.KeyTeam)
//line views/vworkspace/vwutil/TemplateUtils.html:20
		qw422016.N().S(`" value="`)
//line views/vworkspace/vwutil/TemplateUtils.html:20
		qw422016.E().S(teamID.String())
//line views/vworkspace/vwutil/TemplateUtils.html:20
		qw422016.N().S(`" />`)
//line views/vworkspace/vwutil/TemplateUtils.html:21
	}
//line views/vworkspace/vwutil/TemplateUtils.html:22
	if sprintID != nil {
//line views/vworkspace/vwutil/TemplateUtils.html:22
		qw422016.N().S(`<input type="hidden" name="`)
//line views/vworkspace/vwutil/TemplateUtils.html:23
		qw422016.E().S(util.KeySprint)
//line views/vworkspace/vwutil/TemplateUtils.html:23
		qw422016.N().S(`" value="`)
//line views/vworkspace/vwutil/TemplateUtils.html:23
		qw422016.E().S(sprintID.String())
//line views/vworkspace/vwutil/TemplateUtils.html:23
		qw422016.N().S(`" />`)
//line views/vworkspace/vwutil/TemplateUtils.html:24
	}
//line views/vworkspace/vwutil/TemplateUtils.html:24
	qw422016.N().S(`<input type="text" name="title" class="combined" placeholder="`)
//line views/vworkspace/vwutil/TemplateUtils.html:25
	qw422016.E().S(placeholder)
//line views/vworkspace/vwutil/TemplateUtils.html:25
	qw422016.N().S(`" /><button type="submit" class="combined">+</button></form>`)
//line views/vworkspace/vwutil/TemplateUtils.html:28
}

//line views/vworkspace/vwutil/TemplateUtils.html:28
func WriteEditWorkspaceForm(qq422016 qtio422016.Writer, svc string, teamID *uuid.UUID, sprintID *uuid.UUID, placeholder string) {
//line views/vworkspace/vwutil/TemplateUtils.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/TemplateUtils.html:28
	StreamEditWorkspaceForm(qw422016, svc, teamID, sprintID, placeholder)
//line views/vworkspace/vwutil/TemplateUtils.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/TemplateUtils.html:28
}

//line views/vworkspace/vwutil/TemplateUtils.html:28
func EditWorkspaceForm(svc string, teamID *uuid.UUID, sprintID *uuid.UUID, placeholder string) string {
//line views/vworkspace/vwutil/TemplateUtils.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/TemplateUtils.html:28
	WriteEditWorkspaceForm(qb422016, svc, teamID, sprintID, placeholder)
//line views/vworkspace/vwutil/TemplateUtils.html:28
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/TemplateUtils.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/TemplateUtils.html:28
	return qs422016
//line views/vworkspace/vwutil/TemplateUtils.html:28
}
