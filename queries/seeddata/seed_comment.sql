-- {% func CommentSeedData() %}
insert into "comment" (
  "id", "svc", "model_id", "target_type", "target_id", "user_id", "content", "html", "created"
) values (
  '11000000-0000-0000-0000-000000000000', 'sprint', '20000000-0000-0000-0000-000000000000', 'story', '13000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', 'Hello!', '<h1>Hello!</h1>', now()
) on conflict do nothing;
-- {% endfunc %}
