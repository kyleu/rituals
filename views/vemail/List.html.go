// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vemail/List.html:1
package vemail

//line views/vemail/List.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/email"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vemail/List.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vemail/List.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vemail/List.html:11
type List struct {
	layout.Basic
	Models        email.Emails
	UsersByUserID user.Users
	Params        filter.ParamSet
	Paths         []string
}

//line views/vemail/List.html:19
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vemail/List.html:19
	qw422016.N().S(`
  <div class="card">
    <div class="right mrs large-buttons">
`)
//line views/vemail/List.html:22
	if len(p.Models) > 1 {
//line views/vemail/List.html:22
		qw422016.N().S(`<a href="`)
//line views/vemail/List.html:22
		qw422016.E().S(email.Route(p.Paths...))
//line views/vemail/List.html:22
		qw422016.N().S(`/_random"><button>`)
//line views/vemail/List.html:22
		components.StreamSVGButton(qw422016, "gift", ps)
//line views/vemail/List.html:22
		qw422016.N().S(` Random</button></a>`)
//line views/vemail/List.html:22
	}
//line views/vemail/List.html:22
	qw422016.N().S(`      <a href="`)
//line views/vemail/List.html:23
	qw422016.E().S(email.Route(p.Paths...))
//line views/vemail/List.html:23
	qw422016.N().S(`/_new"><button>`)
//line views/vemail/List.html:23
	components.StreamSVGButton(qw422016, "plus", ps)
//line views/vemail/List.html:23
	qw422016.N().S(` New</button></a>
    </div>
    <h3>`)
//line views/vemail/List.html:25
	components.StreamSVGIcon(qw422016, `email`, ps)
//line views/vemail/List.html:25
	qw422016.N().S(` `)
//line views/vemail/List.html:25
	qw422016.E().S(ps.Title)
//line views/vemail/List.html:25
	qw422016.N().S(`</h3>
`)
//line views/vemail/List.html:26
	if len(p.Models) == 0 {
//line views/vemail/List.html:26
		qw422016.N().S(`    <div class="mt"><em>No emails available</em></div>
`)
//line views/vemail/List.html:28
	} else {
//line views/vemail/List.html:28
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vemail/List.html:30
		StreamTable(qw422016, p.Models, p.UsersByUserID, p.Params, as, ps, p.Paths...)
//line views/vemail/List.html:30
		qw422016.N().S(`
    </div>
`)
//line views/vemail/List.html:32
	}
//line views/vemail/List.html:32
	qw422016.N().S(`  </div>
`)
//line views/vemail/List.html:34
}

//line views/vemail/List.html:34
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vemail/List.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vemail/List.html:34
	p.StreamBody(qw422016, as, ps)
//line views/vemail/List.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/vemail/List.html:34
}

//line views/vemail/List.html:34
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vemail/List.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vemail/List.html:34
	p.WriteBody(qb422016, as, ps)
//line views/vemail/List.html:34
	qs422016 := string(qb422016.B)
//line views/vemail/List.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vemail/List.html:34
	return qs422016
//line views/vemail/List.html:34
}
