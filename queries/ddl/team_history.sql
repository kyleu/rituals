-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func TeamHistoryDrop() %}
drop table if exists "team_history";
-- {% endfunc %}

-- {% func TeamHistoryCreate() %}
create table if not exists "team_history" (
  "slug" text not null,
  "team_id" uuid not null,
  "team_name" text not null,
  "created" timestamp not null default now(),
  foreign key ("team_id") references "team" ("id"),
  primary key ("slug")
);

create index if not exists "team_history__team_id_idx" on "team_history" ("team_id");
-- {% endfunc %}
