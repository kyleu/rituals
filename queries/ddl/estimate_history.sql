-- {% func EstimateHistoryDrop() %}
drop table if exists "estimate_history";
-- {% endfunc %}

-- {% func EstimateHistoryCreate() %}
create table if not exists "estimate_history" (
  "slug" text not null,
  "estimate_id" uuid not null,
  "estimate_name" text not null,
  "created" timestamp not null default now(),
  foreign key ("estimate_id") references "estimate" ("id"),
  primary key ("slug")
);

create index if not exists "estimate_history__estimate_id_idx" on "estimate_history" ("estimate_id");
-- {% endfunc %}
