-- {% func SprintMemberSeedData() %}
insert into "sprint_member" (
  "sprint_id", "user_id", "name", "picture", "role", "created", "updated"
) values (
  '20000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', 'Test User', '/assets/logo.png', 'owner', now(), null
), (
  '20000000-0000-0000-0000-000000000000', '90000001-0000-0000-0000-000000000000', 'Test User 2', '/assets/logo.png', 'member', now(), null
) on conflict do nothing;
-- {% endfunc %}
