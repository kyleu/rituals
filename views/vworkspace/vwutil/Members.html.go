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
)

//line views/vworkspace/vwutil/Members.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwutil/Members.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwutil/Members.html:8
func StreamMemberPanels(qw422016 *qt422016.Writer, ms member.Members, admin bool, path string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:8
	qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:9
	self, others, _ := ms.Split(ps.Profile.ID)

//line views/vworkspace/vwutil/Members.html:9
	qw422016.N().S(`  <div id="panel-self">
    <div class="card">`)
//line views/vworkspace/vwutil/Members.html:11
	StreamSelfLink(qw422016, self, ps)
//line views/vworkspace/vwutil/Members.html:11
	qw422016.N().S(`</div>
    `)
//line views/vworkspace/vwutil/Members.html:12
	StreamSelfModal(qw422016, self.UserID, self.Name, self.Picture, self.Role, path, ps)
//line views/vworkspace/vwutil/Members.html:12
	qw422016.N().S(`
  </div>
  <div id="panel-members">
    <div class="card">
      <a href="#modal-invite"><h3>`)
//line views/vworkspace/vwutil/Members.html:16
	components.StreamSVGRefIcon(qw422016, `users`, ps)
//line views/vworkspace/vwutil/Members.html:16
	qw422016.N().S(`Members</h3></a>
      <table class="mt expanded">
        <tbody>
`)
//line views/vworkspace/vwutil/Members.html:19
	for _, m := range others {
//line views/vworkspace/vwutil/Members.html:19
		qw422016.N().S(`          `)
//line views/vworkspace/vwutil/Members.html:20
		StreamMemberRow(qw422016, m, ps)
//line views/vworkspace/vwutil/Members.html:20
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:21
	}
//line views/vworkspace/vwutil/Members.html:21
	qw422016.N().S(`        </tbody>
      </table>
      <div id="member-modals">
`)
//line views/vworkspace/vwutil/Members.html:25
	for _, m := range others {
//line views/vworkspace/vwutil/Members.html:26
		if admin {
//line views/vworkspace/vwutil/Members.html:26
			qw422016.N().S(`        `)
//line views/vworkspace/vwutil/Members.html:27
			StreamMemberModalEdit(qw422016, m, path, ps)
//line views/vworkspace/vwutil/Members.html:27
			qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:28
		} else {
//line views/vworkspace/vwutil/Members.html:28
			qw422016.N().S(`        `)
//line views/vworkspace/vwutil/Members.html:29
			StreamMemberModalView(qw422016, m, path, ps)
//line views/vworkspace/vwutil/Members.html:29
			qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:30
		}
//line views/vworkspace/vwutil/Members.html:31
	}
//line views/vworkspace/vwutil/Members.html:31
	qw422016.N().S(`      </div>
    </div>
    `)
//line views/vworkspace/vwutil/Members.html:34
	StreamInviteModal(qw422016)
//line views/vworkspace/vwutil/Members.html:34
	qw422016.N().S(`
  </div>
`)
//line views/vworkspace/vwutil/Members.html:36
}

//line views/vworkspace/vwutil/Members.html:36
func WriteMemberPanels(qq422016 qtio422016.Writer, ms member.Members, admin bool, path string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Members.html:36
	StreamMemberPanels(qw422016, ms, admin, path, ps)
//line views/vworkspace/vwutil/Members.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Members.html:36
}

//line views/vworkspace/vwutil/Members.html:36
func MemberPanels(ms member.Members, admin bool, path string, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Members.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Members.html:36
	WriteMemberPanels(qb422016, ms, admin, path, ps)
//line views/vworkspace/vwutil/Members.html:36
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Members.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Members.html:36
	return qs422016
//line views/vworkspace/vwutil/Members.html:36
}

//line views/vworkspace/vwutil/Members.html:38
func StreamMemberRow(qw422016 *qt422016.Writer, m *member.Member, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:38
	qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:39
	ps.AddIcon("circle", "check-circle")

//line views/vworkspace/vwutil/Members.html:39
	qw422016.N().S(`  <tr id="member-`)
//line views/vworkspace/vwutil/Members.html:40
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:40
	qw422016.N().S(`" class="member" data-id="`)
//line views/vworkspace/vwutil/Members.html:40
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:40
	qw422016.N().S(`">
    <td>
      <a href="#modal-member-`)
//line views/vworkspace/vwutil/Members.html:42
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:42
	qw422016.N().S(`">
        <div class="left prs member-picture">
`)
//line views/vworkspace/vwutil/Members.html:44
	if m.Picture == "" {
//line views/vworkspace/vwutil/Members.html:44
		qw422016.N().S(`          `)
//line views/vworkspace/vwutil/Members.html:45
		components.StreamSVGRef(qw422016, `profile`, 18, 18, ``, ps)
//line views/vworkspace/vwutil/Members.html:45
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:46
	} else {
//line views/vworkspace/vwutil/Members.html:46
		qw422016.N().S(`          <img style="width: 18px; height: 18px;" src="`)
//line views/vworkspace/vwutil/Members.html:47
		qw422016.E().S(m.Picture)
//line views/vworkspace/vwutil/Members.html:47
		qw422016.N().S(`" />
`)
//line views/vworkspace/vwutil/Members.html:48
	}
//line views/vworkspace/vwutil/Members.html:48
	qw422016.N().S(`        </div>
        <span class="member-name member-`)
//line views/vworkspace/vwutil/Members.html:50
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:50
	qw422016.N().S(`-name">`)
//line views/vworkspace/vwutil/Members.html:50
	qw422016.E().S(m.Name)
//line views/vworkspace/vwutil/Members.html:50
	qw422016.N().S(`</span>
      </a>
    </td>
    <td class="shrink text-align-right"><em class="member-role">`)
//line views/vworkspace/vwutil/Members.html:53
	qw422016.E().S(m.Role.String())
//line views/vworkspace/vwutil/Members.html:53
	qw422016.N().S(`</em></td>
    <td class="shrink online-status" title="`)
//line views/vworkspace/vwutil/Members.html:54
	if m.Online {
//line views/vworkspace/vwutil/Members.html:54
		qw422016.N().S(`online`)
//line views/vworkspace/vwutil/Members.html:54
	} else {
//line views/vworkspace/vwutil/Members.html:54
		qw422016.N().S(`offline`)
//line views/vworkspace/vwutil/Members.html:54
	}
//line views/vworkspace/vwutil/Members.html:54
	qw422016.N().S(`">
`)
//line views/vworkspace/vwutil/Members.html:55
	if m.Online {
//line views/vworkspace/vwutil/Members.html:55
		qw422016.N().S(`      `)
//line views/vworkspace/vwutil/Members.html:56
		components.StreamSVGRef(qw422016, `check-circle`, 18, 18, `right`, ps)
//line views/vworkspace/vwutil/Members.html:56
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:57
	} else {
//line views/vworkspace/vwutil/Members.html:57
		qw422016.N().S(`      `)
//line views/vworkspace/vwutil/Members.html:58
		components.StreamSVGRef(qw422016, `circle`, 18, 18, `right`, ps)
//line views/vworkspace/vwutil/Members.html:58
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:59
	}
//line views/vworkspace/vwutil/Members.html:59
	qw422016.N().S(`    </td>
  </tr>
`)
//line views/vworkspace/vwutil/Members.html:62
}

//line views/vworkspace/vwutil/Members.html:62
func WriteMemberRow(qq422016 qtio422016.Writer, m *member.Member, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:62
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Members.html:62
	StreamMemberRow(qw422016, m, ps)
//line views/vworkspace/vwutil/Members.html:62
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Members.html:62
}

//line views/vworkspace/vwutil/Members.html:62
func MemberRow(m *member.Member, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Members.html:62
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Members.html:62
	WriteMemberRow(qb422016, m, ps)
//line views/vworkspace/vwutil/Members.html:62
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Members.html:62
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Members.html:62
	return qs422016
//line views/vworkspace/vwutil/Members.html:62
}

//line views/vworkspace/vwutil/Members.html:64
func StreamInviteModal(qw422016 *qt422016.Writer) {
//line views/vworkspace/vwutil/Members.html:64
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
//line views/vworkspace/vwutil/Members.html:77
}

//line views/vworkspace/vwutil/Members.html:77
func WriteInviteModal(qq422016 qtio422016.Writer) {
//line views/vworkspace/vwutil/Members.html:77
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Members.html:77
	StreamInviteModal(qw422016)
//line views/vworkspace/vwutil/Members.html:77
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Members.html:77
}

//line views/vworkspace/vwutil/Members.html:77
func InviteModal() string {
//line views/vworkspace/vwutil/Members.html:77
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Members.html:77
	WriteInviteModal(qb422016)
//line views/vworkspace/vwutil/Members.html:77
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Members.html:77
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Members.html:77
	return qs422016
//line views/vworkspace/vwutil/Members.html:77
}

//line views/vworkspace/vwutil/Members.html:79
func StreamMemberModalEdit(qw422016 *qt422016.Writer, m *member.Member, url string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:79
	qw422016.N().S(`
  <div id="modal-member-`)
//line views/vworkspace/vwutil/Members.html:80
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:80
	qw422016.N().S(`" data-id="`)
//line views/vworkspace/vwutil/Members.html:80
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:80
	qw422016.N().S(`" class="modal modal-member" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>
          <span class="member-picture">
`)
//line views/vworkspace/vwutil/Members.html:87
	if m.Picture == "" {
//line views/vworkspace/vwutil/Members.html:87
		qw422016.N().S(`            `)
//line views/vworkspace/vwutil/Members.html:88
		components.StreamSVGRef(qw422016, `profile`, 24, 24, `icon`, ps)
//line views/vworkspace/vwutil/Members.html:88
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:89
	} else {
//line views/vworkspace/vwutil/Members.html:89
		qw422016.N().S(`            <img class="icon" style="width: 24px; height: 24px;" src="`)
//line views/vworkspace/vwutil/Members.html:90
		qw422016.E().S(m.Picture)
//line views/vworkspace/vwutil/Members.html:90
		qw422016.N().S(`" />
`)
//line views/vworkspace/vwutil/Members.html:91
	}
//line views/vworkspace/vwutil/Members.html:91
	qw422016.N().S(`          </span>
          <span class="member-name member-`)
//line views/vworkspace/vwutil/Members.html:93
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:93
	qw422016.N().S(`-name">`)
//line views/vworkspace/vwutil/Members.html:93
	qw422016.E().S(m.Name)
//line views/vworkspace/vwutil/Members.html:93
	qw422016.N().S(`</span>
        </h2>
      </div>
      <div class="modal-body">
        <form action="`)
//line views/vworkspace/vwutil/Members.html:97
	qw422016.E().S(url)
//line views/vworkspace/vwutil/Members.html:97
	qw422016.N().S(`" method="post" class="expanded">
          <input type="hidden" name="userID" value="`)
//line views/vworkspace/vwutil/Members.html:98
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:98
	qw422016.N().S(`" />
          <em>Role</em><br />
          `)
//line views/vworkspace/vwutil/Members.html:100
	components.StreamFormSelect(qw422016, "role", "", m.Role.Key, []string{"owner", "member", "observer"}, []string{"Owner", "Member", "Observer"}, 5)
//line views/vworkspace/vwutil/Members.html:100
	qw422016.N().S(`
          <hr />
          <div class="right"><button class="member-update" type="submit" name="action" value="`)
//line views/vworkspace/vwutil/Members.html:102
	qw422016.E().S(string(action.ActMemberUpdate))
//line views/vworkspace/vwutil/Members.html:102
	qw422016.N().S(`">Save</button></div>
          <button type="submit" class="member-remove" name="action" value="`)
//line views/vworkspace/vwutil/Members.html:103
	qw422016.E().S(string(action.ActMemberRemove))
//line views/vworkspace/vwutil/Members.html:103
	qw422016.N().S(`" onclick="return confirm('Are you sure you wish to remove this user?');">Remove</button>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwutil/Members.html:108
}

//line views/vworkspace/vwutil/Members.html:108
func WriteMemberModalEdit(qq422016 qtio422016.Writer, m *member.Member, url string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:108
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Members.html:108
	StreamMemberModalEdit(qw422016, m, url, ps)
//line views/vworkspace/vwutil/Members.html:108
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Members.html:108
}

//line views/vworkspace/vwutil/Members.html:108
func MemberModalEdit(m *member.Member, url string, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Members.html:108
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Members.html:108
	WriteMemberModalEdit(qb422016, m, url, ps)
//line views/vworkspace/vwutil/Members.html:108
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Members.html:108
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Members.html:108
	return qs422016
//line views/vworkspace/vwutil/Members.html:108
}

//line views/vworkspace/vwutil/Members.html:110
func StreamMemberModalView(qw422016 *qt422016.Writer, m *member.Member, url string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:110
	qw422016.N().S(`
  <div id="modal-member-`)
//line views/vworkspace/vwutil/Members.html:111
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:111
	qw422016.N().S(`" data-id="`)
//line views/vworkspace/vwutil/Members.html:111
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:111
	qw422016.N().S(`" class="modal modal-member" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>
          <span class="member-picture">
`)
//line views/vworkspace/vwutil/Members.html:118
	if m.Picture == "" {
//line views/vworkspace/vwutil/Members.html:118
		qw422016.N().S(`            `)
//line views/vworkspace/vwutil/Members.html:119
		components.StreamSVGRef(qw422016, `profile`, 24, 24, `icon`, ps)
//line views/vworkspace/vwutil/Members.html:119
		qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Members.html:120
	} else {
//line views/vworkspace/vwutil/Members.html:120
		qw422016.N().S(`            <img class="icon" style="width: 24px; height: 24px;" src="`)
//line views/vworkspace/vwutil/Members.html:121
		qw422016.E().S(m.Picture)
//line views/vworkspace/vwutil/Members.html:121
		qw422016.N().S(`" />
`)
//line views/vworkspace/vwutil/Members.html:122
	}
//line views/vworkspace/vwutil/Members.html:122
	qw422016.N().S(`          </span>
          <span class="member-name member-`)
//line views/vworkspace/vwutil/Members.html:124
	qw422016.E().S(m.UserID.String())
//line views/vworkspace/vwutil/Members.html:124
	qw422016.N().S(`-name">`)
//line views/vworkspace/vwutil/Members.html:124
	qw422016.E().S(m.Name)
//line views/vworkspace/vwutil/Members.html:124
	qw422016.N().S(`</span>
        </h2>
      </div>
      <div class="modal-body">
        <em>Role</em><br />
        <span class="member-role">`)
//line views/vworkspace/vwutil/Members.html:129
	qw422016.E().S(m.Role.String())
//line views/vworkspace/vwutil/Members.html:129
	qw422016.N().S(`</span>
      </div>
    </div>
  </div>
`)
//line views/vworkspace/vwutil/Members.html:133
}

//line views/vworkspace/vwutil/Members.html:133
func WriteMemberModalView(qq422016 qtio422016.Writer, m *member.Member, url string, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Members.html:133
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Members.html:133
	StreamMemberModalView(qw422016, m, url, ps)
//line views/vworkspace/vwutil/Members.html:133
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Members.html:133
}

//line views/vworkspace/vwutil/Members.html:133
func MemberModalView(m *member.Member, url string, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Members.html:133
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Members.html:133
	WriteMemberModalView(qb422016, m, url, ps)
//line views/vworkspace/vwutil/Members.html:133
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Members.html:133
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Members.html:133
	return qs422016
//line views/vworkspace/vwutil/Members.html:133
}
