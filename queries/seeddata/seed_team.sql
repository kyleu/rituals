-- {% func TeamSeedData() %}
insert into "team" (
  "id", "slug", "title", "icon", "status", "created", "updated"
) values (
  '10000000-0000-0000-0000-000000000000', 'rituals-team', 'Rituals Team', 'star', 'active', now(), null
), (
  '10000001-0000-0000-0000-000000000000', 'team-2', 'Team 2', 'action', 'active', now(), null
) on conflict do nothing;
-- {% endfunc %}
