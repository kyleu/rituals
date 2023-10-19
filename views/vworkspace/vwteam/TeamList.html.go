// Code generated by qtc from "TeamList.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwteam/TeamList.html:1
package vwteam

//line views/vworkspace/vwteam/TeamList.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/member"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vworkspace/vwutil"
)

//line views/vworkspace/vwteam/TeamList.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwteam/TeamList.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwteam/TeamList.html:14
type TeamList struct {
	layout.Basic
	Teams team.Teams
}

//line views/vworkspace/vwteam/TeamList.html:19
func (p *TeamList) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamList.html:19
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vworkspace/vwteam/TeamList.html:21
	components.StreamSVGRefIcon(qw422016, util.KeyTeam, ps)
//line views/vworkspace/vwteam/TeamList.html:21
	qw422016.E().S(util.StringPlural(len(p.Teams), "Team"))
//line views/vworkspace/vwteam/TeamList.html:21
	qw422016.N().S(`</h3>
    <em>`)
//line views/vworkspace/vwteam/TeamList.html:22
	qw422016.E().S(util.KeyTeamDesc)
//line views/vworkspace/vwteam/TeamList.html:22
	qw422016.N().S(`</em>
    <table class="mt expanded">
      <tbody>
`)
//line views/vworkspace/vwteam/TeamList.html:25
	for _, t := range p.Teams {
//line views/vworkspace/vwteam/TeamList.html:25
		qw422016.N().S(`        <tr>
          <td><a href="`)
//line views/vworkspace/vwteam/TeamList.html:27
		qw422016.E().S(t.PublicWebPath())
//line views/vworkspace/vwteam/TeamList.html:27
		qw422016.N().S(`">`)
//line views/vworkspace/vwteam/TeamList.html:27
		components.StreamSVGRef(qw422016, t.IconSafe(), 16, 16, "icon", ps)
//line views/vworkspace/vwteam/TeamList.html:27
		qw422016.E().S(t.TitleString())
//line views/vworkspace/vwteam/TeamList.html:27
		qw422016.N().S(`</a></td>
          <td class="text-align-right">`)
//line views/vworkspace/vwteam/TeamList.html:28
		qw422016.E().S(t.Status.String())
//line views/vworkspace/vwteam/TeamList.html:28
		qw422016.N().S(`</td>
          <td class="shrink">`)
//line views/vworkspace/vwteam/TeamList.html:29
		components.StreamDisplayTimestamp(qw422016, &t.Created)
//line views/vworkspace/vwteam/TeamList.html:29
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vworkspace/vwteam/TeamList.html:31
	}
//line views/vworkspace/vwteam/TeamList.html:31
	qw422016.N().S(`      </tbody>
    </table>
  </div>
  <div class="card">
    <h3>`)
//line views/vworkspace/vwteam/TeamList.html:36
	components.StreamSVGRefIcon(qw422016, util.KeyTeam, ps)
//line views/vworkspace/vwteam/TeamList.html:36
	qw422016.N().S(`New Team</h3>
    `)
//line views/vworkspace/vwteam/TeamList.html:37
	StreamTeamForm(qw422016, &team.Team{}, as, ps)
//line views/vworkspace/vwteam/TeamList.html:37
	qw422016.N().S(`
  </div>
`)
//line views/vworkspace/vwteam/TeamList.html:39
}

//line views/vworkspace/vwteam/TeamList.html:39
func (p *TeamList) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamList.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwteam/TeamList.html:39
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/vwteam/TeamList.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwteam/TeamList.html:39
}

//line views/vworkspace/vwteam/TeamList.html:39
func (p *TeamList) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwteam/TeamList.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwteam/TeamList.html:39
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/vwteam/TeamList.html:39
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwteam/TeamList.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwteam/TeamList.html:39
	return qs422016
//line views/vworkspace/vwteam/TeamList.html:39
}

//line views/vworkspace/vwteam/TeamList.html:41
func StreamTeamForm(qw422016 *qt422016.Writer, m *team.Team, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamList.html:41
	qw422016.N().S(`
  <form action="" method="post">
    <table class="mt expanded">
      <tbody>
        `)
//line views/vworkspace/vwteam/TeamList.html:45
	components.StreamTableInput(qw422016, "title", "", "Team Title", m.Title, 5, "The name of your team")
//line views/vworkspace/vwteam/TeamList.html:45
	qw422016.N().S(`
        `)
//line views/vworkspace/vwteam/TeamList.html:46
	components.StreamTableInput(qw422016, "name", "", "Your Name", ps.Username(), 5, "Whatever you prefer to be called")
//line views/vworkspace/vwteam/TeamList.html:46
	qw422016.N().S(`
        <tr><td colspan="2"><button type="submit">Add Team</button></td></tr>
      </tbody>
    </table>
  </form>
`)
//line views/vworkspace/vwteam/TeamList.html:51
}

//line views/vworkspace/vwteam/TeamList.html:51
func WriteTeamForm(qq422016 qtio422016.Writer, m *team.Team, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamList.html:51
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwteam/TeamList.html:51
	StreamTeamForm(qw422016, m, as, ps)
//line views/vworkspace/vwteam/TeamList.html:51
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwteam/TeamList.html:51
}

//line views/vworkspace/vwteam/TeamList.html:51
func TeamForm(m *team.Team, as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwteam/TeamList.html:51
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwteam/TeamList.html:51
	WriteTeamForm(qb422016, m, as, ps)
//line views/vworkspace/vwteam/TeamList.html:51
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwteam/TeamList.html:51
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwteam/TeamList.html:51
	return qs422016
//line views/vworkspace/vwteam/TeamList.html:51
}

//line views/vworkspace/vwteam/TeamList.html:53
func StreamTeamListTable(qw422016 *qt422016.Writer, teams team.Teams, showComments bool, comments comment.Comments, members member.Members, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamList.html:53
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vworkspace/vwteam/TeamList.html:55
	vwutil.StreamEditWorkspaceForm(qw422016, util.KeyTeam, nil, nil, "New Team")
//line views/vworkspace/vwteam/TeamList.html:55
	qw422016.N().S(`</div>
    <h3 title="`)
//line views/vworkspace/vwteam/TeamList.html:56
	qw422016.E().S(util.KeyTeamDesc)
//line views/vworkspace/vwteam/TeamList.html:56
	qw422016.N().S(`">`)
//line views/vworkspace/vwteam/TeamList.html:56
	components.StreamSVGRefIcon(qw422016, util.KeyTeam, ps)
//line views/vworkspace/vwteam/TeamList.html:56
	qw422016.N().S(`Teams</h3>
`)
//line views/vworkspace/vwteam/TeamList.html:57
	if len(teams) == 0 {
//line views/vworkspace/vwteam/TeamList.html:57
		qw422016.N().S(`    <div class="mt"><em>no teams</em></div>
`)
//line views/vworkspace/vwteam/TeamList.html:59
	} else {
//line views/vworkspace/vwteam/TeamList.html:59
		qw422016.N().S(`    <table class="mt expanded">
      <tbody>
`)
//line views/vworkspace/vwteam/TeamList.html:62
		for _, x := range teams {
//line views/vworkspace/vwteam/TeamList.html:62
			qw422016.N().S(`        <tr>
          <td>
`)
//line views/vworkspace/vwteam/TeamList.html:65
			if showComments {
//line views/vworkspace/vwteam/TeamList.html:65
				qw422016.N().S(`            <div class="right">
              `)
//line views/vworkspace/vwteam/TeamList.html:67
				vwutil.StreamComments(qw422016, enum.ModelServiceTeam, x.ID, x.TitleString(), comments, members, "member-icon", ps)
//line views/vworkspace/vwteam/TeamList.html:67
				qw422016.N().S(`
            </div>
`)
//line views/vworkspace/vwteam/TeamList.html:69
			}
//line views/vworkspace/vwteam/TeamList.html:69
			qw422016.N().S(`            <a href="`)
//line views/vworkspace/vwteam/TeamList.html:70
			qw422016.E().S(x.PublicWebPath())
//line views/vworkspace/vwteam/TeamList.html:70
			qw422016.N().S(`"><div>
              <span>`)
//line views/vworkspace/vwteam/TeamList.html:71
			components.StreamSVGRef(qw422016, x.IconSafe(), 16, 16, "icon", ps)
//line views/vworkspace/vwteam/TeamList.html:71
			qw422016.N().S(`</span><span>`)
//line views/vworkspace/vwteam/TeamList.html:71
			qw422016.E().S(x.TitleString())
//line views/vworkspace/vwteam/TeamList.html:71
			qw422016.N().S(`</span>
            </div></a>
          </td>
        </tr>
`)
//line views/vworkspace/vwteam/TeamList.html:75
		}
//line views/vworkspace/vwteam/TeamList.html:75
		qw422016.N().S(`      </tbody>
    </table>
`)
//line views/vworkspace/vwteam/TeamList.html:78
	}
//line views/vworkspace/vwteam/TeamList.html:78
	qw422016.N().S(`  </div>
`)
//line views/vworkspace/vwteam/TeamList.html:80
}

//line views/vworkspace/vwteam/TeamList.html:80
func WriteTeamListTable(qq422016 qtio422016.Writer, teams team.Teams, showComments bool, comments comment.Comments, members member.Members, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamList.html:80
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwteam/TeamList.html:80
	StreamTeamListTable(qw422016, teams, showComments, comments, members, ps)
//line views/vworkspace/vwteam/TeamList.html:80
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwteam/TeamList.html:80
}

//line views/vworkspace/vwteam/TeamList.html:80
func TeamListTable(teams team.Teams, showComments bool, comments comment.Comments, members member.Members, ps *cutil.PageState) string {
//line views/vworkspace/vwteam/TeamList.html:80
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwteam/TeamList.html:80
	WriteTeamListTable(qb422016, teams, showComments, comments, members, ps)
//line views/vworkspace/vwteam/TeamList.html:80
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwteam/TeamList.html:80
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwteam/TeamList.html:80
	return qs422016
//line views/vworkspace/vwteam/TeamList.html:80
}
