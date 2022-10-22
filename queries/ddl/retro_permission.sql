-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func RetroPermissionDrop() %}
drop table if exists "retro_permission";
-- {% endfunc %}

-- {% func RetroPermissionCreate() %}
create table if not exists "retro_permission" (
  "retro_id" uuid not null,
  "k" text not null,
  "v" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("retro_id") references "retro" ("id"),
  primary key ("retro_id", "k", "v")
);

create index if not exists "retro_permission__retro_id_idx" on "retro_permission" ("retro_id");

create index if not exists "retro_permission__k_idx" on "retro_permission" ("k");

create index if not exists "retro_permission__v_idx" on "retro_permission" ("v");
-- {% endfunc %}
