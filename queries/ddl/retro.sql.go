// Code generated by qtc from "retro.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/retro.sql:2
package ddl

//line queries/ddl/retro.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/retro.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/retro.sql:2
func StreamRetroDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/retro.sql:2
	qw422016.N().S(`
drop table if exists "retro";
-- `)
//line queries/ddl/retro.sql:4
}

//line queries/ddl/retro.sql:4
func WriteRetroDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/retro.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/retro.sql:4
	StreamRetroDrop(qw422016)
//line queries/ddl/retro.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/retro.sql:4
}

//line queries/ddl/retro.sql:4
func RetroDrop() string {
//line queries/ddl/retro.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/retro.sql:4
	WriteRetroDrop(qb422016)
//line queries/ddl/retro.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/retro.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/retro.sql:4
	return qs422016
//line queries/ddl/retro.sql:4
}

// --

//line queries/ddl/retro.sql:6
func StreamRetroCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/retro.sql:6
	qw422016.N().S(`
create table if not exists "retro" (
  "id" uuid not null,
  "slug" text not null,
  "title" text not null,
  "status" session_status not null,
  "team_id" uuid,
  "sprint_id" uuid,
  "owner" uuid not null,
  "categories" jsonb not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("owner") references "user" ("id"),
  foreign key ("team_id") references "team" ("id"),
  foreign key ("sprint_id") references "sprint" ("id"),
  unique ("slug"),
  primary key ("id")
);

create index if not exists "retro__slug_idx" on "retro" ("slug");

create index if not exists "retro__status_idx" on "retro" ("status");

create index if not exists "retro__owner_idx" on "retro" ("owner");

create index if not exists "retro__team_id_idx" on "retro" ("team_id");

create index if not exists "retro__sprint_id_idx" on "retro" ("sprint_id");
-- `)
//line queries/ddl/retro.sql:34
}

//line queries/ddl/retro.sql:34
func WriteRetroCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/retro.sql:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/retro.sql:34
	StreamRetroCreate(qw422016)
//line queries/ddl/retro.sql:34
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/retro.sql:34
}

//line queries/ddl/retro.sql:34
func RetroCreate() string {
//line queries/ddl/retro.sql:34
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/retro.sql:34
	WriteRetroCreate(qb422016)
//line queries/ddl/retro.sql:34
	qs422016 := string(qb422016.B)
//line queries/ddl/retro.sql:34
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/retro.sql:34
	return qs422016
//line queries/ddl/retro.sql:34
}