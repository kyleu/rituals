-- {% func EstimatePermissionSeedData() %}
insert into "estimate_permission" (
  "estimate_id", "k", "v", "access", "created"
) values (
  '30000000-0000-0000-0000-000000000000', 'key', 'value', 'access', now()
) on conflict do nothing;
-- {% endfunc %}
