// Code generated by qtc from "Nav.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Nav.html:2
package layout

//line views/layout/Nav.html:2
import (
	"strings"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/vutil"
)

//line views/layout/Nav.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Nav.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Nav.html:11
func StreamNav(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:11
	qw422016.N().S(`
<nav id="navbar">
  <a class="logo" href="`)
//line views/layout/Nav.html:13
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:13
	qw422016.N().S(`" title="`)
//line views/layout/Nav.html:13
	qw422016.E().S(ps.RootTitle)
//line views/layout/Nav.html:13
	qw422016.N().S(` `)
//line views/layout/Nav.html:13
	qw422016.E().S(as.BuildInfo.String())
//line views/layout/Nav.html:13
	qw422016.N().S(`">`)
//line views/layout/Nav.html:13
	components.StreamSVGRef(qw422016, ps.RootIcon, 32, 32, ``, ps)
//line views/layout/Nav.html:13
	qw422016.N().S(`</a>
  <div class="breadcrumbs">
    <a class="link" href="`)
//line views/layout/Nav.html:15
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:15
	qw422016.N().S(`">`)
//line views/layout/Nav.html:15
	qw422016.E().S(ps.RootTitle)
//line views/layout/Nav.html:15
	qw422016.N().S(`</a>`)
//line views/layout/Nav.html:15
	StreamNavItems(qw422016, ps)
//line views/layout/Nav.html:15
	qw422016.N().S(`
  </div>
`)
//line views/layout/Nav.html:17
	if ps.SearchPath != "-" {
//line views/layout/Nav.html:17
		qw422016.N().S(`  <form action="`)
//line views/layout/Nav.html:18
		qw422016.E().S(ps.SearchPath)
//line views/layout/Nav.html:18
		qw422016.N().S(`" class="search" title="search">
    <input type="search" name="q" placeholder=" " />
    <div class="search-image" style="display: none;"><svg><use xlink:href="#svg-searchbox" /></svg></div>
  </form>
`)
//line views/layout/Nav.html:22
	}
//line views/layout/Nav.html:22
	qw422016.N().S(`  <a class="profile" title="`)
//line views/layout/Nav.html:23
	qw422016.E().S(ps.AuthString())
//line views/layout/Nav.html:23
	qw422016.N().S(`" href="`)
//line views/layout/Nav.html:23
	qw422016.E().S(ps.ProfilePath)
//line views/layout/Nav.html:23
	qw422016.N().S(`">
`)
//line views/layout/Nav.html:24
	if i := ps.Accounts.Image(); i != "" {
//line views/layout/Nav.html:24
		qw422016.N().S(`    <img style="width: 24px; height: 24px;" src="`)
//line views/layout/Nav.html:25
		qw422016.E().S(i)
//line views/layout/Nav.html:25
		qw422016.N().S(`" />
`)
//line views/layout/Nav.html:26
	} else {
//line views/layout/Nav.html:26
		qw422016.N().S(`    `)
//line views/layout/Nav.html:27
		components.StreamSVGRef(qw422016, `profile`, 24, 24, ``, ps)
//line views/layout/Nav.html:27
		qw422016.N().S(`
`)
//line views/layout/Nav.html:28
	}
//line views/layout/Nav.html:28
	qw422016.N().S(`  </a>
`)
//line views/layout/Nav.html:30
	if !ps.HideMenu {
//line views/layout/Nav.html:30
		qw422016.N().S(`  <input type="checkbox" id="menu-toggle-input" style="display: none;" />
  <label class="menu-toggle" for="menu-toggle-input"><div class="spinner diagonal part-1"></div><div class="spinner horizontal"></div><div class="spinner diagonal part-2"></div></label>
  `)
//line views/layout/Nav.html:33
		StreamMenu(qw422016, ps)
//line views/layout/Nav.html:33
		qw422016.N().S(`
`)
//line views/layout/Nav.html:34
	}
//line views/layout/Nav.html:34
	qw422016.N().S(`</nav>`)
//line views/layout/Nav.html:35
}

//line views/layout/Nav.html:35
func WriteNav(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:35
	StreamNav(qw422016, as, ps)
//line views/layout/Nav.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:35
}

//line views/layout/Nav.html:35
func Nav(as *app.State, ps *cutil.PageState) string {
//line views/layout/Nav.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:35
	WriteNav(qb422016, as, ps)
//line views/layout/Nav.html:35
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:35
	return qs422016
//line views/layout/Nav.html:35
}

//line views/layout/Nav.html:37
func StreamNavItems(qw422016 *qt422016.Writer, ps *cutil.PageState) {
//line views/layout/Nav.html:38
	for idx, bc := range ps.Breadcrumbs {
//line views/layout/Nav.html:40
		i := ps.Menu.GetByPath(ps.Breadcrumbs[:idx+1])

//line views/layout/Nav.html:42
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:42
		qw422016.N().S(`<span class="separator">/</span>`)
//line views/layout/Nav.html:44
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:45
		if i == nil {
//line views/layout/Nav.html:47
			bcLink := ""
			if strings.Contains(bc, "||") {
				bci := strings.Index(bc, "||")
				bcLink = bc[bci+2:]
				bc = bc[:bci]
			}

//line views/layout/Nav.html:53
			qw422016.N().S(`<a class="link" href="`)
//line views/layout/Nav.html:54
			qw422016.E().S(bcLink)
//line views/layout/Nav.html:54
			qw422016.N().S(`">`)
//line views/layout/Nav.html:54
			components.StreamSVGRef(qw422016, "play", 28, 28, "breadcrumb-icon", ps)
//line views/layout/Nav.html:54
			qw422016.E().S(bc)
//line views/layout/Nav.html:54
			qw422016.N().S(`</a>`)
//line views/layout/Nav.html:55
		} else {
//line views/layout/Nav.html:55
			qw422016.N().S(`<a class="link" href="`)
//line views/layout/Nav.html:56
			qw422016.E().S(i.Route)
//line views/layout/Nav.html:56
			qw422016.N().S(`">`)
//line views/layout/Nav.html:56
			components.StreamSVGRef(qw422016, i.Icon, 28, 28, "breadcrumb-icon", ps)
//line views/layout/Nav.html:56
			qw422016.E().S(i.Title)
//line views/layout/Nav.html:56
			qw422016.N().S(`</a>`)
//line views/layout/Nav.html:57
		}
//line views/layout/Nav.html:58
	}
//line views/layout/Nav.html:59
}

//line views/layout/Nav.html:59
func WriteNavItems(qq422016 qtio422016.Writer, ps *cutil.PageState) {
//line views/layout/Nav.html:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:59
	StreamNavItems(qw422016, ps)
//line views/layout/Nav.html:59
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:59
}

//line views/layout/Nav.html:59
func NavItems(ps *cutil.PageState) string {
//line views/layout/Nav.html:59
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:59
	WriteNavItems(qb422016, ps)
//line views/layout/Nav.html:59
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:59
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:59
	return qs422016
//line views/layout/Nav.html:59
}
