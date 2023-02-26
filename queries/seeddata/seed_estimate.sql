-- {% func EstimateSeedData() %}
insert into "estimate" (
  "id", "slug", "title", "icon", "status", "team_id", "sprint_id", "choices", "created", "updated"
) values (
  '30000000-0000-0000-0000-000000000000', 'estimate-1', 'Estimate 1', 'star', 'active', '10000000-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '["0","1","2","3","5","8","13","100"]', now(), null
), (
  '30000001-0000-0000-0000-000000000000', 'estimate-2', 'Estimate 2', 'dot-circle', 'active', '10000001-0000-0000-0000-000000000000', '20000001-0000-0000-0000-000000000000', '["0","1","2","3","5","8","13","100"]', now(), null
) on conflict do nothing;
-- {% endfunc %}
