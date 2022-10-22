-- {% func RetroPermissionSeedData() %}
insert into "retro_permission" (
  "retro_id", "k", "v", "access", "created"
) values (
  '50000000-0000-0000-0000-000000000000', 'key', 'value', 'access', now()
) on conflict do nothing;
-- {% endfunc %}
