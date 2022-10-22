-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func RetroMemberDrop() %}
drop table if exists "retro_member";
-- {% endfunc %}

-- {% func RetroMemberCreate() %}
create table if not exists "retro_member" (
  "retro_id" uuid not null,
  "user_id" uuid not null,
  "name" text not null,
  "picture" text not null,
  "role" text not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("retro_id") references "retro" ("id"),
  foreign key ("user_id") references "user" ("id"),
  primary key ("retro_id", "user_id")
);

create index if not exists "retro_member__retro_id_idx" on "retro_member" ("retro_id");

create index if not exists "retro_member__user_id_idx" on "retro_member" ("user_id");
-- {% endfunc %}
