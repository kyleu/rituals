// Code generated by qtc from "RetroList.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwretro/RetroList.html:1
package vwretro

//line views/vworkspace/vwretro/RetroList.html:1
import (
	"fmt"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vworkspace/vwretro/RetroList.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwretro/RetroList.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwretro/RetroList.html:14
type RetroList struct {
	layout.Basic
	Retros  retro.Retros
	Sprints sprint.Sprints
	Teams   team.Teams
}

//line views/vworkspace/vwretro/RetroList.html:21
func (p *RetroList) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroList.html:21
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vworkspace/vwretro/RetroList.html:23
	components.StreamSVGRefIcon(qw422016, `retro`, ps)
//line views/vworkspace/vwretro/RetroList.html:23
	qw422016.N().D(len(p.Retros))
//line views/vworkspace/vwretro/RetroList.html:23
	qw422016.N().S(` `)
//line views/vworkspace/vwretro/RetroList.html:23
	qw422016.E().S(util.StringPluralMaybe("Retro", len(p.Retros)))
//line views/vworkspace/vwretro/RetroList.html:23
	qw422016.N().S(`</h3>
    <em>Discover improvements and praise for your work</em>
    <table class="mt expanded">
      <tbody>
`)
//line views/vworkspace/vwretro/RetroList.html:27
	for _, r := range p.Retros {
//line views/vworkspace/vwretro/RetroList.html:27
		qw422016.N().S(`        <tr>
          <td><a href="`)
//line views/vworkspace/vwretro/RetroList.html:29
		qw422016.E().S(r.PublicWebPath())
//line views/vworkspace/vwretro/RetroList.html:29
		qw422016.N().S(`">`)
//line views/vworkspace/vwretro/RetroList.html:29
		qw422016.E().S(r.TitleString())
//line views/vworkspace/vwretro/RetroList.html:29
		qw422016.N().S(`</a></td>
          <td style="text-align: right;">`)
//line views/vworkspace/vwretro/RetroList.html:30
		components.StreamDisplayTimestamp(qw422016, &r.Created)
//line views/vworkspace/vwretro/RetroList.html:30
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vworkspace/vwretro/RetroList.html:32
	}
//line views/vworkspace/vwretro/RetroList.html:32
	qw422016.N().S(`      </tbody>
    </table>
  </div>
  <div class="card">
    <h3>`)
//line views/vworkspace/vwretro/RetroList.html:37
	components.StreamSVGRefIcon(qw422016, `retro`, ps)
//line views/vworkspace/vwretro/RetroList.html:37
	qw422016.N().S(`New Retro</h3>
    `)
//line views/vworkspace/vwretro/RetroList.html:38
	StreamRetroForm(qw422016, &retro.Retro{}, p.Teams, p.Sprints, as, ps)
//line views/vworkspace/vwretro/RetroList.html:38
	qw422016.N().S(`
  </div>
`)
//line views/vworkspace/vwretro/RetroList.html:40
}

//line views/vworkspace/vwretro/RetroList.html:40
func (p *RetroList) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroList.html:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroList.html:40
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/vwretro/RetroList.html:40
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroList.html:40
}

//line views/vworkspace/vwretro/RetroList.html:40
func (p *RetroList) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwretro/RetroList.html:40
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroList.html:40
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/vwretro/RetroList.html:40
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroList.html:40
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroList.html:40
	return qs422016
//line views/vworkspace/vwretro/RetroList.html:40
}

//line views/vworkspace/vwretro/RetroList.html:42
func StreamRetroForm(qw422016 *qt422016.Writer, s *retro.Retro, teams team.Teams, sprints sprint.Sprints, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroList.html:42
	qw422016.N().S(`
  <form action="" method="post">
    <table class="mt expanded">
      <tbody>
        `)
//line views/vworkspace/vwretro/RetroList.html:46
	components.StreamTableInput(qw422016, "title", "Team Title", s.Title, 5, "The name of your retro")
//line views/vworkspace/vwretro/RetroList.html:46
	qw422016.N().S(`
        `)
//line views/vworkspace/vwretro/RetroList.html:47
	components.StreamTableInput(qw422016, "name", "Your Name", ps.Profile.Name, 5, "Whatever you prefer to be called")
//line views/vworkspace/vwretro/RetroList.html:47
	qw422016.N().S(`
        `)
//line views/vworkspace/vwretro/RetroList.html:48
	components.StreamTableSelect(qw422016, util.KeyTeam, "Team", fmt.Sprint(s.TeamID), teams.IDStrings(true), teams.TitleStrings("- no team -"), 5, "The team associated to this retro")
//line views/vworkspace/vwretro/RetroList.html:48
	qw422016.N().S(`
        `)
//line views/vworkspace/vwretro/RetroList.html:49
	components.StreamTableSelect(qw422016, util.KeySprint, "Sprint", fmt.Sprint(s.SprintID), sprints.IDStrings(true), sprints.TitleStrings("- no sprint -"), 5, "The sprint associated to this retro")
//line views/vworkspace/vwretro/RetroList.html:49
	qw422016.N().S(`
        <tr><td colspan="2"><button type="submit">Add Retro</button></td></tr>
      </tbody>
    </table>
  </form>
`)
//line views/vworkspace/vwretro/RetroList.html:54
}

//line views/vworkspace/vwretro/RetroList.html:54
func WriteRetroForm(qq422016 qtio422016.Writer, s *retro.Retro, teams team.Teams, sprints sprint.Sprints, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroList.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroList.html:54
	StreamRetroForm(qw422016, s, teams, sprints, as, ps)
//line views/vworkspace/vwretro/RetroList.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroList.html:54
}

//line views/vworkspace/vwretro/RetroList.html:54
func RetroForm(s *retro.Retro, teams team.Teams, sprints sprint.Sprints, as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwretro/RetroList.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroList.html:54
	WriteRetroForm(qb422016, s, teams, sprints, as, ps)
//line views/vworkspace/vwretro/RetroList.html:54
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroList.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroList.html:54
	return qs422016
//line views/vworkspace/vwretro/RetroList.html:54
}
