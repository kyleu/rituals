-- {% func SprintSeedData() %}
insert into "sprint" (
  "id", "slug", "title", "icon", "status", "team_id", "owner", "start_date", "end_date", "created", "updated"
) values (
  '20000000-0000-0000-0000-000000000000', 'rituals-sprint-1', 'Rituals Sprint 1', 'star', 'active', '10000000-0000-0000-0000-000000000000', '90000000-0000-0000-0000-000000000000', '2023-01-01', '2023-02-01', now(), null
) on conflict do nothing;
-- {% endfunc %}
