// Code generated by qtc from "Error.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/verror/Error.html:2
package verror

//line views/verror/Error.html:2
import (
	"strings"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/user"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/layout"
)

//line views/verror/Error.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/verror/Error.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/verror/Error.html:12
func streamerrorStack(qw422016 *qt422016.Writer, ed *util.ErrorDetail) {
//line views/verror/Error.html:12
	qw422016.N().S(`    <div class="overflow full-width">
      <table>
        <tbody>
`)
//line views/verror/Error.html:16
	for _, f := range util.TraceDetail(ed.StackTrace) {
//line views/verror/Error.html:16
		qw422016.N().S(`          <tr>
            <td>
`)
//line views/verror/Error.html:19
		if strings.Contains(f.Key, util.AppKey) {
//line views/verror/Error.html:19
			qw422016.N().S(`              <div class="error-key error-owned">`)
//line views/verror/Error.html:20
			qw422016.E().S(f.Key)
//line views/verror/Error.html:20
			qw422016.N().S(`</div>
`)
//line views/verror/Error.html:21
		} else {
//line views/verror/Error.html:21
			qw422016.N().S(`              <div class="error-key">`)
//line views/verror/Error.html:22
			qw422016.E().S(f.Key)
//line views/verror/Error.html:22
			qw422016.N().S(`</div>
`)
//line views/verror/Error.html:23
		}
//line views/verror/Error.html:23
		qw422016.N().S(`              <div class="error-location">`)
//line views/verror/Error.html:24
		qw422016.E().S(f.Loc)
//line views/verror/Error.html:24
		qw422016.N().S(`</div>
            </td>
          </tr>
`)
//line views/verror/Error.html:27
	}
//line views/verror/Error.html:27
	qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/verror/Error.html:31
}

//line views/verror/Error.html:31
func writeerrorStack(qq422016 qtio422016.Writer, ed *util.ErrorDetail) {
//line views/verror/Error.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/verror/Error.html:31
	streamerrorStack(qw422016, ed)
//line views/verror/Error.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/verror/Error.html:31
}

//line views/verror/Error.html:31
func errorStack(ed *util.ErrorDetail) string {
//line views/verror/Error.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/verror/Error.html:31
	writeerrorStack(qb422016, ed)
//line views/verror/Error.html:31
	qs422016 := string(qb422016.B)
//line views/verror/Error.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/verror/Error.html:31
	return qs422016
//line views/verror/Error.html:31
}

//line views/verror/Error.html:33
type Error struct {
	layout.Basic
	Err *util.ErrorDetail
}

//line views/verror/Error.html:38
func (p *Error) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/verror/Error.html:38
	qw422016.N().S(`
  `)
//line views/verror/Error.html:39
	StreamDetail(qw422016, p.Err, as, ps)
//line views/verror/Error.html:39
	qw422016.N().S(`
`)
//line views/verror/Error.html:40
}

//line views/verror/Error.html:40
func (p *Error) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/verror/Error.html:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/verror/Error.html:40
	p.StreamBody(qw422016, as, ps)
//line views/verror/Error.html:40
	qt422016.ReleaseWriter(qw422016)
//line views/verror/Error.html:40
}

//line views/verror/Error.html:40
func (p *Error) Body(as *app.State, ps *cutil.PageState) string {
//line views/verror/Error.html:40
	qb422016 := qt422016.AcquireByteBuffer()
//line views/verror/Error.html:40
	p.WriteBody(qb422016, as, ps)
//line views/verror/Error.html:40
	qs422016 := string(qb422016.B)
//line views/verror/Error.html:40
	qt422016.ReleaseByteBuffer(qb422016)
//line views/verror/Error.html:40
	return qs422016
//line views/verror/Error.html:40
}

//line views/verror/Error.html:42
func StreamDetail(qw422016 *qt422016.Writer, ed *util.ErrorDetail, as *app.State, ps *cutil.PageState) {
//line views/verror/Error.html:42
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/verror/Error.html:44
	qw422016.E().S(ed.Message)
//line views/verror/Error.html:44
	qw422016.N().S(`</h3>
    <em>Internal Server Error</em>
`)
//line views/verror/Error.html:46
	if user.IsAdmin(ps.Accounts) {
//line views/verror/Error.html:46
		qw422016.N().S(`    `)
//line views/verror/Error.html:47
		streamerrorStack(qw422016, ed)
//line views/verror/Error.html:47
		qw422016.N().S(` `)
//line views/verror/Error.html:47
		cause := ed.Cause

//line views/verror/Error.html:47
		qw422016.N().S(`
`)
//line views/verror/Error.html:48
		for cause != nil {
//line views/verror/Error.html:48
			qw422016.N().S(`    <h3>Caused by</h3>
    <div>`)
//line views/verror/Error.html:50
			qw422016.E().S(cause.Message)
//line views/verror/Error.html:50
			qw422016.N().S(`</div>`)
//line views/verror/Error.html:50
			streamerrorStack(qw422016, cause)
//line views/verror/Error.html:50
			cause = cause.Cause

//line views/verror/Error.html:50
			qw422016.N().S(`
`)
//line views/verror/Error.html:51
		}
//line views/verror/Error.html:52
	}
//line views/verror/Error.html:52
	qw422016.N().S(`  </div>
`)
//line views/verror/Error.html:54
}

//line views/verror/Error.html:54
func WriteDetail(qq422016 qtio422016.Writer, ed *util.ErrorDetail, as *app.State, ps *cutil.PageState) {
//line views/verror/Error.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/verror/Error.html:54
	StreamDetail(qw422016, ed, as, ps)
//line views/verror/Error.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/verror/Error.html:54
}

//line views/verror/Error.html:54
func Detail(ed *util.ErrorDetail, as *app.State, ps *cutil.PageState) string {
//line views/verror/Error.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/verror/Error.html:54
	WriteDetail(qb422016, ed, as, ps)
//line views/verror/Error.html:54
	qs422016 := string(qb422016.B)
//line views/verror/Error.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/verror/Error.html:54
	return qs422016
//line views/verror/Error.html:54
}
