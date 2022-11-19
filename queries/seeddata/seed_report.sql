-- {% func ReportSeedData() %}
insert into "report" (
  "id", "standup_id", "day", "user_id", "content", "html", "created", "updated"
) values (
  '41000000-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2022-10-31', '90000000-0000-0000-0000-000000000000', 'A report!', '<em>A Report!</em>', now(), null
), (
  '41000001-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2022-10-31', '90000001-0000-0000-0000-000000000000', 'A second report!', '<strong>A Report!</strong>', now(), null
) on conflict do nothing;
-- {% endfunc %}
