// Code generated by qtc from "Editor.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vtheme/Editor.html:2
package vtheme

//line views/vtheme/Editor.html:2
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/theme"
)

//line views/vtheme/Editor.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/Editor.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/Editor.html:8
func StreamEditor(qw422016 *qt422016.Writer, title string, navTitle string, t *theme.Theme, icon string, as *app.State, ps *cutil.PageState) {
//line views/vtheme/Editor.html:8
	qw422016.N().S(`
  <div class="card">
    <table class="centered min-200">
      <thead>
      <tr>
        <th class="left-align shrink">`)
//line views/vtheme/Editor.html:13
	qw422016.E().S(title)
//line views/vtheme/Editor.html:13
	qw422016.N().S(`</th>
        <th class="bl" colspan="2">
          <div>Light</div>
          <div id="mockup-light">`)
//line views/vtheme/Editor.html:16
	StreamMockupColors(qw422016, navTitle, "", t.Light, false, icon, 5, ps)
//line views/vtheme/Editor.html:16
	qw422016.N().S(`</div>
        </th>
        <th class="bl" colspan="2">
          <div>Dark</div>
          <div id="mockup-dark">`)
//line views/vtheme/Editor.html:20
	StreamMockupColors(qw422016, navTitle, "", t.Dark, false, icon, 5, ps)
//line views/vtheme/Editor.html:20
	qw422016.N().S(`</div>
        </th>
      </tr>
      <tr>
        <th></th>
        <th class="bl">Background</th>
        <th>Foreground</th>
        <th class="bl">Background</th>
        <th>Foreground</th>
      </tr>
      </thead>
`)
//line views/vtheme/Editor.html:32
	const lp = "light"
	const dp = "dark"

//line views/vtheme/Editor.html:34
	qw422016.N().S(`      <tbody>
        <tr>
          <th class="left-align shrink">Main Content</th>
          <td class="bl">`)
//line views/vtheme/Editor.html:38
	streamcinput(qw422016, lp, "background", t.Light.Background)
//line views/vtheme/Editor.html:38
	qw422016.N().S(`</td>
          <td>`)
//line views/vtheme/Editor.html:39
	streamcinput(qw422016, lp, "foreground", t.Light.Foreground)
//line views/vtheme/Editor.html:39
	qw422016.N().S(`</td>
          <td class="bl">`)
//line views/vtheme/Editor.html:40
	streamcinput(qw422016, dp, "background", t.Dark.Background)
//line views/vtheme/Editor.html:40
	qw422016.N().S(`</td>
          <td>`)
//line views/vtheme/Editor.html:41
	streamcinput(qw422016, dp, "foreground", t.Dark.Foreground)
//line views/vtheme/Editor.html:41
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="left-align shrink">Muted</th>
          <td class="bl">`)
//line views/vtheme/Editor.html:45
	streamcinput(qw422016, lp, "background-muted", t.Light.BackgroundMuted)
//line views/vtheme/Editor.html:45
	qw422016.N().S(`</td>
          <td>`)
//line views/vtheme/Editor.html:46
	streamcinput(qw422016, lp, "foreground-muted", t.Light.ForegroundMuted)
//line views/vtheme/Editor.html:46
	qw422016.N().S(`</td>
          <td class="bl">`)
//line views/vtheme/Editor.html:47
	streamcinput(qw422016, dp, "background-muted", t.Dark.BackgroundMuted)
//line views/vtheme/Editor.html:47
	qw422016.N().S(`</td>
          <td>`)
//line views/vtheme/Editor.html:48
	streamcinput(qw422016, dp, "foreground-muted", t.Dark.ForegroundMuted)
//line views/vtheme/Editor.html:48
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="left-align shrink">Link</th>
          <td class="bl"></td>
          <td>`)
//line views/vtheme/Editor.html:53
	streamcinput(qw422016, lp, "link-foreground", t.Light.LinkForeground)
//line views/vtheme/Editor.html:53
	qw422016.N().S(`</td>
          <td class="bl"></td>
          <td>`)
//line views/vtheme/Editor.html:55
	streamcinput(qw422016, dp, "link-foreground", t.Dark.LinkForeground)
//line views/vtheme/Editor.html:55
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="left-align shrink">Visited Link</th>
          <td class="bl"></td>
          <td>`)
//line views/vtheme/Editor.html:60
	streamcinput(qw422016, lp, "link-visited-foreground", t.Light.LinkVisitedForeground)
//line views/vtheme/Editor.html:60
	qw422016.N().S(`</td>
          <td class="bl"></td>
          <td>`)
//line views/vtheme/Editor.html:62
	streamcinput(qw422016, dp, "link-visited-foreground", t.Dark.LinkVisitedForeground)
//line views/vtheme/Editor.html:62
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="left-align shrink">Navigation</th>
          <td class="bl">`)
//line views/vtheme/Editor.html:66
	streamcinput(qw422016, lp, "nav-background", t.Light.NavBackground)
//line views/vtheme/Editor.html:66
	qw422016.N().S(`</td>
          <td>`)
//line views/vtheme/Editor.html:67
	streamcinput(qw422016, lp, "nav-foreground", t.Light.NavForeground)
//line views/vtheme/Editor.html:67
	qw422016.N().S(`</td>
          <td class="bl">`)
//line views/vtheme/Editor.html:68
	streamcinput(qw422016, dp, "nav-background", t.Dark.NavBackground)
//line views/vtheme/Editor.html:68
	qw422016.N().S(`</td>
          <td>`)
//line views/vtheme/Editor.html:69
	streamcinput(qw422016, dp, "nav-foreground", t.Dark.NavForeground)
//line views/vtheme/Editor.html:69
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="left-align shrink">Menu</th>
          <td class="bl">`)
//line views/vtheme/Editor.html:73
	streamcinput(qw422016, lp, "menu-background", t.Light.MenuBackground)
//line views/vtheme/Editor.html:73
	qw422016.N().S(`</td>
          <td>`)
//line views/vtheme/Editor.html:74
	streamcinput(qw422016, lp, "menu-foreground", t.Light.MenuForeground)
//line views/vtheme/Editor.html:74
	qw422016.N().S(`</td>
          <td class="bl">`)
//line views/vtheme/Editor.html:75
	streamcinput(qw422016, dp, "menu-background", t.Dark.MenuBackground)
//line views/vtheme/Editor.html:75
	qw422016.N().S(`</td>
          <td>`)
//line views/vtheme/Editor.html:76
	streamcinput(qw422016, dp, "menu-foreground", t.Dark.MenuForeground)
//line views/vtheme/Editor.html:76
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="left-align shrink">Selected Menu</th>
          <td class="bl">`)
//line views/vtheme/Editor.html:80
	streamcinput(qw422016, lp, "menu-selected-background", t.Light.MenuSelectedBackground)
//line views/vtheme/Editor.html:80
	qw422016.N().S(`</td>
          <td>`)
//line views/vtheme/Editor.html:81
	streamcinput(qw422016, lp, "menu-selected-foreground", t.Light.MenuSelectedForeground)
//line views/vtheme/Editor.html:81
	qw422016.N().S(`</td>
          <td class="bl">`)
//line views/vtheme/Editor.html:82
	streamcinput(qw422016, dp, "menu-selected-background", t.Dark.MenuSelectedBackground)
//line views/vtheme/Editor.html:82
	qw422016.N().S(`</td>
          <td>`)
//line views/vtheme/Editor.html:83
	streamcinput(qw422016, dp, "menu-selected-foreground", t.Dark.MenuSelectedForeground)
//line views/vtheme/Editor.html:83
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vtheme/Editor.html:88
}

//line views/vtheme/Editor.html:88
func WriteEditor(qq422016 qtio422016.Writer, title string, navTitle string, t *theme.Theme, icon string, as *app.State, ps *cutil.PageState) {
//line views/vtheme/Editor.html:88
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Editor.html:88
	StreamEditor(qw422016, title, navTitle, t, icon, as, ps)
//line views/vtheme/Editor.html:88
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Editor.html:88
}

//line views/vtheme/Editor.html:88
func Editor(title string, navTitle string, t *theme.Theme, icon string, as *app.State, ps *cutil.PageState) string {
//line views/vtheme/Editor.html:88
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Editor.html:88
	WriteEditor(qb422016, title, navTitle, t, icon, as, ps)
//line views/vtheme/Editor.html:88
	qs422016 := string(qb422016.B)
//line views/vtheme/Editor.html:88
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Editor.html:88
	return qs422016
//line views/vtheme/Editor.html:88
}

//line views/vtheme/Editor.html:90
func streamcinput(qw422016 *qt422016.Writer, mode string, k string, v string) {
//line views/vtheme/Editor.html:90
	qw422016.N().S(`<input class="color-var" data-mode="`)
//line views/vtheme/Editor.html:91
	qw422016.E().S(mode)
//line views/vtheme/Editor.html:91
	qw422016.N().S(`" data-var="color-`)
//line views/vtheme/Editor.html:91
	qw422016.E().S(k)
//line views/vtheme/Editor.html:91
	qw422016.N().S(`" type="color" id="`)
//line views/vtheme/Editor.html:91
	qw422016.E().S(mode)
//line views/vtheme/Editor.html:91
	qw422016.N().S(`-`)
//line views/vtheme/Editor.html:91
	qw422016.E().S(k)
//line views/vtheme/Editor.html:91
	qw422016.N().S(`" name="`)
//line views/vtheme/Editor.html:91
	qw422016.E().S(mode)
//line views/vtheme/Editor.html:91
	qw422016.N().S(`-`)
//line views/vtheme/Editor.html:91
	qw422016.E().S(k)
//line views/vtheme/Editor.html:91
	qw422016.N().S(`" value="`)
//line views/vtheme/Editor.html:91
	qw422016.E().S(v)
//line views/vtheme/Editor.html:91
	qw422016.N().S(`" autocomplete="off" />`)
//line views/vtheme/Editor.html:92
}

//line views/vtheme/Editor.html:92
func writecinput(qq422016 qtio422016.Writer, mode string, k string, v string) {
//line views/vtheme/Editor.html:92
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Editor.html:92
	streamcinput(qw422016, mode, k, v)
//line views/vtheme/Editor.html:92
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Editor.html:92
}

//line views/vtheme/Editor.html:92
func cinput(mode string, k string, v string) string {
//line views/vtheme/Editor.html:92
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Editor.html:92
	writecinput(qb422016, mode, k, v)
//line views/vtheme/Editor.html:92
	qs422016 := string(qb422016.B)
//line views/vtheme/Editor.html:92
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Editor.html:92
	return qs422016
//line views/vtheme/Editor.html:92
}
