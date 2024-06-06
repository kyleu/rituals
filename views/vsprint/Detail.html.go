// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsprint/Detail.html:2
package vsprint

//line views/vsprint/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
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
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vestimate"
	"github.com/kyleu/rituals/views/vretro"
	"github.com/kyleu/rituals/views/vsprint/vshistory"
	"github.com/kyleu/rituals/views/vsprint/vsmember"
	"github.com/kyleu/rituals/views/vsprint/vspermission"
	"github.com/kyleu/rituals/views/vstandup"
)

//line views/vsprint/Detail.html:27
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/Detail.html:27
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/Detail.html:27
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

//line views/vsprint/Detail.html:40
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/Detail.html:40
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-sprint"><button type="button">`)
//line views/vsprint/Detail.html:43
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vsprint/Detail.html:43
	qw422016.N().S(`JSON</button></a>
      <a href="`)
//line views/vsprint/Detail.html:44
	qw422016.E().S(p.Model.WebPath())
//line views/vsprint/Detail.html:44
	qw422016.N().S(`/edit"><button>`)
//line views/vsprint/Detail.html:44
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vsprint/Detail.html:44
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vsprint/Detail.html:46
	components.StreamSVGIcon(qw422016, `sprint`, ps)
//line views/vsprint/Detail.html:46
	qw422016.E().S(p.Model.TitleString())
//line views/vsprint/Detail.html:46
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/sprint"><em>Sprint</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
            <td>`)
//line views/vsprint/Detail.html:53
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vsprint/Detail.html:53
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Slug</th>
            <td>`)
//line views/vsprint/Detail.html:57
	view.StreamString(qw422016, p.Model.Slug)
//line views/vsprint/Detail.html:57
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Title</th>
            <td><strong>`)
//line views/vsprint/Detail.html:61
	view.StreamString(qw422016, p.Model.Title)
//line views/vsprint/Detail.html:61
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Icon</th>
            <td>`)
//line views/vsprint/Detail.html:65
	view.StreamString(qw422016, p.Model.Icon)
//line views/vsprint/Detail.html:65
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="`)
//line views/vsprint/Detail.html:68
	qw422016.E().S(enum.AllSessionStatuses.Help())
//line views/vsprint/Detail.html:68
	qw422016.N().S(`">Status</th>
            <td>`)
//line views/vsprint/Detail.html:69
	qw422016.E().S(p.Model.Status.String())
//line views/vsprint/Detail.html:69
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
            <td class="nowrap">
              `)
//line views/vsprint/Detail.html:74
	view.StreamUUID(qw422016, p.Model.TeamID)
//line views/vsprint/Detail.html:74
	if p.TeamByTeamID != nil {
//line views/vsprint/Detail.html:74
		qw422016.N().S(` (`)
//line views/vsprint/Detail.html:74
		qw422016.E().S(p.TeamByTeamID.TitleString())
//line views/vsprint/Detail.html:74
		qw422016.N().S(`)`)
//line views/vsprint/Detail.html:74
	}
//line views/vsprint/Detail.html:74
	qw422016.N().S(`
              `)
//line views/vsprint/Detail.html:75
	if p.Model.TeamID != nil {
//line views/vsprint/Detail.html:75
		qw422016.N().S(`<a title="Team" href="`)
//line views/vsprint/Detail.html:75
		qw422016.E().S(`/admin/db/team` + `/` + p.Model.TeamID.String())
//line views/vsprint/Detail.html:75
		qw422016.N().S(`">`)
//line views/vsprint/Detail.html:75
		components.StreamSVGSimple(qw422016, "team", 18, ps)
//line views/vsprint/Detail.html:75
		qw422016.N().S(`</a>`)
//line views/vsprint/Detail.html:75
	}
//line views/vsprint/Detail.html:75
	qw422016.N().S(`
            </td>
          </tr>
          <tr>
            <th class="shrink" title="Calendar date (optional)">Start Date</th>
            <td>`)
//line views/vsprint/Detail.html:80
	view.StreamTimestampDay(qw422016, p.Model.StartDate)
//line views/vsprint/Detail.html:80
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Calendar date (optional)">End Date</th>
            <td>`)
//line views/vsprint/Detail.html:84
	view.StreamTimestampDay(qw422016, p.Model.EndDate)
//line views/vsprint/Detail.html:84
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vsprint/Detail.html:88
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vsprint/Detail.html:88
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
            <td>`)
//line views/vsprint/Detail.html:92
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vsprint/Detail.html:92
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vsprint/Detail.html:100
	relationHelper := sprint.Sprints{p.Model}

//line views/vsprint/Detail.html:100
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-EstimatesBySprintID" type="checkbox" hidden="hidden"`)
//line views/vsprint/Detail.html:105
	if p.Params.Specifies(`estimate`) {
//line views/vsprint/Detail.html:105
		qw422016.N().S(` checked="checked"`)
//line views/vsprint/Detail.html:105
	}
//line views/vsprint/Detail.html:105
	qw422016.N().S(` />
        <label for="accordion-EstimatesBySprintID">
          `)
//line views/vsprint/Detail.html:107
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsprint/Detail.html:107
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:108
	components.StreamSVGRef(qw422016, `estimate`, 16, 16, `icon`, ps)
//line views/vsprint/Detail.html:108
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:109
	qw422016.E().S(util.StringPlural(len(p.RelEstimatesBySprintID), "Estimate"))
//line views/vsprint/Detail.html:109
	qw422016.N().S(` by [Sprint ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vsprint/Detail.html:112
	if len(p.RelEstimatesBySprintID) == 0 {
//line views/vsprint/Detail.html:112
		qw422016.N().S(`          <em>no related Estimates</em>
`)
//line views/vsprint/Detail.html:114
	} else {
//line views/vsprint/Detail.html:114
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vsprint/Detail.html:116
		vestimate.StreamTable(qw422016, p.RelEstimatesBySprintID, nil, relationHelper, p.Params, as, ps)
//line views/vsprint/Detail.html:116
		qw422016.N().S(`
          </div>
`)
//line views/vsprint/Detail.html:118
	}
//line views/vsprint/Detail.html:118
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-RetrosBySprintID" type="checkbox" hidden="hidden"`)
//line views/vsprint/Detail.html:122
	if p.Params.Specifies(`retro`) {
//line views/vsprint/Detail.html:122
		qw422016.N().S(` checked="checked"`)
//line views/vsprint/Detail.html:122
	}
//line views/vsprint/Detail.html:122
	qw422016.N().S(` />
        <label for="accordion-RetrosBySprintID">
          `)
//line views/vsprint/Detail.html:124
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsprint/Detail.html:124
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:125
	components.StreamSVGRef(qw422016, `retro`, 16, 16, `icon`, ps)
//line views/vsprint/Detail.html:125
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:126
	qw422016.E().S(util.StringPlural(len(p.RelRetrosBySprintID), "Retro"))
//line views/vsprint/Detail.html:126
	qw422016.N().S(` by [Sprint ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vsprint/Detail.html:129
	if len(p.RelRetrosBySprintID) == 0 {
//line views/vsprint/Detail.html:129
		qw422016.N().S(`          <em>no related Retros</em>
`)
//line views/vsprint/Detail.html:131
	} else {
//line views/vsprint/Detail.html:131
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vsprint/Detail.html:133
		vretro.StreamTable(qw422016, p.RelRetrosBySprintID, nil, relationHelper, p.Params, as, ps)
//line views/vsprint/Detail.html:133
		qw422016.N().S(`
          </div>
`)
//line views/vsprint/Detail.html:135
	}
//line views/vsprint/Detail.html:135
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-SprintHistoriesBySprintID" type="checkbox" hidden="hidden"`)
//line views/vsprint/Detail.html:139
	if p.Params.Specifies(`shistory`) {
//line views/vsprint/Detail.html:139
		qw422016.N().S(` checked="checked"`)
//line views/vsprint/Detail.html:139
	}
//line views/vsprint/Detail.html:139
	qw422016.N().S(` />
        <label for="accordion-SprintHistoriesBySprintID">
          `)
//line views/vsprint/Detail.html:141
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsprint/Detail.html:141
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:142
	components.StreamSVGRef(qw422016, `history`, 16, 16, `icon`, ps)
//line views/vsprint/Detail.html:142
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:143
	qw422016.E().S(util.StringPlural(len(p.RelSprintHistoriesBySprintID), "History"))
//line views/vsprint/Detail.html:143
	qw422016.N().S(` by [Sprint ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vsprint/Detail.html:146
	if len(p.RelSprintHistoriesBySprintID) == 0 {
//line views/vsprint/Detail.html:146
		qw422016.N().S(`          <em>no related Histories</em>
`)
//line views/vsprint/Detail.html:148
	} else {
//line views/vsprint/Detail.html:148
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vsprint/Detail.html:150
		vshistory.StreamTable(qw422016, p.RelSprintHistoriesBySprintID, relationHelper, p.Params, as, ps)
//line views/vsprint/Detail.html:150
		qw422016.N().S(`
          </div>
`)
//line views/vsprint/Detail.html:152
	}
//line views/vsprint/Detail.html:152
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-SprintMembersBySprintID" type="checkbox" hidden="hidden"`)
//line views/vsprint/Detail.html:156
	if p.Params.Specifies(`smember`) {
//line views/vsprint/Detail.html:156
		qw422016.N().S(` checked="checked"`)
//line views/vsprint/Detail.html:156
	}
//line views/vsprint/Detail.html:156
	qw422016.N().S(` />
        <label for="accordion-SprintMembersBySprintID">
          `)
//line views/vsprint/Detail.html:158
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsprint/Detail.html:158
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:159
	components.StreamSVGRef(qw422016, `users`, 16, 16, `icon`, ps)
//line views/vsprint/Detail.html:159
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:160
	qw422016.E().S(util.StringPlural(len(p.RelSprintMembersBySprintID), "Member"))
//line views/vsprint/Detail.html:160
	qw422016.N().S(` by [Sprint ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vsprint/Detail.html:163
	if len(p.RelSprintMembersBySprintID) == 0 {
//line views/vsprint/Detail.html:163
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vsprint/Detail.html:165
	} else {
//line views/vsprint/Detail.html:165
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vsprint/Detail.html:167
		vsmember.StreamTable(qw422016, p.RelSprintMembersBySprintID, relationHelper, nil, p.Params, as, ps)
//line views/vsprint/Detail.html:167
		qw422016.N().S(`
          </div>
`)
//line views/vsprint/Detail.html:169
	}
//line views/vsprint/Detail.html:169
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-SprintPermissionsBySprintID" type="checkbox" hidden="hidden"`)
//line views/vsprint/Detail.html:173
	if p.Params.Specifies(`spermission`) {
//line views/vsprint/Detail.html:173
		qw422016.N().S(` checked="checked"`)
//line views/vsprint/Detail.html:173
	}
//line views/vsprint/Detail.html:173
	qw422016.N().S(` />
        <label for="accordion-SprintPermissionsBySprintID">
          `)
//line views/vsprint/Detail.html:175
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsprint/Detail.html:175
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:176
	components.StreamSVGRef(qw422016, `permission`, 16, 16, `icon`, ps)
//line views/vsprint/Detail.html:176
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:177
	qw422016.E().S(util.StringPlural(len(p.RelSprintPermissionsBySprintID), "Permission"))
//line views/vsprint/Detail.html:177
	qw422016.N().S(` by [Sprint ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vsprint/Detail.html:180
	if len(p.RelSprintPermissionsBySprintID) == 0 {
//line views/vsprint/Detail.html:180
		qw422016.N().S(`          <em>no related Permissions</em>
`)
//line views/vsprint/Detail.html:182
	} else {
//line views/vsprint/Detail.html:182
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vsprint/Detail.html:184
		vspermission.StreamTable(qw422016, p.RelSprintPermissionsBySprintID, relationHelper, p.Params, as, ps)
//line views/vsprint/Detail.html:184
		qw422016.N().S(`
          </div>
`)
//line views/vsprint/Detail.html:186
	}
//line views/vsprint/Detail.html:186
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupsBySprintID" type="checkbox" hidden="hidden"`)
//line views/vsprint/Detail.html:190
	if p.Params.Specifies(`standup`) {
//line views/vsprint/Detail.html:190
		qw422016.N().S(` checked="checked"`)
//line views/vsprint/Detail.html:190
	}
//line views/vsprint/Detail.html:190
	qw422016.N().S(` />
        <label for="accordion-StandupsBySprintID">
          `)
//line views/vsprint/Detail.html:192
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vsprint/Detail.html:192
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:193
	components.StreamSVGRef(qw422016, `standup`, 16, 16, `icon`, ps)
//line views/vsprint/Detail.html:193
	qw422016.N().S(`
          `)
//line views/vsprint/Detail.html:194
	qw422016.E().S(util.StringPlural(len(p.RelStandupsBySprintID), "Standup"))
//line views/vsprint/Detail.html:194
	qw422016.N().S(` by [Sprint ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vsprint/Detail.html:197
	if len(p.RelStandupsBySprintID) == 0 {
//line views/vsprint/Detail.html:197
		qw422016.N().S(`          <em>no related Standups</em>
`)
//line views/vsprint/Detail.html:199
	} else {
//line views/vsprint/Detail.html:199
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vsprint/Detail.html:201
		vstandup.StreamTable(qw422016, p.RelStandupsBySprintID, nil, relationHelper, p.Params, as, ps)
//line views/vsprint/Detail.html:201
		qw422016.N().S(`
          </div>
`)
//line views/vsprint/Detail.html:203
	}
//line views/vsprint/Detail.html:203
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
  `)
//line views/vsprint/Detail.html:208
	components.StreamJSONModal(qw422016, "sprint", "Sprint JSON", p.Model, 1)
//line views/vsprint/Detail.html:208
	qw422016.N().S(`
`)
//line views/vsprint/Detail.html:209
}

//line views/vsprint/Detail.html:209
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/Detail.html:209
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/Detail.html:209
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/Detail.html:209
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/Detail.html:209
}

//line views/vsprint/Detail.html:209
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/Detail.html:209
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/Detail.html:209
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/Detail.html:209
	qs422016 := string(qb422016.B)
//line views/vsprint/Detail.html:209
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/Detail.html:209
	return qs422016
//line views/vsprint/Detail.html:209
}
