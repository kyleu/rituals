// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vstandup/Detail.html:1
package vstandup

//line views/vstandup/Detail.html:1
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

//line views/vstandup/Detail.html:23
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vstandup/Detail.html:23
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vstandup/Detail.html:23
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

//line views/vstandup/Detail.html:35
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/Detail.html:35
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-standup"><button type="button">`)
//line views/vstandup/Detail.html:38
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vstandup/Detail.html:38
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vstandup/Detail.html:39
	qw422016.E().S(p.Model.WebPath())
//line views/vstandup/Detail.html:39
	qw422016.N().S(`/edit"><button>`)
//line views/vstandup/Detail.html:39
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vstandup/Detail.html:39
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vstandup/Detail.html:41
	components.StreamSVGIcon(qw422016, `standup`, ps)
//line views/vstandup/Detail.html:41
	qw422016.N().S(` `)
//line views/vstandup/Detail.html:41
	qw422016.E().S(p.Model.TitleString())
//line views/vstandup/Detail.html:41
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/standup"><em>Standup</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
            <td>`)
//line views/vstandup/Detail.html:48
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vstandup/Detail.html:48
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Slug</th>
            <td>`)
//line views/vstandup/Detail.html:52
	view.StreamString(qw422016, p.Model.Slug)
//line views/vstandup/Detail.html:52
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Title</th>
            <td><strong>`)
//line views/vstandup/Detail.html:56
	view.StreamString(qw422016, p.Model.Title)
//line views/vstandup/Detail.html:56
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Icon</th>
            <td>`)
//line views/vstandup/Detail.html:60
	view.StreamString(qw422016, p.Model.Icon)
//line views/vstandup/Detail.html:60
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="`)
//line views/vstandup/Detail.html:63
	qw422016.E().S(enum.AllSessionStatuses.Help())
//line views/vstandup/Detail.html:63
	qw422016.N().S(`">Status</th>
            <td>`)
//line views/vstandup/Detail.html:64
	qw422016.E().S(p.Model.Status.String())
//line views/vstandup/Detail.html:64
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
            <td class="nowrap">
              `)
//line views/vstandup/Detail.html:69
	view.StreamUUID(qw422016, p.Model.TeamID)
//line views/vstandup/Detail.html:69
	if p.TeamByTeamID != nil {
//line views/vstandup/Detail.html:69
		qw422016.N().S(` (`)
//line views/vstandup/Detail.html:69
		qw422016.E().S(p.TeamByTeamID.TitleString())
//line views/vstandup/Detail.html:69
		qw422016.N().S(`)`)
//line views/vstandup/Detail.html:69
	}
//line views/vstandup/Detail.html:69
	qw422016.N().S(`
              `)
//line views/vstandup/Detail.html:70
	if p.Model.TeamID != nil {
//line views/vstandup/Detail.html:70
		qw422016.N().S(`<a title="Team" href="`)
//line views/vstandup/Detail.html:70
		qw422016.E().S(`/admin/db/team` + `/` + p.Model.TeamID.String())
//line views/vstandup/Detail.html:70
		qw422016.N().S(`">`)
//line views/vstandup/Detail.html:70
		components.StreamSVGLink(qw422016, `team`, ps)
//line views/vstandup/Detail.html:70
		qw422016.N().S(`</a>`)
//line views/vstandup/Detail.html:70
	}
//line views/vstandup/Detail.html:70
	qw422016.N().S(`
            </td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Sprint ID</th>
            <td class="nowrap">
              `)
//line views/vstandup/Detail.html:76
	view.StreamUUID(qw422016, p.Model.SprintID)
//line views/vstandup/Detail.html:76
	if p.SprintBySprintID != nil {
//line views/vstandup/Detail.html:76
		qw422016.N().S(` (`)
//line views/vstandup/Detail.html:76
		qw422016.E().S(p.SprintBySprintID.TitleString())
//line views/vstandup/Detail.html:76
		qw422016.N().S(`)`)
//line views/vstandup/Detail.html:76
	}
//line views/vstandup/Detail.html:76
	qw422016.N().S(`
              `)
//line views/vstandup/Detail.html:77
	if p.Model.SprintID != nil {
//line views/vstandup/Detail.html:77
		qw422016.N().S(`<a title="Sprint" href="`)
//line views/vstandup/Detail.html:77
		qw422016.E().S(`/admin/db/sprint` + `/` + p.Model.SprintID.String())
//line views/vstandup/Detail.html:77
		qw422016.N().S(`">`)
//line views/vstandup/Detail.html:77
		components.StreamSVGLink(qw422016, `sprint`, ps)
//line views/vstandup/Detail.html:77
		qw422016.N().S(`</a>`)
//line views/vstandup/Detail.html:77
	}
//line views/vstandup/Detail.html:77
	qw422016.N().S(`
            </td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vstandup/Detail.html:82
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vstandup/Detail.html:82
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
            <td>`)
//line views/vstandup/Detail.html:86
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vstandup/Detail.html:86
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vstandup/Detail.html:94
	relationHelper := standup.Standups{p.Model}

//line views/vstandup/Detail.html:94
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-ReportsByStandupID" type="checkbox" hidden="hidden"`)
//line views/vstandup/Detail.html:99
	if p.Params.Specifies(`report`) {
//line views/vstandup/Detail.html:99
		qw422016.N().S(` checked="checked"`)
//line views/vstandup/Detail.html:99
	}
//line views/vstandup/Detail.html:99
	qw422016.N().S(` />
        <label for="accordion-ReportsByStandupID">
          `)
//line views/vstandup/Detail.html:101
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vstandup/Detail.html:101
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:102
	components.StreamSVGInline(qw422016, `file-alt`, 16, ps)
//line views/vstandup/Detail.html:102
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:103
	qw422016.E().S(util.StringPlural(len(p.RelReportsByStandupID), "Report"))
//line views/vstandup/Detail.html:103
	qw422016.N().S(` by [Standup ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vstandup/Detail.html:106
	if len(p.RelReportsByStandupID) == 0 {
//line views/vstandup/Detail.html:106
		qw422016.N().S(`          <em>no related Reports</em>
`)
//line views/vstandup/Detail.html:108
	} else {
//line views/vstandup/Detail.html:108
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vstandup/Detail.html:110
		vreport.StreamTable(qw422016, p.RelReportsByStandupID, relationHelper, nil, p.Params, as, ps)
//line views/vstandup/Detail.html:110
		qw422016.N().S(`
          </div>
`)
//line views/vstandup/Detail.html:112
	}
//line views/vstandup/Detail.html:112
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupHistoriesByStandupID" type="checkbox" hidden="hidden"`)
//line views/vstandup/Detail.html:116
	if p.Params.Specifies(`uhistory`) {
//line views/vstandup/Detail.html:116
		qw422016.N().S(` checked="checked"`)
//line views/vstandup/Detail.html:116
	}
//line views/vstandup/Detail.html:116
	qw422016.N().S(` />
        <label for="accordion-StandupHistoriesByStandupID">
          `)
//line views/vstandup/Detail.html:118
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vstandup/Detail.html:118
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:119
	components.StreamSVGInline(qw422016, `history`, 16, ps)
//line views/vstandup/Detail.html:119
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:120
	qw422016.E().S(util.StringPlural(len(p.RelStandupHistoriesByStandupID), "History"))
//line views/vstandup/Detail.html:120
	qw422016.N().S(` by [Standup ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vstandup/Detail.html:123
	if len(p.RelStandupHistoriesByStandupID) == 0 {
//line views/vstandup/Detail.html:123
		qw422016.N().S(`          <em>no related Histories</em>
`)
//line views/vstandup/Detail.html:125
	} else {
//line views/vstandup/Detail.html:125
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vstandup/Detail.html:127
		vuhistory.StreamTable(qw422016, p.RelStandupHistoriesByStandupID, relationHelper, p.Params, as, ps)
//line views/vstandup/Detail.html:127
		qw422016.N().S(`
          </div>
`)
//line views/vstandup/Detail.html:129
	}
//line views/vstandup/Detail.html:129
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupMembersByStandupID" type="checkbox" hidden="hidden"`)
//line views/vstandup/Detail.html:133
	if p.Params.Specifies(`umember`) {
//line views/vstandup/Detail.html:133
		qw422016.N().S(` checked="checked"`)
//line views/vstandup/Detail.html:133
	}
//line views/vstandup/Detail.html:133
	qw422016.N().S(` />
        <label for="accordion-StandupMembersByStandupID">
          `)
//line views/vstandup/Detail.html:135
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vstandup/Detail.html:135
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:136
	components.StreamSVGInline(qw422016, `users`, 16, ps)
//line views/vstandup/Detail.html:136
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:137
	qw422016.E().S(util.StringPlural(len(p.RelStandupMembersByStandupID), "Member"))
//line views/vstandup/Detail.html:137
	qw422016.N().S(` by [Standup ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vstandup/Detail.html:140
	if len(p.RelStandupMembersByStandupID) == 0 {
//line views/vstandup/Detail.html:140
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vstandup/Detail.html:142
	} else {
//line views/vstandup/Detail.html:142
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vstandup/Detail.html:144
		vumember.StreamTable(qw422016, p.RelStandupMembersByStandupID, relationHelper, nil, p.Params, as, ps)
//line views/vstandup/Detail.html:144
		qw422016.N().S(`
          </div>
`)
//line views/vstandup/Detail.html:146
	}
//line views/vstandup/Detail.html:146
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupPermissionsByStandupID" type="checkbox" hidden="hidden"`)
//line views/vstandup/Detail.html:150
	if p.Params.Specifies(`upermission`) {
//line views/vstandup/Detail.html:150
		qw422016.N().S(` checked="checked"`)
//line views/vstandup/Detail.html:150
	}
//line views/vstandup/Detail.html:150
	qw422016.N().S(` />
        <label for="accordion-StandupPermissionsByStandupID">
          `)
//line views/vstandup/Detail.html:152
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vstandup/Detail.html:152
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:153
	components.StreamSVGInline(qw422016, `permission`, 16, ps)
//line views/vstandup/Detail.html:153
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:154
	qw422016.E().S(util.StringPlural(len(p.RelStandupPermissionsByStandupID), "Permission"))
//line views/vstandup/Detail.html:154
	qw422016.N().S(` by [Standup ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vstandup/Detail.html:157
	if len(p.RelStandupPermissionsByStandupID) == 0 {
//line views/vstandup/Detail.html:157
		qw422016.N().S(`          <em>no related Permissions</em>
`)
//line views/vstandup/Detail.html:159
	} else {
//line views/vstandup/Detail.html:159
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vstandup/Detail.html:161
		vupermission.StreamTable(qw422016, p.RelStandupPermissionsByStandupID, relationHelper, p.Params, as, ps)
//line views/vstandup/Detail.html:161
		qw422016.N().S(`
          </div>
`)
//line views/vstandup/Detail.html:163
	}
//line views/vstandup/Detail.html:163
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
  `)
//line views/vstandup/Detail.html:168
	components.StreamJSONModal(qw422016, "standup", "Standup JSON", p.Model, 1)
//line views/vstandup/Detail.html:168
	qw422016.N().S(`
`)
//line views/vstandup/Detail.html:169
}

//line views/vstandup/Detail.html:169
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/Detail.html:169
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/Detail.html:169
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/Detail.html:169
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/Detail.html:169
}

//line views/vstandup/Detail.html:169
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/Detail.html:169
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/Detail.html:169
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/Detail.html:169
	qs422016 := string(qb422016.B)
//line views/vstandup/Detail.html:169
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/Detail.html:169
	return qs422016
//line views/vstandup/Detail.html:169
}
