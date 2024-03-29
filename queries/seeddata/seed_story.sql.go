// Code generated by qtc from "seed_story.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/seeddata/seed_story.sql:1
package seeddata

//line queries/seeddata/seed_story.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/seeddata/seed_story.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/seeddata/seed_story.sql:1
func StreamStorySeedData(qw422016 *qt422016.Writer) {
//line queries/seeddata/seed_story.sql:1
	qw422016.N().S(`
insert into "story" (
  "id", "estimate_id", "idx", "user_id", "title", "status", "final_vote", "created", "updated"
) values (
  '31000000-0000-0000-0000-000000000000', '30000000-0000-0000-0000-000000000000', 0, '90000000-0000-0000-0000-000000000000', 'Build rituals.dev', 'new', '100', now(), null
), (
  '31000001-0000-0000-0000-000000000000', '30000000-0000-0000-0000-000000000000', 0, '90000001-0000-0000-0000-000000000000', 'Make it work without JavaScript', 'new', '', now(), null
) on conflict do nothing;
-- `)
//line queries/seeddata/seed_story.sql:9
}

//line queries/seeddata/seed_story.sql:9
func WriteStorySeedData(qq422016 qtio422016.Writer) {
//line queries/seeddata/seed_story.sql:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/seeddata/seed_story.sql:9
	StreamStorySeedData(qw422016)
//line queries/seeddata/seed_story.sql:9
	qt422016.ReleaseWriter(qw422016)
//line queries/seeddata/seed_story.sql:9
}

//line queries/seeddata/seed_story.sql:9
func StorySeedData() string {
//line queries/seeddata/seed_story.sql:9
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/seeddata/seed_story.sql:9
	WriteStorySeedData(qb422016)
//line queries/seeddata/seed_story.sql:9
	qs422016 := string(qb422016.B)
//line queries/seeddata/seed_story.sql:9
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/seeddata/seed_story.sql:9
	return qs422016
//line queries/seeddata/seed_story.sql:9
}
