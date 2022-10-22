-- {% func UserSeedData() %}
insert into "user" (
  "id", "name", "role", "picture", "created", "updated"
) values (
  '90000000-0000-0000-0000-000000000000', 'Kyle', 'admin', 'https://google.com', now(), null
) on conflict do nothing;
-- {% endfunc %}
