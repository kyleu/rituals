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
      <a href="#modal-vote"><button type="button">`)
//line views/vestimate/vstory/vvote/Detail.html:23
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vestimate/vstory/vvote/Detail.html:23
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vestimate/vstory/vvote/Detail.html:24
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vestimate/vstory/vvote/Detail.html:24
	qw422016.N().S(`/edit"><button>`)
//line views/vestimate/vstory/vvote/Detail.html:24
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vestimate/vstory/vvote/Detail.html:24
	qw422016.N().S(` Edit</button></a>
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
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Story ID</th>
            <td class="nowrap">
              `)
//line views/vestimate/vstory/vvote/Detail.html:34
	view.StreamUUID(qw422016, &p.Model.StoryID)
//line views/vestimate/vstory/vvote/Detail.html:34
	if p.StoryByStoryID != nil {
//line views/vestimate/vstory/vvote/Detail.html:34
		qw422016.N().S(` (`)
//line views/vestimate/vstory/vvote/Detail.html:34
		qw422016.E().S(p.StoryByStoryID.TitleString())
//line views/vestimate/vstory/vvote/Detail.html:34
		qw422016.N().S(`)`)
//line views/vestimate/vstory/vvote/Detail.html:34
	}
//line views/vestimate/vstory/vvote/Detail.html:34
	qw422016.N().S(`
              <a title="Story" href="`)
//line views/vestimate/vstory/vvote/Detail.html:35
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vestimate/vstory/vvote/Detail.html:35
	qw422016.N().S(`">`)
//line views/vestimate/vstory/vvote/Detail.html:35
	components.StreamSVGLink(qw422016, `story`, ps)
//line views/vestimate/vstory/vvote/Detail.html:35
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
            <td class="nowrap">
              `)
//line views/vestimate/vstory/vvote/Detail.html:41
	view.StreamUUID(qw422016, &p.Model.UserID)
//line views/vestimate/vstory/vvote/Detail.html:41
	if p.UserByUserID != nil {
//line views/vestimate/vstory/vvote/Detail.html:41
		qw422016.N().S(` (`)
//line views/vestimate/vstory/vvote/Detail.html:41
		qw422016.E().S(p.UserByUserID.TitleString())
//line views/vestimate/vstory/vvote/Detail.html:41
		qw422016.N().S(`)`)
//line views/vestimate/vstory/vvote/Detail.html:41
	}
//line views/vestimate/vstory/vvote/Detail.html:41
	qw422016.N().S(`
              <a title="User" href="`)
//line views/vestimate/vstory/vvote/Detail.html:42
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vestimate/vstory/vvote/Detail.html:42
	qw422016.N().S(`">`)
//line views/vestimate/vstory/vvote/Detail.html:42
	components.StreamSVGLink(qw422016, `profile`, ps)
//line views/vestimate/vstory/vvote/Detail.html:42
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Choice</th>
            <td>`)
//line views/vestimate/vstory/vvote/Detail.html:47
	view.StreamString(qw422016, p.Model.Choice)
//line views/vestimate/vstory/vvote/Detail.html:47
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vestimate/vstory/vvote/Detail.html:51
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vestimate/vstory/vvote/Detail.html:51
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
            <td>`)
//line views/vestimate/vstory/vvote/Detail.html:55
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vestimate/vstory/vvote/Detail.html:55
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vestimate/vstory/vvote/Detail.html:62
	qw422016.N().S(`  `)
//line views/vestimate/vstory/vvote/Detail.html:63
	components.StreamJSONModal(qw422016, "vote", "Vote JSON", p.Model, 1)
//line views/vestimate/vstory/vvote/Detail.html:63
	qw422016.N().S(`
`)
//line views/vestimate/vstory/vvote/Detail.html:64
}

//line views/vestimate/vstory/vvote/Detail.html:64
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/vvote/Detail.html:64
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/vvote/Detail.html:64
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vstory/vvote/Detail.html:64
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/vvote/Detail.html:64
}

//line views/vestimate/vstory/vvote/Detail.html:64
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vstory/vvote/Detail.html:64
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/vvote/Detail.html:64
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vstory/vvote/Detail.html:64
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/vvote/Detail.html:64
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/vvote/Detail.html:64
	return qs422016
//line views/vestimate/vstory/vvote/Detail.html:64
}
