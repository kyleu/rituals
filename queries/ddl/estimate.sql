-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func EstimateDrop() %}
drop table if exists "estimate";
-- {% endfunc %}

-- {% func EstimateCreate() %}
create table if not exists "estimate" (
  "id" uuid not null,
  "slug" text not null,
  "title" text not null,
  "status" session_status not null,
  "team_id" uuid,
  "sprint_id" uuid,
  "owner" uuid not null,
  "choices" jsonb not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("owner") references "user" ("id"),
  foreign key ("team_id") references "team" ("id"),
  foreign key ("sprint_id") references "sprint" ("id"),
  primary key ("id")
);

create index if not exists "estimate__owner_idx" on "estimate" ("owner");

create index if not exists "estimate__team_id_idx" on "estimate" ("team_id");

create index if not exists "estimate__sprint_id_idx" on "estimate" ("sprint_id");
-- {% endfunc %}
