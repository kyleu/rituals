-- {% func SprintHistorySeedData() %}
insert into "sprint_history" (
  "slug", "sprint_id", "sprint_name", "created"
) values (
  'old-name', '20000000-0000-0000-0000-000000000000', 'Old Name', now()
) on conflict do nothing;
-- {% endfunc %}
