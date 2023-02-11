-- {% func TeamPermissionSeedData() %}
insert into "team_permission" (
  "team_id", "key", "value", "access", "created"
) values (
  '10000000-0000-0000-0000-000000000000', 'github', 'kyleu.com', 'member', now()
) on conflict do nothing;
-- {% endfunc %}
