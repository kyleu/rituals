-- {% func SprintPermissionSeedData() %}
insert into "sprint_permission" (
  "sprint_id", "key", "value", "access", "created"
) values (
  '20000000-0000-0000-0000-000000000000', 'github', 'kyleu.com', 'member', now()
) on conflict do nothing;
-- {% endfunc %}
