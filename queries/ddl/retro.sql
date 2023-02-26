-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func RetroDrop() %}
drop table if exists "retro";
-- {% endfunc %}

-- {% func RetroCreate() %}
create table if not exists "retro" (
  "id" uuid not null,
  "slug" text not null,
  "title" text not null,
  "icon" text not null,
  "status" session_status not null,
  "team_id" uuid,
  "sprint_id" uuid,
  "categories" jsonb not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("team_id") references "team" ("id"),
  foreign key ("sprint_id") references "sprint" ("id"),
  unique ("slug"),
  primary key ("id")
);

create index if not exists "retro__slug_idx" on "retro" ("slug");

create index if not exists "retro__status_idx" on "retro" ("status");

create index if not exists "retro__team_id_idx" on "retro" ("team_id");

create index if not exists "retro__sprint_id_idx" on "retro" ("sprint_id");
-- {% endfunc %}
