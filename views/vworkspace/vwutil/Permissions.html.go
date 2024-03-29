// Code generated by qtc from "Permissions.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vworkspace/vwutil/Permissions.html:1
package vwutil

//line views/vworkspace/vwutil/Permissions.html:1
import (
	"github.com/google/uuid"

	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/components"
)

//line views/vworkspace/vwutil/Permissions.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vworkspace/vwutil/Permissions.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vworkspace/vwutil/Permissions.html:12
func StreamPermissionsLink(qw422016 *qt422016.Writer, svc enum.ModelService, id uuid.UUID, permissions util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Permissions.html:12
	qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Permissions.html:14
	title := util.StringPlural(len(permissions), "permission")
	icon := "lock"
	if len(permissions) == 0 {
		icon = "unlock"
	}

//line views/vworkspace/vwutil/Permissions.html:19
	qw422016.N().S(`  <a class="permission-link" href="#modal-`)
//line views/vworkspace/vwutil/Permissions.html:20
	qw422016.E().S(svc.Key)
//line views/vworkspace/vwutil/Permissions.html:20
	qw422016.N().S(`-config" title="`)
//line views/vworkspace/vwutil/Permissions.html:20
	qw422016.E().S(title)
//line views/vworkspace/vwutil/Permissions.html:20
	qw422016.N().S(`">`)
//line views/vworkspace/vwutil/Permissions.html:20
	components.StreamSVGRef(qw422016, icon, 18, 18, "", ps)
//line views/vworkspace/vwutil/Permissions.html:20
	qw422016.N().S(`</a>
`)
//line views/vworkspace/vwutil/Permissions.html:21
}

//line views/vworkspace/vwutil/Permissions.html:21
func WritePermissionsLink(qq422016 qtio422016.Writer, svc enum.ModelService, id uuid.UUID, permissions util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Permissions.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Permissions.html:21
	StreamPermissionsLink(qw422016, svc, id, permissions, ps)
//line views/vworkspace/vwutil/Permissions.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Permissions.html:21
}

//line views/vworkspace/vwutil/Permissions.html:21
func PermissionsLink(svc enum.ModelService, id uuid.UUID, permissions util.Permissions, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Permissions.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Permissions.html:21
	WritePermissionsLink(qb422016, svc, id, permissions, ps)
//line views/vworkspace/vwutil/Permissions.html:21
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Permissions.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Permissions.html:21
	return qs422016
//line views/vworkspace/vwutil/Permissions.html:21
}

//line views/vworkspace/vwutil/Permissions.html:23
func StreamPermissionsForm(qw422016 *qt422016.Writer, key string, perms util.Permissions, showTeam bool, teams team.Teams, showSprint bool, sprints sprint.Sprints, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Permissions.html:23
	qw422016.N().S(`
  <em class="title">Permissions</em>
`)
//line views/vworkspace/vwutil/Permissions.html:25
	if showTeam {
//line views/vworkspace/vwutil/Permissions.html:25
		qw422016.N().S(`    <div class="permission-config-team mt"><label>
`)
//line views/vworkspace/vwutil/Permissions.html:27
		if perms.Get(util.KeyTeam, "true") != nil {
//line views/vworkspace/vwutil/Permissions.html:27
			qw422016.N().S(`      <input class="perm-option" type="checkbox" name="perm-team" value="true" checked="checked">
`)
//line views/vworkspace/vwutil/Permissions.html:29
		} else {
//line views/vworkspace/vwutil/Permissions.html:29
			qw422016.N().S(`      <input class="perm-option" type="checkbox" name="perm-team" value="true">
`)
//line views/vworkspace/vwutil/Permissions.html:31
		}
//line views/vworkspace/vwutil/Permissions.html:31
		qw422016.N().S(`      Must be a member of this `)
//line views/vworkspace/vwutil/Permissions.html:32
		qw422016.E().S(key)
//line views/vworkspace/vwutil/Permissions.html:32
		qw422016.N().S(`'s team
    </label></div>
`)
//line views/vworkspace/vwutil/Permissions.html:34
	}
//line views/vworkspace/vwutil/Permissions.html:35
	if showSprint {
//line views/vworkspace/vwutil/Permissions.html:35
		qw422016.N().S(`    <div class="permission-config-sprint mt"><label>
`)
//line views/vworkspace/vwutil/Permissions.html:37
		if perms.Get(util.KeySprint, "true") != nil {
//line views/vworkspace/vwutil/Permissions.html:37
			qw422016.N().S(`      <input class="perm-option" type="checkbox" name="perm-sprint" value="true" checked="checked">
`)
//line views/vworkspace/vwutil/Permissions.html:39
		} else {
//line views/vworkspace/vwutil/Permissions.html:39
			qw422016.N().S(`      <input class="perm-option" type="checkbox" name="perm-sprint" value="true">
`)
//line views/vworkspace/vwutil/Permissions.html:41
		}
//line views/vworkspace/vwutil/Permissions.html:41
		qw422016.N().S(`      Must be a member of this `)
//line views/vworkspace/vwutil/Permissions.html:42
		qw422016.E().S(key)
//line views/vworkspace/vwutil/Permissions.html:42
		qw422016.N().S(`'s sprint
    </label></div>
`)
//line views/vworkspace/vwutil/Permissions.html:44
	}
//line views/vworkspace/vwutil/Permissions.html:45
	for _, perm := range perms.AuthPerms() {
//line views/vworkspace/vwutil/Permissions.html:46
		if (perm.Value != "*" && len(ps.Accounts.GetByProvider(perm.Key)) == 0) && ps.Accounts.GetByProviderDomain(perm.Key, perm.Value) == nil {
//line views/vworkspace/vwutil/Permissions.html:46
			qw422016.N().S(`
    <div class="mt"><label>
`)
//line views/vworkspace/vwutil/Permissions.html:48
			if perms.Get(perm.Key, "*") != nil {
//line views/vworkspace/vwutil/Permissions.html:48
				qw422016.N().S(`      <input class="perm-option" type="checkbox" name="perm-`)
//line views/vworkspace/vwutil/Permissions.html:49
				qw422016.E().S(perm.Key)
//line views/vworkspace/vwutil/Permissions.html:49
				qw422016.N().S(`" value="true" checked="checked">
`)
//line views/vworkspace/vwutil/Permissions.html:50
			} else {
//line views/vworkspace/vwutil/Permissions.html:50
				qw422016.N().S(`      <input class="perm-option" type="checkbox" name="perm-`)
//line views/vworkspace/vwutil/Permissions.html:51
				qw422016.E().S(perm.Key)
//line views/vworkspace/vwutil/Permissions.html:51
				qw422016.N().S(`" value="true">
`)
//line views/vworkspace/vwutil/Permissions.html:52
			}
//line views/vworkspace/vwutil/Permissions.html:52
			qw422016.N().S(`      Must be signed into [`)
//line views/vworkspace/vwutil/Permissions.html:53
			qw422016.E().S(perm.Key)
//line views/vworkspace/vwutil/Permissions.html:53
			qw422016.N().S(`]
    </label></div>
    <div class="mt"><label>
`)
//line views/vworkspace/vwutil/Permissions.html:56
			if perms.Get(perm.Key, perm.Value) != nil {
//line views/vworkspace/vwutil/Permissions.html:56
				qw422016.N().S(`      <input class="perm-option" type="checkbox" name="perm-`)
//line views/vworkspace/vwutil/Permissions.html:57
				qw422016.E().S(perm.Key)
//line views/vworkspace/vwutil/Permissions.html:57
				qw422016.N().S(`-`)
//line views/vworkspace/vwutil/Permissions.html:57
				qw422016.E().S(perm.Value)
//line views/vworkspace/vwutil/Permissions.html:57
				qw422016.N().S(`" value="true" checked="checked">
`)
//line views/vworkspace/vwutil/Permissions.html:58
			} else {
//line views/vworkspace/vwutil/Permissions.html:58
				qw422016.N().S(`      <input class="perm-option" type="checkbox" name="perm-`)
//line views/vworkspace/vwutil/Permissions.html:59
				qw422016.E().S(perm.Key)
//line views/vworkspace/vwutil/Permissions.html:59
				qw422016.N().S(`-`)
//line views/vworkspace/vwutil/Permissions.html:59
				qw422016.E().S(perm.Value)
//line views/vworkspace/vwutil/Permissions.html:59
				qw422016.N().S(`" value="true">
`)
//line views/vworkspace/vwutil/Permissions.html:60
			}
//line views/vworkspace/vwutil/Permissions.html:60
			qw422016.N().S(`      Must be signed into [`)
//line views/vworkspace/vwutil/Permissions.html:61
			qw422016.E().S(perm.Key)
//line views/vworkspace/vwutil/Permissions.html:61
			qw422016.N().S(`] from [`)
//line views/vworkspace/vwutil/Permissions.html:61
			qw422016.E().S(perm.Value)
//line views/vworkspace/vwutil/Permissions.html:61
			qw422016.N().S(`]
    </label></div>
`)
//line views/vworkspace/vwutil/Permissions.html:63
		}
//line views/vworkspace/vwutil/Permissions.html:64
	}
//line views/vworkspace/vwutil/Permissions.html:65
	if len(ps.Accounts) == 0 {
//line views/vworkspace/vwutil/Permissions.html:65
		qw422016.N().S(`    <div class="mt">Control access to this team by <a href="/profile">signing in</a></div>
`)
//line views/vworkspace/vwutil/Permissions.html:67
	} else {
//line views/vworkspace/vwutil/Permissions.html:68
		for _, acct := range ps.Accounts {
//line views/vworkspace/vwutil/Permissions.html:68
			qw422016.N().S(`    <div class="mt"><label>
`)
//line views/vworkspace/vwutil/Permissions.html:70
			if perms.Get(acct.Provider, "*") != nil {
//line views/vworkspace/vwutil/Permissions.html:70
				qw422016.N().S(`      <input class="perm-option" type="checkbox" name="perm-`)
//line views/vworkspace/vwutil/Permissions.html:71
				qw422016.E().S(acct.Provider)
//line views/vworkspace/vwutil/Permissions.html:71
				qw422016.N().S(`" value="true" checked="checked">
`)
//line views/vworkspace/vwutil/Permissions.html:72
			} else {
//line views/vworkspace/vwutil/Permissions.html:72
				qw422016.N().S(`      <input class="perm-option" type="checkbox" name="perm-`)
//line views/vworkspace/vwutil/Permissions.html:73
				qw422016.E().S(acct.Provider)
//line views/vworkspace/vwutil/Permissions.html:73
				qw422016.N().S(`" value="true">
`)
//line views/vworkspace/vwutil/Permissions.html:74
			}
//line views/vworkspace/vwutil/Permissions.html:74
			qw422016.N().S(`      Must be signed into [`)
//line views/vworkspace/vwutil/Permissions.html:75
			qw422016.E().S(acct.Provider)
//line views/vworkspace/vwutil/Permissions.html:75
			qw422016.N().S(`]
    </label></div>
    <div class="mt"><label>
`)
//line views/vworkspace/vwutil/Permissions.html:78
			if perms.Get(acct.Provider, acct.Domain()) != nil {
//line views/vworkspace/vwutil/Permissions.html:78
				qw422016.N().S(`      <input class="perm-option" type="checkbox" name="perm-`)
//line views/vworkspace/vwutil/Permissions.html:79
				qw422016.E().S(acct.Provider)
//line views/vworkspace/vwutil/Permissions.html:79
				qw422016.N().S(`-`)
//line views/vworkspace/vwutil/Permissions.html:79
				qw422016.E().S(acct.Domain())
//line views/vworkspace/vwutil/Permissions.html:79
				qw422016.N().S(`" value="true" checked="checked">
`)
//line views/vworkspace/vwutil/Permissions.html:80
			} else {
//line views/vworkspace/vwutil/Permissions.html:80
				qw422016.N().S(`      <input class="perm-option" type="checkbox" name="perm-`)
//line views/vworkspace/vwutil/Permissions.html:81
				qw422016.E().S(acct.Provider)
//line views/vworkspace/vwutil/Permissions.html:81
				qw422016.N().S(`-`)
//line views/vworkspace/vwutil/Permissions.html:81
				qw422016.E().S(acct.Domain())
//line views/vworkspace/vwutil/Permissions.html:81
				qw422016.N().S(`" value="true">
`)
//line views/vworkspace/vwutil/Permissions.html:82
			}
//line views/vworkspace/vwutil/Permissions.html:82
			qw422016.N().S(`      Must be signed into [`)
//line views/vworkspace/vwutil/Permissions.html:83
			qw422016.E().S(acct.Provider)
//line views/vworkspace/vwutil/Permissions.html:83
			qw422016.N().S(`] from [`)
//line views/vworkspace/vwutil/Permissions.html:83
			qw422016.E().S(acct.Domain())
//line views/vworkspace/vwutil/Permissions.html:83
			qw422016.N().S(`]
    </label></div>
`)
//line views/vworkspace/vwutil/Permissions.html:85
		}
//line views/vworkspace/vwutil/Permissions.html:86
	}
//line views/vworkspace/vwutil/Permissions.html:87
}

//line views/vworkspace/vwutil/Permissions.html:87
func WritePermissionsForm(qq422016 qtio422016.Writer, key string, perms util.Permissions, showTeam bool, teams team.Teams, showSprint bool, sprints sprint.Sprints, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Permissions.html:87
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Permissions.html:87
	StreamPermissionsForm(qw422016, key, perms, showTeam, teams, showSprint, sprints, ps)
//line views/vworkspace/vwutil/Permissions.html:87
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Permissions.html:87
}

//line views/vworkspace/vwutil/Permissions.html:87
func PermissionsForm(key string, perms util.Permissions, showTeam bool, teams team.Teams, showSprint bool, sprints sprint.Sprints, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Permissions.html:87
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Permissions.html:87
	WritePermissionsForm(qb422016, key, perms, showTeam, teams, showSprint, sprints, ps)
//line views/vworkspace/vwutil/Permissions.html:87
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Permissions.html:87
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Permissions.html:87
	return qs422016
//line views/vworkspace/vwutil/Permissions.html:87
}

//line views/vworkspace/vwutil/Permissions.html:89
func StreamPermissionsList(qw422016 *qt422016.Writer, key string, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Permissions.html:89
	qw422016.N().S(`
`)
//line views/vworkspace/vwutil/Permissions.html:90
	if len(perms) == 0 {
//line views/vworkspace/vwutil/Permissions.html:90
		qw422016.N().S(`    <div>Open</div>
`)
//line views/vworkspace/vwutil/Permissions.html:92
	}
//line views/vworkspace/vwutil/Permissions.html:93
	if perms.Get(util.KeyTeam, "true") != nil {
//line views/vworkspace/vwutil/Permissions.html:93
		qw422016.N().S(`    <div class="permission-config-team">Must be a member of this `)
//line views/vworkspace/vwutil/Permissions.html:94
		qw422016.E().S(key)
//line views/vworkspace/vwutil/Permissions.html:94
		qw422016.N().S(`'s team</div>
`)
//line views/vworkspace/vwutil/Permissions.html:95
	}
//line views/vworkspace/vwutil/Permissions.html:96
	if perms.Get(util.KeySprint, "true") != nil {
//line views/vworkspace/vwutil/Permissions.html:96
		qw422016.N().S(`    <div class="permission-config-sprint">Must be a member of this `)
//line views/vworkspace/vwutil/Permissions.html:97
		qw422016.E().S(key)
//line views/vworkspace/vwutil/Permissions.html:97
		qw422016.N().S(`'s sprint</div>
`)
//line views/vworkspace/vwutil/Permissions.html:98
	}
//line views/vworkspace/vwutil/Permissions.html:99
	for _, perm := range perms.AuthPerms() {
//line views/vworkspace/vwutil/Permissions.html:100
		if perms.Get(perm.Key, "*") != nil {
//line views/vworkspace/vwutil/Permissions.html:100
			qw422016.N().S(`      <div>Must be signed into [`)
//line views/vworkspace/vwutil/Permissions.html:101
			qw422016.E().S(perm.Key)
//line views/vworkspace/vwutil/Permissions.html:101
			qw422016.N().S(`]</div>
`)
//line views/vworkspace/vwutil/Permissions.html:102
		}
//line views/vworkspace/vwutil/Permissions.html:103
		if perms.Get(perm.Key, perm.Value) != nil {
//line views/vworkspace/vwutil/Permissions.html:103
			qw422016.N().S(`      <div>Must be signed into [`)
//line views/vworkspace/vwutil/Permissions.html:104
			qw422016.E().S(perm.Key)
//line views/vworkspace/vwutil/Permissions.html:104
			qw422016.N().S(`] from [`)
//line views/vworkspace/vwutil/Permissions.html:104
			qw422016.E().S(perm.Value)
//line views/vworkspace/vwutil/Permissions.html:104
			qw422016.N().S(`]</div>
`)
//line views/vworkspace/vwutil/Permissions.html:105
		}
//line views/vworkspace/vwutil/Permissions.html:106
	}
//line views/vworkspace/vwutil/Permissions.html:107
}

//line views/vworkspace/vwutil/Permissions.html:107
func WritePermissionsList(qq422016 qtio422016.Writer, key string, perms util.Permissions, ps *cutil.PageState) {
//line views/vworkspace/vwutil/Permissions.html:107
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vworkspace/vwutil/Permissions.html:107
	StreamPermissionsList(qw422016, key, perms, ps)
//line views/vworkspace/vwutil/Permissions.html:107
	qt422016.ReleaseWriter(qw422016)
//line views/vworkspace/vwutil/Permissions.html:107
}

//line views/vworkspace/vwutil/Permissions.html:107
func PermissionsList(key string, perms util.Permissions, ps *cutil.PageState) string {
//line views/vworkspace/vwutil/Permissions.html:107
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vworkspace/vwutil/Permissions.html:107
	WritePermissionsList(qb422016, key, perms, ps)
//line views/vworkspace/vwutil/Permissions.html:107
	qs422016 := string(qb422016.B)
//line views/vworkspace/vwutil/Permissions.html:107
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vworkspace/vwutil/Permissions.html:107
	return qs422016
//line views/vworkspace/vwutil/Permissions.html:107
}
