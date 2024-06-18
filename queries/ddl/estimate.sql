-- {% func EstimateDrop() %}
drop table if exists "estimate";
-- {% endfunc %}

-- {% func EstimateCreate() %}
create table if not exists "estimate" (
  "id" uuid not null,
  "slug" text not null,
  "title" text not null,
  "icon" text not null,
  "status" session_status not null,
  "team_id" uuid,
  "sprint_id" uuid,
  "choices" jsonb not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("team_id") references "team" ("id"),
  foreign key ("sprint_id") references "sprint" ("id"),
  unique ("slug"),
  primary key ("id")
);

create index if not exists "estimate__slug_idx" on "estimate" ("slug");

create index if not exists "estimate__status_idx" on "estimate" ("status");

create index if not exists "estimate__team_id_idx" on "estimate" ("team_id");

create index if not exists "estimate__sprint_id_idx" on "estimate" ("sprint_id");
-- {% endfunc %}
