// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vdatabase/List.html:1
package vdatabase

//line views/vdatabase/List.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/database"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vdatabase/List.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vdatabase/List.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vdatabase/List.html:10
type List struct {
	layout.Basic
	Keys     []string
	Services map[string]*database.Service
}

//line views/vdatabase/List.html:16
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/List.html:16
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vdatabase/List.html:18
	components.StreamSVGIcon(qw422016, `database`, ps)
//line views/vdatabase/List.html:18
	qw422016.N().S(` Databases</h3>
    <em>`)
//line views/vdatabase/List.html:19
	qw422016.E().S(util.StringPlural(len(p.Keys), "database"))
//line views/vdatabase/List.html:19
	qw422016.N().S(` available</em>
  </div>
`)
//line views/vdatabase/List.html:21
	for _, key := range p.Keys {
//line views/vdatabase/List.html:22
		svc := p.Services[key]

//line views/vdatabase/List.html:22
		qw422016.N().S(`  <div class="card">
    <div class="right"><em>`)
//line views/vdatabase/List.html:24
		qw422016.E().S(svc.Type.Title)
//line views/vdatabase/List.html:24
		qw422016.N().S(`</em></div>
    <h3><a href="/admin/database/`)
//line views/vdatabase/List.html:25
		qw422016.E().S(key)
//line views/vdatabase/List.html:25
		qw422016.N().S(`">`)
//line views/vdatabase/List.html:25
		components.StreamSVGIcon(qw422016, `database`, ps)
//line views/vdatabase/List.html:25
		qw422016.N().S(` `)
//line views/vdatabase/List.html:25
		qw422016.E().S(svc.Key)
//line views/vdatabase/List.html:25
		qw422016.N().S(`</a></h3>
  </div>
`)
//line views/vdatabase/List.html:27
	}
//line views/vdatabase/List.html:28
}

//line views/vdatabase/List.html:28
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/List.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdatabase/List.html:28
	p.StreamBody(qw422016, as, ps)
//line views/vdatabase/List.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/vdatabase/List.html:28
}

//line views/vdatabase/List.html:28
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vdatabase/List.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdatabase/List.html:28
	p.WriteBody(qb422016, as, ps)
//line views/vdatabase/List.html:28
	qs422016 := string(qb422016.B)
//line views/vdatabase/List.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdatabase/List.html:28
	return qs422016
//line views/vdatabase/List.html:28
}
