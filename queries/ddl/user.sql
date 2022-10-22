-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func UserDrop() %}
drop table if exists "user";
-- {% endfunc %}

-- {% func UserCreate() %}
create table if not exists "user" (
  "id" uuid not null,
  "name" text not null,
  "role" text not null,
  "picture" text not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  primary key ("id")
);
-- {% endfunc %}
