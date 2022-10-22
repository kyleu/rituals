-- {% func RetroHistorySeedData() %}
insert into "retro_history" (
  "slug", "retro_id", "retro_name", "created"
) values (
  'old-name', '50000000-0000-0000-0000-000000000000', 'Old Name', now()
) on conflict do nothing;
-- {% endfunc %}
