// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vteam/List.html:1
package vteam

//line views/vteam/List.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vteam/List.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vteam/List.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vteam/List.html:11
type List struct {
	layout.Basic
	Models      team.Teams
	Params      filter.ParamSet
	SearchQuery string
}

//line views/vteam/List.html:18
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vteam/List.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vteam/List.html:20
	edit.StreamSearchForm(qw422016, "", "q", "Search Teams", p.SearchQuery, ps)
//line views/vteam/List.html:20
	qw422016.N().S(`</div>
    <div class="right mrs large-buttons">
`)
//line views/vteam/List.html:22
	if len(p.Models) > 0 {
//line views/vteam/List.html:22
		qw422016.N().S(`<a href="/admin/db/team/_random"><button>`)
//line views/vteam/List.html:22
		components.StreamSVGButton(qw422016, "gift", ps)
//line views/vteam/List.html:22
		qw422016.N().S(` Random</button></a>`)
//line views/vteam/List.html:22
	}
//line views/vteam/List.html:22
	qw422016.N().S(`      <a href="/admin/db/team/_new"><button>`)
//line views/vteam/List.html:23
	components.StreamSVGButton(qw422016, "plus", ps)
//line views/vteam/List.html:23
	qw422016.N().S(` New</button></a>
    </div>
    <h3>`)
//line views/vteam/List.html:25
	components.StreamSVGIcon(qw422016, `team`, ps)
//line views/vteam/List.html:25
	qw422016.N().S(` `)
//line views/vteam/List.html:25
	qw422016.E().S(ps.Title)
//line views/vteam/List.html:25
	qw422016.N().S(`</h3>
    <div class="clear"></div>
`)
//line views/vteam/List.html:27
	if p.SearchQuery != "" {
//line views/vteam/List.html:27
		qw422016.N().S(`    <hr />
    <em>Search results for [`)
//line views/vteam/List.html:29
		qw422016.E().S(p.SearchQuery)
//line views/vteam/List.html:29
		qw422016.N().S(`]</em> (<a href="?">clear</a>)
`)
//line views/vteam/List.html:30
	}
//line views/vteam/List.html:31
	if len(p.Models) == 0 {
//line views/vteam/List.html:31
		qw422016.N().S(`    <div class="mt"><em>No teams available</em></div>
`)
//line views/vteam/List.html:33
	} else {
//line views/vteam/List.html:33
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vteam/List.html:35
		StreamTable(qw422016, p.Models, p.Params, as, ps)
//line views/vteam/List.html:35
		qw422016.N().S(`
    </div>
`)
//line views/vteam/List.html:37
	}
//line views/vteam/List.html:37
	qw422016.N().S(`  </div>
`)
//line views/vteam/List.html:39
}

//line views/vteam/List.html:39
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vteam/List.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vteam/List.html:39
	p.StreamBody(qw422016, as, ps)
//line views/vteam/List.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/vteam/List.html:39
}

//line views/vteam/List.html:39
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vteam/List.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vteam/List.html:39
	p.WriteBody(qb422016, as, ps)
//line views/vteam/List.html:39
	qs422016 := string(qb422016.B)
//line views/vteam/List.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vteam/List.html:39
	return qs422016
//line views/vteam/List.html:39
}
