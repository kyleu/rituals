-- {% func StandupPermissionSeedData() %}
insert into "standup_permission" (
  "standup_id", "key", "value", "access", "created"
) values (
  '40000000-0000-0000-0000-000000000000', 'github', 'kyleu.com', 'member', now()
), (
  '40000000-0000-0000-0000-000000000000', 'team', '10000000-0000-0000-0000-000000000000', 'member', now()
), (
  '40000000-0000-0000-0000-000000000000', 'sprint', '20000000-0000-0000-0000-000000000000', 'member', now()
) on conflict do nothing;
-- {% endfunc %}
