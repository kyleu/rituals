-- {% func SprintPermissionSeedData() %}
insert into "sprint_permission" (
  "sprint_id", "k", "v", "access", "created"
) values (
  '20000000-0000-0000-0000-000000000000', 'key', 'value', 'access', now()
) on conflict do nothing;
-- {% endfunc %}
