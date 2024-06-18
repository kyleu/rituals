-- {% func TeamPermissionDrop() %}
drop table if exists "team_permission";
-- {% endfunc %}

-- {% func TeamPermissionCreate() %}
create table if not exists "team_permission" (
  "team_id" uuid not null,
  "key" text not null,
  "value" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("team_id") references "team" ("id"),
  primary key ("team_id", "key", "value")
);

create index if not exists "team_permission__team_id_idx" on "team_permission" ("team_id");

create index if not exists "team_permission__key_idx" on "team_permission" ("key");

create index if not exists "team_permission__value_idx" on "team_permission" ("value");
-- {% endfunc %}
