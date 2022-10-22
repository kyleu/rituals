-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func StandupMemberDrop() %}
drop table if exists "standup_member";
-- {% endfunc %}

-- {% func StandupMemberCreate() %}
create table if not exists "standup_member" (
  "standup_id" uuid not null,
  "user_id" uuid not null,
  "name" text not null,
  "picture" text not null,
  "role" text not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("standup_id") references "standup" ("id"),
  foreign key ("user_id") references "user" ("id"),
  primary key ("standup_id", "user_id")
);

create index if not exists "standup_member__standup_id_idx" on "standup_member" ("standup_id");

create index if not exists "standup_member__user_id_idx" on "standup_member" ("user_id");
-- {% endfunc %}
