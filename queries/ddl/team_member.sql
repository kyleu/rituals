-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func TeamMemberDrop() %}
drop table if exists "team_member";
-- {% endfunc %}

-- {% func TeamMemberCreate() %}
create table if not exists "team_member" (
  "team_id" uuid not null,
  "user_id" uuid not null,
  "name" text not null,
  "picture" text not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("team_id") references "team" ("id"),
  foreign key ("user_id") references "user" ("id"),
  primary key ("team_id", "user_id")
);

create index if not exists "team_member__team_id_idx" on "team_member" ("team_id");

create index if not exists "team_member__user_id_idx" on "team_member" ("user_id");
-- {% endfunc %}
