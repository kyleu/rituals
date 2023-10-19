// Code generated by qtc from "SVG.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/SVG.html:2
package components

//line views/components/SVG.html:2
import (
	"fmt"
	"strings"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
)

//line views/components/SVG.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/SVG.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/SVG.html:11
func StreamSVG(qw422016 *qt422016.Writer, k string) {
//line views/components/SVG.html:11
	qw422016.N().S(util.SVGLibrary[k])
//line views/components/SVG.html:11
}

//line views/components/SVG.html:11
func WriteSVG(qq422016 qtio422016.Writer, k string) {
//line views/components/SVG.html:11
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:11
	StreamSVG(qw422016, k)
//line views/components/SVG.html:11
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:11
}

//line views/components/SVG.html:11
func SVG(k string) string {
//line views/components/SVG.html:11
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:11
	WriteSVG(qb422016, k)
//line views/components/SVG.html:11
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:11
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:11
	return qs422016
//line views/components/SVG.html:11
}

//line views/components/SVG.html:13
func StreamSVGRef(qw422016 *qt422016.Writer, k string, w int, h int, cls string, ps *cutil.PageState) {
//line views/components/SVG.html:14
	if k != "" {
//line views/components/SVG.html:16
		ps.AddIcon(k)
		if w == 0 {
			w = 20
		}
		if h == 0 {
			h = 20
		}
		style := fmt.Sprintf("width: %dpx; height: %dpx;", w, h)

//line views/components/SVG.html:21
		if cls == "" {
//line views/components/SVG.html:21
			qw422016.N().S(`<svg style="`)
//line views/components/SVG.html:22
			qw422016.E().S(style)
//line views/components/SVG.html:22
			qw422016.N().S(`"><use xlink:href="#svg-`)
//line views/components/SVG.html:22
			qw422016.E().S(k)
//line views/components/SVG.html:22
			qw422016.N().S(`" /></svg>`)
//line views/components/SVG.html:23
		} else {
//line views/components/SVG.html:23
			qw422016.N().S(`<svg class="`)
//line views/components/SVG.html:24
			qw422016.E().S(cls)
//line views/components/SVG.html:24
			qw422016.N().S(`" style="`)
//line views/components/SVG.html:24
			qw422016.E().S(style)
//line views/components/SVG.html:24
			qw422016.N().S(`"><use xlink:href="#svg-`)
//line views/components/SVG.html:24
			qw422016.E().S(k)
//line views/components/SVG.html:24
			qw422016.N().S(`" /></svg>`)
//line views/components/SVG.html:25
		}
//line views/components/SVG.html:26
	}
//line views/components/SVG.html:27
}

//line views/components/SVG.html:27
func WriteSVGRef(qq422016 qtio422016.Writer, k string, w int, h int, cls string, ps *cutil.PageState) {
//line views/components/SVG.html:27
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:27
	StreamSVGRef(qw422016, k, w, h, cls, ps)
//line views/components/SVG.html:27
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:27
}

//line views/components/SVG.html:27
func SVGRef(k string, w int, h int, cls string, ps *cutil.PageState) string {
//line views/components/SVG.html:27
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:27
	WriteSVGRef(qb422016, k, w, h, cls, ps)
//line views/components/SVG.html:27
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:27
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:27
	return qs422016
//line views/components/SVG.html:27
}

//line views/components/SVG.html:29
func StreamIcon(qw422016 *qt422016.Writer, k string, size int, cls string, ps *cutil.PageState) {
//line views/components/SVG.html:30
	if strings.Contains(k, "/") {
//line views/components/SVG.html:30
		qw422016.N().S(`<img src="`)
//line views/components/SVG.html:31
		qw422016.E().S(k)
//line views/components/SVG.html:31
		qw422016.N().S(`" style="width:`)
//line views/components/SVG.html:31
		qw422016.N().D(size)
//line views/components/SVG.html:31
		qw422016.N().S(`px; height:`)
//line views/components/SVG.html:31
		qw422016.N().D(size)
//line views/components/SVG.html:31
		qw422016.N().S(`px;" />`)
//line views/components/SVG.html:32
	} else {
//line views/components/SVG.html:33
		StreamSVGRef(qw422016, k, size, size, cls, ps)
//line views/components/SVG.html:34
	}
//line views/components/SVG.html:35
}

//line views/components/SVG.html:35
func WriteIcon(qq422016 qtio422016.Writer, k string, size int, cls string, ps *cutil.PageState) {
//line views/components/SVG.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:35
	StreamIcon(qw422016, k, size, cls, ps)
//line views/components/SVG.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:35
}

//line views/components/SVG.html:35
func Icon(k string, size int, cls string, ps *cutil.PageState) string {
//line views/components/SVG.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:35
	WriteIcon(qb422016, k, size, cls, ps)
//line views/components/SVG.html:35
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:35
	return qs422016
//line views/components/SVG.html:35
}

//line views/components/SVG.html:37
func StreamSVGRefIcon(qw422016 *qt422016.Writer, k string, ps *cutil.PageState) {
//line views/components/SVG.html:38
	StreamSVGRef(qw422016, k, 20, 20, "icon", ps)
//line views/components/SVG.html:39
}

//line views/components/SVG.html:39
func WriteSVGRefIcon(qq422016 qtio422016.Writer, k string, ps *cutil.PageState) {
//line views/components/SVG.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:39
	StreamSVGRefIcon(qw422016, k, ps)
//line views/components/SVG.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:39
}

//line views/components/SVG.html:39
func SVGRefIcon(k string, ps *cutil.PageState) string {
//line views/components/SVG.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:39
	WriteSVGRefIcon(qb422016, k, ps)
//line views/components/SVG.html:39
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:39
	return qs422016
//line views/components/SVG.html:39
}

//line views/components/SVG.html:41
func StreamIconGallery(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/components/SVG.html:41
	qw422016.N().S(`  <div class="flex-wrap mt">
`)
//line views/components/SVG.html:43
	for _, k := range util.SVGIconKeys {
//line views/components/SVG.html:43
		qw422016.N().S(`    <div class="icon-gallery-icon">
      <div class="gallery-svg">`)
//line views/components/SVG.html:45
		StreamSVGRef(qw422016, k, 64, 64, "icon", ps)
//line views/components/SVG.html:45
		qw422016.N().S(`</div>
      <div class="gallery-title">`)
//line views/components/SVG.html:46
		qw422016.E().S(k)
//line views/components/SVG.html:46
		qw422016.N().S(`</div>
    </div>
`)
//line views/components/SVG.html:48
	}
//line views/components/SVG.html:48
	qw422016.N().S(`  </div>
`)
//line views/components/SVG.html:50
}

//line views/components/SVG.html:50
func WriteIconGallery(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/components/SVG.html:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/SVG.html:50
	StreamIconGallery(qw422016, as, ps)
//line views/components/SVG.html:50
	qt422016.ReleaseWriter(qw422016)
//line views/components/SVG.html:50
}

//line views/components/SVG.html:50
func IconGallery(as *app.State, ps *cutil.PageState) string {
//line views/components/SVG.html:50
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/SVG.html:50
	WriteIconGallery(qb422016, as, ps)
//line views/components/SVG.html:50
	qs422016 := string(qb422016.B)
//line views/components/SVG.html:50
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/SVG.html:50
	return qs422016
//line views/components/SVG.html:50
}
