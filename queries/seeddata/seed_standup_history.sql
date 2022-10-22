-- {% func StandupHistorySeedData() %}
insert into "standup_history" (
  "slug", "standup_id", "standup_name", "created"
) values (
  'old-name', '40000000-0000-0000-0000-000000000000', 'Old Name', now()
) on conflict do nothing;
-- {% endfunc %}
