// Code generated by qtc from "TemplateUtils.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/TemplateUtils.html:1
package vworkspace

//line views/vworkspace/TemplateUtils.html:1
import (
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/views/components"
)

//line views/vworkspace/TemplateUtils.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/TemplateUtils.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/TemplateUtils.html:8
func StreamBanner(qw422016 *qt422016.Writer, t *team.Team, s *sprint.Sprint, mdl string) {
//line views/vworkspace/TemplateUtils.html:8
	qw422016.N().S(`<em>`)
//line views/vworkspace/TemplateUtils.html:10
	if s != nil {
//line views/vworkspace/TemplateUtils.html:10
		qw422016.N().S(`<a href="/sprint/`)
//line views/vworkspace/TemplateUtils.html:10
		qw422016.E().S(s.Slug)
//line views/vworkspace/TemplateUtils.html:10
		qw422016.N().S(`">`)
//line views/vworkspace/TemplateUtils.html:10
		qw422016.E().S(s.TitleString())
//line views/vworkspace/TemplateUtils.html:10
		qw422016.N().S(`</a>`)
//line views/vworkspace/TemplateUtils.html:10
		qw422016.N().S(` `)
//line views/vworkspace/TemplateUtils.html:10
	}
//line views/vworkspace/TemplateUtils.html:11
	qw422016.E().S(mdl)
//line views/vworkspace/TemplateUtils.html:12
	if t != nil {
//line views/vworkspace/TemplateUtils.html:12
		qw422016.N().S(` `)
//line views/vworkspace/TemplateUtils.html:12
		qw422016.N().S(`in <a href="/team/`)
//line views/vworkspace/TemplateUtils.html:12
		qw422016.E().S(t.Slug)
//line views/vworkspace/TemplateUtils.html:12
		qw422016.N().S(`">`)
//line views/vworkspace/TemplateUtils.html:12
		qw422016.E().S(t.TitleString())
//line views/vworkspace/TemplateUtils.html:12
		qw422016.N().S(`</a>`)
//line views/vworkspace/TemplateUtils.html:12
	}
//line views/vworkspace/TemplateUtils.html:12
	qw422016.N().S(`</em>`)
//line views/vworkspace/TemplateUtils.html:14
}

//line views/vworkspace/TemplateUtils.html:14
func WriteBanner(qq422016 qtio422016.Writer, t *team.Team, s *sprint.Sprint, mdl string) {
//line views/vworkspace/TemplateUtils.html:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/TemplateUtils.html:14
	StreamBanner(qw422016, t, s, mdl)
//line views/vworkspace/TemplateUtils.html:14
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/TemplateUtils.html:14
}

//line views/vworkspace/TemplateUtils.html:14
func Banner(t *team.Team, s *sprint.Sprint, mdl string) string {
//line views/vworkspace/TemplateUtils.html:14
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/TemplateUtils.html:14
	WriteBanner(qb422016, t, s, mdl)
//line views/vworkspace/TemplateUtils.html:14
	qs422016 := string(qb422016.B)
//line views/vworkspace/TemplateUtils.html:14
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/TemplateUtils.html:14
	return qs422016
//line views/vworkspace/TemplateUtils.html:14
}

//line views/vworkspace/TemplateUtils.html:16
func StreamSelfModal(qw422016 *qt422016.Writer, name string, picture string, role enum.MemberStatus, action string) {
//line views/vworkspace/TemplateUtils.html:16
	qw422016.N().S(`
  <div id="modal-self" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Edit Self</h2>
      </div>
      <div class="modal-body">
        <form action="`)
//line views/vworkspace/TemplateUtils.html:25
	qw422016.E().S(action)
//line views/vworkspace/TemplateUtils.html:25
	qw422016.N().S(`" method="post" class="expanded">
          <input type="hidden" name="action" value="name" />
          <em>Name</em><br />
          `)
//line views/vworkspace/TemplateUtils.html:28
	components.StreamFormInput(qw422016, "title", "input-title", name, "The name you wish to be called")
//line views/vworkspace/TemplateUtils.html:28
	qw422016.N().S(`
          <div><label><input type="radio" name="choice" value="local" checked="checked"> Change for this session only</label></div>
          <div><label><input type="radio" name="choice" value="global"> Change global default</label></div>
          <hr />
          <em>Picture</em>
          <div>To set a profile picture, <a href="/profile">sign in</a></div>
          <hr />
          <div class="right"><button type="submit">Save</button></div>
          <a href="#"><button type="button">Leave</button></a>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/TemplateUtils.html:41
}

//line views/vworkspace/TemplateUtils.html:41
func WriteSelfModal(qq422016 qtio422016.Writer, name string, picture string, role enum.MemberStatus, action string) {
//line views/vworkspace/TemplateUtils.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/TemplateUtils.html:41
	StreamSelfModal(qw422016, name, picture, role, action)
//line views/vworkspace/TemplateUtils.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/TemplateUtils.html:41
}

//line views/vworkspace/TemplateUtils.html:41
func SelfModal(name string, picture string, role enum.MemberStatus, action string) string {
//line views/vworkspace/TemplateUtils.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/TemplateUtils.html:41
	WriteSelfModal(qb422016, name, picture, role, action)
//line views/vworkspace/TemplateUtils.html:41
	qs422016 := string(qb422016.B)
//line views/vworkspace/TemplateUtils.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/TemplateUtils.html:41
	return qs422016
//line views/vworkspace/TemplateUtils.html:41
}