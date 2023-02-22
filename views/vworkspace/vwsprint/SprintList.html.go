// Code generated by qtc from "SprintList.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwsprint/SprintList.html:1
package vwsprint

//line views/vworkspace/vwsprint/SprintList.html:1
import (
	"fmt"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vworkspace/vwsprint/SprintList.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwsprint/SprintList.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwsprint/SprintList.html:13
type SprintList struct {
	layout.Basic
	Sprints sprint.Sprints
	Teams   team.Teams
}

//line views/vworkspace/vwsprint/SprintList.html:19
func (p *SprintList) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwsprint/SprintList.html:19
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vworkspace/vwsprint/SprintList.html:21
	components.StreamSVGRefIcon(qw422016, util.KeySprint, ps)
//line views/vworkspace/vwsprint/SprintList.html:21
	qw422016.N().D(len(p.Sprints))
//line views/vworkspace/vwsprint/SprintList.html:21
	qw422016.N().S(` `)
//line views/vworkspace/vwsprint/SprintList.html:21
	qw422016.E().S(util.StringPluralMaybe("Sprint", len(p.Sprints)))
//line views/vworkspace/vwsprint/SprintList.html:21
	qw422016.N().S(`</h3>
    <em>Plan your time and direct your efforts</em>
    <table class="mt expanded">
      <tbody>
`)
//line views/vworkspace/vwsprint/SprintList.html:25
	for _, s := range p.Sprints {
//line views/vworkspace/vwsprint/SprintList.html:25
		qw422016.N().S(`        <tr>
          <td><a href="`)
//line views/vworkspace/vwsprint/SprintList.html:27
		qw422016.E().S(s.PublicWebPath())
//line views/vworkspace/vwsprint/SprintList.html:27
		qw422016.N().S(`">`)
//line views/vworkspace/vwsprint/SprintList.html:27
		qw422016.E().S(s.TitleString())
//line views/vworkspace/vwsprint/SprintList.html:27
		qw422016.N().S(`</a></td>
          <td class="text-align-right">`)
//line views/vworkspace/vwsprint/SprintList.html:28
		components.StreamDisplayTimestamp(qw422016, &s.Created)
//line views/vworkspace/vwsprint/SprintList.html:28
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vworkspace/vwsprint/SprintList.html:30
	}
//line views/vworkspace/vwsprint/SprintList.html:30
	qw422016.N().S(`      </tbody>
    </table>
  </div>
  <div class="card">
    <h3>`)
//line views/vworkspace/vwsprint/SprintList.html:35
	components.StreamSVGRefIcon(qw422016, util.KeySprint, ps)
//line views/vworkspace/vwsprint/SprintList.html:35
	qw422016.N().S(`New Sprint</h3>
    `)
//line views/vworkspace/vwsprint/SprintList.html:36
	StreamSprintForm(qw422016, &sprint.Sprint{}, p.Teams, as, ps)
//line views/vworkspace/vwsprint/SprintList.html:36
	qw422016.N().S(`
  </div>
`)
//line views/vworkspace/vwsprint/SprintList.html:38
}

//line views/vworkspace/vwsprint/SprintList.html:38
func (p *SprintList) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwsprint/SprintList.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwsprint/SprintList.html:38
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/vwsprint/SprintList.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwsprint/SprintList.html:38
}

//line views/vworkspace/vwsprint/SprintList.html:38
func (p *SprintList) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwsprint/SprintList.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwsprint/SprintList.html:38
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/vwsprint/SprintList.html:38
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwsprint/SprintList.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwsprint/SprintList.html:38
	return qs422016
//line views/vworkspace/vwsprint/SprintList.html:38
}

//line views/vworkspace/vwsprint/SprintList.html:40
func StreamSprintForm(qw422016 *qt422016.Writer, s *sprint.Sprint, teams team.Teams, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwsprint/SprintList.html:40
	qw422016.N().S(`
  <form action="" method="post">
    <table class="mt expanded">
      <tbody>
        `)
//line views/vworkspace/vwsprint/SprintList.html:44
	components.StreamTableInput(qw422016, "title", "", "Sprint Title", s.Title, 5, "The name of your sprint")
//line views/vworkspace/vwsprint/SprintList.html:44
	qw422016.N().S(`
        `)
//line views/vworkspace/vwsprint/SprintList.html:45
	components.StreamTableInput(qw422016, "name", "", "Your Name", ps.Username(), 5, "Whatever you prefer to be called")
//line views/vworkspace/vwsprint/SprintList.html:45
	qw422016.N().S(`
        `)
//line views/vworkspace/vwsprint/SprintList.html:46
	components.StreamTableSelect(qw422016, util.KeyTeam, "", "Team", fmt.Sprint(s.TeamID), teams.IDStrings(true), teams.TitleStrings("- no team -"), 5, "The team associated to this sprint")
//line views/vworkspace/vwsprint/SprintList.html:46
	qw422016.N().S(`
        <tr><td colspan="2"><button type="submit">Add Sprint</button></td></tr>
      </tbody>
    </table>
  </form>
`)
//line views/vworkspace/vwsprint/SprintList.html:51
}

//line views/vworkspace/vwsprint/SprintList.html:51
func WriteSprintForm(qq422016 qtio422016.Writer, s *sprint.Sprint, teams team.Teams, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwsprint/SprintList.html:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwsprint/SprintList.html:51
	StreamSprintForm(qw422016, s, teams, as, ps)
//line views/vworkspace/vwsprint/SprintList.html:51
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwsprint/SprintList.html:51
}

//line views/vworkspace/vwsprint/SprintList.html:51
func SprintForm(s *sprint.Sprint, teams team.Teams, as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwsprint/SprintList.html:51
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwsprint/SprintList.html:51
	WriteSprintForm(qb422016, s, teams, as, ps)
//line views/vworkspace/vwsprint/SprintList.html:51
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwsprint/SprintList.html:51
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwsprint/SprintList.html:51
	return qs422016
//line views/vworkspace/vwsprint/SprintList.html:51
}
