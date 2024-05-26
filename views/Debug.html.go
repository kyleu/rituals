// Code generated by qtc from "Debug.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/Debug.html:2
package views

//line views/Debug.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/Debug.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/Debug.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/Debug.html:10
type Debug struct{ layout.Basic }

//line views/Debug.html:12
func (p *Debug) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Debug.html:12
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/Debug.html:14
	components.StreamSVGRefIcon(qw422016, `graph`, ps)
//line views/Debug.html:14
	qw422016.E().S(ps.Title)
//line views/Debug.html:14
	qw422016.N().S(`</h3>
`)
//line views/Debug.html:15
	if s, ok := ps.Data.(string); ok {
//line views/Debug.html:15
		qw422016.N().S(`    <div class="overflow full-width"><pre class="mt">`)
//line views/Debug.html:16
		qw422016.E().S(s)
//line views/Debug.html:16
		qw422016.N().S(`</pre></div>
`)
//line views/Debug.html:17
	} else {
//line views/Debug.html:18
		if util.ArrayTest(ps.Data) {
//line views/Debug.html:18
			qw422016.N().S(`    <div class="overflow full-width">
      <table class="mt">
        <tbody>
`)
//line views/Debug.html:22
			a := util.ArrayFromAny[any](ps.Data)

//line views/Debug.html:23
			for idx, x := range a {
//line views/Debug.html:24
				if idx < 32 {
//line views/Debug.html:24
					qw422016.N().S(`          <tr>
            <th class="shrink">`)
//line views/Debug.html:26
					qw422016.N().D(idx + 1)
//line views/Debug.html:26
					qw422016.N().S(`</th>
            <td>`)
//line views/Debug.html:27
					qw422016.N().S(components.JSON(x))
//line views/Debug.html:27
					qw422016.N().S(`</td>
          </tr>
`)
//line views/Debug.html:29
				}
//line views/Debug.html:30
			}
//line views/Debug.html:31
			if len(a) > 32 {
//line views/Debug.html:31
				qw422016.N().S(`          <tr>
            <td class="shrink" colspan="2"><em>...and [`)
//line views/Debug.html:33
				qw422016.N().D(len(a) - 32)
//line views/Debug.html:33
				qw422016.N().S(`] more...</em></td>
          </tr>
`)
//line views/Debug.html:35
			}
//line views/Debug.html:35
			qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/Debug.html:39
		} else {
//line views/Debug.html:39
			qw422016.N().S(`    <div class="mt">`)
//line views/Debug.html:40
			qw422016.N().S(components.JSON(ps.Data))
//line views/Debug.html:40
			qw422016.N().S(`</div>
`)
//line views/Debug.html:41
		}
//line views/Debug.html:42
	}
//line views/Debug.html:42
	qw422016.N().S(`  </div>
`)
//line views/Debug.html:44
}

//line views/Debug.html:44
func (p *Debug) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Debug.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Debug.html:44
	p.StreamBody(qw422016, as, ps)
//line views/Debug.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/Debug.html:44
}

//line views/Debug.html:44
func (p *Debug) Body(as *app.State, ps *cutil.PageState) string {
//line views/Debug.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Debug.html:44
	p.WriteBody(qb422016, as, ps)
//line views/Debug.html:44
	qs422016 := string(qb422016.B)
//line views/Debug.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Debug.html:44
	return qs422016
//line views/Debug.html:44
}
