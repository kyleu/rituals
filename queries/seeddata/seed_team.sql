-- {% func TeamSeedData() %}
insert into "team" (
  "id", "slug", "title", "status", "owner", "created", "updated"
) values (
  '10000000-0000-0000-0000-000000000000', 'team-a', 'Team A', 'new', '90000000-0000-0000-0000-000000000000', now(), null
) on conflict do nothing;
-- {% endfunc %}
