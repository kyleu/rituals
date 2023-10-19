// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vteam/Detail.html:2
package vteam

//line views/vteam/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/team/thistory"
	"github.com/kyleu/rituals/app/team/tmember"
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vestimate"
	"github.com/kyleu/rituals/views/vretro"
	"github.com/kyleu/rituals/views/vsprint"
	"github.com/kyleu/rituals/views/vstandup"
	"github.com/kyleu/rituals/views/vteam/vthistory"
	"github.com/kyleu/rituals/views/vteam/vtmember"
	"github.com/kyleu/rituals/views/vteam/vtpermission"
)

//line views/vteam/Detail.html:27
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vteam/Detail.html:27
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vteam/Detail.html:27
type Detail struct {
	layout.Basic
	Model                      *team.Team
	Params                     filter.ParamSet
	RelEstimatesByTeamID       estimate.Estimates
	RelRetrosByTeamID          retro.Retros
	RelSprintsByTeamID         sprint.Sprints
	RelStandupsByTeamID        standup.Standups
	RelTeamHistoriesByTeamID   thistory.TeamHistories
	RelTeamMembersByTeamID     tmember.TeamMembers
	RelTeamPermissionsByTeamID tpermission.TeamPermissions
}

//line views/vteam/Detail.html:40
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vteam/Detail.html:40
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-team"><button type="button">JSON</button></a>
      <a href="`)
//line views/vteam/Detail.html:44
	qw422016.E().S(p.Model.WebPath())
//line views/vteam/Detail.html:44
	qw422016.N().S(`/edit"><button>`)
//line views/vteam/Detail.html:44
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vteam/Detail.html:44
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vteam/Detail.html:46
	components.StreamSVGRefIcon(qw422016, `team`, ps)
//line views/vteam/Detail.html:46
	qw422016.N().S(` `)
//line views/vteam/Detail.html:46
	qw422016.E().S(p.Model.TitleString())
//line views/vteam/Detail.html:46
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/team"><em>Team</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vteam/Detail.html:52
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vteam/Detail.html:52
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Slug</th>
          <td>`)
//line views/vteam/Detail.html:56
	qw422016.E().S(p.Model.Slug)
//line views/vteam/Detail.html:56
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Title</th>
          <td><strong>`)
//line views/vteam/Detail.html:60
	qw422016.E().S(p.Model.Title)
//line views/vteam/Detail.html:60
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Icon</th>
          <td>`)
//line views/vteam/Detail.html:64
	qw422016.E().S(p.Model.Icon)
//line views/vteam/Detail.html:64
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="`)
//line views/vteam/Detail.html:67
	qw422016.E().S(enum.AllSessionStatuses.Help())
//line views/vteam/Detail.html:67
	qw422016.N().S(`">Status</th>
          <td>`)
//line views/vteam/Detail.html:68
	qw422016.E().S(p.Model.Status.String())
//line views/vteam/Detail.html:68
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vteam/Detail.html:72
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vteam/Detail.html:72
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vteam/Detail.html:76
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vteam/Detail.html:76
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vteam/Detail.html:82
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-EstimatesByTeamID" type="checkbox" hidden />
        <label for="accordion-EstimatesByTeamID">
          `)
//line views/vteam/Detail.html:89
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vteam/Detail.html:89
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:90
	components.StreamSVGRef(qw422016, `estimate`, 16, 16, `icon`, ps)
//line views/vteam/Detail.html:90
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:91
	qw422016.E().S(util.StringPlural(len(p.RelEstimatesByTeamID), "Estimate"))
//line views/vteam/Detail.html:91
	qw422016.N().S(` by [Team ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vteam/Detail.html:94
	if len(p.RelEstimatesByTeamID) == 0 {
//line views/vteam/Detail.html:94
		qw422016.N().S(`          <em>no related Estimates</em>
`)
//line views/vteam/Detail.html:96
	} else {
//line views/vteam/Detail.html:96
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vteam/Detail.html:98
		vestimate.StreamTable(qw422016, p.RelEstimatesByTeamID, nil, nil, p.Params, as, ps)
//line views/vteam/Detail.html:98
		qw422016.N().S(`
          </div>
`)
//line views/vteam/Detail.html:100
	}
//line views/vteam/Detail.html:100
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-RetrosByTeamID" type="checkbox" hidden />
        <label for="accordion-RetrosByTeamID">
          `)
//line views/vteam/Detail.html:106
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vteam/Detail.html:106
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:107
	components.StreamSVGRef(qw422016, `retro`, 16, 16, `icon`, ps)
//line views/vteam/Detail.html:107
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:108
	qw422016.E().S(util.StringPlural(len(p.RelRetrosByTeamID), "Retro"))
//line views/vteam/Detail.html:108
	qw422016.N().S(` by [Team ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vteam/Detail.html:111
	if len(p.RelRetrosByTeamID) == 0 {
//line views/vteam/Detail.html:111
		qw422016.N().S(`          <em>no related Retros</em>
`)
//line views/vteam/Detail.html:113
	} else {
//line views/vteam/Detail.html:113
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vteam/Detail.html:115
		vretro.StreamTable(qw422016, p.RelRetrosByTeamID, nil, nil, p.Params, as, ps)
//line views/vteam/Detail.html:115
		qw422016.N().S(`
          </div>
`)
//line views/vteam/Detail.html:117
	}
//line views/vteam/Detail.html:117
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-SprintsByTeamID" type="checkbox" hidden />
        <label for="accordion-SprintsByTeamID">
          `)
//line views/vteam/Detail.html:123
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vteam/Detail.html:123
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:124
	components.StreamSVGRef(qw422016, `sprint`, 16, 16, `icon`, ps)
//line views/vteam/Detail.html:124
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:125
	qw422016.E().S(util.StringPlural(len(p.RelSprintsByTeamID), "Sprint"))
//line views/vteam/Detail.html:125
	qw422016.N().S(` by [Team ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vteam/Detail.html:128
	if len(p.RelSprintsByTeamID) == 0 {
//line views/vteam/Detail.html:128
		qw422016.N().S(`          <em>no related Sprints</em>
`)
//line views/vteam/Detail.html:130
	} else {
//line views/vteam/Detail.html:130
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vteam/Detail.html:132
		vsprint.StreamTable(qw422016, p.RelSprintsByTeamID, nil, p.Params, as, ps)
//line views/vteam/Detail.html:132
		qw422016.N().S(`
          </div>
`)
//line views/vteam/Detail.html:134
	}
//line views/vteam/Detail.html:134
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupsByTeamID" type="checkbox" hidden />
        <label for="accordion-StandupsByTeamID">
          `)
//line views/vteam/Detail.html:140
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vteam/Detail.html:140
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:141
	components.StreamSVGRef(qw422016, `standup`, 16, 16, `icon`, ps)
//line views/vteam/Detail.html:141
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:142
	qw422016.E().S(util.StringPlural(len(p.RelStandupsByTeamID), "Standup"))
//line views/vteam/Detail.html:142
	qw422016.N().S(` by [Team ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vteam/Detail.html:145
	if len(p.RelStandupsByTeamID) == 0 {
//line views/vteam/Detail.html:145
		qw422016.N().S(`          <em>no related Standups</em>
`)
//line views/vteam/Detail.html:147
	} else {
//line views/vteam/Detail.html:147
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vteam/Detail.html:149
		vstandup.StreamTable(qw422016, p.RelStandupsByTeamID, nil, nil, p.Params, as, ps)
//line views/vteam/Detail.html:149
		qw422016.N().S(`
          </div>
`)
//line views/vteam/Detail.html:151
	}
//line views/vteam/Detail.html:151
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-TeamHistoriesByTeamID" type="checkbox" hidden />
        <label for="accordion-TeamHistoriesByTeamID">
          `)
//line views/vteam/Detail.html:157
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vteam/Detail.html:157
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:158
	components.StreamSVGRef(qw422016, `history`, 16, 16, `icon`, ps)
//line views/vteam/Detail.html:158
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:159
	qw422016.E().S(util.StringPlural(len(p.RelTeamHistoriesByTeamID), "History"))
//line views/vteam/Detail.html:159
	qw422016.N().S(` by [Team ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vteam/Detail.html:162
	if len(p.RelTeamHistoriesByTeamID) == 0 {
//line views/vteam/Detail.html:162
		qw422016.N().S(`          <em>no related Histories</em>
`)
//line views/vteam/Detail.html:164
	} else {
//line views/vteam/Detail.html:164
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vteam/Detail.html:166
		vthistory.StreamTable(qw422016, p.RelTeamHistoriesByTeamID, nil, p.Params, as, ps)
//line views/vteam/Detail.html:166
		qw422016.N().S(`
          </div>
`)
//line views/vteam/Detail.html:168
	}
//line views/vteam/Detail.html:168
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-TeamMembersByTeamID" type="checkbox" hidden />
        <label for="accordion-TeamMembersByTeamID">
          `)
//line views/vteam/Detail.html:174
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vteam/Detail.html:174
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:175
	components.StreamSVGRef(qw422016, `users`, 16, 16, `icon`, ps)
//line views/vteam/Detail.html:175
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:176
	qw422016.E().S(util.StringPlural(len(p.RelTeamMembersByTeamID), "Member"))
//line views/vteam/Detail.html:176
	qw422016.N().S(` by [Team ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vteam/Detail.html:179
	if len(p.RelTeamMembersByTeamID) == 0 {
//line views/vteam/Detail.html:179
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vteam/Detail.html:181
	} else {
//line views/vteam/Detail.html:181
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vteam/Detail.html:183
		vtmember.StreamTable(qw422016, p.RelTeamMembersByTeamID, nil, nil, p.Params, as, ps)
//line views/vteam/Detail.html:183
		qw422016.N().S(`
          </div>
`)
//line views/vteam/Detail.html:185
	}
//line views/vteam/Detail.html:185
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-TeamPermissionsByTeamID" type="checkbox" hidden />
        <label for="accordion-TeamPermissionsByTeamID">
          `)
//line views/vteam/Detail.html:191
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vteam/Detail.html:191
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:192
	components.StreamSVGRef(qw422016, `permission`, 16, 16, `icon`, ps)
//line views/vteam/Detail.html:192
	qw422016.N().S(`
          `)
//line views/vteam/Detail.html:193
	qw422016.E().S(util.StringPlural(len(p.RelTeamPermissionsByTeamID), "Permission"))
//line views/vteam/Detail.html:193
	qw422016.N().S(` by [Team ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vteam/Detail.html:196
	if len(p.RelTeamPermissionsByTeamID) == 0 {
//line views/vteam/Detail.html:196
		qw422016.N().S(`          <em>no related Permissions</em>
`)
//line views/vteam/Detail.html:198
	} else {
//line views/vteam/Detail.html:198
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vteam/Detail.html:200
		vtpermission.StreamTable(qw422016, p.RelTeamPermissionsByTeamID, nil, p.Params, as, ps)
//line views/vteam/Detail.html:200
		qw422016.N().S(`
          </div>
`)
//line views/vteam/Detail.html:202
	}
//line views/vteam/Detail.html:202
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
  `)
//line views/vteam/Detail.html:207
	components.StreamJSONModal(qw422016, "team", "Team JSON", p.Model, 1)
//line views/vteam/Detail.html:207
	qw422016.N().S(`
`)
//line views/vteam/Detail.html:208
}

//line views/vteam/Detail.html:208
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vteam/Detail.html:208
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vteam/Detail.html:208
	p.StreamBody(qw422016, as, ps)
//line views/vteam/Detail.html:208
	qt422016.ReleaseWriter(qw422016)
//line views/vteam/Detail.html:208
}

//line views/vteam/Detail.html:208
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vteam/Detail.html:208
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vteam/Detail.html:208
	p.WriteBody(qb422016, as, ps)
//line views/vteam/Detail.html:208
	qs422016 := string(qb422016.B)
//line views/vteam/Detail.html:208
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vteam/Detail.html:208
	return qs422016
//line views/vteam/Detail.html:208
}
