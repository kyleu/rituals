// Code generated by qtc from "team_history.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/team_history.sql:2
package ddl

//line queries/ddl/team_history.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/team_history.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/team_history.sql:2
func StreamTeamHistoryDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/team_history.sql:2
	qw422016.N().S(`
drop table if exists "team_history";
-- `)
//line queries/ddl/team_history.sql:4
}

//line queries/ddl/team_history.sql:4
func WriteTeamHistoryDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/team_history.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/team_history.sql:4
	StreamTeamHistoryDrop(qw422016)
//line queries/ddl/team_history.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/team_history.sql:4
}

//line queries/ddl/team_history.sql:4
func TeamHistoryDrop() string {
//line queries/ddl/team_history.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/team_history.sql:4
	WriteTeamHistoryDrop(qb422016)
//line queries/ddl/team_history.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/team_history.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/team_history.sql:4
	return qs422016
//line queries/ddl/team_history.sql:4
}

// --

//line queries/ddl/team_history.sql:6
func StreamTeamHistoryCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/team_history.sql:6
	qw422016.N().S(`
create table if not exists "team_history" (
  "slug" text not null,
  "team_id" uuid not null,
  "team_name" text not null,
  "created" timestamp not null default now(),
  foreign key ("team_id") references "team" ("id"),
  primary key ("slug")
);

create index if not exists "team_history__team_id_idx" on "team_history" ("team_id");
-- `)
//line queries/ddl/team_history.sql:17
}

//line queries/ddl/team_history.sql:17
func WriteTeamHistoryCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/team_history.sql:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/team_history.sql:17
	StreamTeamHistoryCreate(qw422016)
//line queries/ddl/team_history.sql:17
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/team_history.sql:17
}

//line queries/ddl/team_history.sql:17
func TeamHistoryCreate() string {
//line queries/ddl/team_history.sql:17
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/team_history.sql:17
	WriteTeamHistoryCreate(qb422016)
//line queries/ddl/team_history.sql:17
	qs422016 := string(qb422016.B)
//line queries/ddl/team_history.sql:17
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/team_history.sql:17
	return qs422016
//line queries/ddl/team_history.sql:17
}
