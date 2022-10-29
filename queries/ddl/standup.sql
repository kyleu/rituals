-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func StandupDrop() %}
drop table if exists "standup";
-- {% endfunc %}

-- {% func StandupCreate() %}
create table if not exists "standup" (
  "id" uuid not null,
  "slug" text not null,
  "title" text not null,
  "status" session_status not null,
  "team_id" uuid,
  "sprint_id" uuid,
  "owner" uuid not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("owner") references "user" ("id"),
  foreign key ("team_id") references "team" ("id"),
  foreign key ("sprint_id") references "sprint" ("id"),
  unique ("slug"),
  primary key ("id")
);

create index if not exists "standup__slug_idx" on "standup" ("slug");

create index if not exists "standup__status_idx" on "standup" ("status");

create index if not exists "standup__owner_idx" on "standup" ("owner");

create index if not exists "standup__team_id_idx" on "standup" ("team_id");

create index if not exists "standup__sprint_id_idx" on "standup" ("sprint_id");
-- {% endfunc %}
