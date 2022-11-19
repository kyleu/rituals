// Code generated by qtc from "estimate_permission.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/estimate_permission.sql:2
package ddl

//line queries/ddl/estimate_permission.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/estimate_permission.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/estimate_permission.sql:2
func StreamEstimatePermissionDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/estimate_permission.sql:2
	qw422016.N().S(`
drop table if exists "estimate_permission";
-- `)
//line queries/ddl/estimate_permission.sql:4
}

//line queries/ddl/estimate_permission.sql:4
func WriteEstimatePermissionDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/estimate_permission.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/estimate_permission.sql:4
	StreamEstimatePermissionDrop(qw422016)
//line queries/ddl/estimate_permission.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/estimate_permission.sql:4
}

//line queries/ddl/estimate_permission.sql:4
func EstimatePermissionDrop() string {
//line queries/ddl/estimate_permission.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/estimate_permission.sql:4
	WriteEstimatePermissionDrop(qb422016)
//line queries/ddl/estimate_permission.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/estimate_permission.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/estimate_permission.sql:4
	return qs422016
//line queries/ddl/estimate_permission.sql:4
}

// --

//line queries/ddl/estimate_permission.sql:6
func StreamEstimatePermissionCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/estimate_permission.sql:6
	qw422016.N().S(`
create table if not exists "estimate_permission" (
  "estimate_id" uuid not null,
  "key" text not null,
  "value" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("estimate_id") references "estimate" ("id"),
  primary key ("estimate_id", "key", "value")
);

create index if not exists "estimate_permission__estimate_id_idx" on "estimate_permission" ("estimate_id");

create index if not exists "estimate_permission__key_idx" on "estimate_permission" ("key");

create index if not exists "estimate_permission__value_idx" on "estimate_permission" ("value");
-- `)
//line queries/ddl/estimate_permission.sql:22
}

//line queries/ddl/estimate_permission.sql:22
func WriteEstimatePermissionCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/estimate_permission.sql:22
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/estimate_permission.sql:22
	StreamEstimatePermissionCreate(qw422016)
//line queries/ddl/estimate_permission.sql:22
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/estimate_permission.sql:22
}

//line queries/ddl/estimate_permission.sql:22
func EstimatePermissionCreate() string {
//line queries/ddl/estimate_permission.sql:22
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/estimate_permission.sql:22
	WriteEstimatePermissionCreate(qb422016)
//line queries/ddl/estimate_permission.sql:22
	qs422016 := string(qb422016.B)
//line queries/ddl/estimate_permission.sql:22
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/estimate_permission.sql:22
	return qs422016
//line queries/ddl/estimate_permission.sql:22
}
