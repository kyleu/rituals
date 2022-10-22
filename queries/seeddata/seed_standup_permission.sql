-- {% func StandupPermissionSeedData() %}
insert into "standup_permission" (
  "standup_id", "k", "v", "access", "created"
) values (
  '40000000-0000-0000-0000-000000000000', 'key', 'value', 'access', now()
) on conflict do nothing;
-- {% endfunc %}
