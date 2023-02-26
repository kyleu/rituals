// Code generated by qtc from "seed_estimate_member.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/seeddata/seed_estimate_member.sql:1
package seeddata

//line queries/seeddata/seed_estimate_member.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/seeddata/seed_estimate_member.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/seeddata/seed_estimate_member.sql:1
func StreamEstimateMemberSeedData(qw422016 *qt422016.Writer) {
//line queries/seeddata/seed_estimate_member.sql:1
	qw422016.N().S(`
insert into "estimate_member" (
  "estimate_id", "user_id", "name", "picture", "role", "created", "updated"
) values (
  '30000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', 'Test User', '/assets/logo.png', 'owner', now(), null
), (
  '30000000-0000-0000-0000-000000000000', '90000001-0000-0000-0000-000000000000', 'Test User 2', '/assets/logo.png', 'member', now(), null
) on conflict do nothing;
-- `)
//line queries/seeddata/seed_estimate_member.sql:9
}

//line queries/seeddata/seed_estimate_member.sql:9
func WriteEstimateMemberSeedData(qq422016 qtio422016.Writer) {
//line queries/seeddata/seed_estimate_member.sql:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/seeddata/seed_estimate_member.sql:9
	StreamEstimateMemberSeedData(qw422016)
//line queries/seeddata/seed_estimate_member.sql:9
	qt422016.ReleaseWriter(qw422016)
//line queries/seeddata/seed_estimate_member.sql:9
}

//line queries/seeddata/seed_estimate_member.sql:9
func EstimateMemberSeedData() string {
//line queries/seeddata/seed_estimate_member.sql:9
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/seeddata/seed_estimate_member.sql:9
	WriteEstimateMemberSeedData(qb422016)
//line queries/seeddata/seed_estimate_member.sql:9
	qs422016 := string(qb422016.B)
//line queries/seeddata/seed_estimate_member.sql:9
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/seeddata/seed_estimate_member.sql:9
	return qs422016
//line queries/seeddata/seed_estimate_member.sql:9
}
