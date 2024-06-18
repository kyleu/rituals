// Code generated by qtc from "team_member.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/ddl/team_member.sql:1
package ddl

//line queries/ddl/team_member.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/team_member.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/team_member.sql:1
func StreamTeamMemberDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/team_member.sql:1
	qw422016.N().S(`
drop table if exists "team_member";
-- `)
//line queries/ddl/team_member.sql:3
}

//line queries/ddl/team_member.sql:3
func WriteTeamMemberDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/team_member.sql:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/team_member.sql:3
	StreamTeamMemberDrop(qw422016)
//line queries/ddl/team_member.sql:3
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/team_member.sql:3
}

//line queries/ddl/team_member.sql:3
func TeamMemberDrop() string {
//line queries/ddl/team_member.sql:3
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/team_member.sql:3
	WriteTeamMemberDrop(qb422016)
//line queries/ddl/team_member.sql:3
	qs422016 := string(qb422016.B)
//line queries/ddl/team_member.sql:3
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/team_member.sql:3
	return qs422016
//line queries/ddl/team_member.sql:3
}

// --

//line queries/ddl/team_member.sql:5
func StreamTeamMemberCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/team_member.sql:5
	qw422016.N().S(`
create table if not exists "team_member" (
  "team_id" uuid not null,
  "user_id" uuid not null,
  "name" text not null,
  "picture" text not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("team_id") references "team" ("id"),
  foreign key ("user_id") references "user" ("id"),
  primary key ("team_id", "user_id")
);

create index if not exists "team_member__team_id_idx" on "team_member" ("team_id");

create index if not exists "team_member__user_id_idx" on "team_member" ("user_id");
-- `)
//line queries/ddl/team_member.sql:22
}

//line queries/ddl/team_member.sql:22
func WriteTeamMemberCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/team_member.sql:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/team_member.sql:22
	StreamTeamMemberCreate(qw422016)
//line queries/ddl/team_member.sql:22
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/team_member.sql:22
}

//line queries/ddl/team_member.sql:22
func TeamMemberCreate() string {
//line queries/ddl/team_member.sql:22
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/team_member.sql:22
	WriteTeamMemberCreate(qb422016)
//line queries/ddl/team_member.sql:22
	qs422016 := string(qb422016.B)
//line queries/ddl/team_member.sql:22
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/team_member.sql:22
	return qs422016
//line queries/ddl/team_member.sql:22
}
