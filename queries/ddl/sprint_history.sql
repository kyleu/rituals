-- {% func SprintHistoryDrop() %}
drop table if exists "sprint_history";
-- {% endfunc %}

-- {% func SprintHistoryCreate() %}
create table if not exists "sprint_history" (
  "slug" text not null,
  "sprint_id" uuid not null,
  "sprint_name" text not null,
  "created" timestamp not null default now(),
  foreign key ("sprint_id") references "sprint" ("id"),
  primary key ("slug")
);

create index if not exists "sprint_history__sprint_id_idx" on "sprint_history" ("sprint_id");
-- {% endfunc %}
