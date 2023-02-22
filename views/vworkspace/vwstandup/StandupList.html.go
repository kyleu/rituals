// Code generated by qtc from "StandupList.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwstandup/StandupList.html:1
package vwstandup

//line views/vworkspace/vwstandup/StandupList.html:1
import (
	"fmt"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vworkspace/vwstandup/StandupList.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwstandup/StandupList.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwstandup/StandupList.html:14
type StandupList struct {
	layout.Basic
	Sprints  sprint.Sprints
	Standups standup.Standups
	Teams    team.Teams
}

//line views/vworkspace/vwstandup/StandupList.html:21
func (p *StandupList) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupList.html:21
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vworkspace/vwstandup/StandupList.html:23
	components.StreamSVGRefIcon(qw422016, util.KeyStandup, ps)
//line views/vworkspace/vwstandup/StandupList.html:23
	qw422016.N().D(len(p.Standups))
//line views/vworkspace/vwstandup/StandupList.html:23
	qw422016.N().S(` `)
//line views/vworkspace/vwstandup/StandupList.html:23
	qw422016.E().S(util.StringPluralMaybe("Standup", len(p.Standups)))
//line views/vworkspace/vwstandup/StandupList.html:23
	qw422016.N().S(`</h3>
    <em>Share your progress with your team</em>
    <table class="mt expanded">
      <tbody>
`)
//line views/vworkspace/vwstandup/StandupList.html:27
	for _, u := range p.Standups {
//line views/vworkspace/vwstandup/StandupList.html:27
		qw422016.N().S(`        <tr>
          <td><a href="`)
//line views/vworkspace/vwstandup/StandupList.html:29
		qw422016.E().S(u.PublicWebPath())
//line views/vworkspace/vwstandup/StandupList.html:29
		qw422016.N().S(`">`)
//line views/vworkspace/vwstandup/StandupList.html:29
		qw422016.E().S(u.TitleString())
//line views/vworkspace/vwstandup/StandupList.html:29
		qw422016.N().S(`</a></td>
          <td class="text-align-right">`)
//line views/vworkspace/vwstandup/StandupList.html:30
		components.StreamDisplayTimestamp(qw422016, &u.Created)
//line views/vworkspace/vwstandup/StandupList.html:30
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vworkspace/vwstandup/StandupList.html:32
	}
//line views/vworkspace/vwstandup/StandupList.html:32
	qw422016.N().S(`      </tbody>
    </table>
  </div>
  <div class="card">
    <h3>`)
//line views/vworkspace/vwstandup/StandupList.html:37
	components.StreamSVGRefIcon(qw422016, util.KeyStandup, ps)
//line views/vworkspace/vwstandup/StandupList.html:37
	qw422016.N().S(`New Standup</h3>
    `)
//line views/vworkspace/vwstandup/StandupList.html:38
	StreamStandupForm(qw422016, &standup.Standup{}, p.Teams, p.Sprints, as, ps)
//line views/vworkspace/vwstandup/StandupList.html:38
	qw422016.N().S(`
  </div>
`)
//line views/vworkspace/vwstandup/StandupList.html:40
}

//line views/vworkspace/vwstandup/StandupList.html:40
func (p *StandupList) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupList.html:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupList.html:40
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/vwstandup/StandupList.html:40
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupList.html:40
}

//line views/vworkspace/vwstandup/StandupList.html:40
func (p *StandupList) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwstandup/StandupList.html:40
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupList.html:40
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/vwstandup/StandupList.html:40
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupList.html:40
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupList.html:40
	return qs422016
//line views/vworkspace/vwstandup/StandupList.html:40
}

//line views/vworkspace/vwstandup/StandupList.html:42
func StreamStandupForm(qw422016 *qt422016.Writer, s *standup.Standup, teams team.Teams, sprints sprint.Sprints, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupList.html:42
	qw422016.N().S(`
  <form action="" method="post">
    <table class="mt expanded">
      <tbody>
        `)
//line views/vworkspace/vwstandup/StandupList.html:46
	components.StreamTableInput(qw422016, "title", "", "Standup Title", s.Title, 5, "The name of your standup")
//line views/vworkspace/vwstandup/StandupList.html:46
	qw422016.N().S(`
        `)
//line views/vworkspace/vwstandup/StandupList.html:47
	components.StreamTableInput(qw422016, "name", "", "Your Name", ps.Username(), 5, "Whatever you prefer to be called")
//line views/vworkspace/vwstandup/StandupList.html:47
	qw422016.N().S(`
        `)
//line views/vworkspace/vwstandup/StandupList.html:48
	components.StreamTableSelect(qw422016, util.KeyTeam, "", "Team", fmt.Sprint(s.TeamID), teams.IDStrings(true), teams.TitleStrings("- no team -"), 5, "The team associated to this standup")
//line views/vworkspace/vwstandup/StandupList.html:48
	qw422016.N().S(`
        `)
//line views/vworkspace/vwstandup/StandupList.html:49
	components.StreamTableSelect(qw422016, util.KeySprint, "", "Sprint", fmt.Sprint(s.SprintID), sprints.IDStrings(true), sprints.TitleStrings("- no sprint -"), 5, "The sprint associated to this standup")
//line views/vworkspace/vwstandup/StandupList.html:49
	qw422016.N().S(`
        <tr><td colspan="2"><button type="submit">Add Standup</button></td></tr>
      </tbody>
    </table>
  </form>
`)
//line views/vworkspace/vwstandup/StandupList.html:54
}

//line views/vworkspace/vwstandup/StandupList.html:54
func WriteStandupForm(qq422016 qtio422016.Writer, s *standup.Standup, teams team.Teams, sprints sprint.Sprints, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwstandup/StandupList.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwstandup/StandupList.html:54
	StreamStandupForm(qw422016, s, teams, sprints, as, ps)
//line views/vworkspace/vwstandup/StandupList.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwstandup/StandupList.html:54
}

//line views/vworkspace/vwstandup/StandupList.html:54
func StandupForm(s *standup.Standup, teams team.Teams, sprints sprint.Sprints, as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwstandup/StandupList.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwstandup/StandupList.html:54
	WriteStandupForm(qb422016, s, teams, sprints, as, ps)
//line views/vworkspace/vwstandup/StandupList.html:54
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwstandup/StandupList.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwstandup/StandupList.html:54
	return qs422016
//line views/vworkspace/vwstandup/StandupList.html:54
}
