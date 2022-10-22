-- {% func RetroSeedData() %}
insert into "retro" (
  "id", "slug", "title", "status", "team_id", "sprint_id", "owner", "categories", "created", "updated"
) values (
  '50000000-0000-0000-0000-000000000000', 'retro-a', 'Retro A', 'new', '10000000-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', '["good","bad"]', now(), null
) on conflict do nothing;
-- {% endfunc %}
