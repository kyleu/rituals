-- {% func CommentSeedData() %}
insert into "comment" (
  "id", "svc", "model_id", "user_id", "content", "html", "created"
) values (
  '11000000-0000-0000-0000-000000000000', 'sprint', '20000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', 'Sprint!', '<h1>Sprint!</h1>', now()
), (
  '11000001-0000-0000-0000-000000000000', 'estimate', '30000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', 'Estimate!', '<h1>Estimate!</h1>', now()
) on conflict do nothing;
-- {% endfunc %}
