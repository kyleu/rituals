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
	Paths                            []string
}

//line views/vstandup/Detail.html:36
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/Detail.html:36
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-standup"><button type="button" title="JSON">`)
//line views/vstandup/Detail.html:39
	components.StreamSVGButton(qw422016, "code", ps)
//line views/vstandup/Detail.html:39
	qw422016.N().S(`</button></a>
      <a href="`)
//line views/vstandup/Detail.html:40
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vstandup/Detail.html:40
	qw422016.N().S(`/edit" title="Edit"><button>`)
//line views/vstandup/Detail.html:40
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vstandup/Detail.html:40
	qw422016.N().S(`</button></a>
    </div>
    <h3>`)
//line views/vstandup/Detail.html:42
	components.StreamSVGIcon(qw422016, `standup`, ps)
//line views/vstandup/Detail.html:42
	qw422016.N().S(` `)
//line views/vstandup/Detail.html:42
	qw422016.E().S(p.Model.TitleString())
//line views/vstandup/Detail.html:42
	qw422016.N().S(`</h3>
    <div><a href="`)
//line views/vstandup/Detail.html:43
	qw422016.E().S(standup.Route(p.Paths...))
//line views/vstandup/Detail.html:43
	qw422016.N().S(`"><em>Standup</em></a></div>
    `)
//line views/vstandup/Detail.html:44
	StreamDetailTable(qw422016, p, ps)
//line views/vstandup/Detail.html:44
	qw422016.N().S(`
  </div>
`)
//line views/vstandup/Detail.html:47
	qw422016.N().S(`  `)
//line views/vstandup/Detail.html:48
	StreamDetailRelations(qw422016, as, p, ps)
//line views/vstandup/Detail.html:48
	qw422016.N().S(`
  `)
//line views/vstandup/Detail.html:49
	components.StreamJSONModal(qw422016, "standup", "Standup JSON", p.Model, 1)
//line views/vstandup/Detail.html:49
	qw422016.N().S(`
`)
//line views/vstandup/Detail.html:50
}

//line views/vstandup/Detail.html:50
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vstandup/Detail.html:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/Detail.html:50
	p.StreamBody(qw422016, as, ps)
//line views/vstandup/Detail.html:50
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/Detail.html:50
}

//line views/vstandup/Detail.html:50
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vstandup/Detail.html:50
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/Detail.html:50
	p.WriteBody(qb422016, as, ps)
//line views/vstandup/Detail.html:50
	qs422016 := string(qb422016.B)
//line views/vstandup/Detail.html:50
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/Detail.html:50
	return qs422016
//line views/vstandup/Detail.html:50
}

//line views/vstandup/Detail.html:52
func StreamDetailTable(qw422016 *qt422016.Writer, p *Detail, ps *cutil.PageState) {
//line views/vstandup/Detail.html:52
	qw422016.N().S(`
  <div class="mt overflow full-width">
    <table>
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vstandup/Detail.html:58
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vstandup/Detail.html:58
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Slug</th>
          <td>`)
//line views/vstandup/Detail.html:62
	view.StreamString(qw422016, p.Model.Slug)
//line views/vstandup/Detail.html:62
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Title</th>
          <td><strong>`)
//line views/vstandup/Detail.html:66
	view.StreamString(qw422016, p.Model.Title)
//line views/vstandup/Detail.html:66
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Icon</th>
          <td>`)
//line views/vstandup/Detail.html:70
	view.StreamString(qw422016, p.Model.Icon)
//line views/vstandup/Detail.html:70
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="`)
//line views/vstandup/Detail.html:73
	qw422016.E().S(enum.AllSessionStatuses.Help())
//line views/vstandup/Detail.html:73
	qw422016.N().S(`">Status</th>
          <td>`)
//line views/vstandup/Detail.html:74
	qw422016.E().S(p.Model.Status.String())
//line views/vstandup/Detail.html:74
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
          <td class="nowrap">
            `)
//line views/vstandup/Detail.html:79
	view.StreamUUID(qw422016, p.Model.TeamID)
//line views/vstandup/Detail.html:79
	if p.TeamByTeamID != nil {
//line views/vstandup/Detail.html:79
		qw422016.N().S(` (`)
//line views/vstandup/Detail.html:79
		qw422016.E().S(p.TeamByTeamID.TitleString())
//line views/vstandup/Detail.html:79
		qw422016.N().S(`)`)
//line views/vstandup/Detail.html:79
	}
//line views/vstandup/Detail.html:79
	qw422016.N().S(`
            `)
//line views/vstandup/Detail.html:80
	if p.Model.TeamID != nil {
//line views/vstandup/Detail.html:80
		qw422016.N().S(`<a title="Team" href="`)
//line views/vstandup/Detail.html:80
		if x := p.TeamByTeamID; x != nil {
//line views/vstandup/Detail.html:80
			qw422016.E().S(x.WebPath(p.Paths...))
//line views/vstandup/Detail.html:80
		}
//line views/vstandup/Detail.html:80
		qw422016.N().S(`">`)
//line views/vstandup/Detail.html:80
		components.StreamSVGLink(qw422016, `team`, ps)
//line views/vstandup/Detail.html:80
		qw422016.N().S(`</a>`)
//line views/vstandup/Detail.html:80
	}
//line views/vstandup/Detail.html:80
	qw422016.N().S(`
          </td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Sprint ID</th>
          <td class="nowrap">
            `)
//line views/vstandup/Detail.html:86
	view.StreamUUID(qw422016, p.Model.SprintID)
//line views/vstandup/Detail.html:86
	if p.SprintBySprintID != nil {
//line views/vstandup/Detail.html:86
		qw422016.N().S(` (`)
//line views/vstandup/Detail.html:86
		qw422016.E().S(p.SprintBySprintID.TitleString())
//line views/vstandup/Detail.html:86
		qw422016.N().S(`)`)
//line views/vstandup/Detail.html:86
	}
//line views/vstandup/Detail.html:86
	qw422016.N().S(`
            `)
//line views/vstandup/Detail.html:87
	if p.Model.SprintID != nil {
//line views/vstandup/Detail.html:87
		qw422016.N().S(`<a title="Sprint" href="`)
//line views/vstandup/Detail.html:87
		if x := p.SprintBySprintID; x != nil {
//line views/vstandup/Detail.html:87
			qw422016.E().S(x.WebPath(p.Paths...))
//line views/vstandup/Detail.html:87
		}
//line views/vstandup/Detail.html:87
		qw422016.N().S(`">`)
//line views/vstandup/Detail.html:87
		components.StreamSVGLink(qw422016, `sprint`, ps)
//line views/vstandup/Detail.html:87
		qw422016.N().S(`</a>`)
//line views/vstandup/Detail.html:87
	}
//line views/vstandup/Detail.html:87
	qw422016.N().S(`
          </td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vstandup/Detail.html:92
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vstandup/Detail.html:92
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vstandup/Detail.html:96
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vstandup/Detail.html:96
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vstandup/Detail.html:101
}

//line views/vstandup/Detail.html:101
func WriteDetailTable(qq422016 qtio422016.Writer, p *Detail, ps *cutil.PageState) {
//line views/vstandup/Detail.html:101
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/Detail.html:101
	StreamDetailTable(qw422016, p, ps)
//line views/vstandup/Detail.html:101
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/Detail.html:101
}

//line views/vstandup/Detail.html:101
func DetailTable(p *Detail, ps *cutil.PageState) string {
//line views/vstandup/Detail.html:101
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/Detail.html:101
	WriteDetailTable(qb422016, p, ps)
//line views/vstandup/Detail.html:101
	qs422016 := string(qb422016.B)
//line views/vstandup/Detail.html:101
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/Detail.html:101
	return qs422016
//line views/vstandup/Detail.html:101
}

//line views/vstandup/Detail.html:103
func StreamDetailRelations(qw422016 *qt422016.Writer, as *app.State, p *Detail, ps *cutil.PageState) {
//line views/vstandup/Detail.html:103
	qw422016.N().S(`
`)
//line views/vstandup/Detail.html:104
	relationHelper := standup.Standups{p.Model}

//line views/vstandup/Detail.html:104
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-ReportsByStandupID" type="checkbox" hidden="hidden"`)
//line views/vstandup/Detail.html:109
	if p.Params.Specifies(`report`) {
//line views/vstandup/Detail.html:109
		qw422016.N().S(` checked="checked"`)
//line views/vstandup/Detail.html:109
	}
//line views/vstandup/Detail.html:109
	qw422016.N().S(` />
        <label for="accordion-ReportsByStandupID">
          `)
//line views/vstandup/Detail.html:111
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vstandup/Detail.html:111
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:112
	components.StreamSVGInline(qw422016, `file-alt`, 16, ps)
//line views/vstandup/Detail.html:112
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:113
	qw422016.E().S(util.StringPlural(len(p.RelReportsByStandupID), "Report"))
//line views/vstandup/Detail.html:113
	qw422016.N().S(` by [Standup ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vstandup/Detail.html:116
	if len(p.RelReportsByStandupID) == 0 {
//line views/vstandup/Detail.html:116
		qw422016.N().S(`          <em>no related Reports</em>
`)
//line views/vstandup/Detail.html:118
	} else {
//line views/vstandup/Detail.html:118
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vstandup/Detail.html:120
		vreport.StreamTable(qw422016, p.RelReportsByStandupID, relationHelper, nil, p.Params, as, ps)
//line views/vstandup/Detail.html:120
		qw422016.N().S(`
          </div>
`)
//line views/vstandup/Detail.html:122
	}
//line views/vstandup/Detail.html:122
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupHistoriesByStandupID" type="checkbox" hidden="hidden"`)
//line views/vstandup/Detail.html:126
	if p.Params.Specifies(`uhistory`) {
//line views/vstandup/Detail.html:126
		qw422016.N().S(` checked="checked"`)
//line views/vstandup/Detail.html:126
	}
//line views/vstandup/Detail.html:126
	qw422016.N().S(` />
        <label for="accordion-StandupHistoriesByStandupID">
          `)
//line views/vstandup/Detail.html:128
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vstandup/Detail.html:128
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:129
	components.StreamSVGInline(qw422016, `history`, 16, ps)
//line views/vstandup/Detail.html:129
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:130
	qw422016.E().S(util.StringPlural(len(p.RelStandupHistoriesByStandupID), "History"))
//line views/vstandup/Detail.html:130
	qw422016.N().S(` by [Standup ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vstandup/Detail.html:133
	if len(p.RelStandupHistoriesByStandupID) == 0 {
//line views/vstandup/Detail.html:133
		qw422016.N().S(`          <em>no related Histories</em>
`)
//line views/vstandup/Detail.html:135
	} else {
//line views/vstandup/Detail.html:135
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vstandup/Detail.html:137
		vuhistory.StreamTable(qw422016, p.RelStandupHistoriesByStandupID, relationHelper, p.Params, as, ps)
//line views/vstandup/Detail.html:137
		qw422016.N().S(`
          </div>
`)
//line views/vstandup/Detail.html:139
	}
//line views/vstandup/Detail.html:139
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupMembersByStandupID" type="checkbox" hidden="hidden"`)
//line views/vstandup/Detail.html:143
	if p.Params.Specifies(`umember`) {
//line views/vstandup/Detail.html:143
		qw422016.N().S(` checked="checked"`)
//line views/vstandup/Detail.html:143
	}
//line views/vstandup/Detail.html:143
	qw422016.N().S(` />
        <label for="accordion-StandupMembersByStandupID">
          `)
//line views/vstandup/Detail.html:145
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vstandup/Detail.html:145
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:146
	components.StreamSVGInline(qw422016, `users`, 16, ps)
//line views/vstandup/Detail.html:146
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:147
	qw422016.E().S(util.StringPlural(len(p.RelStandupMembersByStandupID), "Member"))
//line views/vstandup/Detail.html:147
	qw422016.N().S(` by [Standup ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vstandup/Detail.html:150
	if len(p.RelStandupMembersByStandupID) == 0 {
//line views/vstandup/Detail.html:150
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vstandup/Detail.html:152
	} else {
//line views/vstandup/Detail.html:152
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vstandup/Detail.html:154
		vumember.StreamTable(qw422016, p.RelStandupMembersByStandupID, relationHelper, nil, p.Params, as, ps)
//line views/vstandup/Detail.html:154
		qw422016.N().S(`
          </div>
`)
//line views/vstandup/Detail.html:156
	}
//line views/vstandup/Detail.html:156
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StandupPermissionsByStandupID" type="checkbox" hidden="hidden"`)
//line views/vstandup/Detail.html:160
	if p.Params.Specifies(`upermission`) {
//line views/vstandup/Detail.html:160
		qw422016.N().S(` checked="checked"`)
//line views/vstandup/Detail.html:160
	}
//line views/vstandup/Detail.html:160
	qw422016.N().S(` />
        <label for="accordion-StandupPermissionsByStandupID">
          `)
//line views/vstandup/Detail.html:162
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vstandup/Detail.html:162
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:163
	components.StreamSVGInline(qw422016, `permission`, 16, ps)
//line views/vstandup/Detail.html:163
	qw422016.N().S(`
          `)
//line views/vstandup/Detail.html:164
	qw422016.E().S(util.StringPlural(len(p.RelStandupPermissionsByStandupID), "Permission"))
//line views/vstandup/Detail.html:164
	qw422016.N().S(` by [Standup ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vstandup/Detail.html:167
	if len(p.RelStandupPermissionsByStandupID) == 0 {
//line views/vstandup/Detail.html:167
		qw422016.N().S(`          <em>no related Permissions</em>
`)
//line views/vstandup/Detail.html:169
	} else {
//line views/vstandup/Detail.html:169
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vstandup/Detail.html:171
		vupermission.StreamTable(qw422016, p.RelStandupPermissionsByStandupID, relationHelper, p.Params, as, ps)
//line views/vstandup/Detail.html:171
		qw422016.N().S(`
          </div>
`)
//line views/vstandup/Detail.html:173
	}
//line views/vstandup/Detail.html:173
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
`)
//line views/vstandup/Detail.html:178
}

//line views/vstandup/Detail.html:178
func WriteDetailRelations(qq422016 qtio422016.Writer, as *app.State, p *Detail, ps *cutil.PageState) {
//line views/vstandup/Detail.html:178
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vstandup/Detail.html:178
	StreamDetailRelations(qw422016, as, p, ps)
//line views/vstandup/Detail.html:178
	qt422016.ReleaseWriter(qw422016)
//line views/vstandup/Detail.html:178
}

//line views/vstandup/Detail.html:178
func DetailRelations(as *app.State, p *Detail, ps *cutil.PageState) string {
//line views/vstandup/Detail.html:178
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vstandup/Detail.html:178
	WriteDetailRelations(qb422016, as, p, ps)
//line views/vstandup/Detail.html:178
	qs422016 := string(qb422016.B)
//line views/vstandup/Detail.html:178
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vstandup/Detail.html:178
	return qs422016
//line views/vstandup/Detail.html:178
}
