-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func EstimatePermissionDrop() %}
drop table if exists "estimate_permission";
-- {% endfunc %}

-- {% func EstimatePermissionCreate() %}
create table if not exists "estimate_permission" (
  "estimate_id" uuid not null,
  "k" text not null,
  "v" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("estimate_id") references "estimate" ("id"),
  primary key ("estimate_id", "k", "v")
);

create index if not exists "estimate_permission__estimate_id_idx" on "estimate_permission" ("estimate_id");

create index if not exists "estimate_permission__k_idx" on "estimate_permission" ("k");

create index if not exists "estimate_permission__v_idx" on "estimate_permission" ("v");
-- {% endfunc %}
