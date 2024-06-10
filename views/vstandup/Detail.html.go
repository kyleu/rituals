// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vstandup/Detail.html:2
package vstandup

//line views/vstandup/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/standup"
	"github.com/kyleu/rituals/app/standup/report"
	"github.com/kyleu/rituals/app/standup/uhistory"
	"github.com/kyleu/rituals/app/standup/umember"
	"github.com/kyleu/rituals/app/standup/upermission"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vstandup/vreport"
	"github.com/kyleu/rituals/views/vstandup/vuhistory"
	"github.com/kyleu/rituals/views/vstandup/vumember"
	"github.com/kyleu/rituals/views/vstandup/vupermission"
)

//line views/vstandup/Detail.html:24
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/Detail.html:24
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/Detail.html:24
type Detail struct {
	layout.Basic
	Model                            *standup.Standup
	TeamByTeamID                     *team.Team
	SprintBySprintID                 *sprint.Sprint
	Params                           filter.ParamSet
	RelReportsByStandupID            report.Reports
	RelStandupHistoriesByStandupID   uhistory.StandupHistories
	RelStandupMembersByStandupID     umember.StandupMembers
	RelStandupPermissionsByStandupID upermission.StandupPermissions
}

//line views/vstandup/Detail.html:36
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/Detail.html:36
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-standup"><button type="button">`)
//line views/vstandup/Detail.html:39
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vstandup/Detail.html:39
	qw422016.N().S(`JSON</button></a>
      <a href="`)
//line views/vstandup/Detail.html:40
	qw422016.E().S(p.Model.WebPath())
//line views/vstandup/Detail.html:40
	qw422016.N().S(`/edit"><button>`)
//line views/vstandup/Detail.html:40
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vstandup/Detail.html:40
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vstandup/Detail.html:42
	components.StreamSVGIcon(qw422016, `standup`, ps)
//line views/vstandup/Detail.html:42
	qw422016.E().S(p.Model.TitleString())
//line views/vstandup/Detail.html:42
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/standup"><em>Standup</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
            <td>`)
//line views/vstandup/Detail.html:49
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vstandup/Detail.html:49
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Slug</th>
            <td>`)
//line views/vstandup/Detail.html:53
	view.StreamString(qw422016, p.Model.Slug)
//line views/vstandup/Detail.html:53
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Title</th>
            <td><strong>`)
//line views/vstandup/Detail.html:57
	view.StreamString(qw422016, p.Model.Title)
//line views/vstandup/Detail.html:57
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Icon</th>
            <td>`)
//line views/vstandup/Detail.html:61
	view.StreamString(qw422016, p.Model.Icon)
//line views/vstandup/Detail.html:61
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="`)
//line views/vstandup/Detail.html:64
	qw422016.E().S(enum.AllSessionStatuses.Help())
//line views/vstandup/Detail.html:64
	qw422016.N().S(`">Status</th>
            <td>`)
//line views/vstandup/Detail.html:65
	qw422016.E().S(p.Model.Status.String())
//line views/vstandup/Detail.html:65
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
            <td class="nowrap">
              `)
//line views/vstandup/Detail.html:70
	view.StreamUUID(qw422016, p.Model.TeamID)
//line views/vstandup/Detail.html:70
	if p.TeamByTeamID != nil {
//line views/vstandup/Detail.html:70
		qw422016.N().S(` (`)
//line views/vstandup/Detail.html:70
		qw422016.E().S(p.TeamByTeamID.TitleString())
//line views/vstandup/Detail.html:70
		qw422016.N().S(`)`)
//line views/vstandup/Detail.html:70
	}
//line views/vstandup/Detail.html:70
	qw422016.N().S(`
              `)
//line views/vstandup/Detail.html:71
	if p.Model.TeamID != nil {
//line views/vstandup/Detail.html:71
		qw422016.N().S(`<a title="Team" href="`)
//line views/vstandup/Detail.html:71
		qw422016.E().S(`/admin/db/team` + `/` + p.Model.TeamID.String())
//line views/vstandup/Detail.html:71
		qw422016.N().S(`">`)
//line views/vstandup/Detail.html:71
		components.StreamSVGSimple(qw422016, "team", 18, ps)
//line views/vstandup/Detail.html:71
		qw422016.N().S(`</a>`)
//line views/vstandup/Detail.html:71
	}
//line views/vstandup/Detail.html:71
	qw422016.N().S(`
            </td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Sprint ID</th>
            <td class="nowrap">
              `)
//line views/vstandup/Detail.html:77
	view.StreamUUID(qw422016, p.Model.SprintID)
//line views/vstandup/Detail.html:77
	if p.SprintBySprintID != nil {
//line views/vstandup/Detail.html:77
		qw422016.N().S(` (`)
//line views/vstandup/Detail.html:77
		qw422016.E().S(p.SprintBySprintID.TitleString())
//line views/vstandup/Detail.html:77
		qw422016.N().S(`)`)
//line views/vstandup/Detail.html:77
	}
//line views/vstandup/Detail.html:77
	qw422016.N().S(`
              `)
//line views/vstandup/Detail.html:78
	if p.Model.SprintID != nil {
//line views/vstandup/Detail.html:78
		qw422016.N().S(`<a title="Sprint" href="`)
//line views/vstandup/Detail.html:78
		qw422016.E().S(`/admin/db/sprint` + `/` + p.Model.SprintID.String())
//line views/vstandup/Detail.html:78
		qw422016.N().S(`">`)
//line views/vstandup/Detail.html:78
		components.StreamSVGSimple(qw422016, "sprint", 18, ps)
//line views/vstandup/Detail.html:78
		qw422016.N().S(`</a>`)
//line views/vstandup/Detail.html:78
	}
//line views/vstandup/Detail.html:78
	qw422016.N().S(`
            </td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vstandup/Detail.html:83
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vstandup/Detail.html:83
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
            <td>`)
//line views/vstandup/Detail.html:87
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vstandup/Detail.html:87
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vstandup/Detail.html:95
	relationHelper := standup.Standups{p.Model}

//line views/vstandup/Detail.html:95
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-ReportsByStandupID" type="checkbox" hidden="hidden"`)
//line views/vstandup/Detail.html:100
	if p.Params.Specifies(`report`) {
//line views/vstandup/Detail.html:100
		qw422016.N().S(` checked="checked"`)
//line views/vstandup/Detail.html:100
	}
//line views/vstandup/Detail.html:100
	qw422016.N().S(` />
        <label for="accordion-ReportsByStandupID">
          `)
//line views/vstandup/Detail.html:102
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vstandup/Detail.html:102
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:103
	components.StreamSVGInline(qw422016, `file-alt`, 16, ps)
//line views/vstandup/Detail.html:103
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:104
	qw422016.E().S(util.StringPlural(len(p.RelReportsByStandupID), "Report"))
//line views/vstandup/Detail.html:104
	qw422016.N().S(` by [Standup ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vstandup/Detail.html:107
	if len(p.RelReportsByStandupID) == 0 {
//line views/vstandup/Detail.html:107
		qw422016.N().S(`          <em>no related Reports</em>
`)
//line views/vstandup/Detail.html:109
	} else {
//line views/vstandup/Detail.html:109
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vstandup/Detail.html:111
		vreport.StreamTable(qw422016, p.RelReportsByStandupID, relationHelper, nil, p.Params, as, ps)
//line views/vstandup/Detail.html:111
		qw422016.N().S(`
          </div>
`)
//line views/vstandup/Detail.html:113
	}
//line views/vstandup/Detail.html:113
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupHistoriesByStandupID" type="checkbox" hidden="hidden"`)
//line views/vstandup/Detail.html:117
	if p.Params.Specifies(`uhistory`) {
//line views/vstandup/Detail.html:117
		qw422016.N().S(` checked="checked"`)
//line views/vstandup/Detail.html:117
	}
//line views/vstandup/Detail.html:117
	qw422016.N().S(` />
        <label for="accordion-StandupHistoriesByStandupID">
          `)
//line views/vstandup/Detail.html:119
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vstandup/Detail.html:119
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:120
	components.StreamSVGInline(qw422016, `history`, 16, ps)
//line views/vstandup/Detail.html:120
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:121
	qw422016.E().S(util.StringPlural(len(p.RelStandupHistoriesByStandupID), "History"))
//line views/vstandup/Detail.html:121
	qw422016.N().S(` by [Standup ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vstandup/Detail.html:124
	if len(p.RelStandupHistoriesByStandupID) == 0 {
//line views/vstandup/Detail.html:124
		qw422016.N().S(`          <em>no related Histories</em>
`)
//line views/vstandup/Detail.html:126
	} else {
//line views/vstandup/Detail.html:126
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vstandup/Detail.html:128
		vuhistory.StreamTable(qw422016, p.RelStandupHistoriesByStandupID, relationHelper, p.Params, as, ps)
//line views/vstandup/Detail.html:128
		qw422016.N().S(`
          </div>
`)
//line views/vstandup/Detail.html:130
	}
//line views/vstandup/Detail.html:130
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupMembersByStandupID" type="checkbox" hidden="hidden"`)
//line views/vstandup/Detail.html:134
	if p.Params.Specifies(`umember`) {
//line views/vstandup/Detail.html:134
		qw422016.N().S(` checked="checked"`)
//line views/vstandup/Detail.html:134
	}
//line views/vstandup/Detail.html:134
	qw422016.N().S(` />
        <label for="accordion-StandupMembersByStandupID">
          `)
//line views/vstandup/Detail.html:136
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vstandup/Detail.html:136
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:137
	components.StreamSVGInline(qw422016, `users`, 16, ps)
//line views/vstandup/Detail.html:137
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:138
	qw422016.E().S(util.StringPlural(len(p.RelStandupMembersByStandupID), "Member"))
//line views/vstandup/Detail.html:138
	qw422016.N().S(` by [Standup ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vstandup/Detail.html:141
	if len(p.RelStandupMembersByStandupID) == 0 {
//line views/vstandup/Detail.html:141
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vstandup/Detail.html:143
	} else {
//line views/vstandup/Detail.html:143
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vstandup/Detail.html:145
		vumember.StreamTable(qw422016, p.RelStandupMembersByStandupID, relationHelper, nil, p.Params, as, ps)
//line views/vstandup/Detail.html:145
		qw422016.N().S(`
          </div>
`)
//line views/vstandup/Detail.html:147
	}
//line views/vstandup/Detail.html:147
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupPermissionsByStandupID" type="checkbox" hidden="hidden"`)
//line views/vstandup/Detail.html:151
	if p.Params.Specifies(`upermission`) {
//line views/vstandup/Detail.html:151
		qw422016.N().S(` checked="checked"`)
//line views/vstandup/Detail.html:151
	}
//line views/vstandup/Detail.html:151
	qw422016.N().S(` />
        <label for="accordion-StandupPermissionsByStandupID">
          `)
//line views/vstandup/Detail.html:153
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vstandup/Detail.html:153
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:154
	components.StreamSVGInline(qw422016, `permission`, 16, ps)
//line views/vstandup/Detail.html:154
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:155
	qw422016.E().S(util.StringPlural(len(p.RelStandupPermissionsByStandupID), "Permission"))
//line views/vstandup/Detail.html:155
	qw422016.N().S(` by [Standup ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vstandup/Detail.html:158
	if len(p.RelStandupPermissionsByStandupID) == 0 {
//line views/vstandup/Detail.html:158
		qw422016.N().S(`          <em>no related Permissions</em>
`)
//line views/vstandup/Detail.html:160
	} else {
//line views/vstandup/Detail.html:160
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vstandup/Detail.html:162
		vupermission.StreamTable(qw422016, p.RelStandupPermissionsByStandupID, relationHelper, p.Params, as, ps)
//line views/vstandup/Detail.html:162
		qw422016.N().S(`
          </div>
`)
//line views/vstandup/Detail.html:164
	}
//line views/vstandup/Detail.html:164
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
  `)
//line views/vstandup/Detail.html:169
	components.StreamJSONModal(qw422016, "standup", "Standup JSON", p.Model, 1)
//line views/vstandup/Detail.html:169
	qw422016.N().S(`
`)
//line views/vstandup/Detail.html:170
}

//line views/vstandup/Detail.html:170
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/Detail.html:170
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/Detail.html:170
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/Detail.html:170
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/Detail.html:170
}

//line views/vstandup/Detail.html:170
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/Detail.html:170
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/Detail.html:170
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/Detail.html:170
	qs422016 := string(qb422016.B)
//line views/vstandup/Detail.html:170
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/Detail.html:170
	return qs422016
//line views/vstandup/Detail.html:170
}
