-- {% func StandupSeedData() %}
insert into "standup" (
  "id", "slug", "title", "status", "team_id", "sprint_id", "owner", "created", "updated"
) values (
  '40000000-0000-0000-0000-000000000000', 'standup-a', 'Standup A', 'new', '10000000-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', now(), null
) on conflict do nothing;
-- {% endfunc %}
