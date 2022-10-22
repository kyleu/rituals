-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func SprintMemberDrop() %}
drop table if exists "sprint_member";
-- {% endfunc %}

-- {% func SprintMemberCreate() %}
create table if not exists "sprint_member" (
  "sprint_id" uuid not null,
  "user_id" uuid not null,
  "name" text not null,
  "picture" text not null,
  "role" text not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("sprint_id") references "sprint" ("id"),
  foreign key ("user_id") references "user" ("id"),
  primary key ("sprint_id", "user_id")
);

create index if not exists "sprint_member__sprint_id_idx" on "sprint_member" ("sprint_id");

create index if not exists "sprint_member__user_id_idx" on "sprint_member" ("user_id");
-- {% endfunc %}
