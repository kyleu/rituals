-- {% func SprintSeedData() %}
insert into "sprint" (
  "id", "slug", "title", "status", "team_id", "owner", "start_date", "end_date", "created", "updated"
) values (
  '20000000-0000-0000-0000-000000000000', 'sprint-a', 'Sprint A', 'new', '10000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', '2022-01-01', '2022-02-01', now(), null
) on conflict do nothing;
-- {% endfunc %}
