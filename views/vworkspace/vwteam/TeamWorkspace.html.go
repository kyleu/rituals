// Code generated by qtc from "TeamWorkspace.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwteam/TeamWorkspace.html:1
package vwteam

//line views/vworkspace/vwteam/TeamWorkspace.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vworkspace/vwestimate"
	"github.com/kyleu/rituals/views/vworkspace/vwretro"
	"github.com/kyleu/rituals/views/vworkspace/vwsprint"
	"github.com/kyleu/rituals/views/vworkspace/vwstandup"
	"github.com/kyleu/rituals/views/vworkspace/vwutil"
)

//line views/vworkspace/vwteam/TeamWorkspace.html:19
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwteam/TeamWorkspace.html:19
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwteam/TeamWorkspace.html:19
type TeamWorkspace struct {
	layout.Basic
	FullTeam *workspace.FullTeam
}

//line views/vworkspace/vwteam/TeamWorkspace.html:24
func (p *TeamWorkspace) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamWorkspace.html:24
	qw422016.N().S(`
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:26
	w := p.FullTeam
	t := w.Team
	self, others, _ := w.UtilMembers.Split(ps.Profile.ID)

//line views/vworkspace/vwteam/TeamWorkspace.html:29
	qw422016.N().S(`  <div style="display: flex; flex-wrap: wrap;">
    <div id="panel-summary">
      <div class="card">
        <div class="right">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:33
	vwutil.StreamComments(qw422016, enum.ModelServiceTeam, t.ID, t.TitleString(), w.Comments, w.UtilMembers, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:33
	qw422016.N().S(`</div>
        <h3><a href="#modal-team-config">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:34
	components.StreamSVGRefIcon(qw422016, util.KeyTeam, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:34
	qw422016.E().S(t.TitleString())
//line views/vworkspace/vwteam/TeamWorkspace.html:34
	qw422016.N().S(`</a></h3>
        `)
//line views/vworkspace/vwteam/TeamWorkspace.html:35
	vwutil.StreamBanner(qw422016, nil, nil, util.KeyTeam)
//line views/vworkspace/vwteam/TeamWorkspace.html:35
	qw422016.N().S(`
        `)
//line views/vworkspace/vwteam/TeamWorkspace.html:36
	StreamTeamWorkspaceModal(qw422016, t, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:36
	qw422016.N().S(`
      </div>
    </div>
    <div id="panel-detail">
      `)
//line views/vworkspace/vwteam/TeamWorkspace.html:40
	vwsprint.StreamSprintWorkspaceList(qw422016, w.Sprints, &t.ID, true, w.Comments, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:40
	qw422016.N().S(`
      `)
//line views/vworkspace/vwteam/TeamWorkspace.html:41
	vwestimate.StreamEstimateWorkspaceList(qw422016, w.Estimates, &t.ID, nil, true, w.Comments, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:41
	qw422016.N().S(`
      `)
//line views/vworkspace/vwteam/TeamWorkspace.html:42
	vwstandup.StreamStandupWorkspaceList(qw422016, w.Standups, &t.ID, nil, true, w.Comments, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:42
	qw422016.N().S(`
      `)
//line views/vworkspace/vwteam/TeamWorkspace.html:43
	vwretro.StreamRetroWorkspaceList(qw422016, w.Retros, &t.ID, nil, true, w.Comments, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:43
	qw422016.N().S(`
    </div>
    <div id="panel-self">
      <div class="card">
        <span id="self-id" style="display: none;">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:47
	qw422016.E().S(self.UserID.String())
//line views/vworkspace/vwteam/TeamWorkspace.html:47
	qw422016.N().S(`</span>
        <h3><a href="#modal-self">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:48
	components.StreamSVGRefIcon(qw422016, `profile`, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:48
	qw422016.N().S(`<span id="self-name">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:48
	qw422016.E().S(self.Name)
//line views/vworkspace/vwteam/TeamWorkspace.html:48
	qw422016.N().S(`</span></a></h3>
        <em id="self-role">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:49
	qw422016.E().S(string(self.Role))
//line views/vworkspace/vwteam/TeamWorkspace.html:49
	qw422016.N().S(`</em>
      </div>
      `)
//line views/vworkspace/vwteam/TeamWorkspace.html:51
	vwutil.StreamSelfModal(qw422016, self.Name, self.Picture, self.Role, t.PublicWebPath())
//line views/vworkspace/vwteam/TeamWorkspace.html:51
	qw422016.N().S(`
    </div>
    <div id="panel-members">
      <div class="card">
        <h3><a href="#modal-invite">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:55
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:55
	qw422016.N().S(`Members</a></h3>
        <table class="mt expanded">
          <tbody>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:58
	for _, m := range others {
//line views/vworkspace/vwteam/TeamWorkspace.html:58
		qw422016.N().S(`            `)
//line views/vworkspace/vwteam/TeamWorkspace.html:59
		vwutil.StreamMemberRow(qw422016, m, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:59
		qw422016.N().S(`
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:60
	}
//line views/vworkspace/vwteam/TeamWorkspace.html:60
	qw422016.N().S(`          </tbody>
        </table>
        <div id="member-modals">
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:64
	for _, m := range others {
//line views/vworkspace/vwteam/TeamWorkspace.html:64
		qw422016.N().S(`          `)
//line views/vworkspace/vwteam/TeamWorkspace.html:65
		vwutil.StreamMemberModal(qw422016, m, t.PublicWebPath())
//line views/vworkspace/vwteam/TeamWorkspace.html:65
		qw422016.N().S(`
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:66
	}
//line views/vworkspace/vwteam/TeamWorkspace.html:66
	qw422016.N().S(`        </div>
      </div>
      `)
//line views/vworkspace/vwteam/TeamWorkspace.html:69
	vwutil.StreamInviteModal(qw422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:69
	qw422016.N().S(`
    </div>
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      initWorkspace("`)
//line views/vworkspace/vwteam/TeamWorkspace.html:74
	qw422016.E().S(util.KeyTeam)
//line views/vworkspace/vwteam/TeamWorkspace.html:74
	qw422016.N().S(`", "`)
//line views/vworkspace/vwteam/TeamWorkspace.html:74
	qw422016.E().S(t.ID.String())
//line views/vworkspace/vwteam/TeamWorkspace.html:74
	qw422016.N().S(`");
    });
  </script>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:77
}

//line views/vworkspace/vwteam/TeamWorkspace.html:77
func (p *TeamWorkspace) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamWorkspace.html:77
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:77
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:77
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:77
}

//line views/vworkspace/vwteam/TeamWorkspace.html:77
func (p *TeamWorkspace) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwteam/TeamWorkspace.html:77
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwteam/TeamWorkspace.html:77
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:77
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwteam/TeamWorkspace.html:77
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:77
	return qs422016
//line views/vworkspace/vwteam/TeamWorkspace.html:77
}

//line views/vworkspace/vwteam/TeamWorkspace.html:79
func StreamTeamWorkspaceModal(qw422016 *qt422016.Writer, t *team.Team, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamWorkspace.html:79
	qw422016.N().S(`
  <div id="modal-team-config" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Team</h2>
      </div>
      <div class="modal-body">
        <form action="`)
//line views/vworkspace/vwteam/TeamWorkspace.html:88
	qw422016.E().S(t.PublicWebPath())
//line views/vworkspace/vwteam/TeamWorkspace.html:88
	qw422016.N().S(`" method="post" class="expanded">
          <input type="hidden" name="action" value="`)
//line views/vworkspace/vwteam/TeamWorkspace.html:89
	qw422016.E().S(string(action.ActUpdate))
//line views/vworkspace/vwteam/TeamWorkspace.html:89
	qw422016.N().S(`" />
          `)
//line views/vworkspace/vwteam/TeamWorkspace.html:90
	components.StreamFormVerticalInput(qw422016, "title", "", "Title", t.TitleString(), 5, "The name of your team")
//line views/vworkspace/vwteam/TeamWorkspace.html:90
	qw422016.N().S(`
          `)
//line views/vworkspace/vwteam/TeamWorkspace.html:91
	components.StreamFormVerticalIconPicker(qw422016, "icon", "Icon", t.IconSafe(), ps, 5)
//line views/vworkspace/vwteam/TeamWorkspace.html:91
	qw422016.N().S(`
          <em class="title">Permissions</em>
          <div class="mt">Control access to this team by <a href="/profile">signing in</a></div>
          <hr />
          <div class="right"><button type="submit">Save</button></div>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:100
}

//line views/vworkspace/vwteam/TeamWorkspace.html:100
func WriteTeamWorkspaceModal(qq422016 qtio422016.Writer, t *team.Team, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamWorkspace.html:100
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:100
	StreamTeamWorkspaceModal(qw422016, t, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:100
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:100
}

//line views/vworkspace/vwteam/TeamWorkspace.html:100
func TeamWorkspaceModal(t *team.Team, ps *cutil.PageState) string {
//line views/vworkspace/vwteam/TeamWorkspace.html:100
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwteam/TeamWorkspace.html:100
	WriteTeamWorkspaceModal(qb422016, t, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:100
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwteam/TeamWorkspace.html:100
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:100
	return qs422016
//line views/vworkspace/vwteam/TeamWorkspace.html:100
}

//line views/vworkspace/vwteam/TeamWorkspace.html:102
func StreamTeamWorkspaceList(qw422016 *qt422016.Writer, teams team.Teams, showComments bool, comments comment.Comments, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamWorkspace.html:102
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:104
	vwutil.StreamEditWorkspaceForm(qw422016, util.KeyTeam, nil, nil, "New Team")
//line views/vworkspace/vwteam/TeamWorkspace.html:104
	qw422016.N().S(`</div>
    <h3>`)
//line views/vworkspace/vwteam/TeamWorkspace.html:105
	components.StreamSVGRefIcon(qw422016, util.KeyTeam, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:105
	qw422016.N().S(`Teams</h3>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:106
	if len(teams) == 0 {
//line views/vworkspace/vwteam/TeamWorkspace.html:106
		qw422016.N().S(`    <div class="mt"><em>no teams</em></div>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:108
	} else {
//line views/vworkspace/vwteam/TeamWorkspace.html:108
		qw422016.N().S(`    <table class="mt expanded">
      <tbody>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:111
		for _, x := range teams {
//line views/vworkspace/vwteam/TeamWorkspace.html:111
			qw422016.N().S(`        <tr>
          <td>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:114
			if showComments {
//line views/vworkspace/vwteam/TeamWorkspace.html:114
				qw422016.N().S(`            <div class="right">
              `)
//line views/vworkspace/vwteam/TeamWorkspace.html:116
				vwutil.StreamComments(qw422016, enum.ModelServiceTeam, x.ID, x.TitleString(), comments, nil, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:116
				qw422016.N().S(`
            </div>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:118
			}
//line views/vworkspace/vwteam/TeamWorkspace.html:118
			qw422016.N().S(`            <a href="`)
//line views/vworkspace/vwteam/TeamWorkspace.html:119
			qw422016.E().S(x.PublicWebPath())
//line views/vworkspace/vwteam/TeamWorkspace.html:119
			qw422016.N().S(`">`)
//line views/vworkspace/vwteam/TeamWorkspace.html:119
			qw422016.E().S(x.TitleString())
//line views/vworkspace/vwteam/TeamWorkspace.html:119
			qw422016.N().S(`</a>
          </td>
        </tr>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:122
		}
//line views/vworkspace/vwteam/TeamWorkspace.html:122
		qw422016.N().S(`      </tbody>
    </table>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:125
	}
//line views/vworkspace/vwteam/TeamWorkspace.html:125
	qw422016.N().S(`  </div>
`)
//line views/vworkspace/vwteam/TeamWorkspace.html:127
}

//line views/vworkspace/vwteam/TeamWorkspace.html:127
func WriteTeamWorkspaceList(qq422016 qtio422016.Writer, teams team.Teams, showComments bool, comments comment.Comments, ps *cutil.PageState) {
//line views/vworkspace/vwteam/TeamWorkspace.html:127
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:127
	StreamTeamWorkspaceList(qw422016, teams, showComments, comments, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:127
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:127
}

//line views/vworkspace/vwteam/TeamWorkspace.html:127
func TeamWorkspaceList(teams team.Teams, showComments bool, comments comment.Comments, ps *cutil.PageState) string {
//line views/vworkspace/vwteam/TeamWorkspace.html:127
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwteam/TeamWorkspace.html:127
	WriteTeamWorkspaceList(qb422016, teams, showComments, comments, ps)
//line views/vworkspace/vwteam/TeamWorkspace.html:127
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwteam/TeamWorkspace.html:127
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwteam/TeamWorkspace.html:127
	return qs422016
//line views/vworkspace/vwteam/TeamWorkspace.html:127
}