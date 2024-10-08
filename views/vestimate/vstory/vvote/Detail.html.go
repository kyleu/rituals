// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vestimate/vstory/vvote/Detail.html:1
package vvote

//line views/vestimate/vstory/vvote/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vestimate/vstory/vvote/Detail.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vstory/vvote/Detail.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vstory/vvote/Detail.html:12
type Detail struct {
	layout.Basic
	Model          *vote.Vote
	StoryByStoryID *story.Story
	UserByUserID   *user.User
	Paths          []string
}

//line views/vestimate/vstory/vvote/Detail.html:20
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/vvote/Detail.html:20
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-vote"><button type="button" title="JSON">`)
//line views/vestimate/vstory/vvote/Detail.html:23
	components.StreamSVGButton(qw422016, "code", ps)
//line views/vestimate/vstory/vvote/Detail.html:23
	qw422016.N().S(`</button></a>
      <a href="`)
//line views/vestimate/vstory/vvote/Detail.html:24
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vestimate/vstory/vvote/Detail.html:24
	qw422016.N().S(`/edit" title="Edit"><button>`)
//line views/vestimate/vstory/vvote/Detail.html:24
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vestimate/vstory/vvote/Detail.html:24
	qw422016.N().S(`</button></a>
    </div>
    <h3>`)
//line views/vestimate/vstory/vvote/Detail.html:26
	components.StreamSVGIcon(qw422016, `vote-yea`, ps)
//line views/vestimate/vstory/vvote/Detail.html:26
	qw422016.N().S(` `)
//line views/vestimate/vstory/vvote/Detail.html:26
	qw422016.E().S(p.Model.TitleString())
//line views/vestimate/vstory/vvote/Detail.html:26
	qw422016.N().S(`</h3>
    <div><a href="`)
//line views/vestimate/vstory/vvote/Detail.html:27
	qw422016.E().S(vote.Route(p.Paths...))
//line views/vestimate/vstory/vvote/Detail.html:27
	qw422016.N().S(`"><em>Vote</em></a></div>
    `)
//line views/vestimate/vstory/vvote/Detail.html:28
	StreamDetailTable(qw422016, p, ps)
//line views/vestimate/vstory/vvote/Detail.html:28
	qw422016.N().S(`
  </div>
`)
//line views/vestimate/vstory/vvote/Detail.html:31
	qw422016.N().S(`  `)
//line views/vestimate/vstory/vvote/Detail.html:32
	components.StreamJSONModal(qw422016, "vote", "Vote JSON", p.Model, 1)
//line views/vestimate/vstory/vvote/Detail.html:32
	qw422016.N().S(`
`)
//line views/vestimate/vstory/vvote/Detail.html:33
}

//line views/vestimate/vstory/vvote/Detail.html:33
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/vvote/Detail.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/vvote/Detail.html:33
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vstory/vvote/Detail.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/vvote/Detail.html:33
}

//line views/vestimate/vstory/vvote/Detail.html:33
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vstory/vvote/Detail.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/vvote/Detail.html:33
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vstory/vvote/Detail.html:33
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/vvote/Detail.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/vvote/Detail.html:33
	return qs422016
//line views/vestimate/vstory/vvote/Detail.html:33
}

//line views/vestimate/vstory/vvote/Detail.html:35
func StreamDetailTable(qw422016 *qt422016.Writer, p *Detail, ps *cutil.PageState) {
//line views/vestimate/vstory/vvote/Detail.html:35
	qw422016.N().S(`
  <div class="mt overflow full-width">
    <table>
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Story ID</th>
          <td class="nowrap">
            `)
//line views/vestimate/vstory/vvote/Detail.html:42
	view.StreamUUID(qw422016, &p.Model.StoryID)
//line views/vestimate/vstory/vvote/Detail.html:42
	if p.StoryByStoryID != nil {
//line views/vestimate/vstory/vvote/Detail.html:42
		qw422016.N().S(` (`)
//line views/vestimate/vstory/vvote/Detail.html:42
		qw422016.E().S(p.StoryByStoryID.TitleString())
//line views/vestimate/vstory/vvote/Detail.html:42
		qw422016.N().S(`)`)
//line views/vestimate/vstory/vvote/Detail.html:42
	}
//line views/vestimate/vstory/vvote/Detail.html:42
	qw422016.N().S(`
            <a title="Story" href="`)
//line views/vestimate/vstory/vvote/Detail.html:43
	if x := p.StoryByStoryID; x != nil {
//line views/vestimate/vstory/vvote/Detail.html:43
		qw422016.E().S(x.WebPath(p.Paths...))
//line views/vestimate/vstory/vvote/Detail.html:43
	}
//line views/vestimate/vstory/vvote/Detail.html:43
	qw422016.N().S(`">`)
//line views/vestimate/vstory/vvote/Detail.html:43
	components.StreamSVGLink(qw422016, `story`, ps)
//line views/vestimate/vstory/vvote/Detail.html:43
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
          <td class="nowrap">
            `)
//line views/vestimate/vstory/vvote/Detail.html:49
	view.StreamUUID(qw422016, &p.Model.UserID)
//line views/vestimate/vstory/vvote/Detail.html:49
	if p.UserByUserID != nil {
//line views/vestimate/vstory/vvote/Detail.html:49
		qw422016.N().S(` (`)
//line views/vestimate/vstory/vvote/Detail.html:49
		qw422016.E().S(p.UserByUserID.TitleString())
//line views/vestimate/vstory/vvote/Detail.html:49
		qw422016.N().S(`)`)
//line views/vestimate/vstory/vvote/Detail.html:49
	}
//line views/vestimate/vstory/vvote/Detail.html:49
	qw422016.N().S(`
            <a title="User" href="`)
//line views/vestimate/vstory/vvote/Detail.html:50
	if x := p.UserByUserID; x != nil {
//line views/vestimate/vstory/vvote/Detail.html:50
		qw422016.E().S(x.WebPath(p.Paths...))
//line views/vestimate/vstory/vvote/Detail.html:50
	}
//line views/vestimate/vstory/vvote/Detail.html:50
	qw422016.N().S(`">`)
//line views/vestimate/vstory/vvote/Detail.html:50
	components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vestimate/vstory/vvote/Detail.html:50
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Choice</th>
          <td>`)
//line views/vestimate/vstory/vvote/Detail.html:55
	view.StreamString(qw422016, p.Model.Choice)
//line views/vestimate/vstory/vvote/Detail.html:55
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vestimate/vstory/vvote/Detail.html:59
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vestimate/vstory/vvote/Detail.html:59
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vestimate/vstory/vvote/Detail.html:63
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vestimate/vstory/vvote/Detail.html:63
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vestimate/vstory/vvote/Detail.html:68
}

//line views/vestimate/vstory/vvote/Detail.html:68
func WriteDetailTable(qq422016 qtio422016.Writer, p *Detail, ps *cutil.PageState) {
//line views/vestimate/vstory/vvote/Detail.html:68
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/vvote/Detail.html:68
	StreamDetailTable(qw422016, p, ps)
//line views/vestimate/vstory/vvote/Detail.html:68
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/vvote/Detail.html:68
}

//line views/vestimate/vstory/vvote/Detail.html:68
func DetailTable(p *Detail, ps *cutil.PageState) string {
//line views/vestimate/vstory/vvote/Detail.html:68
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/vvote/Detail.html:68
	WriteDetailTable(qb422016, p, ps)
//line views/vestimate/vstory/vvote/Detail.html:68
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/vvote/Detail.html:68
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/vvote/Detail.html:68
	return qs422016
//line views/vestimate/vstory/vvote/Detail.html:68
}
