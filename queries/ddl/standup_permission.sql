-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func StandupPermissionDrop() %}
drop table if exists "standup_permission";
-- {% endfunc %}

-- {% func StandupPermissionCreate() %}
create table if not exists "standup_permission" (
  "standup_id" uuid not null,
  "k" text not null,
  "v" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("standup_id") references "standup" ("id"),
  primary key ("standup_id", "k", "v")
);

create index if not exists "standup_permission__standup_id_idx" on "standup_permission" ("standup_id");

create index if not exists "standup_permission__k_idx" on "standup_permission" ("k");

create index if not exists "standup_permission__v_idx" on "standup_permission" ("v");
-- {% endfunc %}
