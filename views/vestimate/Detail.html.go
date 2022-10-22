// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/Detail.html:2
package vestimate

//line views/vestimate/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vestimate/vehistory"
	"github.com/kyleu/rituals/views/vestimate/vemember"
	"github.com/kyleu/rituals/views/vestimate/vepermission"
	"github.com/kyleu/rituals/views/vestimate/vstory"
)

//line views/vestimate/Detail.html:22
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/Detail.html:22
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/Detail.html:22
type Detail struct {
	layout.Basic
	Model                           *estimate.Estimate
	Users                           user.Users
	Teams                           team.Teams
	Sprints                         sprint.Sprints
	Params                          filter.ParamSet
	EstimateHistoriesByEstimateID   ehistory.EstimateHistories
	EstimateMembersByEstimateID     emember.EstimateMembers
	EstimatePermissionsByEstimateID epermission.EstimatePermissions
	StoriesByEstimateID             story.Stories
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
	components.StreamSVGRefIcon(qw422016, `star`, ps)
//line views/vestimate/Detail.html:41
	qw422016.N().S(` `)
//line views/vestimate/Detail.html:41
	qw422016.E().S(p.Model.TitleString())
//line views/vestimate/Detail.html:41
	qw422016.N().S(`</h3>
    <div><a href="/estimate"><em>Estimate</em></a></div>
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
          <th class="shrink" title="Available options: [new, active, complete, deleted]">Status</th>
          <td>`)
//line views/vestimate/Detail.html:59
	qw422016.E().V(p.Model.Status)
//line views/vestimate/Detail.html:59
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Team ID</th>
          <td>
            <div class="icon">`)
//line views/vestimate/Detail.html:64
	components.StreamDisplayUUID(qw422016, p.Model.TeamID)
//line views/vestimate/Detail.html:64
	if p.Model.TeamID != nil {
//line views/vestimate/Detail.html:64
		if x := p.Teams.Get(*p.Model.TeamID); x != nil {
//line views/vestimate/Detail.html:64
			qw422016.N().S(` (`)
//line views/vestimate/Detail.html:64
			qw422016.E().S(x.TitleString())
//line views/vestimate/Detail.html:64
			qw422016.N().S(`)`)
//line views/vestimate/Detail.html:64
		}
//line views/vestimate/Detail.html:64
	}
//line views/vestimate/Detail.html:64
	qw422016.N().S(`</div>
            <a title="Team" href="`)
//line views/vestimate/Detail.html:65
	qw422016.E().S(`/team` + `/` + p.Model.TeamID.String())
//line views/vestimate/Detail.html:65
	qw422016.N().S(`">`)
//line views/vestimate/Detail.html:65
	components.StreamSVGRefIcon(qw422016, "users", ps)
//line views/vestimate/Detail.html:65
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000) (optional)">Sprint ID</th>
          <td>
            <div class="icon">`)
//line views/vestimate/Detail.html:71
	components.StreamDisplayUUID(qw422016, p.Model.SprintID)
//line views/vestimate/Detail.html:71
	if p.Model.SprintID != nil {
//line views/vestimate/Detail.html:71
		if x := p.Sprints.Get(*p.Model.SprintID); x != nil {
//line views/vestimate/Detail.html:71
			qw422016.N().S(` (`)
//line views/vestimate/Detail.html:71
			qw422016.E().S(x.TitleString())
//line views/vestimate/Detail.html:71
			qw422016.N().S(`)`)
//line views/vestimate/Detail.html:71
		}
//line views/vestimate/Detail.html:71
	}
//line views/vestimate/Detail.html:71
	qw422016.N().S(`</div>
            <a title="Sprint" href="`)
//line views/vestimate/Detail.html:72
	qw422016.E().S(`/sprint` + `/` + p.Model.SprintID.String())
//line views/vestimate/Detail.html:72
	qw422016.N().S(`">`)
//line views/vestimate/Detail.html:72
	components.StreamSVGRefIcon(qw422016, "star", ps)
//line views/vestimate/Detail.html:72
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Owner</th>
          <td>
            <div class="icon">`)
//line views/vestimate/Detail.html:78
	components.StreamDisplayUUID(qw422016, &p.Model.Owner)
//line views/vestimate/Detail.html:78
	if x := p.Users.Get(p.Model.Owner); x != nil {
//line views/vestimate/Detail.html:78
		qw422016.N().S(` (`)
//line views/vestimate/Detail.html:78
		qw422016.E().S(x.TitleString())
//line views/vestimate/Detail.html:78
		qw422016.N().S(`)`)
//line views/vestimate/Detail.html:78
	}
//line views/vestimate/Detail.html:78
	qw422016.N().S(`</div>
            <a title="User" href="`)
//line views/vestimate/Detail.html:79
	qw422016.E().S(`/user` + `/` + p.Model.Owner.String())
//line views/vestimate/Detail.html:79
	qw422016.N().S(`">`)
//line views/vestimate/Detail.html:79
	components.StreamSVGRefIcon(qw422016, "profile", ps)
//line views/vestimate/Detail.html:79
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="Comma-separated list of values">Choices</th>
          <td>`)
//line views/vestimate/Detail.html:84
	components.StreamDisplayStringArray(qw422016, p.Model.Choices)
//line views/vestimate/Detail.html:84
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vestimate/Detail.html:88
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vestimate/Detail.html:88
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vestimate/Detail.html:92
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vestimate/Detail.html:92
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vestimate/Detail.html:99
	if len(p.EstimateHistoriesByEstimateID) > 0 {
//line views/vestimate/Detail.html:99
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vestimate/Detail.html:101
		components.StreamSVGRefIcon(qw422016, `clock`, ps)
//line views/vestimate/Detail.html:101
		qw422016.N().S(` Related histories by [estimate id]</h3>
    <div class="overflow clear">
      `)
//line views/vestimate/Detail.html:103
		vehistory.StreamTable(qw422016, p.EstimateHistoriesByEstimateID, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:103
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vestimate/Detail.html:106
	}
//line views/vestimate/Detail.html:107
	if len(p.EstimateMembersByEstimateID) > 0 {
//line views/vestimate/Detail.html:107
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vestimate/Detail.html:109
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vestimate/Detail.html:109
		qw422016.N().S(` Related members by [estimate id]</h3>
    <div class="overflow clear">
      `)
//line views/vestimate/Detail.html:111
		vemember.StreamTable(qw422016, p.EstimateMembersByEstimateID, nil, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:111
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vestimate/Detail.html:114
	}
//line views/vestimate/Detail.html:115
	if len(p.EstimatePermissionsByEstimateID) > 0 {
//line views/vestimate/Detail.html:115
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vestimate/Detail.html:117
		components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vestimate/Detail.html:117
		qw422016.N().S(` Related permissions by [estimate id]</h3>
    <div class="overflow clear">
      `)
//line views/vestimate/Detail.html:119
		vepermission.StreamTable(qw422016, p.EstimatePermissionsByEstimateID, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:119
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vestimate/Detail.html:122
	}
//line views/vestimate/Detail.html:123
	if len(p.StoriesByEstimateID) > 0 {
//line views/vestimate/Detail.html:123
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vestimate/Detail.html:125
		components.StreamSVGRefIcon(qw422016, `star`, ps)
//line views/vestimate/Detail.html:125
		qw422016.N().S(` Related stories by [estimate id]</h3>
    <div class="overflow clear">
      `)
//line views/vestimate/Detail.html:127
		vstory.StreamTable(qw422016, p.StoriesByEstimateID, nil, nil, p.Params, as, ps)
//line views/vestimate/Detail.html:127
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vestimate/Detail.html:130
	}
//line views/vestimate/Detail.html:130
	qw422016.N().S(`  `)
//line views/vestimate/Detail.html:131
	components.StreamJSONModal(qw422016, "estimate", "Estimate JSON", p.Model, 1)
//line views/vestimate/Detail.html:131
	qw422016.N().S(`
`)
//line views/vestimate/Detail.html:132
}

//line views/vestimate/Detail.html:132
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/Detail.html:132
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/Detail.html:132
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/Detail.html:132
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/Detail.html:132
}

//line views/vestimate/Detail.html:132
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/Detail.html:132
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/Detail.html:132
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/Detail.html:132
	qs422016 := string(qb422016.B)
//line views/vestimate/Detail.html:132
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/Detail.html:132
	return qs422016
//line views/vestimate/Detail.html:132
}
