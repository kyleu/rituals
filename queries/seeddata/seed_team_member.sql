-- {% func TeamMemberSeedData() %}
insert into "team_member" (
  "team_id", "user_id", "name", "picture", "role", "created", "updated"
) values (
  '10000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', 'Member', 'https://google.com', 'owner', now(), null
) on conflict do nothing;
-- {% endfunc %}
