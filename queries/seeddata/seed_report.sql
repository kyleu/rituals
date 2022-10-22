-- {% func ReportSeedData() %}
insert into "report" (
  "id", "standup_id", "d", "user_id", "content", "html", "created", "updated"
) values (
  '41000000-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2022-10-31', '90000000-0000-0000-0000-000000000000', 'A Report!', '<h1>A Report!</h1>', now(), null
) on conflict do nothing;
-- {% endfunc %}
