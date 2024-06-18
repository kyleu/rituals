// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaction/Detail.html:1
package vaction

//line views/vaction/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vaction/Detail.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/Detail.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/Detail.html:12
type Detail struct {
	layout.Basic
	Model        *action.Action
	UserByUserID *user.User
}

//line views/vaction/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-action"><button type="button">`)
//line views/vaction/Detail.html:21
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vaction/Detail.html:21
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vaction/Detail.html:22
	qw422016.E().S(p.Model.WebPath())
//line views/vaction/Detail.html:22
	qw422016.N().S(`/edit"><button>`)
//line views/vaction/Detail.html:22
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vaction/Detail.html:22
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vaction/Detail.html:24
	components.StreamSVGIcon(qw422016, `action`, ps)
//line views/vaction/Detail.html:24
	qw422016.N().S(` `)
//line views/vaction/Detail.html:24
	qw422016.E().S(p.Model.TitleString())
//line views/vaction/Detail.html:24
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/action"><em>Action</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
            <td>`)
//line views/vaction/Detail.html:31
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vaction/Detail.html:31
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="`)
//line views/vaction/Detail.html:34
	qw422016.E().S(enum.AllModelServices.Help())
//line views/vaction/Detail.html:34
	qw422016.N().S(`">Svc</th>
            <td>`)
//line views/vaction/Detail.html:35
	qw422016.E().S(p.Model.Svc.String())
//line views/vaction/Detail.html:35
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Model ID</th>
            <td>`)
//line views/vaction/Detail.html:39
	view.StreamUUID(qw422016, &p.Model.ModelID)
//line views/vaction/Detail.html:39
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
            <td class="nowrap">
              `)
//line views/vaction/Detail.html:44
	view.StreamUUID(qw422016, &p.Model.UserID)
//line views/vaction/Detail.html:44
	if p.UserByUserID != nil {
//line views/vaction/Detail.html:44
		qw422016.N().S(` (`)
//line views/vaction/Detail.html:44
		qw422016.E().S(p.UserByUserID.TitleString())
//line views/vaction/Detail.html:44
		qw422016.N().S(`)`)
//line views/vaction/Detail.html:44
	}
//line views/vaction/Detail.html:44
	qw422016.N().S(`
              <a title="User" href="`)
//line views/vaction/Detail.html:45
	qw422016.E().S(`/admin/db/user` + `/` + p.Model.UserID.String())
//line views/vaction/Detail.html:45
	qw422016.N().S(`">`)
//line views/vaction/Detail.html:45
	components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vaction/Detail.html:45
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Act</th>
            <td>`)
//line views/vaction/Detail.html:50
	view.StreamString(qw422016, p.Model.Act)
//line views/vaction/Detail.html:50
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="JSON object">Content</th>
            <td>`)
//line views/vaction/Detail.html:54
	components.StreamJSON(qw422016, p.Model.Content)
//line views/vaction/Detail.html:54
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Note</th>
            <td>`)
//line views/vaction/Detail.html:58
	view.StreamString(qw422016, p.Model.Note)
//line views/vaction/Detail.html:58
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vaction/Detail.html:62
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vaction/Detail.html:62
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vaction/Detail.html:69
	qw422016.N().S(`  `)
//line views/vaction/Detail.html:70
	components.StreamJSONModal(qw422016, "action", "Action JSON", p.Model, 1)
//line views/vaction/Detail.html:70
	qw422016.N().S(`
`)
//line views/vaction/Detail.html:71
}

//line views/vaction/Detail.html:71
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Detail.html:71
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Detail.html:71
	p.StreamBody(qw422016, as, ps)
//line views/vaction/Detail.html:71
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Detail.html:71
}

//line views/vaction/Detail.html:71
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaction/Detail.html:71
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Detail.html:71
	p.WriteBody(qb422016, as, ps)
//line views/vaction/Detail.html:71
	qs422016 := string(qb422016.B)
//line views/vaction/Detail.html:71
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Detail.html:71
	return qs422016
//line views/vaction/Detail.html:71
}
