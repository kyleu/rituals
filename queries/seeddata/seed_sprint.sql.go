// Code generated by qtc from "seed_sprint.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/seeddata/seed_sprint.sql:1
package seeddata

//line queries/seeddata/seed_sprint.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/seeddata/seed_sprint.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/seeddata/seed_sprint.sql:1
func StreamSprintSeedData(qw422016 *qt422016.Writer) {
//line queries/seeddata/seed_sprint.sql:1
	qw422016.N().S(`
insert into "sprint" (
  "id", "slug", "title", "icon", "status", "team_id", "owner", "start_date", "end_date", "created", "updated"
) values (
  '20000000-0000-0000-0000-000000000000', 'rituals-sprint-1', 'Rituals Sprint 1', 'star', 'active', '10000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', '2023-01-01', '2023-02-01', now(), null
) on conflict do nothing;
-- `)
//line queries/seeddata/seed_sprint.sql:7
}

//line queries/seeddata/seed_sprint.sql:7
func WriteSprintSeedData(qq422016 qtio422016.Writer) {
//line queries/seeddata/seed_sprint.sql:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/seeddata/seed_sprint.sql:7
	StreamSprintSeedData(qw422016)
//line queries/seeddata/seed_sprint.sql:7
	qt422016.ReleaseWriter(qw422016)
//line queries/seeddata/seed_sprint.sql:7
}

//line queries/seeddata/seed_sprint.sql:7
func SprintSeedData() string {
//line queries/seeddata/seed_sprint.sql:7
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/seeddata/seed_sprint.sql:7
	WriteSprintSeedData(qb422016)
//line queries/seeddata/seed_sprint.sql:7
	qs422016 := string(qb422016.B)
//line queries/seeddata/seed_sprint.sql:7
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/seeddata/seed_sprint.sql:7
	return qs422016
//line queries/seeddata/seed_sprint.sql:7
}
