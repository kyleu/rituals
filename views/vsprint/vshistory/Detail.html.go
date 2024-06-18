// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsprint/vshistory/Detail.html:1
package vshistory

//line views/vsprint/vshistory/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/sprint/shistory"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsprint/vshistory/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/vshistory/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/vshistory/Detail.html:11
type Detail struct {
	layout.Basic
	Model            *shistory.SprintHistory
	SprintBySprintID *sprint.Sprint
}

//line views/vsprint/vshistory/Detail.html:17
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vshistory/Detail.html:17
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-sprintHistory"><button type="button">`)
//line views/vsprint/vshistory/Detail.html:20
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vsprint/vshistory/Detail.html:20
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vsprint/vshistory/Detail.html:21
	qw422016.E().S(p.Model.WebPath())
//line views/vsprint/vshistory/Detail.html:21
	qw422016.N().S(`/edit"><button>`)
//line views/vsprint/vshistory/Detail.html:21
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vsprint/vshistory/Detail.html:21
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vsprint/vshistory/Detail.html:23
	components.StreamSVGIcon(qw422016, `history`, ps)
//line views/vsprint/vshistory/Detail.html:23
	qw422016.N().S(` `)
//line views/vsprint/vshistory/Detail.html:23
	qw422016.E().S(p.Model.TitleString())
//line views/vsprint/vshistory/Detail.html:23
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/sprint/history"><em>History</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="String text">Slug</th>
            <td>`)
//line views/vsprint/vshistory/Detail.html:30
	view.StreamString(qw422016, p.Model.Slug)
//line views/vsprint/vshistory/Detail.html:30
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Sprint ID</th>
            <td class="nowrap">
              `)
//line views/vsprint/vshistory/Detail.html:35
	view.StreamUUID(qw422016, &p.Model.SprintID)
//line views/vsprint/vshistory/Detail.html:35
	if p.SprintBySprintID != nil {
//line views/vsprint/vshistory/Detail.html:35
		qw422016.N().S(` (`)
//line views/vsprint/vshistory/Detail.html:35
		qw422016.E().S(p.SprintBySprintID.TitleString())
//line views/vsprint/vshistory/Detail.html:35
		qw422016.N().S(`)`)
//line views/vsprint/vshistory/Detail.html:35
	}
//line views/vsprint/vshistory/Detail.html:35
	qw422016.N().S(`
              <a title="Sprint" href="`)
//line views/vsprint/vshistory/Detail.html:36
	qw422016.E().S(`/admin/db/sprint` + `/` + p.Model.SprintID.String())
//line views/vsprint/vshistory/Detail.html:36
	qw422016.N().S(`">`)
//line views/vsprint/vshistory/Detail.html:36
	components.StreamSVGLink(qw422016, `sprint`, ps)
//line views/vsprint/vshistory/Detail.html:36
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Sprint Name</th>
            <td>`)
//line views/vsprint/vshistory/Detail.html:41
	view.StreamString(qw422016, p.Model.SprintName)
//line views/vsprint/vshistory/Detail.html:41
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vsprint/vshistory/Detail.html:45
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vsprint/vshistory/Detail.html:45
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vsprint/vshistory/Detail.html:52
	qw422016.N().S(`  `)
//line views/vsprint/vshistory/Detail.html:53
	components.StreamJSONModal(qw422016, "sprintHistory", "History JSON", p.Model, 1)
//line views/vsprint/vshistory/Detail.html:53
	qw422016.N().S(`
`)
//line views/vsprint/vshistory/Detail.html:54
}

//line views/vsprint/vshistory/Detail.html:54
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/vshistory/Detail.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/vshistory/Detail.html:54
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/vshistory/Detail.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/vshistory/Detail.html:54
}

//line views/vsprint/vshistory/Detail.html:54
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/vshistory/Detail.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/vshistory/Detail.html:54
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/vshistory/Detail.html:54
	qs422016 := string(qb422016.B)
//line views/vsprint/vshistory/Detail.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/vshistory/Detail.html:54
	return qs422016
//line views/vsprint/vshistory/Detail.html:54
}
