// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsprint/Detail.html:2
package vsprint

//line views/vsprint/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/app/sprint/smember"
	"github.com/kyleu/rituals/app/sprint/spermission"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vestimate"
	"github.com/kyleu/rituals/views/vretro"
	"github.com/kyleu/rituals/views/vsprint/vshistory"
	"github.com/kyleu/rituals/views/vsprint/vsmember"
	"github.com/kyleu/rituals/views/vsprint/vspermission"
	"github.com/kyleu/rituals/views/vstandup"
)

//line views/vsprint/Detail.html:25
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/Detail.html:25
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/Detail.html:25
type Detail struct {
	layout.Basic
	Model                          *sprint.Sprint
	TeamByTeamID                   *team.Team
	Params                         filter.ParamSet
	RelEstimatesBySprintID         estimate.Estimates
	RelRetrosBySprintID            retro.Retros
	RelSprintHistoriesBySprintID   shistory.SprintHistories
	RelSprintMembersBySprintID     smember.SprintMembers
	RelSprintPermissionsBySprintID spermission.SprintPermissions
	RelStandupsBySprintID          standup.Standups
}

//line views/vsprint/Detail.html:38
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/Detail.html:38
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-sprint"><button type="button">JSON</button></a>
      <a href="`)
//line views/vsprint/Detail.html:42
	qw422016.E().S(p.Model.WebPath())
//line views/vsprint/Detail.html:42
	qw422016.N().S(`/edit"><button>`)
//line views/vsprint/Detail.html:42
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vsprint/Detail.html:42
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vsprint/Detail.html:44
	components.StreamSVGRefIcon(qw422016, `sprint`, ps)
//line views/vsprint/Detail.html:44
	qw422016.N().S(` `)
//line views/vsprint/Detail.html:44
	qw422016.E().S(p.Model.TitleString())
//line views/vsprint/Detail.html:44
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/sprint"><em>Sprint</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vsprint/Detail.html:50
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vsprint/Detail.html:50
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Slug</th>
          <td>`)
//line views/vsprint/Detail.html:54
	qw422016.E().S(p.Model.Slug)
//line views/vsprint/Detail.html:54
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Title</th>
          <td><strong>`)
//line views/vsprint/Detail.html:58
	qw422016.E().S(p.Model.Title)
//line views/vsprint/Detail.html:58
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Icon</th>
          <td>`)
//line views/vsprint/Detail.html:62
	qw422016.E().S(p.Model.Icon)
//line views/vsprint/Detail.html:62
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Available options: [new, active, complete]">Status</th>
          <td>`)
//line views/vsprint/Detail.html:66
	qw422016.E().V(p.Model.Status)
//line views/vsprint/Detail.html:66
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
          <td class="nowrap">
            `)
//line views/vsprint/Detail.html:71
	components.StreamDisplayUUID(qw422016, p.Model.TeamID)
//line views/vsprint/Detail.html:71
	if p.TeamByTeamID != nil {
//line views/vsprint/Detail.html:71
		qw422016.N().S(` (`)
//line views/vsprint/Detail.html:71
		qw422016.E().S(p.TeamByTeamID.TitleString())
//line views/vsprint/Detail.html:71
		qw422016.N().S(`)`)
//line views/vsprint/Detail.html:71
	}
//line views/vsprint/Detail.html:71
	qw422016.N().S(`
            `)
//line views/vsprint/Detail.html:72
	if p.Model.TeamID != nil {
//line views/vsprint/Detail.html:72
		qw422016.N().S(`<a title="Team" href="`)
//line views/vsprint/Detail.html:72
		qw422016.E().S(`/admin/db/team` + `/` + p.Model.TeamID.String())
//line views/vsprint/Detail.html:72
		qw422016.N().S(`">`)
//line views/vsprint/Detail.html:72
		components.StreamSVGRef(qw422016, "team", 18, 18, "", ps)
//line views/vsprint/Detail.html:72
		qw422016.N().S(`</a>`)
//line views/vsprint/Detail.html:72
	}
//line views/vsprint/Detail.html:72
	qw422016.N().S(`
          </td>
        </tr>
        <tr>
          <th class="shrink" title="Calendar date (optional)">Start Date</th>
          <td>`)
//line views/vsprint/Detail.html:77
	components.StreamDisplayTimestampDay(qw422016, p.Model.StartDate)
//line views/vsprint/Detail.html:77
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Calendar date (optional)">End Date</th>
          <td>`)
//line views/vsprint/Detail.html:81
	components.StreamDisplayTimestampDay(qw422016, p.Model.EndDate)
//line views/vsprint/Detail.html:81
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vsprint/Detail.html:85
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vsprint/Detail.html:85
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vsprint/Detail.html:89
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vsprint/Detail.html:89
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vsprint/Detail.html:95
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-EstimatesBySprintID" type="checkbox" hidden />
        <label for="accordion-EstimatesBySprintID">
          `)
//line views/vsprint/Detail.html:102
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsprint/Detail.html:102
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:103
	components.StreamSVGRefIcon(qw422016, `estimate`, ps)
//line views/vsprint/Detail.html:103
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:104
	qw422016.N().D(len(p.RelEstimatesBySprintID))
//line views/vsprint/Detail.html:104
	qw422016.N().S(` `)
//line views/vsprint/Detail.html:104
	qw422016.E().S(util.StringPluralMaybe("Estimate", len(p.RelEstimatesBySprintID)))
//line views/vsprint/Detail.html:104
	qw422016.N().S(` by [sprint_id]
        </label>
        <div class="bd">
`)
//line views/vsprint/Detail.html:107
	if len(p.RelEstimatesBySprintID) == 0 {
//line views/vsprint/Detail.html:107
		qw422016.N().S(`          <em>no related Estimates</em>
`)
//line views/vsprint/Detail.html:109
	} else {
//line views/vsprint/Detail.html:109
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vsprint/Detail.html:111
		vestimate.StreamTable(qw422016, p.RelEstimatesBySprintID, nil, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:111
		qw422016.N().S(`
          </div>
`)
//line views/vsprint/Detail.html:113
	}
//line views/vsprint/Detail.html:113
	qw422016.N().S(`        </div>
      </li>
      <li>
        <input id="accordion-RetrosBySprintID" type="checkbox" hidden />
        <label for="accordion-RetrosBySprintID">
          `)
//line views/vsprint/Detail.html:119
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsprint/Detail.html:119
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:120
	components.StreamSVGRefIcon(qw422016, `retro`, ps)
//line views/vsprint/Detail.html:120
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:121
	qw422016.N().D(len(p.RelRetrosBySprintID))
//line views/vsprint/Detail.html:121
	qw422016.N().S(` `)
//line views/vsprint/Detail.html:121
	qw422016.E().S(util.StringPluralMaybe("Retro", len(p.RelRetrosBySprintID)))
//line views/vsprint/Detail.html:121
	qw422016.N().S(` by [sprint_id]
        </label>
        <div class="bd">
`)
//line views/vsprint/Detail.html:124
	if len(p.RelRetrosBySprintID) == 0 {
//line views/vsprint/Detail.html:124
		qw422016.N().S(`          <em>no related Retros</em>
`)
//line views/vsprint/Detail.html:126
	} else {
//line views/vsprint/Detail.html:126
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vsprint/Detail.html:128
		vretro.StreamTable(qw422016, p.RelRetrosBySprintID, nil, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:128
		qw422016.N().S(`
          </div>
`)
//line views/vsprint/Detail.html:130
	}
//line views/vsprint/Detail.html:130
	qw422016.N().S(`        </div>
      </li>
      <li>
        <input id="accordion-SprintHistoriesBySprintID" type="checkbox" hidden />
        <label for="accordion-SprintHistoriesBySprintID">
          `)
//line views/vsprint/Detail.html:136
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsprint/Detail.html:136
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:137
	components.StreamSVGRefIcon(qw422016, `history`, ps)
//line views/vsprint/Detail.html:137
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:138
	qw422016.N().D(len(p.RelSprintHistoriesBySprintID))
//line views/vsprint/Detail.html:138
	qw422016.N().S(` `)
//line views/vsprint/Detail.html:138
	qw422016.E().S(util.StringPluralMaybe("History", len(p.RelSprintHistoriesBySprintID)))
//line views/vsprint/Detail.html:138
	qw422016.N().S(` by [sprint_id]
        </label>
        <div class="bd">
`)
//line views/vsprint/Detail.html:141
	if len(p.RelSprintHistoriesBySprintID) == 0 {
//line views/vsprint/Detail.html:141
		qw422016.N().S(`          <em>no related Histories</em>
`)
//line views/vsprint/Detail.html:143
	} else {
//line views/vsprint/Detail.html:143
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vsprint/Detail.html:145
		vshistory.StreamTable(qw422016, p.RelSprintHistoriesBySprintID, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:145
		qw422016.N().S(`
          </div>
`)
//line views/vsprint/Detail.html:147
	}
//line views/vsprint/Detail.html:147
	qw422016.N().S(`        </div>
      </li>
      <li>
        <input id="accordion-SprintMembersBySprintID" type="checkbox" hidden />
        <label for="accordion-SprintMembersBySprintID">
          `)
//line views/vsprint/Detail.html:153
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsprint/Detail.html:153
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:154
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vsprint/Detail.html:154
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:155
	qw422016.N().D(len(p.RelSprintMembersBySprintID))
//line views/vsprint/Detail.html:155
	qw422016.N().S(` `)
//line views/vsprint/Detail.html:155
	qw422016.E().S(util.StringPluralMaybe("Member", len(p.RelSprintMembersBySprintID)))
//line views/vsprint/Detail.html:155
	qw422016.N().S(` by [sprint_id]
        </label>
        <div class="bd">
`)
//line views/vsprint/Detail.html:158
	if len(p.RelSprintMembersBySprintID) == 0 {
//line views/vsprint/Detail.html:158
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vsprint/Detail.html:160
	} else {
//line views/vsprint/Detail.html:160
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vsprint/Detail.html:162
		vsmember.StreamTable(qw422016, p.RelSprintMembersBySprintID, nil, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:162
		qw422016.N().S(`
          </div>
`)
//line views/vsprint/Detail.html:164
	}
//line views/vsprint/Detail.html:164
	qw422016.N().S(`        </div>
      </li>
      <li>
        <input id="accordion-SprintPermissionsBySprintID" type="checkbox" hidden />
        <label for="accordion-SprintPermissionsBySprintID">
          `)
//line views/vsprint/Detail.html:170
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsprint/Detail.html:170
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:171
	components.StreamSVGRefIcon(qw422016, `permission`, ps)
//line views/vsprint/Detail.html:171
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:172
	qw422016.N().D(len(p.RelSprintPermissionsBySprintID))
//line views/vsprint/Detail.html:172
	qw422016.N().S(` `)
//line views/vsprint/Detail.html:172
	qw422016.E().S(util.StringPluralMaybe("Permission", len(p.RelSprintPermissionsBySprintID)))
//line views/vsprint/Detail.html:172
	qw422016.N().S(` by [sprint_id]
        </label>
        <div class="bd">
`)
//line views/vsprint/Detail.html:175
	if len(p.RelSprintPermissionsBySprintID) == 0 {
//line views/vsprint/Detail.html:175
		qw422016.N().S(`          <em>no related Permissions</em>
`)
//line views/vsprint/Detail.html:177
	} else {
//line views/vsprint/Detail.html:177
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vsprint/Detail.html:179
		vspermission.StreamTable(qw422016, p.RelSprintPermissionsBySprintID, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:179
		qw422016.N().S(`
          </div>
`)
//line views/vsprint/Detail.html:181
	}
//line views/vsprint/Detail.html:181
	qw422016.N().S(`        </div>
      </li>
      <li>
        <input id="accordion-StandupsBySprintID" type="checkbox" hidden />
        <label for="accordion-StandupsBySprintID">
          `)
//line views/vsprint/Detail.html:187
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsprint/Detail.html:187
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:188
	components.StreamSVGRefIcon(qw422016, `standup`, ps)
//line views/vsprint/Detail.html:188
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:189
	qw422016.N().D(len(p.RelStandupsBySprintID))
//line views/vsprint/Detail.html:189
	qw422016.N().S(` `)
//line views/vsprint/Detail.html:189
	qw422016.E().S(util.StringPluralMaybe("Standup", len(p.RelStandupsBySprintID)))
//line views/vsprint/Detail.html:189
	qw422016.N().S(` by [sprint_id]
        </label>
        <div class="bd">
`)
//line views/vsprint/Detail.html:192
	if len(p.RelStandupsBySprintID) == 0 {
//line views/vsprint/Detail.html:192
		qw422016.N().S(`          <em>no related Standups</em>
`)
//line views/vsprint/Detail.html:194
	} else {
//line views/vsprint/Detail.html:194
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vsprint/Detail.html:196
		vstandup.StreamTable(qw422016, p.RelStandupsBySprintID, nil, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:196
		qw422016.N().S(`
          </div>
`)
//line views/vsprint/Detail.html:198
	}
//line views/vsprint/Detail.html:198
	qw422016.N().S(`        </div>
      </li>
    </ul>
  </div>
  `)
//line views/vsprint/Detail.html:203
	components.StreamJSONModal(qw422016, "sprint", "Sprint JSON", p.Model, 1)
//line views/vsprint/Detail.html:203
	qw422016.N().S(`
`)
//line views/vsprint/Detail.html:204
}

//line views/vsprint/Detail.html:204
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/Detail.html:204
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/Detail.html:204
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/Detail.html:204
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/Detail.html:204
}

//line views/vsprint/Detail.html:204
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/Detail.html:204
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/Detail.html:204
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/Detail.html:204
	qs422016 := string(qb422016.B)
//line views/vsprint/Detail.html:204
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/Detail.html:204
	return qs422016
//line views/vsprint/Detail.html:204
}
