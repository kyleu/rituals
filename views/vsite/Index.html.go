// Code generated by qtc from "Index.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsite/Index.html:2
package vsite

//line views/vsite/Index.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vsite/Index.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsite/Index.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsite/Index.html:10
type Index struct{ layout.Basic }

//line views/vsite/Index.html:12
func (p *Index) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/Index.html:12
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vsite/Index.html:14
	components.StreamSVGRefIcon(qw422016, `app`, ps)
//line views/vsite/Index.html:14
	qw422016.E().S(util.AppName)
//line views/vsite/Index.html:14
	qw422016.N().S(`</h3>
    <p>TODO</p>
    <p>
      <a href="/download"><button>Download</button></a>
      <a href="https://github.com/kyleu/rituals"><button>Source Code</button></a>
    </p>
  </div>
`)
//line views/vsite/Index.html:21
}

//line views/vsite/Index.html:21
func (p *Index) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/Index.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsite/Index.html:21
	p.StreamBody(qw422016, as, ps)
//line views/vsite/Index.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/vsite/Index.html:21
}

//line views/vsite/Index.html:21
func (p *Index) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsite/Index.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsite/Index.html:21
	p.WriteBody(qb422016, as, ps)
//line views/vsite/Index.html:21
	qs422016 := string(qb422016.B)
//line views/vsite/Index.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsite/Index.html:21
	return qs422016
//line views/vsite/Index.html:21
}
