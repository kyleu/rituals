// Code generated by qtc from "EstimateWorkspace.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwestimate/EstimateWorkspace.html:1
package vwestimate

//line views/vworkspace/vwestimate/EstimateWorkspace.html:1
import (
	"strings"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vworkspace/vwutil"
)

//line views/vworkspace/vwestimate/EstimateWorkspace.html:19
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwestimate/EstimateWorkspace.html:19
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwestimate/EstimateWorkspace.html:19
type EstimateWorkspace struct {
	layout.Basic
	FullEstimate *workspace.FullEstimate
	Teams        team.Teams
	Sprints      sprint.Sprints
}

//line views/vworkspace/vwestimate/EstimateWorkspace.html:26
func (p *EstimateWorkspace) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:26
	qw422016.N().S(`
`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:28
	w := p.FullEstimate
	e := w.Estimate

//line views/vworkspace/vwestimate/EstimateWorkspace.html:30
	qw422016.N().S(`  <div class="flex-wrap">
    <div id="panel-summary">
      <div class="card">
        <div class="right">
          `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:35
	vwutil.StreamPermissionsLink(qw422016, enum.ModelServiceEstimate, e.ID, w.Permissions.ToPermissions(), ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:35
	qw422016.N().S(`
          `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:36
	vwutil.StreamComments(qw422016, enum.ModelServiceEstimate, e.ID, e.TitleString(), w.Comments, w.UtilMembers, "", ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:36
	qw422016.N().S(`
        </div>
        <a href="#modal-estimate-config" id="modal-estimate-config-link"><h3>
          <span id="model-icon">`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:39
	components.StreamSVGIcon(qw422016, e.IconSafe(), ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:39
	qw422016.N().S(`</span>
          <span id="model-title">`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:40
	qw422016.E().S(e.TitleString())
//line views/vworkspace/vwestimate/EstimateWorkspace.html:40
	qw422016.N().S(`</span>
        </h3></a>
        `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:42
	vwutil.StreamBanner(qw422016, w.Team, w.Sprint, util.KeyEstimate)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:42
	qw422016.N().S(`
`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:43
	if w.Admin() {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:43
		qw422016.N().S(`        `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:44
		StreamEstimateWorkspaceModalEdit(qw422016, e, p.Teams, p.Sprints, w.Permissions.ToPermissions(), ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:44
		qw422016.N().S(`
`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:45
	} else {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:45
		qw422016.N().S(`        `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:46
		StreamEstimateWorkspaceModalView(qw422016, e, p.Teams, p.Sprints, w.Permissions.ToPermissions(), ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:46
		qw422016.N().S(`
`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:47
	}
//line views/vworkspace/vwestimate/EstimateWorkspace.html:47
	qw422016.N().S(`      </div>
    </div>
    <div id="panel-detail">
      <div class="card">
        <div class="right"><a class="add-story-link" href="#modal-story--add"><button>Add Story</button></a></div>
        <a class="add-story-link" href="#modal-story--add"><h3>`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:53
	components.StreamSVGIcon(qw422016, util.KeyStory, ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:53
	qw422016.N().S(`Stories</h3></a>
        `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:54
	StreamEstimateWorkspaceStories(qw422016, w, ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:54
	qw422016.N().S(`
        `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:55
	StreamEstimateWorkspaceStoryModalAdd(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:55
	qw422016.N().S(`
        `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:56
	StreamEstimateWorkspaceStoryModalEmpty(qw422016, ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:56
	qw422016.N().S(`
        `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:57
	StreamEstimateWorkspaceStoryModalEditPanel(qw422016, "", "", ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:57
	qw422016.N().S(`
      </div>
    </div>
    `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:60
	vwutil.StreamMemberPanels(qw422016, w.UtilMembers, w.Admin(), e.PublicWebPath(), ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:60
	qw422016.N().S(`
  </div>
  <script>
    document.addEventListener("DOMContentLoaded", function() {
      initWorkspace("`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:64
	qw422016.E().S(util.KeyEstimate)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:64
	qw422016.N().S(`", "`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:64
	qw422016.E().S(e.ID.String())
//line views/vworkspace/vwestimate/EstimateWorkspace.html:64
	qw422016.N().S(`");
    });
  </script>
`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
}

//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
func (p *EstimateWorkspace) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
	p.StreamBody(qw422016, as, ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
}

//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
func (p *EstimateWorkspace) Body(as *app.State, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
	p.WriteBody(qb422016, as, ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspace.html:67
}

//line views/vworkspace/vwestimate/EstimateWorkspace.html:69
func StreamEstimateWorkspaceModalEdit(qw422016 *qt422016.Writer, e *estimate.Estimate, teams team.Teams, sprints sprint.Sprints, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:69
	qw422016.N().S(`
  <div id="modal-estimate-config" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:75
	components.StreamSVGRef(qw422016, util.KeyEstimate, 24, 24, "icon", ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:75
	qw422016.N().S(` Estimate</h2>
      </div>
      <div class="modal-body">
        <form action="`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:78
	qw422016.E().S(e.PublicWebPath())
//line views/vworkspace/vwestimate/EstimateWorkspace.html:78
	qw422016.N().S(`" method="post" class="expanded">
          <input type="hidden" name="action" value="`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:79
	qw422016.E().S(string(action.ActUpdate))
//line views/vworkspace/vwestimate/EstimateWorkspace.html:79
	qw422016.N().S(`" />
          `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:80
	edit.StreamStringVertical(qw422016, "title", "", "Title", e.TitleString(), 5, "The name of your estimate")
//line views/vworkspace/vwestimate/EstimateWorkspace.html:80
	qw422016.N().S(`
          `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:81
	edit.StreamIconPickerVertical(qw422016, "icon", "Icon", e.IconSafe(), ps, 5)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:81
	qw422016.N().S(`
          `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:82
	edit.StreamTagsVertical(qw422016, "choices", "", "Choices", e.Choices, ps, 5, "The available options for stories in this estimate")
//line views/vworkspace/vwestimate/EstimateWorkspace.html:82
	qw422016.N().S(`
          `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:83
	edit.StreamSelectVertical(qw422016, util.KeyTeam, "config-team-input", "Team", util.UUIDString(e.TeamID), teams.IDStrings(true), teams.TitleStrings("- no team -"), 5)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:83
	qw422016.N().S(`
          `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:84
	edit.StreamSelectVertical(qw422016, util.KeySprint, "config-sprint-input", "Sprint", util.UUIDString(e.SprintID), sprints.IDStrings(true), sprints.TitleStrings("- no sprint -"), 5)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:84
	qw422016.N().S(`
          <hr />
          `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:86
	vwutil.StreamPermissionsForm(qw422016, util.KeyEstimate, perms, true, teams, true, sprints, ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:86
	qw422016.N().S(`
          <hr />
          <div class="right"><button type="submit">Save</button></div>
          <a href="`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:89
	qw422016.E().S(e.PublicWebPath())
//line views/vworkspace/vwestimate/EstimateWorkspace.html:89
	qw422016.N().S(`/delete" onclick="return confirm('are you sure you wish to delete this estimate?')"><button type="button">Delete</button></a>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
}

//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
func WriteEstimateWorkspaceModalEdit(qq422016 qtio422016.Writer, e *estimate.Estimate, teams team.Teams, sprints sprint.Sprints, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
	StreamEstimateWorkspaceModalEdit(qw422016, e, teams, sprints, perms, ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
}

//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
func EstimateWorkspaceModalEdit(e *estimate.Estimate, teams team.Teams, sprints sprint.Sprints, perms util.Permissions, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
	WriteEstimateWorkspaceModalEdit(qb422016, e, teams, sprints, perms, ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspace.html:94
}

//line views/vworkspace/vwestimate/EstimateWorkspace.html:96
func StreamEstimateWorkspaceModalView(qw422016 *qt422016.Writer, e *estimate.Estimate, teams team.Teams, sprints sprint.Sprints, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:96
	qw422016.N().S(`
  <div id="modal-estimate-config" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>
          <span class="view-icon">`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:103
	components.StreamSVGRef(qw422016, e.IconSafe(), 24, 24, "icon", ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:103
	qw422016.N().S(`</span>
          <span class="view-title">`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:104
	qw422016.E().S(e.TitleString())
//line views/vworkspace/vwestimate/EstimateWorkspace.html:104
	qw422016.N().S(`</span>
        </h2>
        <em>`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:106
	qw422016.E().S(util.KeyEstimate)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:106
	qw422016.N().S(`</em>
      </div>
      <div class="modal-body">
        <div style="display: none;">
          `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:110
	edit.StreamIconPicker(qw422016, "icon", e.IconSafe(), ps, 5)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:110
	qw422016.N().S(`
          `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:111
	edit.StreamSelect(qw422016, util.KeyTeam, "config-team-input", util.UUIDString(e.TeamID), teams.IDStrings(true), teams.TitleStrings("- no team -"), 5)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:111
	qw422016.N().S(`
          `)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:112
	edit.StreamSelect(qw422016, util.KeySprint, "config-sprint-input", util.UUIDString(e.SprintID), sprints.IDStrings(true), sprints.TitleStrings("- no sprint -"), 5)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:112
	qw422016.N().S(`
        </div>
        <table>
          <tbody>
          <tr>
            <th class="shrink">Choices</th>
            <td class="config-panel-choices">`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:118
	qw422016.E().S(strings.Join(e.Choices, ", "))
//line views/vworkspace/vwestimate/EstimateWorkspace.html:118
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Team</th>
            <td class="config-panel-team">`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:122
	if e.TeamID == nil {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:122
		qw422016.N().S(`-`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:122
	} else {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:122
		qw422016.N().S(`<a href="/team/`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:122
		qw422016.E().S(e.TeamID.String())
//line views/vworkspace/vwestimate/EstimateWorkspace.html:122
		qw422016.N().S(`">`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:122
		qw422016.E().S(teams.TitleFor(e.TeamID))
//line views/vworkspace/vwestimate/EstimateWorkspace.html:122
		qw422016.N().S(`</a>`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:122
	}
//line views/vworkspace/vwestimate/EstimateWorkspace.html:122
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Sprint</th>
            <td class="config-panel-sprint">`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:126
	if e.SprintID == nil {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:126
		qw422016.N().S(`-`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:126
	} else {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:126
		qw422016.N().S(`<a href="/sprint/`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:126
		qw422016.E().S(e.SprintID.String())
//line views/vworkspace/vwestimate/EstimateWorkspace.html:126
		qw422016.N().S(`">`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:126
		qw422016.E().S(sprints.TitleFor(e.SprintID))
//line views/vworkspace/vwestimate/EstimateWorkspace.html:126
		qw422016.N().S(`</a>`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:126
	}
//line views/vworkspace/vwestimate/EstimateWorkspace.html:126
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th>Permissions</th>
            <td class="config-panel-perms">`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:130
	vwutil.StreamPermissionsList(qw422016, util.KeyEstimate, perms, ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:130
	qw422016.N().S(`</td>
          </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
}

//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
func WriteEstimateWorkspaceModalView(qq422016 qtio422016.Writer, e *estimate.Estimate, teams team.Teams, sprints sprint.Sprints, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
	StreamEstimateWorkspaceModalView(qw422016, e, teams, sprints, perms, ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
}

//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
func EstimateWorkspaceModalView(e *estimate.Estimate, teams team.Teams, sprints sprint.Sprints, perms util.Permissions, ps *cutil.PageState) string {
//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
	WriteEstimateWorkspaceModalView(qb422016, e, teams, sprints, perms, ps)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
	return qs422016
//line views/vworkspace/vwestimate/EstimateWorkspace.html:137
}
