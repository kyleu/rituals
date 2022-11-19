-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func TeamDrop() %}
drop table if exists "team";
-- {% endfunc %}

-- {% func TeamCreate() %}
create table if not exists "team" (
  "id" uuid not null,
  "slug" text not null,
  "title" text not null,
  "icon" text not null,
  "status" session_status not null,
  "owner" uuid not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("owner") references "user" ("id"),
  unique ("slug"),
  primary key ("id")
);

create index if not exists "team__slug_idx" on "team" ("slug");

create index if not exists "team__status_idx" on "team" ("status");

create index if not exists "team__owner_idx" on "team" ("owner");
-- {% endfunc %}
