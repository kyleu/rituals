-- {% func UserSeedData() %}
insert into "user" (
  "id", "name", "picture", "created", "updated"
) values (
  '90000000-0000-0000-0000-000000000000', 'Test User', 'https://electricfrankfurter.com/index.png', now(), null
) on conflict do nothing;
-- {% endfunc %}
