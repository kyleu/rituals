// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vestimate/vepermission/Detail.html:1
package vepermission

//line views/vestimate/vepermission/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vestimate/vepermission/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vepermission/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vepermission/Detail.html:11
type Detail struct {
	layout.Basic
	Model                *epermission.EstimatePermission
	EstimateByEstimateID *estimate.Estimate
	Paths                []string
}

//line views/vestimate/vepermission/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vepermission/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-estimatePermission"><button type="button" title="JSON">`)
//line views/vestimate/vepermission/Detail.html:21
	components.StreamSVGButton(qw422016, "code", ps)
//line views/vestimate/vepermission/Detail.html:21
	qw422016.N().S(`</button></a>
      <a href="`)
//line views/vestimate/vepermission/Detail.html:22
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vestimate/vepermission/Detail.html:22
	qw422016.N().S(`/edit" title="Edit"><button>`)
//line views/vestimate/vepermission/Detail.html:22
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vestimate/vepermission/Detail.html:22
	qw422016.N().S(`</button></a>
    </div>
    <h3>`)
//line views/vestimate/vepermission/Detail.html:24
	components.StreamSVGIcon(qw422016, `permission`, ps)
//line views/vestimate/vepermission/Detail.html:24
	qw422016.N().S(` `)
//line views/vestimate/vepermission/Detail.html:24
	qw422016.E().S(p.Model.TitleString())
//line views/vestimate/vepermission/Detail.html:24
	qw422016.N().S(`</h3>
    <div><a href="`)
//line views/vestimate/vepermission/Detail.html:25
	qw422016.E().S(epermission.Route(p.Paths...))
//line views/vestimate/vepermission/Detail.html:25
	qw422016.N().S(`"><em>Permission</em></a></div>
    `)
//line views/vestimate/vepermission/Detail.html:26
	StreamDetailTable(qw422016, p, ps)
//line views/vestimate/vepermission/Detail.html:26
	qw422016.N().S(`
  </div>
`)
//line views/vestimate/vepermission/Detail.html:29
	qw422016.N().S(`  `)
//line views/vestimate/vepermission/Detail.html:30
	components.StreamJSONModal(qw422016, "estimatePermission", "Permission JSON", p.Model, 1)
//line views/vestimate/vepermission/Detail.html:30
	qw422016.N().S(`
`)
//line views/vestimate/vepermission/Detail.html:31
}

//line views/vestimate/vepermission/Detail.html:31
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vepermission/Detail.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vepermission/Detail.html:31
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vepermission/Detail.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vepermission/Detail.html:31
}

//line views/vestimate/vepermission/Detail.html:31
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vepermission/Detail.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vepermission/Detail.html:31
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vepermission/Detail.html:31
	qs422016 := string(qb422016.B)
//line views/vestimate/vepermission/Detail.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vepermission/Detail.html:31
	return qs422016
//line views/vestimate/vepermission/Detail.html:31
}

//line views/vestimate/vepermission/Detail.html:33
func StreamDetailTable(qw422016 *qt422016.Writer, p *Detail, ps *cutil.PageState) {
//line views/vestimate/vepermission/Detail.html:33
	qw422016.N().S(`
  <div class="mt overflow full-width">
    <table>
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Estimate ID</th>
          <td class="nowrap">
            `)
//line views/vestimate/vepermission/Detail.html:40
	if x := p.EstimateByEstimateID; x != nil {
//line views/vestimate/vepermission/Detail.html:40
		qw422016.N().S(`
            <a href="`)
//line views/vestimate/vepermission/Detail.html:41
		qw422016.E().S(p.Model.WebPath())
//line views/vestimate/vepermission/Detail.html:41
		qw422016.N().S(`">`)
//line views/vestimate/vepermission/Detail.html:41
		qw422016.E().S(x.TitleString())
//line views/vestimate/vepermission/Detail.html:41
		qw422016.N().S(`</a> <a title="Estimate" href="`)
//line views/vestimate/vepermission/Detail.html:41
		qw422016.E().S(x.WebPath(p.Paths...))
//line views/vestimate/vepermission/Detail.html:41
		qw422016.N().S(`">`)
//line views/vestimate/vepermission/Detail.html:41
		components.StreamSVGLink(qw422016, `estimate`, ps)
//line views/vestimate/vepermission/Detail.html:41
		qw422016.N().S(`</a>
            `)
//line views/vestimate/vepermission/Detail.html:42
	} else {
//line views/vestimate/vepermission/Detail.html:42
		qw422016.N().S(`
            <a href="`)
//line views/vestimate/vepermission/Detail.html:43
		qw422016.E().S(p.Model.WebPath())
//line views/vestimate/vepermission/Detail.html:43
		qw422016.N().S(`">`)
//line views/vestimate/vepermission/Detail.html:43
		view.StreamUUID(qw422016, &p.Model.EstimateID)
//line views/vestimate/vepermission/Detail.html:43
		qw422016.N().S(`</a>
            `)
//line views/vestimate/vepermission/Detail.html:44
	}
//line views/vestimate/vepermission/Detail.html:44
	qw422016.N().S(`
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Key</th>
          <td>`)
//line views/vestimate/vepermission/Detail.html:49
	view.StreamString(qw422016, p.Model.Key)
//line views/vestimate/vepermission/Detail.html:49
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Value</th>
          <td>`)
//line views/vestimate/vepermission/Detail.html:53
	view.StreamString(qw422016, p.Model.Value)
//line views/vestimate/vepermission/Detail.html:53
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Access</th>
          <td>`)
//line views/vestimate/vepermission/Detail.html:57
	view.StreamString(qw422016, p.Model.Access)
//line views/vestimate/vepermission/Detail.html:57
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vestimate/vepermission/Detail.html:61
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vestimate/vepermission/Detail.html:61
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vestimate/vepermission/Detail.html:66
}

//line views/vestimate/vepermission/Detail.html:66
func WriteDetailTable(qq422016 qtio422016.Writer, p *Detail, ps *cutil.PageState) {
//line views/vestimate/vepermission/Detail.html:66
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vepermission/Detail.html:66
	StreamDetailTable(qw422016, p, ps)
//line views/vestimate/vepermission/Detail.html:66
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vepermission/Detail.html:66
}

//line views/vestimate/vepermission/Detail.html:66
func DetailTable(p *Detail, ps *cutil.PageState) string {
//line views/vestimate/vepermission/Detail.html:66
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vepermission/Detail.html:66
	WriteDetailTable(qb422016, p, ps)
//line views/vestimate/vepermission/Detail.html:66
	qs422016 := string(qb422016.B)
//line views/vestimate/vepermission/Detail.html:66
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vepermission/Detail.html:66
	return qs422016
//line views/vestimate/vepermission/Detail.html:66
}
