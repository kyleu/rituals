// Code generated by qtc from "sprint_permission.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/sprint_permission.sql:2
package ddl

//line queries/ddl/sprint_permission.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/sprint_permission.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/sprint_permission.sql:2
func StreamSprintPermissionDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/sprint_permission.sql:2
	qw422016.N().S(`
drop table if exists "sprint_permission";
-- `)
//line queries/ddl/sprint_permission.sql:4
}

//line queries/ddl/sprint_permission.sql:4
func WriteSprintPermissionDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/sprint_permission.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/sprint_permission.sql:4
	StreamSprintPermissionDrop(qw422016)
//line queries/ddl/sprint_permission.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/sprint_permission.sql:4
}

//line queries/ddl/sprint_permission.sql:4
func SprintPermissionDrop() string {
//line queries/ddl/sprint_permission.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/sprint_permission.sql:4
	WriteSprintPermissionDrop(qb422016)
//line queries/ddl/sprint_permission.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/sprint_permission.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/sprint_permission.sql:4
	return qs422016
//line queries/ddl/sprint_permission.sql:4
}

// --

//line queries/ddl/sprint_permission.sql:6
func StreamSprintPermissionCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/sprint_permission.sql:6
	qw422016.N().S(`
create table if not exists "sprint_permission" (
  "sprint_id" uuid not null,
  "k" text not null,
  "v" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("sprint_id") references "sprint" ("id"),
  primary key ("sprint_id", "k", "v")
);

create index if not exists "sprint_permission__sprint_id_idx" on "sprint_permission" ("sprint_id");

create index if not exists "sprint_permission__k_idx" on "sprint_permission" ("k");

create index if not exists "sprint_permission__v_idx" on "sprint_permission" ("v");
-- `)
//line queries/ddl/sprint_permission.sql:22
}

//line queries/ddl/sprint_permission.sql:22
func WriteSprintPermissionCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/sprint_permission.sql:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/sprint_permission.sql:22
	StreamSprintPermissionCreate(qw422016)
//line queries/ddl/sprint_permission.sql:22
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/sprint_permission.sql:22
}

//line queries/ddl/sprint_permission.sql:22
func SprintPermissionCreate() string {
//line queries/ddl/sprint_permission.sql:22
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/sprint_permission.sql:22
	WriteSprintPermissionCreate(qb422016)
//line queries/ddl/sprint_permission.sql:22
	qs422016 := string(qb422016.B)
//line queries/ddl/sprint_permission.sql:22
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/sprint_permission.sql:22
	return qs422016
//line queries/ddl/sprint_permission.sql:22
}
