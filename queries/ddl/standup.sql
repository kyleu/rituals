-- {% func StandupDrop() %}
drop table if exists "standup";
-- {% endfunc %}

-- {% func StandupCreate() %}
create table if not exists "standup" (
  "id" uuid not null,
  "slug" text not null,
  "title" text not null,
  "icon" text not null,
  "status" session_status not null,
  "team_id" uuid,
  "sprint_id" uuid,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("team_id") references "team" ("id"),
  foreign key ("sprint_id") references "sprint" ("id"),
  unique ("slug"),
  primary key ("id")
);

create index if not exists "standup__slug_idx" on "standup" ("slug");

create index if not exists "standup__status_idx" on "standup" ("status");

create index if not exists "standup__team_id_idx" on "standup" ("team_id");

create index if not exists "standup__sprint_id_idx" on "standup" ("sprint_id");
-- {% endfunc %}
