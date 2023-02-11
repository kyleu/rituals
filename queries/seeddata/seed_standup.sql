-- {% func StandupSeedData() %}
insert into "standup" (
  "id", "slug", "title", "icon", "status", "team_id", "sprint_id", "owner", "created", "updated"
) values (
  '40000000-0000-0000-0000-000000000000', 'standup-1', 'Standup 1', 'star', 'active', '10000000-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', now(), null
), (
  '40000001-0000-0000-0000-000000000000', 'standup-2', 'Standup 2', 'apple', 'active', '10000001-0000-0000-0000-000000000000', '20000001-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', now(), null
) on conflict do nothing;
-- {% endfunc %}
