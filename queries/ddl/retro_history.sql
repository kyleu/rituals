-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func RetroHistoryDrop() %}
drop table if exists "retro_history";
-- {% endfunc %}

-- {% func RetroHistoryCreate() %}
create table if not exists "retro_history" (
  "slug" text not null,
  "retro_id" uuid not null,
  "retro_name" text not null,
  "created" timestamp not null default now(),
  foreign key ("retro_id") references "retro" ("id"),
  primary key ("slug")
);

create index if not exists "retro_history__retro_id_idx" on "retro_history" ("retro_id");
-- {% endfunc %}
