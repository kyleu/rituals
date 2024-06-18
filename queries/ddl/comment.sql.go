// Code generated by qtc from "comment.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/ddl/comment.sql:1
package ddl

//line queries/ddl/comment.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/comment.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/comment.sql:1
func StreamCommentDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/comment.sql:1
	qw422016.N().S(`
drop table if exists "comment";
-- `)
//line queries/ddl/comment.sql:3
}

//line queries/ddl/comment.sql:3
func WriteCommentDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/comment.sql:3
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/comment.sql:3
	StreamCommentDrop(qw422016)
//line queries/ddl/comment.sql:3
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/comment.sql:3
}

//line queries/ddl/comment.sql:3
func CommentDrop() string {
//line queries/ddl/comment.sql:3
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/comment.sql:3
	WriteCommentDrop(qb422016)
//line queries/ddl/comment.sql:3
	qs422016 := string(qb422016.B)
//line queries/ddl/comment.sql:3
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/comment.sql:3
	return qs422016
//line queries/ddl/comment.sql:3
}

// --

//line queries/ddl/comment.sql:5
func StreamCommentCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/comment.sql:5
	qw422016.N().S(`
create table if not exists "comment" (
  "id" uuid not null,
  "svc" model_service not null,
  "model_id" uuid not null,
  "user_id" uuid not null,
  "content" text not null,
  "html" text not null,
  "created" timestamp not null default now(),
  foreign key ("user_id") references "user" ("id"),
  primary key ("id")
);

create index if not exists "comment__user_id_idx" on "comment" ("user_id");
create index if not exists "comment__svc_model_id_idx" on "comment" ("svc", "model_id");
-- `)
//line queries/ddl/comment.sql:20
}

//line queries/ddl/comment.sql:20
func WriteCommentCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/comment.sql:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/comment.sql:20
	StreamCommentCreate(qw422016)
//line queries/ddl/comment.sql:20
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/comment.sql:20
}

//line queries/ddl/comment.sql:20
func CommentCreate() string {
//line queries/ddl/comment.sql:20
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/comment.sql:20
	WriteCommentCreate(qb422016)
//line queries/ddl/comment.sql:20
	qs422016 := string(qb422016.B)
//line queries/ddl/comment.sql:20
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/comment.sql:20
	return qs422016
//line queries/ddl/comment.sql:20
}
