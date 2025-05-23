// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vretro/vrhistory/Detail.html:1
package vrhistory

//line views/vretro/vrhistory/Detail.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/retro"
	"github.com/kyleu/rituals/app/retro/rhistory"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/view"
	"github.com/kyleu/rituals/views/layout"
)

//line views/vretro/vrhistory/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vretro/vrhistory/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vretro/vrhistory/Detail.html:11
type Detail struct {
	layout.Basic
	Model          *rhistory.RetroHistory
	RetroByRetroID *retro.Retro
	Paths          []string
}

//line views/vretro/vrhistory/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrhistory/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-retroHistory"><button type="button" title="JSON">`)
//line views/vretro/vrhistory/Detail.html:21
	components.StreamSVGButton(qw422016, "code", ps)
//line views/vretro/vrhistory/Detail.html:21
	qw422016.N().S(`</button></a>
      <a href="`)
//line views/vretro/vrhistory/Detail.html:22
	qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/vretro/vrhistory/Detail.html:22
	qw422016.N().S(`/edit" title="Edit"><button>`)
//line views/vretro/vrhistory/Detail.html:22
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vretro/vrhistory/Detail.html:22
	qw422016.N().S(`</button></a>
    </div>
    <h3>`)
//line views/vretro/vrhistory/Detail.html:24
	components.StreamSVGIcon(qw422016, `history`, ps)
//line views/vretro/vrhistory/Detail.html:24
	qw422016.N().S(` `)
//line views/vretro/vrhistory/Detail.html:24
	qw422016.E().S(p.Model.TitleString())
//line views/vretro/vrhistory/Detail.html:24
	qw422016.N().S(`</h3>
    <div><a href="`)
//line views/vretro/vrhistory/Detail.html:25
	qw422016.E().S(rhistory.Route(p.Paths...))
//line views/vretro/vrhistory/Detail.html:25
	qw422016.N().S(`"><em>History</em></a></div>
    `)
//line views/vretro/vrhistory/Detail.html:26
	StreamDetailTable(qw422016, p, ps)
//line views/vretro/vrhistory/Detail.html:26
	qw422016.N().S(`
  </div>
`)
//line views/vretro/vrhistory/Detail.html:29
	qw422016.N().S(`  `)
//line views/vretro/vrhistory/Detail.html:30
	components.StreamJSONModal(qw422016, "retroHistory", "History JSON", p.Model, 1)
//line views/vretro/vrhistory/Detail.html:30
	qw422016.N().S(`
`)
//line views/vretro/vrhistory/Detail.html:31
}

//line views/vretro/vrhistory/Detail.html:31
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vretro/vrhistory/Detail.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrhistory/Detail.html:31
	p.StreamBody(qw422016, as, ps)
//line views/vretro/vrhistory/Detail.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrhistory/Detail.html:31
}

//line views/vretro/vrhistory/Detail.html:31
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vretro/vrhistory/Detail.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrhistory/Detail.html:31
	p.WriteBody(qb422016, as, ps)
//line views/vretro/vrhistory/Detail.html:31
	qs422016 := string(qb422016.B)
//line views/vretro/vrhistory/Detail.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrhistory/Detail.html:31
	return qs422016
//line views/vretro/vrhistory/Detail.html:31
}

//line views/vretro/vrhistory/Detail.html:33
func StreamDetailTable(qw422016 *qt422016.Writer, p *Detail, ps *cutil.PageState) {
//line views/vretro/vrhistory/Detail.html:33
	qw422016.N().S(`
  <div class="mt overflow full-width">
    <table>
      <tbody>
        <tr>
          <th class="shrink" title="String text">Slug</th>
          <td>`)
//line views/vretro/vrhistory/Detail.html:39
	view.StreamString(qw422016, p.Model.Slug)
//line views/vretro/vrhistory/Detail.html:39
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Retro ID</th>
          <td class="nowrap">
            `)
//line views/vretro/vrhistory/Detail.html:44
	if x := p.RetroByRetroID; x != nil {
//line views/vretro/vrhistory/Detail.html:44
		qw422016.N().S(`
            `)
//line views/vretro/vrhistory/Detail.html:45
		qw422016.E().S(x.TitleString())
//line views/vretro/vrhistory/Detail.html:45
		qw422016.N().S(` <a title="Retro" href="`)
//line views/vretro/vrhistory/Detail.html:45
		qw422016.E().S(x.WebPath(p.Paths...))
//line views/vretro/vrhistory/Detail.html:45
		qw422016.N().S(`">`)
//line views/vretro/vrhistory/Detail.html:45
		components.StreamSVGLink(qw422016, `retro`, ps)
//line views/vretro/vrhistory/Detail.html:45
		qw422016.N().S(`</a>
            `)
//line views/vretro/vrhistory/Detail.html:46
	} else {
//line views/vretro/vrhistory/Detail.html:46
		qw422016.N().S(`
            `)
//line views/vretro/vrhistory/Detail.html:47
		view.StreamUUID(qw422016, &p.Model.RetroID)
//line views/vretro/vrhistory/Detail.html:47
		qw422016.N().S(`
            `)
//line views/vretro/vrhistory/Detail.html:48
	}
//line views/vretro/vrhistory/Detail.html:48
	qw422016.N().S(`
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Retro Name</th>
          <td>`)
//line views/vretro/vrhistory/Detail.html:53
	view.StreamString(qw422016, p.Model.RetroName)
//line views/vretro/vrhistory/Detail.html:53
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vretro/vrhistory/Detail.html:57
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vretro/vrhistory/Detail.html:57
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vretro/vrhistory/Detail.html:62
}

//line views/vretro/vrhistory/Detail.html:62
func WriteDetailTable(qq422016 qtio422016.Writer, p *Detail, ps *cutil.PageState) {
//line views/vretro/vrhistory/Detail.html:62
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vretro/vrhistory/Detail.html:62
	StreamDetailTable(qw422016, p, ps)
//line views/vretro/vrhistory/Detail.html:62
	qt422016.ReleaseWriter(qw422016)
//line views/vretro/vrhistory/Detail.html:62
}

//line views/vretro/vrhistory/Detail.html:62
func DetailTable(p *Detail, ps *cutil.PageState) string {
//line views/vretro/vrhistory/Detail.html:62
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vretro/vrhistory/Detail.html:62
	WriteDetailTable(qb422016, p, ps)
//line views/vretro/vrhistory/Detail.html:62
	qs422016 := string(qb422016.B)
//line views/vretro/vrhistory/Detail.html:62
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vretro/vrhistory/Detail.html:62
	return qs422016
//line views/vretro/vrhistory/Detail.html:62
}
