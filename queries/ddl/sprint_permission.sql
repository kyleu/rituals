-- {% func SprintPermissionDrop() %}
drop table if exists "sprint_permission";
-- {% endfunc %}

-- {% func SprintPermissionCreate() %}
create table if not exists "sprint_permission" (
  "sprint_id" uuid not null,
  "key" text not null,
  "value" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("sprint_id") references "sprint" ("id"),
  primary key ("sprint_id", "key", "value")
);

create index if not exists "sprint_permission__sprint_id_idx" on "sprint_permission" ("sprint_id");

create index if not exists "sprint_permission__key_idx" on "sprint_permission" ("key");

create index if not exists "sprint_permission__value_idx" on "sprint_permission" ("value");
-- {% endfunc %}
