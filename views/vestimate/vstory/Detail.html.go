// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vstory/Detail.html:2
package vstory

//line views/vestimate/vstory/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vestimate/vstory/vvote"
)

//line views/vestimate/vstory/Detail.html:15
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vstory/Detail.html:15
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vstory/Detail.html:15
type Detail struct {
	layout.Basic
	Model          *story.Story
	Estimates      estimate.Estimates
	Users          user.Users
	Params         filter.ParamSet
	VotesByStoryID vote.Votes
}

//line views/vestimate/vstory/Detail.html:24
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/Detail.html:24
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-story"><button type="button">JSON</button></a>
      <a href="`)
//line views/vestimate/vstory/Detail.html:28
	qw422016.E().S(p.Model.WebPath())
//line views/vestimate/vstory/Detail.html:28
	qw422016.N().S(`/edit"><button>`)
//line views/vestimate/vstory/Detail.html:28
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vestimate/vstory/Detail.html:28
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vestimate/vstory/Detail.html:30
	components.StreamSVGRefIcon(qw422016, `book-reader`, ps)
//line views/vestimate/vstory/Detail.html:30
	qw422016.N().S(` `)
//line views/vestimate/vstory/Detail.html:30
	qw422016.E().S(p.Model.TitleString())
//line views/vestimate/vstory/Detail.html:30
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/estimate/story"><em>Story</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vestimate/vstory/Detail.html:36
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vestimate/vstory/Detail.html:36
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Estimate ID</th>
          <td>
            <div class="icon">`)
//line views/vestimate/vstory/Detail.html:41
	components.StreamDisplayUUID(qw422016, &p.Model.EstimateID)
//line views/vestimate/vstory/Detail.html:41
	if x := p.Estimates.Get(p.Model.EstimateID); x != nil {
//line views/vestimate/vstory/Detail.html:41
		qw422016.N().S(` (`)
//line views/vestimate/vstory/Detail.html:41
		qw422016.E().S(x.TitleString())
//line views/vestimate/vstory/Detail.html:41
		qw422016.N().S(`)`)
//line views/vestimate/vstory/Detail.html:41
	}
//line views/vestimate/vstory/Detail.html:41
	qw422016.N().S(`</div>
            <a title="Estimate" href="`)
//line views/vestimate/vstory/Detail.html:42
	qw422016.E().S(`/estimate` + `/` + p.Model.EstimateID.String())
//line views/vestimate/vstory/Detail.html:42
	qw422016.N().S(`">`)
//line views/vestimate/vstory/Detail.html:42
	components.StreamSVGRefIcon(qw422016, "estimate", ps)
//line views/vestimate/vstory/Detail.html:42
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="Integer">Idx</th>
          <td>`)
//line views/vestimate/vstory/Detail.html:47
	qw422016.N().D(p.Model.Idx)
//line views/vestimate/vstory/Detail.html:47
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
          <td>
            <div class="icon">`)
//line views/vestimate/vstory/Detail.html:52
	components.StreamDisplayUUID(qw422016, &p.Model.UserID)
//line views/vestimate/vstory/Detail.html:52
	if x := p.Users.Get(p.Model.UserID); x != nil {
//line views/vestimate/vstory/Detail.html:52
		qw422016.N().S(` (`)
//line views/vestimate/vstory/Detail.html:52
		qw422016.E().S(x.TitleString())
//line views/vestimate/vstory/Detail.html:52
		qw422016.N().S(`)`)
//line views/vestimate/vstory/Detail.html:52
	}
//line views/vestimate/vstory/Detail.html:52
	qw422016.N().S(`</div>
            <a title="User" href="`)
//line views/vestimate/vstory/Detail.html:53
	qw422016.E().S(`/user` + `/` + p.Model.UserID.String())
//line views/vestimate/vstory/Detail.html:53
	qw422016.N().S(`">`)
//line views/vestimate/vstory/Detail.html:53
	components.StreamSVGRefIcon(qw422016, "profile", ps)
//line views/vestimate/vstory/Detail.html:53
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Title</th>
          <td><strong>`)
//line views/vestimate/vstory/Detail.html:58
	qw422016.E().S(p.Model.Title)
//line views/vestimate/vstory/Detail.html:58
	qw422016.N().S(`</strong></td>
        </tr>
        <tr>
          <th class="shrink" title="Available options: [new, active, complete, deleted]">Status</th>
          <td>`)
//line views/vestimate/vstory/Detail.html:62
	qw422016.E().V(p.Model.Status)
//line views/vestimate/vstory/Detail.html:62
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Final Vote</th>
          <td>`)
//line views/vestimate/vstory/Detail.html:66
	qw422016.E().S(p.Model.FinalVote)
//line views/vestimate/vstory/Detail.html:66
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vestimate/vstory/Detail.html:70
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vestimate/vstory/Detail.html:70
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vestimate/vstory/Detail.html:74
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vestimate/vstory/Detail.html:74
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vestimate/vstory/Detail.html:81
	if len(p.VotesByStoryID) > 0 {
//line views/vestimate/vstory/Detail.html:81
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vestimate/vstory/Detail.html:83
		components.StreamSVGRefIcon(qw422016, `vote-yea`, ps)
//line views/vestimate/vstory/Detail.html:83
		qw422016.N().S(` Related votes by [story id]</h3>
    <div class="overflow clear">
      `)
//line views/vestimate/vstory/Detail.html:85
		vvote.StreamTable(qw422016, p.VotesByStoryID, nil, nil, p.Params, as, ps)
//line views/vestimate/vstory/Detail.html:85
		qw422016.N().S(`
    </div>
  </div>
`)
//line views/vestimate/vstory/Detail.html:88
	}
//line views/vestimate/vstory/Detail.html:88
	qw422016.N().S(`  `)
//line views/vestimate/vstory/Detail.html:89
	components.StreamJSONModal(qw422016, "story", "Story JSON", p.Model, 1)
//line views/vestimate/vstory/Detail.html:89
	qw422016.N().S(`
`)
//line views/vestimate/vstory/Detail.html:90
}

//line views/vestimate/vstory/Detail.html:90
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/Detail.html:90
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/Detail.html:90
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vstory/Detail.html:90
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/Detail.html:90
}

//line views/vestimate/vstory/Detail.html:90
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vstory/Detail.html:90
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/Detail.html:90
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vstory/Detail.html:90
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/Detail.html:90
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/Detail.html:90
	return qs422016
//line views/vestimate/vstory/Detail.html:90
}