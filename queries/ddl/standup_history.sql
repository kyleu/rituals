-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func StandupHistoryDrop() %}
drop table if exists "standup_history";
-- {% endfunc %}

-- {% func StandupHistoryCreate() %}
create table if not exists "standup_history" (
  "slug" text not null,
  "standup_id" uuid not null,
  "standup_name" text not null,
  "created" timestamp not null default now(),
  foreign key ("standup_id") references "standup" ("id"),
  primary key ("slug")
);

create index if not exists "standup_history__standup_id_idx" on "standup_history" ("standup_id");
-- {% endfunc %}
