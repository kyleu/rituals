// Code generated by qtc from "team_history.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/ddl/team_history.sql:1
package ddl

//line queries/ddl/team_history.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/team_history.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/team_history.sql:1
func StreamTeamHistoryDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/team_history.sql:1
	qw422016.N().S(`
drop table if exists "team_history";
-- `)
//line queries/ddl/team_history.sql:3
}

//line queries/ddl/team_history.sql:3
func WriteTeamHistoryDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/team_history.sql:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/team_history.sql:3
	StreamTeamHistoryDrop(qw422016)
//line queries/ddl/team_history.sql:3
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/team_history.sql:3
}

//line queries/ddl/team_history.sql:3
func TeamHistoryDrop() string {
//line queries/ddl/team_history.sql:3
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/team_history.sql:3
	WriteTeamHistoryDrop(qb422016)
//line queries/ddl/team_history.sql:3
	qs422016 := string(qb422016.B)
//line queries/ddl/team_history.sql:3
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/team_history.sql:3
	return qs422016
//line queries/ddl/team_history.sql:3
}

// --

//line queries/ddl/team_history.sql:5
func StreamTeamHistoryCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/team_history.sql:5
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
//line queries/ddl/team_history.sql:16
}

//line queries/ddl/team_history.sql:16
func WriteTeamHistoryCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/team_history.sql:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/team_history.sql:16
	StreamTeamHistoryCreate(qw422016)
//line queries/ddl/team_history.sql:16
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/team_history.sql:16
}

//line queries/ddl/team_history.sql:16
func TeamHistoryCreate() string {
//line queries/ddl/team_history.sql:16
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/team_history.sql:16
	WriteTeamHistoryCreate(qb422016)
//line queries/ddl/team_history.sql:16
	qs422016 := string(qb422016.B)
//line queries/ddl/team_history.sql:16
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/team_history.sql:16
	return qs422016
//line queries/ddl/team_history.sql:16
}
