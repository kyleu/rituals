-- {% func SprintMemberSeedData() %}
insert into "sprint_member" (
  "sprint_id", "user_id", "name", "picture", "role", "created", "updated"
) values (
  '20000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', 'Member', 'https://google.com', 'owner', now(), null
) on conflict do nothing;
-- {% endfunc %}
