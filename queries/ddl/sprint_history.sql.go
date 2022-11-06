// Code generated by qtc from "sprint_history.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/sprint_history.sql:2
package ddl

//line queries/ddl/sprint_history.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/sprint_history.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/sprint_history.sql:2
func StreamSprintHistoryDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/sprint_history.sql:2
	qw422016.N().S(`
drop table if exists "sprint_history";
-- `)
//line queries/ddl/sprint_history.sql:4
}

//line queries/ddl/sprint_history.sql:4
func WriteSprintHistoryDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/sprint_history.sql:4
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/sprint_history.sql:4
	StreamSprintHistoryDrop(qw422016)
//line queries/ddl/sprint_history.sql:4
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/sprint_history.sql:4
}

//line queries/ddl/sprint_history.sql:4
func SprintHistoryDrop() string {
//line queries/ddl/sprint_history.sql:4
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/sprint_history.sql:4
	WriteSprintHistoryDrop(qb422016)
//line queries/ddl/sprint_history.sql:4
	qs422016 := string(qb422016.B)
//line queries/ddl/sprint_history.sql:4
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/sprint_history.sql:4
	return qs422016
//line queries/ddl/sprint_history.sql:4
}

// --

//line queries/ddl/sprint_history.sql:6
func StreamSprintHistoryCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/sprint_history.sql:6
	qw422016.N().S(`
create table if not exists "sprint_history" (
  "slug" text not null,
  "sprint_id" uuid not null,
  "sprint_name" text not null,
  "created" timestamp not null default now(),
  foreign key ("sprint_id") references "sprint" ("id"),
  primary key ("slug")
);

create index if not exists "sprint_history__sprint_id_idx" on "sprint_history" ("sprint_id");
-- `)
//line queries/ddl/sprint_history.sql:17
}

//line queries/ddl/sprint_history.sql:17
func WriteSprintHistoryCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/sprint_history.sql:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/sprint_history.sql:17
	StreamSprintHistoryCreate(qw422016)
//line queries/ddl/sprint_history.sql:17
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/sprint_history.sql:17
}

//line queries/ddl/sprint_history.sql:17
func SprintHistoryCreate() string {
//line queries/ddl/sprint_history.sql:17
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/sprint_history.sql:17
	WriteSprintHistoryCreate(qb422016)
//line queries/ddl/sprint_history.sql:17
	qs422016 := string(qb422016.B)
//line queries/ddl/sprint_history.sql:17
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/sprint_history.sql:17
	return qs422016
//line queries/ddl/sprint_history.sql:17
}