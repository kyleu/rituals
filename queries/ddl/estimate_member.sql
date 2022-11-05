-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func EstimateMemberDrop() %}
drop table if exists "estimate_member";
-- {% endfunc %}

-- {% func EstimateMemberCreate() %}
create table if not exists "estimate_member" (
  "estimate_id" uuid not null,
  "user_id" uuid not null,
  "name" text not null,
  "picture" text not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("estimate_id") references "estimate" ("id"),
  foreign key ("user_id") references "user" ("id"),
  primary key ("estimate_id", "user_id")
);

create index if not exists "estimate_member__estimate_id_idx" on "estimate_member" ("estimate_id");

create index if not exists "estimate_member__user_id_idx" on "estimate_member" ("user_id");
-- {% endfunc %}
