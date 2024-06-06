// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vtheme/List.html:2
package vtheme

//line views/vtheme/List.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/theme"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vtheme/List.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/List.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/List.html:10
type List struct {
	layout.Basic
	Themes theme.Themes
}

//line views/vtheme/List.html:15
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/List.html:15
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vtheme/List.html:17
	components.StreamSVGIcon(qw422016, `eye`, ps)
//line views/vtheme/List.html:17
	qw422016.N().S(`Add Theme</h3>
    <div class="mt">
      <a href="/theme/new" title="add new theme"><button>New Theme</button></a>
    </div>
  </div>
  <div class="card">
    <h3>`)
//line views/vtheme/List.html:23
	components.StreamSVGIcon(qw422016, `play`, ps)
//line views/vtheme/List.html:23
	qw422016.N().S(`Current Themes</h3>
    <div class="overflow full-width">
      <div class="theme-container mt">
`)
//line views/vtheme/List.html:26
	for _, t := range p.Themes {
//line views/vtheme/List.html:26
		qw422016.N().S(`        <div class="theme-item">
          <a href="/theme/`)
//line views/vtheme/List.html:28
		qw422016.N().U(t.Key)
//line views/vtheme/List.html:28
		qw422016.N().S(`">
            `)
//line views/vtheme/List.html:29
		StreamMockupTheme(qw422016, t, true, "app", 5, ps)
//line views/vtheme/List.html:29
		qw422016.N().S(`
          </a>
        </div>
`)
//line views/vtheme/List.html:32
	}
//line views/vtheme/List.html:32
	qw422016.N().S(`      </div>
    </div>
  </div>
`)
//line views/vtheme/List.html:36
}

//line views/vtheme/List.html:36
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/List.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/List.html:36
	p.StreamBody(qw422016, as, ps)
//line views/vtheme/List.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/List.html:36
}

//line views/vtheme/List.html:36
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vtheme/List.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/List.html:36
	p.WriteBody(qb422016, as, ps)
//line views/vtheme/List.html:36
	qs422016 := string(qb422016.B)
//line views/vtheme/List.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/List.html:36
	return qs422016
//line views/vtheme/List.html:36
}
