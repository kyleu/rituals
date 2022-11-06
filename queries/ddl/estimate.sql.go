// Code generated by qtc from "estimate.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/estimate.sql:2
package ddl

//line queries/ddl/estimate.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/estimate.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/estimate.sql:2
func StreamEstimateDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/estimate.sql:2
	qw422016.N().S(`
drop table if exists "estimate";
-- `)
//line queries/ddl/estimate.sql:4
}

//line queries/ddl/estimate.sql:4
func WriteEstimateDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/estimate.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/estimate.sql:4
	StreamEstimateDrop(qw422016)
//line queries/ddl/estimate.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/estimate.sql:4
}

//line queries/ddl/estimate.sql:4
func EstimateDrop() string {
//line queries/ddl/estimate.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/estimate.sql:4
	WriteEstimateDrop(qb422016)
//line queries/ddl/estimate.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/estimate.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/estimate.sql:4
	return qs422016
//line queries/ddl/estimate.sql:4
}

// --

//line queries/ddl/estimate.sql:6
func StreamEstimateCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/estimate.sql:6
	qw422016.N().S(`
create table if not exists "estimate" (
  "id" uuid not null,
  "slug" text not null,
  "title" text not null,
  "status" session_status not null,
  "team_id" uuid,
  "sprint_id" uuid,
  "owner" uuid not null,
  "choices" jsonb not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("owner") references "user" ("id"),
  foreign key ("team_id") references "team" ("id"),
  foreign key ("sprint_id") references "sprint" ("id"),
  unique ("slug"),
  primary key ("id")
);

create index if not exists "estimate__slug_idx" on "estimate" ("slug");

create index if not exists "estimate__status_idx" on "estimate" ("status");

create index if not exists "estimate__owner_idx" on "estimate" ("owner");

create index if not exists "estimate__team_id_idx" on "estimate" ("team_id");

create index if not exists "estimate__sprint_id_idx" on "estimate" ("sprint_id");
-- `)
//line queries/ddl/estimate.sql:34
}

//line queries/ddl/estimate.sql:34
func WriteEstimateCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/estimate.sql:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/estimate.sql:34
	StreamEstimateCreate(qw422016)
//line queries/ddl/estimate.sql:34
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/estimate.sql:34
}

//line queries/ddl/estimate.sql:34
func EstimateCreate() string {
//line queries/ddl/estimate.sql:34
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/estimate.sql:34
	WriteEstimateCreate(qb422016)
//line queries/ddl/estimate.sql:34
	qs422016 := string(qb422016.B)
//line queries/ddl/estimate.sql:34
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/estimate.sql:34
	return qs422016
//line queries/ddl/estimate.sql:34
}