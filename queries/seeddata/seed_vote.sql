-- {% func VoteSeedData() %}
insert into "vote" (
  "story_id", "user_id", "choice", "created", "updated"
) values (
  '31000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', '100', now(), null
) on conflict do nothing;
-- {% endfunc %}
