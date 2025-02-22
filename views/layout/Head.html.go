// Code generated by qtc from "Head.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/layout/Head.html:1
package layout

//line views/layout/Head.html:1
import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/assets"
)

//line views/layout/Head.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Head.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Head.html:8
func StreamHead(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Head.html:8
	qw422016.N().S(`
`)
//line views/layout/Head.html:9
	thm := as.Themes.Get(ps.Profile.Theme, ps.Logger)

//line views/layout/Head.html:9
	qw422016.N().S(`  <meta charset="UTF-8">
  <title>`)
//line views/layout/Head.html:11
	qw422016.E().S(ps.TitleString())
//line views/layout/Head.html:11
	qw422016.N().S(`</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0, viewport-fit=cover">
  `)
//line views/layout/Head.html:13
	if ps.Description != "" {
//line views/layout/Head.html:13
		qw422016.N().S(`<meta property="description" content="`)
//line views/layout/Head.html:13
		qw422016.E().S(ps.Description)
//line views/layout/Head.html:13
		qw422016.N().S(`">
  `)
//line views/layout/Head.html:14
	}
//line views/layout/Head.html:14
	qw422016.N().S(`<meta property="og:title" content="`)
//line views/layout/Head.html:14
	qw422016.E().S(ps.TitleString())
//line views/layout/Head.html:14
	qw422016.N().S(`">
  <meta property="og:type" content="website">
  <meta property="og:image" content="/assets/`)
//line views/layout/Head.html:16
	qw422016.N().U(util.AppKey)
//line views/layout/Head.html:16
	qw422016.N().S(`.svg">
  <meta property="og:locale" content="en_US">
  <meta name="theme-color" content="`)
//line views/layout/Head.html:18
	qw422016.E().S(thm.Light.NavBackground)
//line views/layout/Head.html:18
	qw422016.N().S(`" media="(prefers-color-scheme: light)">
  <meta name="theme-color" content="`)
//line views/layout/Head.html:19
	qw422016.E().S(thm.Dark.NavBackground)
//line views/layout/Head.html:19
	qw422016.N().S(`" media="(prefers-color-scheme: dark)">`)
//line views/layout/Head.html:19
	qw422016.N().S(ps.HeaderContent)
//line views/layout/Head.html:19
	qw422016.N().S(`
  <link rel="icon" href="`)
//line views/layout/Head.html:20
	qw422016.E().S(assets.URL(`logo.svg`))
//line views/layout/Head.html:20
	qw422016.N().S(`" type="image/svg+xml">
  <style>
    `)
//line views/layout/Head.html:22
	qw422016.N().S(thm.CSS(2))
//line views/layout/Head.html:22
	qw422016.N().S(`  </style>`)
//line views/layout/Head.html:22
	if ps.HideHeader && ps.HideMenu {
//line views/layout/Head.html:22
		streaminlineResources(qw422016)
//line views/layout/Head.html:22
	} else {
//line views/layout/Head.html:22
		qw422016.N().S(`  `)
//line views/layout/Head.html:23
		qw422016.N().S(assets.StylesheetElement(`client.css`))
//line views/layout/Head.html:23
		if !ps.NoScript {
//line views/layout/Head.html:23
			qw422016.N().S(`
  `)
//line views/layout/Head.html:24
			qw422016.N().S(assets.ScriptElement(`client.js`, false))
//line views/layout/Head.html:24
		}
//line views/layout/Head.html:24
	}
//line views/layout/Head.html:25
}

//line views/layout/Head.html:25
func WriteHead(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/layout/Head.html:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Head.html:25
	StreamHead(qw422016, as, ps)
//line views/layout/Head.html:25
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Head.html:25
}

//line views/layout/Head.html:25
func Head(as *app.State, ps *cutil.PageState) string {
//line views/layout/Head.html:25
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Head.html:25
	WriteHead(qb422016, as, ps)
//line views/layout/Head.html:25
	qs422016 := string(qb422016.B)
//line views/layout/Head.html:25
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Head.html:25
	return qs422016
//line views/layout/Head.html:25
}

//line views/layout/Head.html:27
func streaminlineResources(qw422016 *qt422016.Writer) {
//line views/layout/Head.html:29
	csv, err := assets.Embed("client.css")
	if err != nil {
		panic(err)
	}
	js, err := assets.Embed("client.js")
	if err != nil {
		panic(err)
	}

//line views/layout/Head.html:37
	qw422016.N().S(`<style>`)
//line views/layout/Head.html:38
	qw422016.N().S(string(csv.Bytes))
//line views/layout/Head.html:38
	qw422016.N().S(`</style><script>`)
//line views/layout/Head.html:39
	qw422016.N().S(string(js.Bytes))
//line views/layout/Head.html:39
	qw422016.N().S(`</script>`)
//line views/layout/Head.html:40
}

//line views/layout/Head.html:40
func writeinlineResources(qq422016 qtio422016.Writer) {
//line views/layout/Head.html:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Head.html:40
	streaminlineResources(qw422016)
//line views/layout/Head.html:40
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Head.html:40
}

//line views/layout/Head.html:40
func inlineResources() string {
//line views/layout/Head.html:40
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Head.html:40
	writeinlineResources(qb422016)
//line views/layout/Head.html:40
	qs422016 := string(qb422016.B)
//line views/layout/Head.html:40
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Head.html:40
	return qs422016
//line views/layout/Head.html:40
}
