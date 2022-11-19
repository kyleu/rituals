-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func RetroPermissionDrop() %}
drop table if exists "retro_permission";
-- {% endfunc %}

-- {% func RetroPermissionCreate() %}
create table if not exists "retro_permission" (
  "retro_id" uuid not null,
  "key" text not null,
  "value" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("retro_id") references "retro" ("id"),
  primary key ("retro_id", "key", "value")
);

create index if not exists "retro_permission__retro_id_idx" on "retro_permission" ("retro_id");

create index if not exists "retro_permission__key_idx" on "retro_permission" ("key");

create index if not exists "retro_permission__value_idx" on "retro_permission" ("value");
-- {% endfunc %}
