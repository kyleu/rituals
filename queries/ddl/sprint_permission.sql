-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func SprintPermissionDrop() %}
drop table if exists "sprint_permission";
-- {% endfunc %}

-- {% func SprintPermissionCreate() %}
create table if not exists "sprint_permission" (
  "sprint_id" uuid not null,
  "k" text not null,
  "v" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("sprint_id") references "sprint" ("id"),
  primary key ("sprint_id", "k", "v")
);

create index if not exists "sprint_permission__sprint_id_idx" on "sprint_permission" ("sprint_id");

create index if not exists "sprint_permission__k_idx" on "sprint_permission" ("k");

create index if not exists "sprint_permission__v_idx" on "sprint_permission" ("v");
-- {% endfunc %}
