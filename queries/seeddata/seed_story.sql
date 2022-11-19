-- {% func StorySeedData() %}
insert into "story" (
  "id", "estimate_id", "idx", "user_id", "title", "status", "final_vote", "created", "updated"
) values (
  '31000000-0000-0000-0000-000000000000', '30000000-0000-0000-0000-000000000000', 0, '90000000-0000-0000-0000-000000000000', 'Build rituals.dev', 'new', '100', now(), null
), (
  '31000001-0000-0000-0000-000000000000', '30000000-0000-0000-0000-000000000000', 0, '90000001-0000-0000-0000-000000000000', 'Make it work without JavaScript', 'new', '', now(), null
) on conflict do nothing;
-- {% endfunc %}
