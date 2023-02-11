-- {% func RetroPermissionSeedData() %}
insert into "retro_permission" (
  "retro_id", "key", "value", "access", "created"
) values (
  '50000000-0000-0000-0000-000000000000', 'github', 'kyleu.com', 'member', now()
), (
  '50000000-0000-0000-0000-000000000000', 'team', '10000000-0000-0000-0000-000000000000', 'member', now()
), (
  '50000000-0000-0000-0000-000000000000', 'sprint', '20000000-0000-0000-0000-000000000000', 'member', now()
) on conflict do nothing;
-- {% endfunc %}
