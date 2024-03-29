// Code generated by qtc from "seed_team_history.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/seeddata/seed_team_history.sql:1
package seeddata

//line queries/seeddata/seed_team_history.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/seeddata/seed_team_history.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/seeddata/seed_team_history.sql:1
func StreamTeamHistorySeedData(qw422016 *qt422016.Writer) {
//line queries/seeddata/seed_team_history.sql:1
	qw422016.N().S(`
insert into "team_history" (
  "slug", "team_id", "team_name", "created"
) values (
  'old-name', '10000000-0000-0000-0000-000000000000', 'Old Name', now()
) on conflict do nothing;
-- `)
//line queries/seeddata/seed_team_history.sql:7
}

//line queries/seeddata/seed_team_history.sql:7
func WriteTeamHistorySeedData(qq422016 qtio422016.Writer) {
//line queries/seeddata/seed_team_history.sql:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/seeddata/seed_team_history.sql:7
	StreamTeamHistorySeedData(qw422016)
//line queries/seeddata/seed_team_history.sql:7
	qt422016.ReleaseWriter(qw422016)
//line queries/seeddata/seed_team_history.sql:7
}

//line queries/seeddata/seed_team_history.sql:7
func TeamHistorySeedData() string {
//line queries/seeddata/seed_team_history.sql:7
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/seeddata/seed_team_history.sql:7
	WriteTeamHistorySeedData(qb422016)
//line queries/seeddata/seed_team_history.sql:7
	qs422016 := string(qb422016.B)
//line queries/seeddata/seed_team_history.sql:7
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/seeddata/seed_team_history.sql:7
	return qs422016
//line queries/seeddata/seed_team_history.sql:7
}
