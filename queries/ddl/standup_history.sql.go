// Code generated by qtc from "standup_history.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/standup_history.sql:2
package ddl

//line queries/ddl/standup_history.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/standup_history.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/standup_history.sql:2
func StreamStandupHistoryDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/standup_history.sql:2
	qw422016.N().S(`
drop table if exists "standup_history";
-- `)
//line queries/ddl/standup_history.sql:4
}

//line queries/ddl/standup_history.sql:4
func WriteStandupHistoryDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/standup_history.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/standup_history.sql:4
	StreamStandupHistoryDrop(qw422016)
//line queries/ddl/standup_history.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/standup_history.sql:4
}

//line queries/ddl/standup_history.sql:4
func StandupHistoryDrop() string {
//line queries/ddl/standup_history.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/standup_history.sql:4
	WriteStandupHistoryDrop(qb422016)
//line queries/ddl/standup_history.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/standup_history.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/standup_history.sql:4
	return qs422016
//line queries/ddl/standup_history.sql:4
}

// --

//line queries/ddl/standup_history.sql:6
func StreamStandupHistoryCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/standup_history.sql:6
	qw422016.N().S(`
create table if not exists "standup_history" (
  "slug" text not null,
  "standup_id" uuid not null,
  "standup_name" text not null,
  "created" timestamp not null default now(),
  foreign key ("standup_id") references "standup" ("id"),
  primary key ("slug")
);

create index if not exists "standup_history__standup_id_idx" on "standup_history" ("standup_id");
-- `)
//line queries/ddl/standup_history.sql:17
}

//line queries/ddl/standup_history.sql:17
func WriteStandupHistoryCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/standup_history.sql:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/standup_history.sql:17
	StreamStandupHistoryCreate(qw422016)
//line queries/ddl/standup_history.sql:17
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/standup_history.sql:17
}

//line queries/ddl/standup_history.sql:17
func StandupHistoryCreate() string {
//line queries/ddl/standup_history.sql:17
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/standup_history.sql:17
	WriteStandupHistoryCreate(qb422016)
//line queries/ddl/standup_history.sql:17
	qs422016 := string(qb422016.B)
//line queries/ddl/standup_history.sql:17
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/standup_history.sql:17
	return qs422016
//line queries/ddl/standup_history.sql:17
}