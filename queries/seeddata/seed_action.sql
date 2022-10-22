-- {% func ActionSeedData() %}
insert into "action" (
  "id", "svc", "model_id", "user_id", "act", "content", "note", "created"
) values (
  '12000000-0000-0000-0000-000000000000', 'sprint', '20000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', 'create-stuff', '{}', 'A note!', now()
) on conflict do nothing;
-- {% endfunc %}
