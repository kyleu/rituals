-- {% func FeedbackSeedData() %}
insert into "feedback" (
  "id", "retro_id", "idx", "user_id", "category", "content", "html", "created", "updated"
) values (
  '51000000-0000-0000-0000-000000000000', '50000000-0000-0000-0000-000000000000', 0, '90000000-0000-0000-0000-000000000000', 'Category', 'Content', 'HTML', now(), null
) on conflict do nothing;
-- {% endfunc %}
