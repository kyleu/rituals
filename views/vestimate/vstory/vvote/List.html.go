// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vestimate/vstory/vvote/List.html:2
package vvote

//line views/vestimate/vstory/vvote/List.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vestimate/vstory/vvote/List.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vestimate/vstory/vvote/List.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vestimate/vstory/vvote/List.html:13
type List struct {
	layout.Basic
	Models           vote.Votes
	StoriesByStoryID story.Stories
	UsersByUserID    user.Users
	Params           filter.ParamSet
}

//line views/vestimate/vstory/vvote/List.html:21
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/vvote/List.html:21
	qw422016.N().S(`
  <div class="card">
    <div class="right"><a href="/admin/db/estimate/story/vote/_new"><button>New</button></a></div>
    <h3>`)
//line views/vestimate/vstory/vvote/List.html:24
	components.StreamSVGRefIcon(qw422016, `vote-yea`, ps)
//line views/vestimate/vstory/vvote/List.html:24
	qw422016.E().S(ps.Title)
//line views/vestimate/vstory/vvote/List.html:24
	qw422016.N().S(`</h3>
`)
//line views/vestimate/vstory/vvote/List.html:25
	if len(p.Models) == 0 {
//line views/vestimate/vstory/vvote/List.html:25
		qw422016.N().S(`    <div class="mt"><em>No votes available</em></div>
`)
//line views/vestimate/vstory/vvote/List.html:27
	} else {
//line views/vestimate/vstory/vvote/List.html:27
		qw422016.N().S(`    <div class="overflow clear mt">
      `)
//line views/vestimate/vstory/vvote/List.html:29
		StreamTable(qw422016, p.Models, p.StoriesByStoryID, p.UsersByUserID, p.Params, as, ps)
//line views/vestimate/vstory/vvote/List.html:29
		qw422016.N().S(`
    </div>
`)
//line views/vestimate/vstory/vvote/List.html:31
	}
//line views/vestimate/vstory/vvote/List.html:31
	qw422016.N().S(`  </div>
`)
//line views/vestimate/vstory/vvote/List.html:33
}

//line views/vestimate/vstory/vvote/List.html:33
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vestimate/vstory/vvote/List.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vestimate/vstory/vvote/List.html:33
	p.StreamBody(qw422016, as, ps)
//line views/vestimate/vstory/vvote/List.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/vestimate/vstory/vvote/List.html:33
}

//line views/vestimate/vstory/vvote/List.html:33
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vestimate/vstory/vvote/List.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vestimate/vstory/vvote/List.html:33
	p.WriteBody(qb422016, as, ps)
//line views/vestimate/vstory/vvote/List.html:33
	qs422016 := string(qb422016.B)
//line views/vestimate/vstory/vvote/List.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vestimate/vstory/vvote/List.html:33
	return qs422016
//line views/vestimate/vstory/vvote/List.html:33
}
