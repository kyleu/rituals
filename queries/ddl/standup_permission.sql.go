// Code generated by qtc from "standup_permission.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/standup_permission.sql:2
package ddl

//line queries/ddl/standup_permission.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/standup_permission.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/standup_permission.sql:2
func StreamStandupPermissionDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/standup_permission.sql:2
	qw422016.N().S(`
drop table if exists "standup_permission";
-- `)
//line queries/ddl/standup_permission.sql:4
}

//line queries/ddl/standup_permission.sql:4
func WriteStandupPermissionDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/standup_permission.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/standup_permission.sql:4
	StreamStandupPermissionDrop(qw422016)
//line queries/ddl/standup_permission.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/standup_permission.sql:4
}

//line queries/ddl/standup_permission.sql:4
func StandupPermissionDrop() string {
//line queries/ddl/standup_permission.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/standup_permission.sql:4
	WriteStandupPermissionDrop(qb422016)
//line queries/ddl/standup_permission.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/standup_permission.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/standup_permission.sql:4
	return qs422016
//line queries/ddl/standup_permission.sql:4
}

// --

//line queries/ddl/standup_permission.sql:6
func StreamStandupPermissionCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/standup_permission.sql:6
	qw422016.N().S(`
create table if not exists "standup_permission" (
  "standup_id" uuid not null,
  "k" text not null,
  "v" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("standup_id") references "standup" ("id"),
  primary key ("standup_id", "k", "v")
);

create index if not exists "standup_permission__standup_id_idx" on "standup_permission" ("standup_id");

create index if not exists "standup_permission__k_idx" on "standup_permission" ("k");

create index if not exists "standup_permission__v_idx" on "standup_permission" ("v");
-- `)
//line queries/ddl/standup_permission.sql:22
}

//line queries/ddl/standup_permission.sql:22
func WriteStandupPermissionCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/standup_permission.sql:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/standup_permission.sql:22
	StreamStandupPermissionCreate(qw422016)
//line queries/ddl/standup_permission.sql:22
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/standup_permission.sql:22
}

//line queries/ddl/standup_permission.sql:22
func StandupPermissionCreate() string {
//line queries/ddl/standup_permission.sql:22
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/standup_permission.sql:22
	WriteStandupPermissionCreate(qb422016)
//line queries/ddl/standup_permission.sql:22
	qs422016 := string(qb422016.B)
//line queries/ddl/standup_permission.sql:22
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/standup_permission.sql:22
	return qs422016
//line queries/ddl/standup_permission.sql:22
}