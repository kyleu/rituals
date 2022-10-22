// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vcomment/Detail.html:2
package vcomment

//line views/vcomment/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/comment"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vcomment/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vcomment/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vcomment/Detail.html:11
type Detail struct {
	layout.Basic
	Model *comment.Comment
	Users user.Users
}

//line views/vcomment/Detail.html:17
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcomment/Detail.html:17
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-comment"><button type="button">JSON</button></a>
      <a href="`)
//line views/vcomment/Detail.html:21
	qw422016.E().S(p.Model.WebPath())
//line views/vcomment/Detail.html:21
	qw422016.N().S(`/edit"><button>`)
//line views/vcomment/Detail.html:21
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vcomment/Detail.html:21
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vcomment/Detail.html:23
	components.StreamSVGRefIcon(qw422016, `comments`, ps)
//line views/vcomment/Detail.html:23
	qw422016.N().S(` `)
//line views/vcomment/Detail.html:23
	qw422016.E().S(p.Model.TitleString())
//line views/vcomment/Detail.html:23
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/comment"><em>Comment</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vcomment/Detail.html:29
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vcomment/Detail.html:29
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Available options: [team, sprint, estimate, standup, retro, story, feedback, report]">Svc</th>
          <td>`)
//line views/vcomment/Detail.html:33
	qw422016.E().V(p.Model.Svc)
//line views/vcomment/Detail.html:33
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Model ID</th>
          <td>`)
//line views/vcomment/Detail.html:37
	components.StreamDisplayUUID(qw422016, &p.Model.ModelID)
//line views/vcomment/Detail.html:37
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Target Type</th>
          <td>`)
//line views/vcomment/Detail.html:41
	qw422016.E().S(p.Model.TargetType)
//line views/vcomment/Detail.html:41
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Target ID</th>
          <td>`)
//line views/vcomment/Detail.html:45
	components.StreamDisplayUUID(qw422016, &p.Model.TargetID)
//line views/vcomment/Detail.html:45
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
          <td>
            <div class="icon">`)
//line views/vcomment/Detail.html:50
	components.StreamDisplayUUID(qw422016, &p.Model.UserID)
//line views/vcomment/Detail.html:50
	if x := p.Users.Get(p.Model.UserID); x != nil {
//line views/vcomment/Detail.html:50
		qw422016.N().S(` (`)
//line views/vcomment/Detail.html:50
		qw422016.E().S(x.TitleString())
//line views/vcomment/Detail.html:50
		qw422016.N().S(`)`)
//line views/vcomment/Detail.html:50
	}
//line views/vcomment/Detail.html:50
	qw422016.N().S(`</div>
            <a title="User" href="`)
//line views/vcomment/Detail.html:51
	qw422016.E().S(`/user` + `/` + p.Model.UserID.String())
//line views/vcomment/Detail.html:51
	qw422016.N().S(`">`)
//line views/vcomment/Detail.html:51
	components.StreamSVGRefIcon(qw422016, "profile", ps)
//line views/vcomment/Detail.html:51
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Content</th>
          <td>`)
//line views/vcomment/Detail.html:56
	qw422016.E().S(p.Model.Content)
//line views/vcomment/Detail.html:56
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">HTML</th>
          <td>`)
//line views/vcomment/Detail.html:60
	qw422016.E().S(p.Model.HTML)
//line views/vcomment/Detail.html:60
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vcomment/Detail.html:64
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vcomment/Detail.html:64
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vcomment/Detail.html:70
	qw422016.N().S(`  `)
//line views/vcomment/Detail.html:71
	components.StreamJSONModal(qw422016, "comment", "Comment JSON", p.Model, 1)
//line views/vcomment/Detail.html:71
	qw422016.N().S(`
`)
//line views/vcomment/Detail.html:72
}

//line views/vcomment/Detail.html:72
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcomment/Detail.html:72
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vcomment/Detail.html:72
	p.StreamBody(qw422016, as, ps)
//line views/vcomment/Detail.html:72
	qt422016.ReleaseWriter(qw422016)
//line views/vcomment/Detail.html:72
}

//line views/vcomment/Detail.html:72
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vcomment/Detail.html:72
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vcomment/Detail.html:72
	p.WriteBody(qb422016, as, ps)
//line views/vcomment/Detail.html:72
	qs422016 := string(qb422016.B)
//line views/vcomment/Detail.html:72
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vcomment/Detail.html:72
	return qs422016
//line views/vcomment/Detail.html:72
}
