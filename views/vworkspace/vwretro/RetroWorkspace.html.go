// Code generated by qtc from "RetroWorkspace.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwretro/RetroWorkspace.html:1
package vwretro

//line views/vworkspace/vwretro/RetroWorkspace.html:1
import (
	"strings"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vworkspace/vwutil"
)

//line views/vworkspace/vwretro/RetroWorkspace.html:19
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwretro/RetroWorkspace.html:19
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwretro/RetroWorkspace.html:19
type RetroWorkspace struct {
	layout.Basic
	FullRetro *workspace.FullRetro
	Teams     team.Teams
	Sprints   sprint.Sprints
}

//line views/vworkspace/vwretro/RetroWorkspace.html:26
func (p *RetroWorkspace) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroWorkspace.html:26
	qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspace.html:28
	w := p.FullRetro
	r := w.Retro

//line views/vworkspace/vwretro/RetroWorkspace.html:30
	qw422016.N().S(`  <div class="flex-wrap">
    <div id="panel-summary">
      <div class="card">
        <div class="right">
          `)
//line views/vworkspace/vwretro/RetroWorkspace.html:35
	vwutil.StreamPermissionsLink(qw422016, enum.ModelServiceRetro, r.ID, w.Permissions.ToPermissions(), ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:35
	qw422016.N().S(`
          `)
//line views/vworkspace/vwretro/RetroWorkspace.html:36
	vwutil.StreamComments(qw422016, enum.ModelServiceRetro, r.ID, r.TitleString(), w.Comments, w.UtilMembers, "", ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:36
	qw422016.N().S(`
        </div>
        <a href="#modal-retro-config" id="modal-retro-config-link"><h3>
          <span id="model-icon">`)
//line views/vworkspace/vwretro/RetroWorkspace.html:39
	components.StreamSVGIcon(qw422016, r.IconSafe(), ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:39
	qw422016.N().S(`</span>
          <span id="model-title">`)
//line views/vworkspace/vwretro/RetroWorkspace.html:40
	qw422016.E().S(r.TitleString())
//line views/vworkspace/vwretro/RetroWorkspace.html:40
	qw422016.N().S(`</span>
        </h3></a>
        `)
//line views/vworkspace/vwretro/RetroWorkspace.html:42
	vwutil.StreamBanner(qw422016, w.Team, w.Sprint, "retrospective")
//line views/vworkspace/vwretro/RetroWorkspace.html:42
	qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspace.html:43
	if w.Admin() {
//line views/vworkspace/vwretro/RetroWorkspace.html:43
		qw422016.N().S(`        `)
//line views/vworkspace/vwretro/RetroWorkspace.html:44
		StreamRetroWorkspaceModalEdit(qw422016, r, p.Teams, p.Sprints, w.Permissions.ToPermissions(), ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:44
		qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspace.html:45
	} else {
//line views/vworkspace/vwretro/RetroWorkspace.html:45
		qw422016.N().S(`        `)
//line views/vworkspace/vwretro/RetroWorkspace.html:46
		StreamRetroWorkspaceModalView(qw422016, r, p.Teams, p.Sprints, w.Permissions.ToPermissions(), ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:46
		qw422016.N().S(`
`)
//line views/vworkspace/vwretro/RetroWorkspace.html:47
	}
//line views/vworkspace/vwretro/RetroWorkspace.html:47
	qw422016.N().S(`      </div>
    </div>
    <div id="panel-detail">
      <div class="card">
        <h3>`)
//line views/vworkspace/vwretro/RetroWorkspace.html:52
	components.StreamSVGIcon(qw422016, `comment`, ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:52
	qw422016.N().S(`Feedback</h3>
        <div class="clear"></div>
        `)
//line views/vworkspace/vwretro/RetroWorkspace.html:54
	StreamRetroWorkspaceFeedbacks(qw422016, ps.Profile.ID, ps.Username(), w, ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:54
	qw422016.N().S(`
      </div>
    </div>
    `)
//line views/vworkspace/vwretro/RetroWorkspace.html:57
	vwutil.StreamMemberPanels(qw422016, w.UtilMembers, w.Admin(), r.PublicWebPath(), ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:57
	qw422016.N().S(`
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      initWorkspace("`)
//line views/vworkspace/vwretro/RetroWorkspace.html:61
	qw422016.E().S(util.KeyRetro)
//line views/vworkspace/vwretro/RetroWorkspace.html:61
	qw422016.N().S(`", "`)
//line views/vworkspace/vwretro/RetroWorkspace.html:61
	qw422016.E().S(r.ID.String())
//line views/vworkspace/vwretro/RetroWorkspace.html:61
	qw422016.N().S(`");
    });
  </script>
`)
//line views/vworkspace/vwretro/RetroWorkspace.html:64
}

//line views/vworkspace/vwretro/RetroWorkspace.html:64
func (p *RetroWorkspace) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroWorkspace.html:64
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspace.html:64
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:64
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspace.html:64
}

//line views/vworkspace/vwretro/RetroWorkspace.html:64
func (p *RetroWorkspace) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwretro/RetroWorkspace.html:64
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspace.html:64
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:64
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspace.html:64
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspace.html:64
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspace.html:64
}

//line views/vworkspace/vwretro/RetroWorkspace.html:66
func StreamRetroWorkspaceModalEdit(qw422016 *qt422016.Writer, r *retro.Retro, teams team.Teams, sprints sprint.Sprints, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroWorkspace.html:66
	qw422016.N().S(`
  <div id="modal-retro-config" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>`)
//line views/vworkspace/vwretro/RetroWorkspace.html:72
	components.StreamSVGRef(qw422016, util.KeyRetro, 24, 24, "icon", ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:72
	qw422016.N().S(` Retro</h2>
      </div>
      <div class="modal-body">
        <form action="`)
//line views/vworkspace/vwretro/RetroWorkspace.html:75
	qw422016.E().S(r.PublicWebPath())
//line views/vworkspace/vwretro/RetroWorkspace.html:75
	qw422016.N().S(`" method="post" class="expanded">
          <input type="hidden" name="action" value="`)
//line views/vworkspace/vwretro/RetroWorkspace.html:76
	qw422016.E().S(string(action.ActUpdate))
//line views/vworkspace/vwretro/RetroWorkspace.html:76
	qw422016.N().S(`" />
          `)
//line views/vworkspace/vwretro/RetroWorkspace.html:77
	edit.StreamStringVertical(qw422016, "title", "", "Title", r.TitleString(), 5, "The name of your retro")
//line views/vworkspace/vwretro/RetroWorkspace.html:77
	qw422016.N().S(`
          `)
//line views/vworkspace/vwretro/RetroWorkspace.html:78
	edit.StreamIconPickerVertical(qw422016, "icon", "Icon", r.IconSafe(), ps, 5)
//line views/vworkspace/vwretro/RetroWorkspace.html:78
	qw422016.N().S(`
          `)
//line views/vworkspace/vwretro/RetroWorkspace.html:79
	edit.StreamTagsVertical(qw422016, "categories", "", "Categories", r.Categories, ps, 5, "The available categories for this retro")
//line views/vworkspace/vwretro/RetroWorkspace.html:79
	qw422016.N().S(`
          `)
//line views/vworkspace/vwretro/RetroWorkspace.html:80
	edit.StreamSelectVertical(qw422016, util.KeyTeam, "config-team-input", "Team", util.UUIDString(r.TeamID), teams.IDStrings(true), teams.TitleStrings("- no team -"), 5)
//line views/vworkspace/vwretro/RetroWorkspace.html:80
	qw422016.N().S(`
          `)
//line views/vworkspace/vwretro/RetroWorkspace.html:81
	edit.StreamSelectVertical(qw422016, util.KeySprint, "config-sprint-input", "Sprint", util.UUIDString(r.SprintID), sprints.IDStrings(true), sprints.TitleStrings("- no sprint -"), 5)
//line views/vworkspace/vwretro/RetroWorkspace.html:81
	qw422016.N().S(`
          <hr />
          `)
//line views/vworkspace/vwretro/RetroWorkspace.html:83
	vwutil.StreamPermissionsForm(qw422016, util.KeyRetro, perms, true, teams, true, sprints, ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:83
	qw422016.N().S(`
          <hr />
          <div class="right"><button type="submit">Save</button></div>
          <a href="`)
//line views/vworkspace/vwretro/RetroWorkspace.html:86
	qw422016.E().S(r.PublicWebPath())
//line views/vworkspace/vwretro/RetroWorkspace.html:86
	qw422016.N().S(`/delete" onclick="return confirm('are you sure you wish to delete this retro?')"><button type="button">Delete</button></a>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwretro/RetroWorkspace.html:91
}

//line views/vworkspace/vwretro/RetroWorkspace.html:91
func WriteRetroWorkspaceModalEdit(qq422016 qtio422016.Writer, r *retro.Retro, teams team.Teams, sprints sprint.Sprints, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroWorkspace.html:91
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspace.html:91
	StreamRetroWorkspaceModalEdit(qw422016, r, teams, sprints, perms, ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:91
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspace.html:91
}

//line views/vworkspace/vwretro/RetroWorkspace.html:91
func RetroWorkspaceModalEdit(r *retro.Retro, teams team.Teams, sprints sprint.Sprints, perms util.Permissions, ps *cutil.PageState) string {
//line views/vworkspace/vwretro/RetroWorkspace.html:91
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspace.html:91
	WriteRetroWorkspaceModalEdit(qb422016, r, teams, sprints, perms, ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:91
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspace.html:91
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspace.html:91
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspace.html:91
}

//line views/vworkspace/vwretro/RetroWorkspace.html:93
func StreamRetroWorkspaceModalView(qw422016 *qt422016.Writer, r *retro.Retro, teams team.Teams, sprints sprint.Sprints, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroWorkspace.html:93
	qw422016.N().S(`
  <div id="modal-retro-config" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>
          <span class="view-icon">`)
//line views/vworkspace/vwretro/RetroWorkspace.html:100
	components.StreamSVGRef(qw422016, r.IconSafe(), 24, 24, "icon", ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:100
	qw422016.N().S(`</span>
          <span class="view-title">`)
//line views/vworkspace/vwretro/RetroWorkspace.html:101
	qw422016.E().S(r.TitleString())
//line views/vworkspace/vwretro/RetroWorkspace.html:101
	qw422016.N().S(`</span>
        </h2>
        <em>`)
//line views/vworkspace/vwretro/RetroWorkspace.html:103
	qw422016.E().S(util.KeyRetro)
//line views/vworkspace/vwretro/RetroWorkspace.html:103
	qw422016.N().S(`</em>
      </div>
      <div class="modal-body">
        <div style="display: none;">
          `)
//line views/vworkspace/vwretro/RetroWorkspace.html:107
	edit.StreamIconPicker(qw422016, "icon", r.IconSafe(), ps, 5)
//line views/vworkspace/vwretro/RetroWorkspace.html:107
	qw422016.N().S(`
          `)
//line views/vworkspace/vwretro/RetroWorkspace.html:108
	edit.StreamSelect(qw422016, util.KeyTeam, "config-team-input", util.UUIDString(r.TeamID), teams.IDStrings(true), teams.TitleStrings("- no team -"), 5)
//line views/vworkspace/vwretro/RetroWorkspace.html:108
	qw422016.N().S(`
          `)
//line views/vworkspace/vwretro/RetroWorkspace.html:109
	edit.StreamSelect(qw422016, util.KeySprint, "config-sprint-input", util.UUIDString(r.SprintID), sprints.IDStrings(true), sprints.TitleStrings("- no sprint -"), 5)
//line views/vworkspace/vwretro/RetroWorkspace.html:109
	qw422016.N().S(`
        </div>
        <table>
          <tbody>
          <tr>
            <th class="shrink">Categories</th>
            <td class="config-panel-categories">`)
//line views/vworkspace/vwretro/RetroWorkspace.html:115
	qw422016.E().S(strings.Join(r.Categories, ", "))
//line views/vworkspace/vwretro/RetroWorkspace.html:115
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Team</th>
            <td class="config-panel-team">`)
//line views/vworkspace/vwretro/RetroWorkspace.html:119
	if r.TeamID == nil {
//line views/vworkspace/vwretro/RetroWorkspace.html:119
		qw422016.N().S(`-`)
//line views/vworkspace/vwretro/RetroWorkspace.html:119
	} else {
//line views/vworkspace/vwretro/RetroWorkspace.html:119
		qw422016.N().S(`<a href="/team/`)
//line views/vworkspace/vwretro/RetroWorkspace.html:119
		qw422016.E().S(r.TeamID.String())
//line views/vworkspace/vwretro/RetroWorkspace.html:119
		qw422016.N().S(`">`)
//line views/vworkspace/vwretro/RetroWorkspace.html:119
		qw422016.E().S(teams.TitleFor(r.TeamID))
//line views/vworkspace/vwretro/RetroWorkspace.html:119
		qw422016.N().S(`</a>`)
//line views/vworkspace/vwretro/RetroWorkspace.html:119
	}
//line views/vworkspace/vwretro/RetroWorkspace.html:119
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Sprint</th>
            <td class="config-panel-sprint">`)
//line views/vworkspace/vwretro/RetroWorkspace.html:123
	if r.SprintID == nil {
//line views/vworkspace/vwretro/RetroWorkspace.html:123
		qw422016.N().S(`-`)
//line views/vworkspace/vwretro/RetroWorkspace.html:123
	} else {
//line views/vworkspace/vwretro/RetroWorkspace.html:123
		qw422016.N().S(`<a href="/sprint/`)
//line views/vworkspace/vwretro/RetroWorkspace.html:123
		qw422016.E().S(r.SprintID.String())
//line views/vworkspace/vwretro/RetroWorkspace.html:123
		qw422016.N().S(`">`)
//line views/vworkspace/vwretro/RetroWorkspace.html:123
		qw422016.E().S(sprints.TitleFor(r.SprintID))
//line views/vworkspace/vwretro/RetroWorkspace.html:123
		qw422016.N().S(`</a>`)
//line views/vworkspace/vwretro/RetroWorkspace.html:123
	}
//line views/vworkspace/vwretro/RetroWorkspace.html:123
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Permissions</th>
            <td class="config-panel-perms">`)
//line views/vworkspace/vwretro/RetroWorkspace.html:127
	vwutil.StreamPermissionsList(qw422016, util.KeyRetro, perms, ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:127
	qw422016.N().S(`</td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwretro/RetroWorkspace.html:134
}

//line views/vworkspace/vwretro/RetroWorkspace.html:134
func WriteRetroWorkspaceModalView(qq422016 qtio422016.Writer, r *retro.Retro, teams team.Teams, sprints sprint.Sprints, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwretro/RetroWorkspace.html:134
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwretro/RetroWorkspace.html:134
	StreamRetroWorkspaceModalView(qw422016, r, teams, sprints, perms, ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:134
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwretro/RetroWorkspace.html:134
}

//line views/vworkspace/vwretro/RetroWorkspace.html:134
func RetroWorkspaceModalView(r *retro.Retro, teams team.Teams, sprints sprint.Sprints, perms util.Permissions, ps *cutil.PageState) string {
//line views/vworkspace/vwretro/RetroWorkspace.html:134
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwretro/RetroWorkspace.html:134
	WriteRetroWorkspaceModalView(qb422016, r, teams, sprints, perms, ps)
//line views/vworkspace/vwretro/RetroWorkspace.html:134
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwretro/RetroWorkspace.html:134
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwretro/RetroWorkspace.html:134
	return qs422016
//line views/vworkspace/vwretro/RetroWorkspace.html:134
}
