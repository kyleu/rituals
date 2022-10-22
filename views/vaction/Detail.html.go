// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vaction/Detail.html:2
package vaction

//line views/vaction/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vaction/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaction/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaction/Detail.html:11
type Detail struct {
	layout.Basic
	Model *action.Action
	Users user.Users
}

//line views/vaction/Detail.html:17
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Detail.html:17
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-action"><button type="button">JSON</button></a>
      <a href="`)
//line views/vaction/Detail.html:21
	qw422016.E().S(p.Model.WebPath())
//line views/vaction/Detail.html:21
	qw422016.N().S(`/edit"><button>`)
//line views/vaction/Detail.html:21
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vaction/Detail.html:21
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vaction/Detail.html:23
	components.StreamSVGRefIcon(qw422016, `gamepad`, ps)
//line views/vaction/Detail.html:23
	qw422016.N().S(` `)
//line views/vaction/Detail.html:23
	qw422016.E().S(p.Model.TitleString())
//line views/vaction/Detail.html:23
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/action"><em>Action</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vaction/Detail.html:29
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vaction/Detail.html:29
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Available options: [team, sprint, estimate, standup, retro, story, feedback, report]">Svc</th>
          <td>`)
//line views/vaction/Detail.html:33
	qw422016.E().V(p.Model.Svc)
//line views/vaction/Detail.html:33
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Model ID</th>
          <td>`)
//line views/vaction/Detail.html:37
	components.StreamDisplayUUID(qw422016, &p.Model.ModelID)
//line views/vaction/Detail.html:37
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
          <td>
            <div class="icon">`)
//line views/vaction/Detail.html:42
	components.StreamDisplayUUID(qw422016, &p.Model.UserID)
//line views/vaction/Detail.html:42
	if x := p.Users.Get(p.Model.UserID); x != nil {
//line views/vaction/Detail.html:42
		qw422016.N().S(` (`)
//line views/vaction/Detail.html:42
		qw422016.E().S(x.TitleString())
//line views/vaction/Detail.html:42
		qw422016.N().S(`)`)
//line views/vaction/Detail.html:42
	}
//line views/vaction/Detail.html:42
	qw422016.N().S(`</div>
            <a title="User" href="`)
//line views/vaction/Detail.html:43
	qw422016.E().S(`/user` + `/` + p.Model.UserID.String())
//line views/vaction/Detail.html:43
	qw422016.N().S(`">`)
//line views/vaction/Detail.html:43
	components.StreamSVGRefIcon(qw422016, "profile", ps)
//line views/vaction/Detail.html:43
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Act</th>
          <td>`)
//line views/vaction/Detail.html:48
	qw422016.E().S(p.Model.Act)
//line views/vaction/Detail.html:48
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Content</th>
          <td>`)
//line views/vaction/Detail.html:52
	qw422016.E().S(p.Model.Content)
//line views/vaction/Detail.html:52
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Note</th>
          <td>`)
//line views/vaction/Detail.html:56
	qw422016.E().S(p.Model.Note)
//line views/vaction/Detail.html:56
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vaction/Detail.html:60
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vaction/Detail.html:60
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vaction/Detail.html:66
	qw422016.N().S(`  `)
//line views/vaction/Detail.html:67
	components.StreamJSONModal(qw422016, "action", "Action JSON", p.Model, 1)
//line views/vaction/Detail.html:67
	qw422016.N().S(`
`)
//line views/vaction/Detail.html:68
}

//line views/vaction/Detail.html:68
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaction/Detail.html:68
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaction/Detail.html:68
	p.StreamBody(qw422016, as, ps)
//line views/vaction/Detail.html:68
	qt422016.ReleaseWriter(qw422016)
//line views/vaction/Detail.html:68
}

//line views/vaction/Detail.html:68
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaction/Detail.html:68
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaction/Detail.html:68
	p.WriteBody(qb422016, as, ps)
//line views/vaction/Detail.html:68
	qs422016 := string(qb422016.B)
//line views/vaction/Detail.html:68
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaction/Detail.html:68
	return qs422016
//line views/vaction/Detail.html:68
}
