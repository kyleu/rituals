// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/Detail.html:2
package vestimate

//line views/vestimate/Detail.html:2
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
}

//line views/vestimate/Detail.html:35
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/Detail.html:35
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-estimate"><button type="button">JSON</button></a>
      <a href="`)
//line views/vestimate/Detail.html:39
	qw422016.E().S(p.Model.WebPath())
//line views/vestimate/Detail.html:39
	qw422016.N().S(`/edit"><button>`)
//line views/vestimate/Detail.html:39
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vestimate/Detail.html:39
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vestimate/Detail.html:41
	components.StreamSVGRefIcon(qw422016, `estimate`, ps)
//line views/vestimate/Detail.html:41
	qw422016.N().S(` `)
//line views/vestimate/Detail.html:41
	qw422016.E().S(p.Model.TitleString())
//line views/vestimate/Detail.html:41
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/estimate"><em>Estimate</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vestimate/Detail.html:47
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vestimate/Detail.html:47
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Slug</th>
          <td>`)
//line views/vestimate/Detail.html:51
	qw422016.E().S(p.Model.Slug)
//line views/vestimate/Detail.html:51
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Title</th>
          <td><strong>`)
//line views/vestimate/Detail.html:55
	qw422016.E().S(p.Model.Title)
//line views/vestimate/Detail.html:55
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Icon</th>
          <td>`)
//line views/vestimate/Detail.html:59
	qw422016.E().S(p.Model.Icon)
//line views/vestimate/Detail.html:59
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="`)
//line views/vestimate/Detail.html:62
	qw422016.E().S(enum.AllSessionStatuses.Help())
//line views/vestimate/Detail.html:62
	qw422016.N().S(`">Status</th>
          <td>`)
//line views/vestimate/Detail.html:63
	qw422016.E().S(p.Model.Status.String())
//line views/vestimate/Detail.html:63
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
          <td class="nowrap">
            `)
//line views/vestimate/Detail.html:68
	components.StreamDisplayUUID(qw422016, p.Model.TeamID)
//line views/vestimate/Detail.html:68
	if p.TeamByTeamID != nil {
//line views/vestimate/Detail.html:68
		qw422016.N().S(` (`)
//line views/vestimate/Detail.html:68
		qw422016.E().S(p.TeamByTeamID.TitleString())
//line views/vestimate/Detail.html:68
		qw422016.N().S(`)`)
//line views/vestimate/Detail.html:68
	}
//line views/vestimate/Detail.html:68
	qw422016.N().S(`
            `)
//line views/vestimate/Detail.html:69
	if p.Model.TeamID != nil {
//line views/vestimate/Detail.html:69
		qw422016.N().S(`<a title="Team" href="`)
//line views/vestimate/Detail.html:69
		qw422016.E().S(`/admin/db/team` + `/` + p.Model.TeamID.String())
//line views/vestimate/Detail.html:69
		qw422016.N().S(`">`)
//line views/vestimate/Detail.html:69
		components.StreamSVGRef(qw422016, "team", 18, 18, "", ps)
//line views/vestimate/Detail.html:69
		qw422016.N().S(`</a>`)
//line views/vestimate/Detail.html:69
	}
//line views/vestimate/Detail.html:69
	qw422016.N().S(`
          </td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Sprint ID</th>
          <td class="nowrap">
            `)
//line views/vestimate/Detail.html:75
	components.StreamDisplayUUID(qw422016, p.Model.SprintID)
//line views/vestimate/Detail.html:75
	if p.SprintBySprintID != nil {
//line views/vestimate/Detail.html:75
		qw422016.N().S(` (`)
//line views/vestimate/Detail.html:75
		qw422016.E().S(p.SprintBySprintID.TitleString())
//line views/vestimate/Detail.html:75
		qw422016.N().S(`)`)
//line views/vestimate/Detail.html:75
	}
//line views/vestimate/Detail.html:75
	qw422016.N().S(`
            `)
//line views/vestimate/Detail.html:76
	if p.Model.SprintID != nil {
//line views/vestimate/Detail.html:76
		qw422016.N().S(`<a title="Sprint" href="`)
//line views/vestimate/Detail.html:76
		qw422016.E().S(`/admin/db/sprint` + `/` + p.Model.SprintID.String())
//line views/vestimate/Detail.html:76
		qw422016.N().S(`">`)
//line views/vestimate/Detail.html:76
		components.StreamSVGRef(qw422016, "sprint", 18, 18, "", ps)
//line views/vestimate/Detail.html:76
		qw422016.N().S(`</a>`)
//line views/vestimate/Detail.html:76
	}
//line views/vestimate/Detail.html:76
	qw422016.N().S(`
          </td>
        </tr>
        <tr>
          <th class="shrink" title="Comma-separated list of values">Choices</th>
          <td>`)
//line views/vestimate/Detail.html:81
	components.StreamDisplayStringArray(qw422016, p.Model.Choices)
//line views/vestimate/Detail.html:81
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vestimate/Detail.html:85
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vestimate/Detail.html:85
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vestimate/Detail.html:89
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vestimate/Detail.html:89
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vestimate/Detail.html:95
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-EstimateHistoriesByEstimateID" type="checkbox" hidden />
        <label for="accordion-EstimateHistoriesByEstimateID">
          `)
//line views/vestimate/Detail.html:102
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vestimate/Detail.html:102
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:103
	components.StreamSVGRef(qw422016, `history`, 16, 16, `icon`, ps)
//line views/vestimate/Detail.html:103
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:104
	qw422016.E().S(util.StringPlural(len(p.RelEstimateHistoriesByEstimateID), "History"))
//line views/vestimate/Detail.html:104
	qw422016.N().S(` by [Estimate ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vestimate/Detail.html:107
	if len(p.RelEstimateHistoriesByEstimateID) == 0 {
//line views/vestimate/Detail.html:107
		qw422016.N().S(`          <em>no related Histories</em>
`)
//line views/vestimate/Detail.html:109
	} else {
//line views/vestimate/Detail.html:109
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vestimate/Detail.html:111
		vehistory.StreamTable(qw422016, p.RelEstimateHistoriesByEstimateID, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:111
		qw422016.N().S(`
          </div>
`)
//line views/vestimate/Detail.html:113
	}
//line views/vestimate/Detail.html:113
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-EstimateMembersByEstimateID" type="checkbox" hidden />
        <label for="accordion-EstimateMembersByEstimateID">
          `)
//line views/vestimate/Detail.html:119
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vestimate/Detail.html:119
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:120
	components.StreamSVGRef(qw422016, `users`, 16, 16, `icon`, ps)
//line views/vestimate/Detail.html:120
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:121
	qw422016.E().S(util.StringPlural(len(p.RelEstimateMembersByEstimateID), "Member"))
//line views/vestimate/Detail.html:121
	qw422016.N().S(` by [Estimate ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vestimate/Detail.html:124
	if len(p.RelEstimateMembersByEstimateID) == 0 {
//line views/vestimate/Detail.html:124
		qw422016.N().S(`          <em>no related Members</em>
`)
//line views/vestimate/Detail.html:126
	} else {
//line views/vestimate/Detail.html:126
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vestimate/Detail.html:128
		vemember.StreamTable(qw422016, p.RelEstimateMembersByEstimateID, nil, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:128
		qw422016.N().S(`
          </div>
`)
//line views/vestimate/Detail.html:130
	}
//line views/vestimate/Detail.html:130
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-EstimatePermissionsByEstimateID" type="checkbox" hidden />
        <label for="accordion-EstimatePermissionsByEstimateID">
          `)
//line views/vestimate/Detail.html:136
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vestimate/Detail.html:136
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:137
	components.StreamSVGRef(qw422016, `permission`, 16, 16, `icon`, ps)
//line views/vestimate/Detail.html:137
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:138
	qw422016.E().S(util.StringPlural(len(p.RelEstimatePermissionsByEstimateID), "Permission"))
//line views/vestimate/Detail.html:138
	qw422016.N().S(` by [Estimate ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vestimate/Detail.html:141
	if len(p.RelEstimatePermissionsByEstimateID) == 0 {
//line views/vestimate/Detail.html:141
		qw422016.N().S(`          <em>no related Permissions</em>
`)
//line views/vestimate/Detail.html:143
	} else {
//line views/vestimate/Detail.html:143
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vestimate/Detail.html:145
		vepermission.StreamTable(qw422016, p.RelEstimatePermissionsByEstimateID, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:145
		qw422016.N().S(`
          </div>
`)
//line views/vestimate/Detail.html:147
	}
//line views/vestimate/Detail.html:147
	qw422016.N().S(`        </div></div></div>
      </li>
      <li>
        <input id="accordion-StoriesByEstimateID" type="checkbox" hidden />
        <label for="accordion-StoriesByEstimateID">
          `)
//line views/vestimate/Detail.html:153
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vestimate/Detail.html:153
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:154
	components.StreamSVGRef(qw422016, `story`, 16, 16, `icon`, ps)
//line views/vestimate/Detail.html:154
	qw422016.N().S(`
          `)
//line views/vestimate/Detail.html:155
	qw422016.E().S(util.StringPlural(len(p.RelStoriesByEstimateID), "Story"))
//line views/vestimate/Detail.html:155
	qw422016.N().S(` by [Estimate ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vestimate/Detail.html:158
	if len(p.RelStoriesByEstimateID) == 0 {
//line views/vestimate/Detail.html:158
		qw422016.N().S(`          <em>no related Stories</em>
`)
//line views/vestimate/Detail.html:160
	} else {
//line views/vestimate/Detail.html:160
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vestimate/Detail.html:162
		vstory.StreamTable(qw422016, p.RelStoriesByEstimateID, nil, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:162
		qw422016.N().S(`
          </div>
`)
//line views/vestimate/Detail.html:164
	}
//line views/vestimate/Detail.html:164
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
  `)
//line views/vestimate/Detail.html:169
	components.StreamJSONModal(qw422016, "estimate", "Estimate JSON", p.Model, 1)
//line views/vestimate/Detail.html:169
	qw422016.N().S(`
`)
//line views/vestimate/Detail.html:170
}

//line views/vestimate/Detail.html:170
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/Detail.html:170
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/Detail.html:170
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/Detail.html:170
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/Detail.html:170
}

//line views/vestimate/Detail.html:170
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/Detail.html:170
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/Detail.html:170
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/Detail.html:170
	qs422016 := string(qb422016.B)
//line views/vestimate/Detail.html:170
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/Detail.html:170
	return qs422016
//line views/vestimate/Detail.html:170
}
