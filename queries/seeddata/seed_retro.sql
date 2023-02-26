-- {% func RetroSeedData() %}
insert into "retro" (
  "id", "slug", "title", "icon", "status", "team_id", "sprint_id", "categories", "created", "updated"
) values (
  '50000000-0000-0000-0000-000000000000', 'retro-1', 'Retro 1', 'star', 'active', '10000000-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '["good","bad"]', now(), null
), (
  '50000001-0000-0000-0000-000000000000', 'retro-2', 'Retro 2', 'bolt', 'active', '10000001-0000-0000-0000-000000000000', '20000001-0000-0000-0000-000000000000', '["good","bad"]', now(), null
) on conflict do nothing;
-- {% endfunc %}
