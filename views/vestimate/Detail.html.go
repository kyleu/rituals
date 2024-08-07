// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vestimate/Detail.html:1
package vestimate

//line views/vestimate/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vestimate/vehistory"
	"github.com/kyleu/rituals/views/vestimate/vemember"
	"github.com/kyleu/rituals/views/vestimate/vepermission"
	"github.com/kyleu/rituals/views/vestimate/vstory"
)

//line views/vestimate/Detail.html:23
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/Detail.html:23
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/Detail.html:23
type Detail struct {
	layout.Basic
	Model                              *estimate.Estimate
	TeamByTeamID                       *team.Team
	SprintBySprintID                   *sprint.Sprint
	Params                             filter.ParamSet
	RelEstimateHistoriesByEstimateID   ehistory.EstimateHistories
	RelEstimateMembersByEstimateID     emember.EstimateMembers
	RelEstimatePermissionsByEstimateID epermission.EstimatePermissions
	RelStoriesByEstimateID             story.Stories
	Paths                              []string
}

//line views/vestimate/Detail.html:36
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/Detail.html:36
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-estimate"><button type="button">`)
//line views/vestimate/Detail.html:39
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vestimate/Detail.html:39
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vestimate/Detail.html:40
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vestimate/Detail.html:40
	qw422016.N().S(`/edit"><button>`)
//line views/vestimate/Detail.html:40
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vestimate/Detail.html:40
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vestimate/Detail.html:42
	components.StreamSVGIcon(qw422016, `estimate`, ps)
//line views/vestimate/Detail.html:42
	qw422016.N().S(` `)
//line views/vestimate/Detail.html:42
	qw422016.E().S(p.Model.TitleString())
//line views/vestimate/Detail.html:42
	qw422016.N().S(`</h3>
    <div><a href="`)
//line views/vestimate/Detail.html:43
	qw422016.E().S(estimate.Route(p.Paths...))
//line views/vestimate/Detail.html:43
	qw422016.N().S(`"><em>Estimate</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
            <td>`)
//line views/vestimate/Detail.html:49
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vestimate/Detail.html:49
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Slug</th>
            <td>`)
//line views/vestimate/Detail.html:53
	view.StreamString(qw422016, p.Model.Slug)
//line views/vestimate/Detail.html:53
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Title</th>
            <td><strong>`)
//line views/vestimate/Detail.html:57
	view.StreamString(qw422016, p.Model.Title)
//line views/vestimate/Detail.html:57
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Icon</th>
            <td>`)
//line views/vestimate/Detail.html:61
	view.StreamString(qw422016, p.Model.Icon)
//line views/vestimate/Detail.html:61
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="`)
//line views/vestimate/Detail.html:64
	qw422016.E().S(enum.AllSessionStatuses.Help())
//line views/vestimate/Detail.html:64
	qw422016.N().S(`">Status</th>
            <td>`)
//line views/vestimate/Detail.html:65
	qw422016.E().S(p.Model.Status.String())
//line views/vestimate/Detail.html:65
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
            <td class="nowrap">
              `)
//line views/vestimate/Detail.html:70
	view.StreamUUID(qw422016, p.Model.TeamID)
//line views/vestimate/Detail.html:70
	if p.TeamByTeamID != nil {
//line views/vestimate/Detail.html:70
		qw422016.N().S(` (`)
//line views/vestimate/Detail.html:70
		qw422016.E().S(p.TeamByTeamID.TitleString())
//line views/vestimate/Detail.html:70
		qw422016.N().S(`)`)
//line views/vestimate/Detail.html:70
	}
//line views/vestimate/Detail.html:70
	qw422016.N().S(`
              `)
//line views/vestimate/Detail.html:71
	if p.Model.TeamID != nil {
//line views/vestimate/Detail.html:71
		qw422016.N().S(`<a title="Team" href="`)
//line views/vestimate/Detail.html:71
		qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vestimate/Detail.html:71
		qw422016.N().S(`">`)
//line views/vestimate/Detail.html:71
		components.StreamSVGLink(qw422016, `team`, ps)
//line views/vestimate/Detail.html:71
		qw422016.N().S(`</a>`)
//line views/vestimate/Detail.html:71
	}
//line views/vestimate/Detail.html:71
	qw422016.N().S(`
            </td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Sprint ID</th>
            <td class="nowrap">
              `)
//line views/vestimate/Detail.html:77
	view.StreamUUID(qw422016, p.Model.SprintID)
//line views/vestimate/Detail.html:77
	if p.SprintBySprintID != nil {
//line views/vestimate/Detail.html:77
		qw422016.N().S(` (`)
//line views/vestimate/Detail.html:77
		qw422016.E().S(p.SprintBySprintID.TitleString())
//line views/vestimate/Detail.html:77
		qw422016.N().S(`)`)
//line views/vestimate/Detail.html:77
	}
//line views/vestimate/Detail.html:77
	qw422016.N().S(`
              `)
//line views/vestimate/Detail.html:78
	if p.Model.SprintID != nil {
//line views/vestimate/Detail.html:78
		qw422016.N().S(`<a title="Sprint" href="`)
//line views/vestimate/Detail.html:78
		qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vestimate/Detail.html:78
		qw422016.N().S(`">`)
//line views/vestimate/Detail.html:78
		components.StreamSVGLink(qw422016, `sprint`, ps)
//line views/vestimate/Detail.html:78
		qw422016.N().S(`</a>`)
//line views/vestimate/Detail.html:78
	}
//line views/vestimate/Detail.html:78
	qw422016.N().S(`
            </td>
          </tr>
          <tr>
            <th class="shrink" title="Comma-separated list of values">Choices</th>
            <td>`)
//line views/vestimate/Detail.html:83
	view.StreamStringArray(qw422016, p.Model.Choices)
//line views/vestimate/Detail.html:83
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vestimate/Detail.html:87
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vestimate/Detail.html:87
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
            <td>`)
//line views/vestimate/Detail.html:91
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vestimate/Detail.html:91
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vestimate/Detail.html:99
	relationHelper := estimate.Estimates{p.Model}

//line views/vestimate/Detail.html:99
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-EstimateHistoriesByEstimateID" type="checkbox" hidden="hidden"`)
//line views/vestimate/Detail.html:104
	if p.Params.Specifies(`ehistory`) {
//line views/vestimate/Detail.html:104
		qw422016.N().S(` checked="checked"`)
//line views/vestimate/Detail.html:104
	}
//line views/vestimate/Detail.html:104
	qw422016.N().S(` />
        <label for="accordion-EstimateHistoriesByEstimateID">
          `)
//line views/vestimate/Detail.html:106
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vestimate/Detail.html:106
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:107
	components.StreamSVGInline(qw422016, `history`, 16, ps)
//line views/vestimate/Detail.html:107
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:108
	qw422016.E().S(util.StringPlural(len(p.RelEstimateHistoriesByEstimateID), "History"))
//line views/vestimate/Detail.html:108
	qw422016.N().S(` by [Estimate ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vestimate/Detail.html:111
	if len(p.RelEstimateHistoriesByEstimateID) == 0 {
//line views/vestimate/Detail.html:111
		qw422016.N().S(`          <em>no related Histories</em>
`)
//line views/vestimate/Detail.html:113
	} else {
//line views/vestimate/Detail.html:113
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vestimate/Detail.html:115
		vehistory.StreamTable(qw422016, p.RelEstimateHistoriesByEstimateID, relationHelper, p.Params, as, ps)
//line views/vestimate/Detail.html:115
		qw422016.N().S(`
          </div>
`)
//line views/vestimate/Detail.html:117
	}
//line views/vestimate/Detail.html:117
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-EstimateMembersByEstimateID" type="checkbox" hidden="hidden"`)
//line views/vestimate/Detail.html:121
	if p.Params.Specifies(`emember`) {
//line views/vestimate/Detail.html:121
		qw422016.N().S(` checked="checked"`)
//line views/vestimate/Detail.html:121
	}
//line views/vestimate/Detail.html:121
	qw422016.N().S(` />
        <label for="accordion-EstimateMembersByEstimateID">
          `)
//line views/vestimate/Detail.html:123
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vestimate/Detail.html:123
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:124
	components.StreamSVGInline(qw422016, `users`, 16, ps)
//line views/vestimate/Detail.html:124
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:125
	qw422016.E().S(util.StringPlural(len(p.RelEstimateMembersByEstimateID), "Member"))
//line views/vestimate/Detail.html:125
	qw422016.N().S(` by [Estimate ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vestimate/Detail.html:128
	if len(p.RelEstimateMembersByEstimateID) == 0 {
//line views/vestimate/Detail.html:128
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vestimate/Detail.html:130
	} else {
//line views/vestimate/Detail.html:130
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vestimate/Detail.html:132
		vemember.StreamTable(qw422016, p.RelEstimateMembersByEstimateID, relationHelper, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:132
		qw422016.N().S(`
          </div>
`)
//line views/vestimate/Detail.html:134
	}
//line views/vestimate/Detail.html:134
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-EstimatePermissionsByEstimateID" type="checkbox" hidden="hidden"`)
//line views/vestimate/Detail.html:138
	if p.Params.Specifies(`epermission`) {
//line views/vestimate/Detail.html:138
		qw422016.N().S(` checked="checked"`)
//line views/vestimate/Detail.html:138
	}
//line views/vestimate/Detail.html:138
	qw422016.N().S(` />
        <label for="accordion-EstimatePermissionsByEstimateID">
          `)
//line views/vestimate/Detail.html:140
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vestimate/Detail.html:140
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:141
	components.StreamSVGInline(qw422016, `permission`, 16, ps)
//line views/vestimate/Detail.html:141
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:142
	qw422016.E().S(util.StringPlural(len(p.RelEstimatePermissionsByEstimateID), "Permission"))
//line views/vestimate/Detail.html:142
	qw422016.N().S(` by [Estimate ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vestimate/Detail.html:145
	if len(p.RelEstimatePermissionsByEstimateID) == 0 {
//line views/vestimate/Detail.html:145
		qw422016.N().S(`          <em>no related Permissions</em>
`)
//line views/vestimate/Detail.html:147
	} else {
//line views/vestimate/Detail.html:147
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vestimate/Detail.html:149
		vepermission.StreamTable(qw422016, p.RelEstimatePermissionsByEstimateID, relationHelper, p.Params, as, ps)
//line views/vestimate/Detail.html:149
		qw422016.N().S(`
          </div>
`)
//line views/vestimate/Detail.html:151
	}
//line views/vestimate/Detail.html:151
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StoriesByEstimateID" type="checkbox" hidden="hidden"`)
//line views/vestimate/Detail.html:155
	if p.Params.Specifies(`story`) {
//line views/vestimate/Detail.html:155
		qw422016.N().S(` checked="checked"`)
//line views/vestimate/Detail.html:155
	}
//line views/vestimate/Detail.html:155
	qw422016.N().S(` />
        <label for="accordion-StoriesByEstimateID">
          `)
//line views/vestimate/Detail.html:157
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vestimate/Detail.html:157
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:158
	components.StreamSVGInline(qw422016, `story`, 16, ps)
//line views/vestimate/Detail.html:158
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:159
	qw422016.E().S(util.StringPlural(len(p.RelStoriesByEstimateID), "Story"))
//line views/vestimate/Detail.html:159
	qw422016.N().S(` by [Estimate ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vestimate/Detail.html:162
	if len(p.RelStoriesByEstimateID) == 0 {
//line views/vestimate/Detail.html:162
		qw422016.N().S(`          <em>no related Stories</em>
`)
//line views/vestimate/Detail.html:164
	} else {
//line views/vestimate/Detail.html:164
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vestimate/Detail.html:166
		vstory.StreamTable(qw422016, p.RelStoriesByEstimateID, relationHelper, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:166
		qw422016.N().S(`
          </div>
`)
//line views/vestimate/Detail.html:168
	}
//line views/vestimate/Detail.html:168
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
  `)
//line views/vestimate/Detail.html:173
	components.StreamJSONModal(qw422016, "estimate", "Estimate JSON", p.Model, 1)
//line views/vestimate/Detail.html:173
	qw422016.N().S(`
`)
//line views/vestimate/Detail.html:174
}

//line views/vestimate/Detail.html:174
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/Detail.html:174
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/Detail.html:174
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/Detail.html:174
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/Detail.html:174
}

//line views/vestimate/Detail.html:174
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/Detail.html:174
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/Detail.html:174
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/Detail.html:174
	qs422016 := string(qb422016.B)
//line views/vestimate/Detail.html:174
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/Detail.html:174
	return qs422016
//line views/vestimate/Detail.html:174
}
