-- {% func EstimateHistorySeedData() %}
insert into "estimate_history" (
  "slug", "estimate_id", "estimate_name", "created"
) values (
  'old-name', '30000000-0000-0000-0000-000000000000', 'Old Name', now()
) on conflict do nothing;
-- {% endfunc %}
