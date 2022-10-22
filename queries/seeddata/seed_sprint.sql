-- {% func SprintSeedData() %}
insert into "sprint" (
  "id", "slug", "title", "status", "team_id", "owner", "created", "updated"
) values (
  '20000000-0000-0000-0000-000000000000', 'team-a', 'Team A', 'new', '10000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', now(), null
) on conflict do nothing;
-- {% endfunc %}
