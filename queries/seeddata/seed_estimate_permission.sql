-- {% func EstimatePermissionSeedData() %}
insert into "estimate_permission" (
  "estimate_id", "key", "value", "access", "created"
) values (
  '30000000-0000-0000-0000-000000000000', 'github', 'kyleu.com', 'member', now()
) on conflict do nothing;
-- {% endfunc %}
