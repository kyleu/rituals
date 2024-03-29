// Code generated by qtc from "seed_report.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/seeddata/seed_report.sql:1
package seeddata

//line queries/seeddata/seed_report.sql:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/seeddata/seed_report.sql:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/seeddata/seed_report.sql:1
func StreamReportSeedData(qw422016 *qt422016.Writer) {
//line queries/seeddata/seed_report.sql:1
	qw422016.N().S(`
insert into "report" (
  "id", "standup_id", "day", "user_id", "content", "html", "created", "updated"
) values (
  '41000000-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2022-10-31', '90000000-0000-0000-0000-000000000000', 'A report!', '<em>A Report!</em>', now(), null
), (
  '41000001-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2022-10-31', '90000001-0000-0000-0000-000000000000', 'A second report!', '<strong>A Report!</strong>', now(), null
) on conflict do nothing;
-- `)
//line queries/seeddata/seed_report.sql:9
}

//line queries/seeddata/seed_report.sql:9
func WriteReportSeedData(qq422016 qtio422016.Writer) {
//line queries/seeddata/seed_report.sql:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/seeddata/seed_report.sql:9
	StreamReportSeedData(qw422016)
//line queries/seeddata/seed_report.sql:9
	qt422016.ReleaseWriter(qw422016)
//line queries/seeddata/seed_report.sql:9
}

//line queries/seeddata/seed_report.sql:9
func ReportSeedData() string {
//line queries/seeddata/seed_report.sql:9
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/seeddata/seed_report.sql:9
	WriteReportSeedData(qb422016)
//line queries/seeddata/seed_report.sql:9
	qs422016 := string(qb422016.B)
//line queries/seeddata/seed_report.sql:9
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/seeddata/seed_report.sql:9
	return qs422016
//line queries/seeddata/seed_report.sql:9
}
