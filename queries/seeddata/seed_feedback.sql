-- {% func FeedbackSeedData() %}
insert into "feedback" (
  "id", "retro_id", "idx", "user_id", "category", "content", "html", "created", "updated"
) values (
  '51000000-0000-0000-0000-000000000000', '50000000-0000-0000-0000-000000000000', 0, '90000000-0000-0000-0000-000000000000', 'good', 'First feedback', '<em>First feedback</em>', now(), null
), (
  '51000001-0000-0000-0000-000000000000', '50000000-0000-0000-0000-000000000000', 0, '90000000-0000-0000-0000-000000000000', 'bad', 'Second feedback', '<em>Second feedback</em>', now(), null
), (
  '51000002-0000-0000-0000-000000000000', '50000000-0000-0000-0000-000000000000', 0, '90000001-0000-0000-0000-000000000000', 'extra', 'Third feedback', '<em>Third feedback</em>', now(), null
) on conflict do nothing;
-- {% endfunc %}
