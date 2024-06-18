// Code generated by qtc from "estimate_history.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/ddl/estimate_history.sql:1
package ddl

//line queries/ddl/estimate_history.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/estimate_history.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/estimate_history.sql:1
func StreamEstimateHistoryDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/estimate_history.sql:1
	qw422016.N().S(`
drop table if exists "estimate_history";
-- `)
//line queries/ddl/estimate_history.sql:3
}

//line queries/ddl/estimate_history.sql:3
func WriteEstimateHistoryDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/estimate_history.sql:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/estimate_history.sql:3
	StreamEstimateHistoryDrop(qw422016)
//line queries/ddl/estimate_history.sql:3
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/estimate_history.sql:3
}

//line queries/ddl/estimate_history.sql:3
func EstimateHistoryDrop() string {
//line queries/ddl/estimate_history.sql:3
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/estimate_history.sql:3
	WriteEstimateHistoryDrop(qb422016)
//line queries/ddl/estimate_history.sql:3
	qs422016 := string(qb422016.B)
//line queries/ddl/estimate_history.sql:3
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/estimate_history.sql:3
	return qs422016
//line queries/ddl/estimate_history.sql:3
}

// --

//line queries/ddl/estimate_history.sql:5
func StreamEstimateHistoryCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/estimate_history.sql:5
	qw422016.N().S(`
create table if not exists "estimate_history" (
  "slug" text not null,
  "estimate_id" uuid not null,
  "estimate_name" text not null,
  "created" timestamp not null default now(),
  foreign key ("estimate_id") references "estimate" ("id"),
  primary key ("slug")
);

create index if not exists "estimate_history__estimate_id_idx" on "estimate_history" ("estimate_id");
-- `)
//line queries/ddl/estimate_history.sql:16
}

//line queries/ddl/estimate_history.sql:16
func WriteEstimateHistoryCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/estimate_history.sql:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/estimate_history.sql:16
	StreamEstimateHistoryCreate(qw422016)
//line queries/ddl/estimate_history.sql:16
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/estimate_history.sql:16
}

//line queries/ddl/estimate_history.sql:16
func EstimateHistoryCreate() string {
//line queries/ddl/estimate_history.sql:16
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/estimate_history.sql:16
	WriteEstimateHistoryCreate(qb422016)
//line queries/ddl/estimate_history.sql:16
	qs422016 := string(qb422016.B)
//line queries/ddl/estimate_history.sql:16
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/estimate_history.sql:16
	return qs422016
//line queries/ddl/estimate_history.sql:16
}
