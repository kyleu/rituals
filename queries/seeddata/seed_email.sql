-- {% func EmailSeedData() %}
insert into "email" (
  "id", "recipients", "subject", "data", "plain", "html", "user_id", "status", "created"
) values (
  '12000000-0000-0000-0000-000000000000', '["a","b","c"]', 'An Email!', '"{\"x\": 1}"', 'Hello!', '<h1>Hello!</h1>', '90000000-0000-0000-0000-000000000000', 'sent', now()
) on conflict do nothing;
-- {% endfunc %}
