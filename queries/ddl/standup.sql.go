// Code generated by qtc from "standup.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/ddl/standup.sql:1
package ddl

//line queries/ddl/standup.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/standup.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/standup.sql:1
func StreamStandupDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/standup.sql:1
	qw422016.N().S(`
drop table if exists "standup";
-- `)
//line queries/ddl/standup.sql:3
}

//line queries/ddl/standup.sql:3
func WriteStandupDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/standup.sql:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/standup.sql:3
	StreamStandupDrop(qw422016)
//line queries/ddl/standup.sql:3
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/standup.sql:3
}

//line queries/ddl/standup.sql:3
func StandupDrop() string {
//line queries/ddl/standup.sql:3
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/standup.sql:3
	WriteStandupDrop(qb422016)
//line queries/ddl/standup.sql:3
	qs422016 := string(qb422016.B)
//line queries/ddl/standup.sql:3
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/standup.sql:3
	return qs422016
//line queries/ddl/standup.sql:3
}

// --

//line queries/ddl/standup.sql:5
func StreamStandupCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/standup.sql:5
	qw422016.N().S(`
create table if not exists "standup" (
  "id" uuid not null,
  "slug" text not null,
  "title" text not null,
  "icon" text not null,
  "status" session_status not null,
  "team_id" uuid,
  "sprint_id" uuid,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("team_id") references "team" ("id"),
  foreign key ("sprint_id") references "sprint" ("id"),
  unique ("slug"),
  primary key ("id")
);

create index if not exists "standup__slug_idx" on "standup" ("slug");

create index if not exists "standup__status_idx" on "standup" ("status");

create index if not exists "standup__team_id_idx" on "standup" ("team_id");

create index if not exists "standup__sprint_id_idx" on "standup" ("sprint_id");
-- `)
//line queries/ddl/standup.sql:29
}

//line queries/ddl/standup.sql:29
func WriteStandupCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/standup.sql:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/standup.sql:29
	StreamStandupCreate(qw422016)
//line queries/ddl/standup.sql:29
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/standup.sql:29
}

//line queries/ddl/standup.sql:29
func StandupCreate() string {
//line queries/ddl/standup.sql:29
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/standup.sql:29
	WriteStandupCreate(qb422016)
//line queries/ddl/standup.sql:29
	qs422016 := string(qb422016.B)
//line queries/ddl/standup.sql:29
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/standup.sql:29
	return qs422016
//line queries/ddl/standup.sql:29
}
