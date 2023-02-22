// Code generated by qtc from "Nav.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Nav.html:2
package layout

//line views/layout/Nav.html:2
import (
	"strings"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cmenu"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/menu"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/vutil"
)

//line views/layout/Nav.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Nav.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Nav.html:13
func StreamNav(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:13
	qw422016.N().S(`
<nav id="navbar">
  <a class="logo" href="`)
//line views/layout/Nav.html:15
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:15
	qw422016.N().S(`" title="`)
//line views/layout/Nav.html:15
	qw422016.E().S(ps.RootTitle)
//line views/layout/Nav.html:15
	qw422016.N().S(` `)
//line views/layout/Nav.html:15
	qw422016.E().S(as.BuildInfo.String())
//line views/layout/Nav.html:15
	qw422016.N().S(`">`)
//line views/layout/Nav.html:15
	components.StreamSVGRef(qw422016, ps.RootIcon, 32, 32, ``, ps)
//line views/layout/Nav.html:15
	qw422016.N().S(`</a>
  <div class="breadcrumbs">
    <a class="link" href="`)
//line views/layout/Nav.html:17
	qw422016.E().S(ps.RootPath)
//line views/layout/Nav.html:17
	qw422016.N().S(`">`)
//line views/layout/Nav.html:17
	qw422016.E().S(ps.RootTitle)
//line views/layout/Nav.html:17
	qw422016.N().S(`</a>`)
//line views/layout/Nav.html:17
	StreamNavItems(qw422016, ps.Menu, ps.Breadcrumbs)
//line views/layout/Nav.html:17
	qw422016.N().S(`
  </div>
`)
//line views/layout/Nav.html:19
	if ps.SearchPath != "-" {
//line views/layout/Nav.html:19
		qw422016.N().S(`  <form action="`)
//line views/layout/Nav.html:20
		qw422016.E().S(ps.SearchPath)
//line views/layout/Nav.html:20
		qw422016.N().S(`" class="search" title="search">
    <input type="search" name="q" placeholder=" " />
    <div class="search-image" style="display: none;"><svg><use xlink:href="#svg-searchbox" /></svg></div>
  </form>
`)
//line views/layout/Nav.html:24
	}
//line views/layout/Nav.html:24
	qw422016.N().S(`  <a class="profile" title="`)
//line views/layout/Nav.html:25
	qw422016.E().S(ps.AuthString())
//line views/layout/Nav.html:25
	qw422016.N().S(`" href="`)
//line views/layout/Nav.html:25
	qw422016.E().S(ps.ProfilePath)
//line views/layout/Nav.html:25
	qw422016.N().S(`">
`)
//line views/layout/Nav.html:26
	if i := ps.Accounts.Image(); i != "" {
//line views/layout/Nav.html:26
		qw422016.N().S(`    <img style="width: 24px; height: 24px;" src="`)
//line views/layout/Nav.html:27
		qw422016.E().S(i)
//line views/layout/Nav.html:27
		qw422016.N().S(`" />
`)
//line views/layout/Nav.html:28
	} else {
//line views/layout/Nav.html:28
		qw422016.N().S(`    `)
//line views/layout/Nav.html:29
		components.StreamSVGRef(qw422016, `profile`, 24, 24, ``, ps)
//line views/layout/Nav.html:29
		qw422016.N().S(`
`)
//line views/layout/Nav.html:30
	}
//line views/layout/Nav.html:30
	qw422016.N().S(`  </a>`)
//line views/layout/Nav.html:31
	if !ps.HideMenu {
//line views/layout/Nav.html:31
		qw422016.N().S(`

  <input type="checkbox" id="menu-toggle-input" style="display: none;" />
  <label class="menu-toggle" for="menu-toggle-input"><div class="spinner diagonal part-1"></div><div class="spinner horizontal"></div><div class="spinner diagonal part-2"></div></label>
  `)
//line views/layout/Nav.html:35
		StreamMenu(qw422016, ps)
//line views/layout/Nav.html:35
	}
//line views/layout/Nav.html:35
	qw422016.N().S(`
</nav>`)
//line views/layout/Nav.html:36
}

//line views/layout/Nav.html:36
func WriteNav(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Nav.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:36
	StreamNav(qw422016, as, ps)
//line views/layout/Nav.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:36
}

//line views/layout/Nav.html:36
func Nav(as *app.State, ps *cutil.PageState) string {
//line views/layout/Nav.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:36
	WriteNav(qb422016, as, ps)
//line views/layout/Nav.html:36
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:36
	return qs422016
//line views/layout/Nav.html:36
}

//line views/layout/Nav.html:38
func StreamNavItems(qw422016 *qt422016.Writer, m menu.Items, breadcrumbs cmenu.Breadcrumbs) {
//line views/layout/Nav.html:39
	for idx, bc := range breadcrumbs {
//line views/layout/Nav.html:41
		i := m.GetByPath(breadcrumbs[:idx+1])

//line views/layout/Nav.html:43
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:43
		qw422016.N().S(`<span class="separator">/</span>`)
//line views/layout/Nav.html:45
		vutil.StreamIndent(qw422016, true, 2)
//line views/layout/Nav.html:46
		if i == nil {
//line views/layout/Nav.html:48
			bcLink := ""
			if strings.Contains(bc, "||") {
				bci := strings.Index(bc, "||")
				bcLink = bc[bci+2:]
				bc = bc[:bci]
			}

//line views/layout/Nav.html:54
			qw422016.N().S(`<a class="link" href="`)
//line views/layout/Nav.html:55
			qw422016.E().S(bcLink)
//line views/layout/Nav.html:55
			qw422016.N().S(`">`)
//line views/layout/Nav.html:55
			qw422016.E().S(bc)
//line views/layout/Nav.html:55
			qw422016.N().S(`</a>`)
//line views/layout/Nav.html:56
		} else {
//line views/layout/Nav.html:56
			qw422016.N().S(`<a class="link" href="`)
//line views/layout/Nav.html:57
			qw422016.E().S(i.Route)
//line views/layout/Nav.html:57
			qw422016.N().S(`">`)
//line views/layout/Nav.html:57
			qw422016.E().S(i.Title)
//line views/layout/Nav.html:57
			qw422016.N().S(`</a>`)
//line views/layout/Nav.html:58
		}
//line views/layout/Nav.html:59
	}
//line views/layout/Nav.html:60
}

//line views/layout/Nav.html:60
func WriteNavItems(qq422016 qtio422016.Writer, m menu.Items, breadcrumbs cmenu.Breadcrumbs) {
//line views/layout/Nav.html:60
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Nav.html:60
	StreamNavItems(qw422016, m, breadcrumbs)
//line views/layout/Nav.html:60
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Nav.html:60
}

//line views/layout/Nav.html:60
func NavItems(m menu.Items, breadcrumbs cmenu.Breadcrumbs) string {
//line views/layout/Nav.html:60
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Nav.html:60
	WriteNavItems(qb422016, m, breadcrumbs)
//line views/layout/Nav.html:60
	qs422016 := string(qb422016.B)
//line views/layout/Nav.html:60
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Nav.html:60
	return qs422016
//line views/layout/Nav.html:60
}
