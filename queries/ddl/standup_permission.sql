-- {% func StandupPermissionDrop() %}
drop table if exists "standup_permission";
-- {% endfunc %}

-- {% func StandupPermissionCreate() %}
create table if not exists "standup_permission" (
  "standup_id" uuid not null,
  "key" text not null,
  "value" text not null,
  "access" text not null,
  "created" timestamp not null default now(),
  foreign key ("standup_id") references "standup" ("id"),
  primary key ("standup_id", "key", "value")
);

create index if not exists "standup_permission__standup_id_idx" on "standup_permission" ("standup_id");

create index if not exists "standup_permission__key_idx" on "standup_permission" ("key");

create index if not exists "standup_permission__value_idx" on "standup_permission" ("value");
-- {% endfunc %}
