-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func TeamPermissionDrop() %}
drop table if exists "team_permission";
-- {% endfunc %}

-- {% func TeamPermissionCreate() %}
create table if not exists "team_permission" (
  "team_id" uuid not null,
  "k" text not null,
  "v" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("team_id") references "team" ("id"),
  primary key ("team_id", "k", "v")
);

create index if not exists "team_permission__team_id_idx" on "team_permission" ("team_id");

create index if not exists "team_permission__k_idx" on "team_permission" ("k");

create index if not exists "team_permission__v_idx" on "team_permission" ("v");
-- {% endfunc %}
