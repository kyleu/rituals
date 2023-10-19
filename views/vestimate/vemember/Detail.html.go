// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vemember/Detail.html:2
package vemember

//line views/vestimate/vemember/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vestimate/vemember/Detail.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vemember/Detail.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vemember/Detail.html:13
type Detail struct {
	layout.Basic
	Model                *emember.EstimateMember
	EstimateByEstimateID *estimate.Estimate
	UserByUserID         *user.User
}

//line views/vestimate/vemember/Detail.html:20
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vemember/Detail.html:20
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-estimateMember"><button type="button">JSON</button></a>
      <a href="`)
//line views/vestimate/vemember/Detail.html:24
	qw422016.E().S(p.Model.WebPath())
//line views/vestimate/vemember/Detail.html:24
	qw422016.N().S(`/edit"><button>`)
//line views/vestimate/vemember/Detail.html:24
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vestimate/vemember/Detail.html:24
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vestimate/vemember/Detail.html:26
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vestimate/vemember/Detail.html:26
	qw422016.N().S(` `)
//line views/vestimate/vemember/Detail.html:26
	qw422016.E().S(p.Model.TitleString())
//line views/vestimate/vemember/Detail.html:26
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/estimate/member"><em>Member</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Estimate ID</th>
          <td class="nowrap">
            `)
//line views/vestimate/vemember/Detail.html:33
	components.StreamDisplayUUID(qw422016, &p.Model.EstimateID)
//line views/vestimate/vemember/Detail.html:33
	if p.EstimateByEstimateID != nil {
//line views/vestimate/vemember/Detail.html:33
		qw422016.N().S(` (`)
//line views/vestimate/vemember/Detail.html:33
		qw422016.E().S(p.EstimateByEstimateID.TitleString())
//line views/vestimate/vemember/Detail.html:33
		qw422016.N().S(`)`)
//line views/vestimate/vemember/Detail.html:33
	}
//line views/vestimate/vemember/Detail.html:33
	qw422016.N().S(`
            <a title="Estimate" href="`)
//line views/vestimate/vemember/Detail.html:34
	qw422016.E().S(`/admin/db/estimate` + `/` + p.Model.EstimateID.String())
//line views/vestimate/vemember/Detail.html:34
	qw422016.N().S(`">`)
//line views/vestimate/vemember/Detail.html:34
	components.StreamSVGRef(qw422016, "estimate", 18, 18, "", ps)
//line views/vestimate/vemember/Detail.html:34
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
          <td class="nowrap">
            `)
//line views/vestimate/vemember/Detail.html:40
	components.StreamDisplayUUID(qw422016, &p.Model.UserID)
//line views/vestimate/vemember/Detail.html:40
	if p.UserByUserID != nil {
//line views/vestimate/vemember/Detail.html:40
		qw422016.N().S(` (`)
//line views/vestimate/vemember/Detail.html:40
		qw422016.E().S(p.UserByUserID.TitleString())
//line views/vestimate/vemember/Detail.html:40
		qw422016.N().S(`)`)
//line views/vestimate/vemember/Detail.html:40
	}
//line views/vestimate/vemember/Detail.html:40
	qw422016.N().S(`
            <a title="User" href="`)
//line views/vestimate/vemember/Detail.html:41
	qw422016.E().S(`/admin/db/user` + `/` + p.Model.UserID.String())
//line views/vestimate/vemember/Detail.html:41
	qw422016.N().S(`">`)
//line views/vestimate/vemember/Detail.html:41
	components.StreamSVGRef(qw422016, "profile", 18, 18, "", ps)
//line views/vestimate/vemember/Detail.html:41
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Name</th>
          <td><strong>`)
//line views/vestimate/vemember/Detail.html:46
	qw422016.E().S(p.Model.Name)
//line views/vestimate/vemember/Detail.html:46
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="URL in string form">Picture</th>
          <td><a href="`)
//line views/vestimate/vemember/Detail.html:50
	qw422016.E().S(p.Model.Picture)
//line views/vestimate/vemember/Detail.html:50
	qw422016.N().S(`" target="_blank" rel="noopener noreferrer">`)
//line views/vestimate/vemember/Detail.html:50
	qw422016.E().S(p.Model.Picture)
//line views/vestimate/vemember/Detail.html:50
	qw422016.N().S(`</a></td>
        </tr>
        <tr>
          <th class="shrink" title="`)
//line views/vestimate/vemember/Detail.html:53
	qw422016.E().S(enum.AllMemberStatuses.Help())
//line views/vestimate/vemember/Detail.html:53
	qw422016.N().S(`">Role</th>
          <td>`)
//line views/vestimate/vemember/Detail.html:54
	qw422016.E().S(p.Model.Role.String())
//line views/vestimate/vemember/Detail.html:54
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vestimate/vemember/Detail.html:58
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vestimate/vemember/Detail.html:58
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vestimate/vemember/Detail.html:62
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vestimate/vemember/Detail.html:62
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vestimate/vemember/Detail.html:68
	qw422016.N().S(`  `)
//line views/vestimate/vemember/Detail.html:69
	components.StreamJSONModal(qw422016, "estimateMember", "Member JSON", p.Model, 1)
//line views/vestimate/vemember/Detail.html:69
	qw422016.N().S(`
`)
//line views/vestimate/vemember/Detail.html:70
}

//line views/vestimate/vemember/Detail.html:70
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vemember/Detail.html:70
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vemember/Detail.html:70
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vemember/Detail.html:70
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vemember/Detail.html:70
}

//line views/vestimate/vemember/Detail.html:70
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vemember/Detail.html:70
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vemember/Detail.html:70
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vemember/Detail.html:70
	qs422016 := string(qb422016.B)
//line views/vestimate/vemember/Detail.html:70
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vemember/Detail.html:70
	return qs422016
//line views/vestimate/vemember/Detail.html:70
}
