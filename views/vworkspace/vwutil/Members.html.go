// Code generated by qtc from "Members.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwutil/Members.html:1
package vwutil

//line views/vworkspace/vwutil/Members.html:1
import (
	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/member"
	"github.com/kyleu/rituals/views/components"
	"github.com/kyleu/rituals/views/components/edit"
)

//line views/vworkspace/vwutil/Members.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwutil/Members.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwutil/Members.html:9
func StreamMemberPanels(qw422016 *qt422016.Writer, ms member.Members, admin bool, path string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:9
	qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:10
	self, others, _ := ms.Split(ps.Profile.ID)

//line views/vworkspace/vwutil/Members.html:10
	qw422016.N().S(`  <div id="panel-self">
    <div class="card">`)
//line views/vworkspace/vwutil/Members.html:12
	StreamSelfLink(qw422016, self, ps)
//line views/vworkspace/vwutil/Members.html:12
	qw422016.N().S(`</div>
    `)
//line views/vworkspace/vwutil/Members.html:13
	StreamSelfModal(qw422016, self.UserID, self.Name, self.Picture, self.Role, path, ps)
//line views/vworkspace/vwutil/Members.html:13
	qw422016.N().S(`
  </div>
  <div id="panel-members">
    <div class="card">
      <a href="#modal-invite"><h3>`)
//line views/vworkspace/vwutil/Members.html:17
	components.StreamSVGIcon(qw422016, `users`, ps)
//line views/vworkspace/vwutil/Members.html:17
	qw422016.N().S(`Members</h3></a>
      <table class="mt expanded">
        <tbody>
`)
//line views/vworkspace/vwutil/Members.html:20
	for _, m := range others {
//line views/vworkspace/vwutil/Members.html:20
		qw422016.N().S(`          `)
//line views/vworkspace/vwutil/Members.html:21
		StreamMemberRow(qw422016, m, ps)
//line views/vworkspace/vwutil/Members.html:21
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:22
	}
//line views/vworkspace/vwutil/Members.html:22
	qw422016.N().S(`        </tbody>
      </table>
      <div id="member-modals">
`)
//line views/vworkspace/vwutil/Members.html:26
	for _, m := range others {
//line views/vworkspace/vwutil/Members.html:27
		if admin {
//line views/vworkspace/vwutil/Members.html:27
			qw422016.N().S(`        `)
//line views/vworkspace/vwutil/Members.html:28
			StreamMemberModalEdit(qw422016, m, path, ps)
//line views/vworkspace/vwutil/Members.html:28
			qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:29
		} else {
//line views/vworkspace/vwutil/Members.html:29
			qw422016.N().S(`        `)
//line views/vworkspace/vwutil/Members.html:30
			StreamMemberModalView(qw422016, m, path, ps)
//line views/vworkspace/vwutil/Members.html:30
			qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:31
		}
//line views/vworkspace/vwutil/Members.html:32
	}
//line views/vworkspace/vwutil/Members.html:32
	qw422016.N().S(`      </div>
    </div>
    `)
//line views/vworkspace/vwutil/Members.html:35
	StreamInviteModal(qw422016)
//line views/vworkspace/vwutil/Members.html:35
	qw422016.N().S(`
  </div>
`)
//line views/vworkspace/vwutil/Members.html:37
}

//line views/vworkspace/vwutil/Members.html:37
func WriteMemberPanels(qq422016 qtio422016.Writer, ms member.Members, admin bool, path string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Members.html:37
	StreamMemberPanels(qw422016, ms, admin, path, ps)
//line views/vworkspace/vwutil/Members.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Members.html:37
}

//line views/vworkspace/vwutil/Members.html:37
func MemberPanels(ms member.Members, admin bool, path string, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Members.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Members.html:37
	WriteMemberPanels(qb422016, ms, admin, path, ps)
//line views/vworkspace/vwutil/Members.html:37
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Members.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Members.html:37
	return qs422016
//line views/vworkspace/vwutil/Members.html:37
}

//line views/vworkspace/vwutil/Members.html:39
func StreamMemberRow(qw422016 *qt422016.Writer, m *member.Member, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:39
	qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:40
	ps.AddIcon("circle", "check-circle")

//line views/vworkspace/vwutil/Members.html:40
	qw422016.N().S(`  <tr id="member-`)
//line views/vworkspace/vwutil/Members.html:41
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:41
	qw422016.N().S(`" class="member" data-id="`)
//line views/vworkspace/vwutil/Members.html:41
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:41
	qw422016.N().S(`">
    <td>
      <a href="#modal-member-`)
//line views/vworkspace/vwutil/Members.html:43
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:43
	qw422016.N().S(`">
        <div class="left prs member-picture">
`)
//line views/vworkspace/vwutil/Members.html:45
	if m.Picture == "" {
//line views/vworkspace/vwutil/Members.html:45
		qw422016.N().S(`          `)
//line views/vworkspace/vwutil/Members.html:46
		components.StreamSVGRef(qw422016, `profile`, 18, 18, ``, ps)
//line views/vworkspace/vwutil/Members.html:46
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:47
	} else {
//line views/vworkspace/vwutil/Members.html:47
		qw422016.N().S(`          <img style="width: 18px; height: 18px;" src="`)
//line views/vworkspace/vwutil/Members.html:48
		qw422016.E().S(m.Picture)
//line views/vworkspace/vwutil/Members.html:48
		qw422016.N().S(`" />
`)
//line views/vworkspace/vwutil/Members.html:49
	}
//line views/vworkspace/vwutil/Members.html:49
	qw422016.N().S(`        </div>
        <span class="member-name member-`)
//line views/vworkspace/vwutil/Members.html:51
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:51
	qw422016.N().S(`-name">`)
//line views/vworkspace/vwutil/Members.html:51
	qw422016.E().S(m.Name)
//line views/vworkspace/vwutil/Members.html:51
	qw422016.N().S(`</span>
      </a>
    </td>
    <td class="shrink text-align-right"><em class="member-role">`)
//line views/vworkspace/vwutil/Members.html:54
	qw422016.E().S(m.Role.String())
//line views/vworkspace/vwutil/Members.html:54
	qw422016.N().S(`</em></td>
    <td class="shrink online-status" title="`)
//line views/vworkspace/vwutil/Members.html:55
	if m.Online {
//line views/vworkspace/vwutil/Members.html:55
		qw422016.N().S(`online`)
//line views/vworkspace/vwutil/Members.html:55
	} else {
//line views/vworkspace/vwutil/Members.html:55
		qw422016.N().S(`offline`)
//line views/vworkspace/vwutil/Members.html:55
	}
//line views/vworkspace/vwutil/Members.html:55
	qw422016.N().S(`">
`)
//line views/vworkspace/vwutil/Members.html:56
	if m.Online {
//line views/vworkspace/vwutil/Members.html:56
		qw422016.N().S(`      `)
//line views/vworkspace/vwutil/Members.html:57
		components.StreamSVGRef(qw422016, `check-circle`, 18, 18, `right`, ps)
//line views/vworkspace/vwutil/Members.html:57
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:58
	} else {
//line views/vworkspace/vwutil/Members.html:58
		qw422016.N().S(`      `)
//line views/vworkspace/vwutil/Members.html:59
		components.StreamSVGRef(qw422016, `circle`, 18, 18, `right`, ps)
//line views/vworkspace/vwutil/Members.html:59
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:60
	}
//line views/vworkspace/vwutil/Members.html:60
	qw422016.N().S(`    </td>
  </tr>
`)
//line views/vworkspace/vwutil/Members.html:63
}

//line views/vworkspace/vwutil/Members.html:63
func WriteMemberRow(qq422016 qtio422016.Writer, m *member.Member, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Members.html:63
	StreamMemberRow(qw422016, m, ps)
//line views/vworkspace/vwutil/Members.html:63
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Members.html:63
}

//line views/vworkspace/vwutil/Members.html:63
func MemberRow(m *member.Member, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Members.html:63
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Members.html:63
	WriteMemberRow(qb422016, m, ps)
//line views/vworkspace/vwutil/Members.html:63
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Members.html:63
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Members.html:63
	return qs422016
//line views/vworkspace/vwutil/Members.html:63
}

//line views/vworkspace/vwutil/Members.html:65
func StreamInviteModal(qw422016 *qt422016.Writer) {
//line views/vworkspace/vwutil/Members.html:65
	qw422016.N().S(`
  <div id="modal-invite" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Invite Member</h2>
      </div>
      <div class="modal-body">
        For now, just share the url of this page from your address bar
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwutil/Members.html:78
}

//line views/vworkspace/vwutil/Members.html:78
func WriteInviteModal(qq422016 qtio422016.Writer) {
//line views/vworkspace/vwutil/Members.html:78
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Members.html:78
	StreamInviteModal(qw422016)
//line views/vworkspace/vwutil/Members.html:78
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Members.html:78
}

//line views/vworkspace/vwutil/Members.html:78
func InviteModal() string {
//line views/vworkspace/vwutil/Members.html:78
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Members.html:78
	WriteInviteModal(qb422016)
//line views/vworkspace/vwutil/Members.html:78
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Members.html:78
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Members.html:78
	return qs422016
//line views/vworkspace/vwutil/Members.html:78
}

//line views/vworkspace/vwutil/Members.html:80
func StreamMemberModalEdit(qw422016 *qt422016.Writer, m *member.Member, url string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:80
	qw422016.N().S(`
  <div id="modal-member-`)
//line views/vworkspace/vwutil/Members.html:81
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:81
	qw422016.N().S(`" data-id="`)
//line views/vworkspace/vwutil/Members.html:81
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:81
	qw422016.N().S(`" class="modal modal-member" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>
          <span class="member-picture">
`)
//line views/vworkspace/vwutil/Members.html:88
	if m.Picture == "" {
//line views/vworkspace/vwutil/Members.html:88
		qw422016.N().S(`            `)
//line views/vworkspace/vwutil/Members.html:89
		components.StreamSVGRef(qw422016, `profile`, 24, 24, `icon`, ps)
//line views/vworkspace/vwutil/Members.html:89
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:90
	} else {
//line views/vworkspace/vwutil/Members.html:90
		qw422016.N().S(`            <img class="icon" style="width: 24px; height: 24px;" src="`)
//line views/vworkspace/vwutil/Members.html:91
		qw422016.E().S(m.Picture)
//line views/vworkspace/vwutil/Members.html:91
		qw422016.N().S(`" />
`)
//line views/vworkspace/vwutil/Members.html:92
	}
//line views/vworkspace/vwutil/Members.html:92
	qw422016.N().S(`          </span>
          <span class="member-name member-`)
//line views/vworkspace/vwutil/Members.html:94
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:94
	qw422016.N().S(`-name">`)
//line views/vworkspace/vwutil/Members.html:94
	qw422016.E().S(m.Name)
//line views/vworkspace/vwutil/Members.html:94
	qw422016.N().S(`</span>
        </h2>
      </div>
      <div class="modal-body">
        <form action="`)
//line views/vworkspace/vwutil/Members.html:98
	qw422016.E().S(url)
//line views/vworkspace/vwutil/Members.html:98
	qw422016.N().S(`" method="post" class="expanded">
          <input type="hidden" name="userID" value="`)
//line views/vworkspace/vwutil/Members.html:99
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:99
	qw422016.N().S(`" />
          <em>Role</em><br />
          `)
//line views/vworkspace/vwutil/Members.html:101
	edit.StreamSelect(qw422016, "role", "", m.Role.Key, []string{"owner", "member", "observer"}, []string{"Owner", "Member", "Observer"}, 5)
//line views/vworkspace/vwutil/Members.html:101
	qw422016.N().S(`
          <hr />
          <div class="right"><button class="member-update" type="submit" name="action" value="`)
//line views/vworkspace/vwutil/Members.html:103
	qw422016.E().S(string(action.ActMemberUpdate))
//line views/vworkspace/vwutil/Members.html:103
	qw422016.N().S(`">Save</button></div>
          <button type="submit" class="member-remove" name="action" value="`)
//line views/vworkspace/vwutil/Members.html:104
	qw422016.E().S(string(action.ActMemberRemove))
//line views/vworkspace/vwutil/Members.html:104
	qw422016.N().S(`" onclick="return confirm('Are you sure you wish to remove this user?');">Remove</button>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwutil/Members.html:109
}

//line views/vworkspace/vwutil/Members.html:109
func WriteMemberModalEdit(qq422016 qtio422016.Writer, m *member.Member, url string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:109
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Members.html:109
	StreamMemberModalEdit(qw422016, m, url, ps)
//line views/vworkspace/vwutil/Members.html:109
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Members.html:109
}

//line views/vworkspace/vwutil/Members.html:109
func MemberModalEdit(m *member.Member, url string, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Members.html:109
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Members.html:109
	WriteMemberModalEdit(qb422016, m, url, ps)
//line views/vworkspace/vwutil/Members.html:109
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Members.html:109
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Members.html:109
	return qs422016
//line views/vworkspace/vwutil/Members.html:109
}

//line views/vworkspace/vwutil/Members.html:111
func StreamMemberModalView(qw422016 *qt422016.Writer, m *member.Member, url string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:111
	qw422016.N().S(`
  <div id="modal-member-`)
//line views/vworkspace/vwutil/Members.html:112
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:112
	qw422016.N().S(`" data-id="`)
//line views/vworkspace/vwutil/Members.html:112
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:112
	qw422016.N().S(`" class="modal modal-member" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>
          <span class="member-picture">
`)
//line views/vworkspace/vwutil/Members.html:119
	if m.Picture == "" {
//line views/vworkspace/vwutil/Members.html:119
		qw422016.N().S(`            `)
//line views/vworkspace/vwutil/Members.html:120
		components.StreamSVGRef(qw422016, `profile`, 24, 24, `icon`, ps)
//line views/vworkspace/vwutil/Members.html:120
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:121
	} else {
//line views/vworkspace/vwutil/Members.html:121
		qw422016.N().S(`            <img class="icon" style="width: 24px; height: 24px;" src="`)
//line views/vworkspace/vwutil/Members.html:122
		qw422016.E().S(m.Picture)
//line views/vworkspace/vwutil/Members.html:122
		qw422016.N().S(`" />
`)
//line views/vworkspace/vwutil/Members.html:123
	}
//line views/vworkspace/vwutil/Members.html:123
	qw422016.N().S(`          </span>
          <span class="member-name member-`)
//line views/vworkspace/vwutil/Members.html:125
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:125
	qw422016.N().S(`-name">`)
//line views/vworkspace/vwutil/Members.html:125
	qw422016.E().S(m.Name)
//line views/vworkspace/vwutil/Members.html:125
	qw422016.N().S(`</span>
        </h2>
      </div>
      <div class="modal-body">
        <em>Role</em><br />
        <span class="member-role">`)
//line views/vworkspace/vwutil/Members.html:130
	qw422016.E().S(m.Role.String())
//line views/vworkspace/vwutil/Members.html:130
	qw422016.N().S(`</span>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwutil/Members.html:134
}

//line views/vworkspace/vwutil/Members.html:134
func WriteMemberModalView(qq422016 qtio422016.Writer, m *member.Member, url string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:134
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Members.html:134
	StreamMemberModalView(qw422016, m, url, ps)
//line views/vworkspace/vwutil/Members.html:134
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Members.html:134
}

//line views/vworkspace/vwutil/Members.html:134
func MemberModalView(m *member.Member, url string, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Members.html:134
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Members.html:134
	WriteMemberModalView(qb422016, m, url, ps)
//line views/vworkspace/vwutil/Members.html:134
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Members.html:134
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Members.html:134
	return qs422016
//line views/vworkspace/vwutil/Members.html:134
}
