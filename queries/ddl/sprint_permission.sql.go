// Code generated by qtc from "sprint_permission.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/ddl/sprint_permission.sql:1
package ddl

//line queries/ddl/sprint_permission.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/sprint_permission.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/sprint_permission.sql:1
func StreamSprintPermissionDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/sprint_permission.sql:1
	qw422016.N().S(`
drop table if exists "sprint_permission";
-- `)
//line queries/ddl/sprint_permission.sql:3
}

//line queries/ddl/sprint_permission.sql:3
func WriteSprintPermissionDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/sprint_permission.sql:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/sprint_permission.sql:3
	StreamSprintPermissionDrop(qw422016)
//line queries/ddl/sprint_permission.sql:3
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/sprint_permission.sql:3
}

//line queries/ddl/sprint_permission.sql:3
func SprintPermissionDrop() string {
//line queries/ddl/sprint_permission.sql:3
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/sprint_permission.sql:3
	WriteSprintPermissionDrop(qb422016)
//line queries/ddl/sprint_permission.sql:3
	qs422016 := string(qb422016.B)
//line queries/ddl/sprint_permission.sql:3
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/sprint_permission.sql:3
	return qs422016
//line queries/ddl/sprint_permission.sql:3
}

// --

//line queries/ddl/sprint_permission.sql:5
func StreamSprintPermissionCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/sprint_permission.sql:5
	qw422016.N().S(`
create table if not exists "sprint_permission" (
  "sprint_id" uuid not null,
  "key" text not null,
  "value" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("sprint_id") references "sprint" ("id"),
  primary key ("sprint_id", "key", "value")
);

create index if not exists "sprint_permission__sprint_id_idx" on "sprint_permission" ("sprint_id");

create index if not exists "sprint_permission__key_idx" on "sprint_permission" ("key");

create index if not exists "sprint_permission__value_idx" on "sprint_permission" ("value");
-- `)
//line queries/ddl/sprint_permission.sql:21
}

//line queries/ddl/sprint_permission.sql:21
func WriteSprintPermissionCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/sprint_permission.sql:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/sprint_permission.sql:21
	StreamSprintPermissionCreate(qw422016)
//line queries/ddl/sprint_permission.sql:21
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/sprint_permission.sql:21
}

//line queries/ddl/sprint_permission.sql:21
func SprintPermissionCreate() string {
//line queries/ddl/sprint_permission.sql:21
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/sprint_permission.sql:21
	WriteSprintPermissionCreate(qb422016)
//line queries/ddl/sprint_permission.sql:21
	qs422016 := string(qb422016.B)
//line queries/ddl/sprint_permission.sql:21
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/sprint_permission.sql:21
	return qs422016
//line queries/ddl/sprint_permission.sql:21
}
