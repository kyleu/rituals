// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsprint/List.html:1
package vsprint

//line views/vsprint/List.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsprint/List.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsprint/List.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsprint/List.html:12
type List struct {
	layout.Basic
	Models        sprint.Sprints
	TeamsByTeamID team.Teams
	Params        filter.ParamSet
	SearchQuery   string
}

//line views/vsprint/List.html:20
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/List.html:20
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vsprint/List.html:22
	edit.StreamSearchForm(qw422016, "", "q", "Search Sprints", p.SearchQuery, ps)
//line views/vsprint/List.html:22
	qw422016.N().S(`</div>
    <div class="right mrs large-buttons">
`)
//line views/vsprint/List.html:24
	if len(p.Models) > 0 {
//line views/vsprint/List.html:24
		qw422016.N().S(`<a href="/admin/db/sprint/_random"><button>`)
//line views/vsprint/List.html:24
		components.StreamSVGButton(qw422016, "gift", ps)
//line views/vsprint/List.html:24
		qw422016.N().S(` Random</button></a>`)
//line views/vsprint/List.html:24
	}
//line views/vsprint/List.html:24
	qw422016.N().S(`      <a href="/admin/db/sprint/_new"><button>`)
//line views/vsprint/List.html:25
	components.StreamSVGButton(qw422016, "plus", ps)
//line views/vsprint/List.html:25
	qw422016.N().S(` New</button></a>
    </div>
    <h3>`)
//line views/vsprint/List.html:27
	components.StreamSVGIcon(qw422016, `sprint`, ps)
//line views/vsprint/List.html:27
	qw422016.N().S(` `)
//line views/vsprint/List.html:27
	qw422016.E().S(ps.Title)
//line views/vsprint/List.html:27
	qw422016.N().S(`</h3>
    <div class="clear"></div>
`)
//line views/vsprint/List.html:29
	if p.SearchQuery != "" {
//line views/vsprint/List.html:29
		qw422016.N().S(`    <hr />
    <em>Search results for [`)
//line views/vsprint/List.html:31
		qw422016.E().S(p.SearchQuery)
//line views/vsprint/List.html:31
		qw422016.N().S(`]</em> (<a href="?">clear</a>)
`)
//line views/vsprint/List.html:32
	}
//line views/vsprint/List.html:33
	if len(p.Models) == 0 {
//line views/vsprint/List.html:33
		qw422016.N().S(`    <div class="mt"><em>No sprints available</em></div>
`)
//line views/vsprint/List.html:35
	} else {
//line views/vsprint/List.html:35
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vsprint/List.html:37
		StreamTable(qw422016, p.Models, p.TeamsByTeamID, p.Params, as, ps)
//line views/vsprint/List.html:37
		qw422016.N().S(`
    </div>
`)
//line views/vsprint/List.html:39
	}
//line views/vsprint/List.html:39
	qw422016.N().S(`  </div>
`)
//line views/vsprint/List.html:41
}

//line views/vsprint/List.html:41
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsprint/List.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsprint/List.html:41
	p.StreamBody(qw422016, as, ps)
//line views/vsprint/List.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/vsprint/List.html:41
}

//line views/vsprint/List.html:41
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsprint/List.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsprint/List.html:41
	p.WriteBody(qb422016, as, ps)
//line views/vsprint/List.html:41
	qs422016 := string(qb422016.B)
//line views/vsprint/List.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsprint/List.html:41
	return qs422016
//line views/vsprint/List.html:41
}
