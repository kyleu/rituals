// Code generated by qtc from "standup_member.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/standup_member.sql:2
package ddl

//line queries/ddl/standup_member.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/standup_member.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/standup_member.sql:2
func StreamStandupMemberDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/standup_member.sql:2
	qw422016.N().S(`
drop table if exists "standup_member";
-- `)
//line queries/ddl/standup_member.sql:4
}

//line queries/ddl/standup_member.sql:4
func WriteStandupMemberDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/standup_member.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/standup_member.sql:4
	StreamStandupMemberDrop(qw422016)
//line queries/ddl/standup_member.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/standup_member.sql:4
}

//line queries/ddl/standup_member.sql:4
func StandupMemberDrop() string {
//line queries/ddl/standup_member.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/standup_member.sql:4
	WriteStandupMemberDrop(qb422016)
//line queries/ddl/standup_member.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/standup_member.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/standup_member.sql:4
	return qs422016
//line queries/ddl/standup_member.sql:4
}

// --

//line queries/ddl/standup_member.sql:6
func StreamStandupMemberCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/standup_member.sql:6
	qw422016.N().S(`
create table if not exists "standup_member" (
  "standup_id" uuid not null,
  "user_id" uuid not null,
  "name" text not null,
  "picture" text not null,
  "role" text not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("standup_id") references "standup" ("id"),
  foreign key ("user_id") references "user" ("id"),
  primary key ("standup_id", "user_id")
);

create index if not exists "standup_member__standup_id_idx" on "standup_member" ("standup_id");

create index if not exists "standup_member__user_id_idx" on "standup_member" ("user_id");
-- `)
//line queries/ddl/standup_member.sql:23
}

//line queries/ddl/standup_member.sql:23
func WriteStandupMemberCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/standup_member.sql:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/standup_member.sql:23
	StreamStandupMemberCreate(qw422016)
//line queries/ddl/standup_member.sql:23
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/standup_member.sql:23
}

//line queries/ddl/standup_member.sql:23
func StandupMemberCreate() string {
//line queries/ddl/standup_member.sql:23
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/standup_member.sql:23
	WriteStandupMemberCreate(qb422016)
//line queries/ddl/standup_member.sql:23
	qs422016 := string(qb422016.B)
//line queries/ddl/standup_member.sql:23
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/standup_member.sql:23
	return qs422016
//line queries/ddl/standup_member.sql:23
}
