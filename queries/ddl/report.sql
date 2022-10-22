-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func ReportDrop() %}
drop table if exists "report";
-- {% endfunc %}

-- {% func ReportCreate() %}
create table if not exists "report" (
  "id" uuid not null,
  "standup_id" uuid not null,
  "d" timestamp not null,
  "user_id" uuid not null,
  "content" text not null,
  "html" text not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("standup_id") references "standup" ("id"),
  foreign key ("user_id") references "user" ("id"),
  primary key ("id")
);

create index if not exists "report__standup_id_idx" on "report" ("standup_id");

create index if not exists "report__user_id_idx" on "report" ("user_id");
-- {% endfunc %}
