// Code generated by qtc from "seed_sprint_history.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/seeddata/seed_sprint_history.sql:1
package seeddata

//line queries/seeddata/seed_sprint_history.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/seeddata/seed_sprint_history.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/seeddata/seed_sprint_history.sql:1
func StreamSprintHistorySeedData(qw422016 *qt422016.Writer) {
//line queries/seeddata/seed_sprint_history.sql:1
	qw422016.N().S(`
insert into "sprint_history" (
  "slug", "sprint_id", "sprint_name", "created"
) values (
  'old-name', '20000000-0000-0000-0000-000000000000', 'Old Name', now()
) on conflict do nothing;
-- `)
//line queries/seeddata/seed_sprint_history.sql:7
}

//line queries/seeddata/seed_sprint_history.sql:7
func WriteSprintHistorySeedData(qq422016 qtio422016.Writer) {
//line queries/seeddata/seed_sprint_history.sql:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/seeddata/seed_sprint_history.sql:7
	StreamSprintHistorySeedData(qw422016)
//line queries/seeddata/seed_sprint_history.sql:7
	qt422016.ReleaseWriter(qw422016)
//line queries/seeddata/seed_sprint_history.sql:7
}

//line queries/seeddata/seed_sprint_history.sql:7
func SprintHistorySeedData() string {
//line queries/seeddata/seed_sprint_history.sql:7
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/seeddata/seed_sprint_history.sql:7
	WriteSprintHistorySeedData(qb422016)
//line queries/seeddata/seed_sprint_history.sql:7
	qs422016 := string(qb422016.B)
//line queries/seeddata/seed_sprint_history.sql:7
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/seeddata/seed_sprint_history.sql:7
	return qs422016
//line queries/seeddata/seed_sprint_history.sql:7
}
