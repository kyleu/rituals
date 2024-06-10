// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vstory/Detail.html:2
package vstory

//line views/vestimate/vstory/Detail.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/vestimate/vstory/vvote"
)

//line views/vestimate/vstory/Detail.html:18
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vstory/Detail.html:18
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vstory/Detail.html:18
type Detail struct {
	layout.Basic
	Model                *story.Story
	EstimateByEstimateID *estimate.Estimate
	UserByUserID         *user.User
	Params               filter.ParamSet
	RelVotesByStoryID    vote.Votes
}

//line views/vestimate/vstory/Detail.html:27
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/Detail.html:27
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-story"><button type="button">`)
//line views/vestimate/vstory/Detail.html:30
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vestimate/vstory/Detail.html:30
	qw422016.N().S(`JSON</button></a>
      <a href="`)
//line views/vestimate/vstory/Detail.html:31
	qw422016.E().S(p.Model.WebPath())
//line views/vestimate/vstory/Detail.html:31
	qw422016.N().S(`/edit"><button>`)
//line views/vestimate/vstory/Detail.html:31
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vestimate/vstory/Detail.html:31
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vestimate/vstory/Detail.html:33
	components.StreamSVGIcon(qw422016, `story`, ps)
//line views/vestimate/vstory/Detail.html:33
	qw422016.E().S(p.Model.TitleString())
//line views/vestimate/vstory/Detail.html:33
	qw422016.N().S(`</h3>
    <div><a href="/admin/db/estimate/story"><em>Story</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
            <td>`)
//line views/vestimate/vstory/Detail.html:40
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vestimate/vstory/Detail.html:40
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Estimate ID</th>
            <td class="nowrap">
              `)
//line views/vestimate/vstory/Detail.html:45
	view.StreamUUID(qw422016, &p.Model.EstimateID)
//line views/vestimate/vstory/Detail.html:45
	if p.EstimateByEstimateID != nil {
//line views/vestimate/vstory/Detail.html:45
		qw422016.N().S(` (`)
//line views/vestimate/vstory/Detail.html:45
		qw422016.E().S(p.EstimateByEstimateID.TitleString())
//line views/vestimate/vstory/Detail.html:45
		qw422016.N().S(`)`)
//line views/vestimate/vstory/Detail.html:45
	}
//line views/vestimate/vstory/Detail.html:45
	qw422016.N().S(`
              <a title="Estimate" href="`)
//line views/vestimate/vstory/Detail.html:46
	qw422016.E().S(`/admin/db/estimate` + `/` + p.Model.EstimateID.String())
//line views/vestimate/vstory/Detail.html:46
	qw422016.N().S(`">`)
//line views/vestimate/vstory/Detail.html:46
	components.StreamSVGSimple(qw422016, "estimate", 18, ps)
//line views/vestimate/vstory/Detail.html:46
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="Integer">Idx</th>
            <td>`)
//line views/vestimate/vstory/Detail.html:51
	qw422016.N().D(p.Model.Idx)
//line views/vestimate/vstory/Detail.html:51
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">User ID</th>
            <td class="nowrap">
              `)
//line views/vestimate/vstory/Detail.html:56
	view.StreamUUID(qw422016, &p.Model.UserID)
//line views/vestimate/vstory/Detail.html:56
	if p.UserByUserID != nil {
//line views/vestimate/vstory/Detail.html:56
		qw422016.N().S(` (`)
//line views/vestimate/vstory/Detail.html:56
		qw422016.E().S(p.UserByUserID.TitleString())
//line views/vestimate/vstory/Detail.html:56
		qw422016.N().S(`)`)
//line views/vestimate/vstory/Detail.html:56
	}
//line views/vestimate/vstory/Detail.html:56
	qw422016.N().S(`
              <a title="User" href="`)
//line views/vestimate/vstory/Detail.html:57
	qw422016.E().S(`/admin/db/user` + `/` + p.Model.UserID.String())
//line views/vestimate/vstory/Detail.html:57
	qw422016.N().S(`">`)
//line views/vestimate/vstory/Detail.html:57
	components.StreamSVGSimple(qw422016, "profile", 18, ps)
//line views/vestimate/vstory/Detail.html:57
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Title</th>
            <td><strong>`)
//line views/vestimate/vstory/Detail.html:62
	view.StreamString(qw422016, p.Model.Title)
//line views/vestimate/vstory/Detail.html:62
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="`)
//line views/vestimate/vstory/Detail.html:65
	qw422016.E().S(enum.AllSessionStatuses.Help())
//line views/vestimate/vstory/Detail.html:65
	qw422016.N().S(`">Status</th>
            <td>`)
//line views/vestimate/vstory/Detail.html:66
	qw422016.E().S(p.Model.Status.String())
//line views/vestimate/vstory/Detail.html:66
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Final Vote</th>
            <td>`)
//line views/vestimate/vstory/Detail.html:70
	view.StreamString(qw422016, p.Model.FinalVote)
//line views/vestimate/vstory/Detail.html:70
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vestimate/vstory/Detail.html:74
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vestimate/vstory/Detail.html:74
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
            <td>`)
//line views/vestimate/vstory/Detail.html:78
	view.StreamTimestamp(qw422016, p.Model.Updated)
//line views/vestimate/vstory/Detail.html:78
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vestimate/vstory/Detail.html:86
	relationHelper := story.Stories{p.Model}

//line views/vestimate/vstory/Detail.html:86
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-VotesByStoryID" type="checkbox" hidden="hidden"`)
//line views/vestimate/vstory/Detail.html:91
	if p.Params.Specifies(`vote`) {
//line views/vestimate/vstory/Detail.html:91
		qw422016.N().S(` checked="checked"`)
//line views/vestimate/vstory/Detail.html:91
	}
//line views/vestimate/vstory/Detail.html:91
	qw422016.N().S(` />
        <label for="accordion-VotesByStoryID">
          `)
//line views/vestimate/vstory/Detail.html:93
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vestimate/vstory/Detail.html:93
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Detail.html:94
	components.StreamSVGInline(qw422016, `vote-yea`, 16, ps)
//line views/vestimate/vstory/Detail.html:94
	qw422016.N().S(`
          `)
//line views/vestimate/vstory/Detail.html:95
	qw422016.E().S(util.StringPlural(len(p.RelVotesByStoryID), "Vote"))
//line views/vestimate/vstory/Detail.html:95
	qw422016.N().S(` by [Story ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vestimate/vstory/Detail.html:98
	if len(p.RelVotesByStoryID) == 0 {
//line views/vestimate/vstory/Detail.html:98
		qw422016.N().S(`          <em>no related Votes</em>
`)
//line views/vestimate/vstory/Detail.html:100
	} else {
//line views/vestimate/vstory/Detail.html:100
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vestimate/vstory/Detail.html:102
		vvote.StreamTable(qw422016, p.RelVotesByStoryID, relationHelper, nil, p.Params, as, ps)
//line views/vestimate/vstory/Detail.html:102
		qw422016.N().S(`
          </div>
`)
//line views/vestimate/vstory/Detail.html:104
	}
//line views/vestimate/vstory/Detail.html:104
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
  `)
//line views/vestimate/vstory/Detail.html:109
	components.StreamJSONModal(qw422016, "story", "Story JSON", p.Model, 1)
//line views/vestimate/vstory/Detail.html:109
	qw422016.N().S(`
`)
//line views/vestimate/vstory/Detail.html:110
}

//line views/vestimate/vstory/Detail.html:110
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/Detail.html:110
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/Detail.html:110
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vstory/Detail.html:110
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/Detail.html:110
}

//line views/vestimate/vstory/Detail.html:110
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vstory/Detail.html:110
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/Detail.html:110
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vstory/Detail.html:110
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/Detail.html:110
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/Detail.html:110
	return qs422016
//line views/vestimate/vstory/Detail.html:110
}
