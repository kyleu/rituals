// Code generated by qtc from "StandupWorkspace.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/StandupWorkspace.html:1
package vworkspace

//line views/vworkspace/StandupWorkspace.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vworkspace/StandupWorkspace.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/StandupWorkspace.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/StandupWorkspace.html:12
type StandupWorkspace struct {
	layout.Basic
	Standup *workspace.FullStandup
	Teams   team.Teams
	Sprints sprint.Sprints
}

//line views/vworkspace/StandupWorkspace.html:19
func (p *StandupWorkspace) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/StandupWorkspace.html:19
	qw422016.N().S(`
`)
//line views/vworkspace/StandupWorkspace.html:20
	w := p.Standup

//line views/vworkspace/StandupWorkspace.html:21
	t := w.Standup

//line views/vworkspace/StandupWorkspace.html:23
	self, others, err := w.Members.Split(ps.Profile.ID)
	if err != nil {
		panic(err)
	}

//line views/vworkspace/StandupWorkspace.html:27
	qw422016.N().S(`  <div style="display: flex; flex-wrap: wrap;">
    <div id="panel-summary">
      <div class="card">
        <div class="right"><a href="#modal-standup"><button type="button">JSON</button></a></div>
        <a href="#modal-standup-config"><h3>`)
//line views/vworkspace/StandupWorkspace.html:32
	components.StreamSVGRefIcon(qw422016, `standup`, ps)
//line views/vworkspace/StandupWorkspace.html:32
	qw422016.E().S(t.TitleString())
//line views/vworkspace/StandupWorkspace.html:32
	qw422016.N().S(`</h3></a>
        `)
//line views/vworkspace/StandupWorkspace.html:33
	StreamBanner(qw422016, w.Team, w.Sprint, "standup")
//line views/vworkspace/StandupWorkspace.html:33
	qw422016.N().S(`
        `)
//line views/vworkspace/StandupWorkspace.html:34
	StreamStandupWorkspaceModal(qw422016, w, p.Teams, p.Sprints)
//line views/vworkspace/StandupWorkspace.html:34
	qw422016.N().S(`
      </div>
    </div>
    <div id="panel-detail">
      <div class="card">
        <h3>`)
//line views/vworkspace/StandupWorkspace.html:39
	components.StreamSVGRefIcon(qw422016, `file-alt`, ps)
//line views/vworkspace/StandupWorkspace.html:39
	qw422016.N().S(`Reports</h3>
        <table class="mt expanded">
          <tbody>
`)
//line views/vworkspace/StandupWorkspace.html:42
	for _, r := range w.Reports {
//line views/vworkspace/StandupWorkspace.html:42
		qw422016.N().S(`            <tr>
              <td><a href="#modal-report-`)
//line views/vworkspace/StandupWorkspace.html:44
		qw422016.E().S(r.ID.String())
//line views/vworkspace/StandupWorkspace.html:44
		qw422016.N().S(`">`)
//line views/vworkspace/StandupWorkspace.html:44
		qw422016.E().S(r.TitleString())
//line views/vworkspace/StandupWorkspace.html:44
		qw422016.N().S(`</a></td>
              <td class="shrink">0</td>
            </tr>
`)
//line views/vworkspace/StandupWorkspace.html:47
	}
//line views/vworkspace/StandupWorkspace.html:47
	qw422016.N().S(`          </tbody>
        </table>
`)
//line views/vworkspace/StandupWorkspace.html:50
	for _, r := range w.Reports {
//line views/vworkspace/StandupWorkspace.html:50
		qw422016.N().S(`        <div id="modal-report-`)
//line views/vworkspace/StandupWorkspace.html:51
		qw422016.E().S(r.ID.String())
//line views/vworkspace/StandupWorkspace.html:51
		qw422016.N().S(`" class="modal" style="display: none;">
          <a class="backdrop" href="#"></a>
          <div class="modal-content">
            <div class="modal-header">
              <a href="#" class="modal-close">×</a>
              <h2>`)
//line views/vworkspace/StandupWorkspace.html:56
		qw422016.E().S(r.TitleString())
//line views/vworkspace/StandupWorkspace.html:56
		qw422016.N().S(`</h2>
            </div>
            <div class="modal-body">
              TODO
            </div>
          </div>
        </div>
`)
//line views/vworkspace/StandupWorkspace.html:63
	}
//line views/vworkspace/StandupWorkspace.html:63
	qw422016.N().S(`      </div>
    </div>
    <div id="panel-self">
      <div class="card">
        <a href="#modal-self"><h3>`)
//line views/vworkspace/StandupWorkspace.html:68
	components.StreamSVGRefIcon(qw422016, `profile`, ps)
//line views/vworkspace/StandupWorkspace.html:68
	qw422016.E().S(self.Name)
//line views/vworkspace/StandupWorkspace.html:68
	qw422016.N().S(`</h3></a>
        <em>`)
//line views/vworkspace/StandupWorkspace.html:69
	qw422016.E().S(string(self.Role))
//line views/vworkspace/StandupWorkspace.html:69
	qw422016.N().S(`</em>
      </div>
      `)
//line views/vworkspace/StandupWorkspace.html:71
	StreamSelfModal(qw422016, self.Name, self.Picture, self.Role, "/standup/"+w.Standup.Slug)
//line views/vworkspace/StandupWorkspace.html:71
	qw422016.N().S(`
    </div>
    <div id="panel-members">
      <div class="card">
        <a href="#modal-invite"><h3>`)
//line views/vworkspace/StandupWorkspace.html:75
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vworkspace/StandupWorkspace.html:75
	qw422016.N().S(`Members</h3></a>
        <table class="mt expanded">
          <tbody>
`)
//line views/vworkspace/StandupWorkspace.html:78
	for _, m := range others {
//line views/vworkspace/StandupWorkspace.html:78
		qw422016.N().S(`            `)
//line views/vworkspace/StandupWorkspace.html:79
		StreamMemberRow(qw422016, m.UserID, m.Name, m.Picture, m.Role, m.Updated, ps)
//line views/vworkspace/StandupWorkspace.html:79
		qw422016.N().S(`
`)
//line views/vworkspace/StandupWorkspace.html:80
	}
//line views/vworkspace/StandupWorkspace.html:80
	qw422016.N().S(`          </tbody>
        </table>
`)
//line views/vworkspace/StandupWorkspace.html:83
	for _, m := range others {
//line views/vworkspace/StandupWorkspace.html:83
		qw422016.N().S(`        `)
//line views/vworkspace/StandupWorkspace.html:84
		StreamMemberModal(qw422016, m.UserID, m.Name, m.Picture, m.Role, m.Updated, "/standup/"+w.Standup.Slug)
//line views/vworkspace/StandupWorkspace.html:84
		qw422016.N().S(`
`)
//line views/vworkspace/StandupWorkspace.html:85
	}
//line views/vworkspace/StandupWorkspace.html:85
	qw422016.N().S(`      </div>
      `)
//line views/vworkspace/StandupWorkspace.html:87
	StreamInviteModal(qw422016)
//line views/vworkspace/StandupWorkspace.html:87
	qw422016.N().S(`
    </div>
  </div>
  `)
//line views/vworkspace/StandupWorkspace.html:90
	components.StreamJSONModal(qw422016, "standup", "Standup JSON", w, 1)
//line views/vworkspace/StandupWorkspace.html:90
	qw422016.N().S(`
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      const standup = `)
//line views/vworkspace/StandupWorkspace.html:93
	qw422016.N().S(util.ToJSONCompact(t))
//line views/vworkspace/StandupWorkspace.html:93
	qw422016.N().S(`;
      const members = `)
//line views/vworkspace/StandupWorkspace.html:94
	qw422016.N().S(util.ToJSONCompact(w.Members))
//line views/vworkspace/StandupWorkspace.html:94
	qw422016.N().S(`;
      const permissions = `)
//line views/vworkspace/StandupWorkspace.html:95
	qw422016.N().S(util.ToJSONCompact(w.Permissions))
//line views/vworkspace/StandupWorkspace.html:95
	qw422016.N().S(`;
      rituals.initWorkspace("standup", standup, members, permissions);
    });
  </script>
`)
//line views/vworkspace/StandupWorkspace.html:99
}

//line views/vworkspace/StandupWorkspace.html:99
func (p *StandupWorkspace) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/StandupWorkspace.html:99
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/StandupWorkspace.html:99
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/StandupWorkspace.html:99
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/StandupWorkspace.html:99
}

//line views/vworkspace/StandupWorkspace.html:99
func (p *StandupWorkspace) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/StandupWorkspace.html:99
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/StandupWorkspace.html:99
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/StandupWorkspace.html:99
	qs422016 := string(qb422016.B)
//line views/vworkspace/StandupWorkspace.html:99
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/StandupWorkspace.html:99
	return qs422016
//line views/vworkspace/StandupWorkspace.html:99
}

//line views/vworkspace/StandupWorkspace.html:101
func StreamStandupWorkspaceModal(qw422016 *qt422016.Writer, w *workspace.FullStandup, teams team.Teams, sprints sprint.Sprints) {
//line views/vworkspace/StandupWorkspace.html:101
	qw422016.N().S(`
  <div id="modal-standup-config" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Standup</h2>
      </div>
      <div class="modal-body">
        <form action="/team/`)
//line views/vworkspace/StandupWorkspace.html:110
	qw422016.E().S(w.Standup.Slug)
//line views/vworkspace/StandupWorkspace.html:110
	qw422016.N().S(`" method="post" class="expanded">
          <input type="hidden" name="action" value="edit" />
          <em>Name</em><br />
          `)
//line views/vworkspace/StandupWorkspace.html:113
	components.StreamFormInput(qw422016, "title", "input-title", w.Standup.TitleString(), "The name of your standup")
//line views/vworkspace/StandupWorkspace.html:113
	qw422016.N().S(`
          <hr />
          <em>Team</em><br />
          `)
//line views/vworkspace/StandupWorkspace.html:116
	components.StreamFormSelect(qw422016, "team", "input-team", util.UUIDString(w.Standup.TeamID), teams.IDStrings(true), teams.TitleStrings("- no team -"), 5)
//line views/vworkspace/StandupWorkspace.html:116
	qw422016.N().S(`
          <hr />
          <em>Sprint</em><br />
          `)
//line views/vworkspace/StandupWorkspace.html:119
	components.StreamFormSelect(qw422016, "sprint", "input-sprint", util.UUIDString(w.Standup.SprintID), sprints.IDStrings(true), sprints.TitleStrings("- no sprint -"), 5)
//line views/vworkspace/StandupWorkspace.html:119
	qw422016.N().S(`
          <hr />
          <em>Permissions</em>
          <div><label><input type="checkbox" name="perm-team" value="true"> Must be a member of this standup's team</label></div>
          <div><label><input type="checkbox" name="perm-sprint" value="true"> Must be a member of this standup's sprint</label></div>
          <hr />
          <div class="right"><button type="submit">Save</button></div>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/StandupWorkspace.html:130
}

//line views/vworkspace/StandupWorkspace.html:130
func WriteStandupWorkspaceModal(qq422016 qtio422016.Writer, w *workspace.FullStandup, teams team.Teams, sprints sprint.Sprints) {
//line views/vworkspace/StandupWorkspace.html:130
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/StandupWorkspace.html:130
	StreamStandupWorkspaceModal(qw422016, w, teams, sprints)
//line views/vworkspace/StandupWorkspace.html:130
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/StandupWorkspace.html:130
}

//line views/vworkspace/StandupWorkspace.html:130
func StandupWorkspaceModal(w *workspace.FullStandup, teams team.Teams, sprints sprint.Sprints) string {
//line views/vworkspace/StandupWorkspace.html:130
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/StandupWorkspace.html:130
	WriteStandupWorkspaceModal(qb422016, w, teams, sprints)
//line views/vworkspace/StandupWorkspace.html:130
	qs422016 := string(qb422016.B)
//line views/vworkspace/StandupWorkspace.html:130
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/StandupWorkspace.html:130
	return qs422016
//line views/vworkspace/StandupWorkspace.html:130
}