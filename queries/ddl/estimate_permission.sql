-- {% func EstimatePermissionDrop() %}
drop table if exists "estimate_permission";
-- {% endfunc %}

-- {% func EstimatePermissionCreate() %}
create table if not exists "estimate_permission" (
  "estimate_id" uuid not null,
  "key" text not null,
  "value" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("estimate_id") references "estimate" ("id"),
  primary key ("estimate_id", "key", "value")
);

create index if not exists "estimate_permission__estimate_id_idx" on "estimate_permission" ("estimate_id");

create index if not exists "estimate_permission__key_idx" on "estimate_permission" ("key");

create index if not exists "estimate_permission__value_idx" on "estimate_permission" ("value");
-- {% endfunc %}
