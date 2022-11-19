-- {% func RetroMemberSeedData() %}
insert into "retro_member" (
  "retro_id", "user_id", "name", "picture", "role", "created", "updated"
) values (
  '50000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', 'Test User', 'https://google.com', 'owner', now(), null
), (
  '50000000-0000-0000-0000-000000000000', '90000001-0000-0000-0000-000000000000', 'Test User 2', 'https://google.com', 'member', now(), null
) on conflict do nothing;
-- {% endfunc %}
