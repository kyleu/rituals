-- {% func TeamHistorySeedData() %}
insert into "team_history" (
  "slug", "team_id", "team_name", "created"
) values (
  'old-name', '10000000-0000-0000-0000-000000000000', 'Old Name', now()
) on conflict do nothing;
-- {% endfunc %}
