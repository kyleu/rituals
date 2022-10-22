-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func FeedbackDrop() %}
drop table if exists "feedback";
-- {% endfunc %}

-- {% func FeedbackCreate() %}
create table if not exists "feedback" (
  "id" uuid not null,
  "retro_id" uuid not null,
  "idx" int not null,
  "user_id" uuid not null,
  "category" text not null,
  "content" text not null,
  "html" text not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("retro_id") references "retro" ("id"),
  foreign key ("user_id") references "user" ("id"),
  primary key ("id")
);

create index if not exists "feedback__retro_id_idx" on "feedback" ("retro_id");

create index if not exists "feedback__user_id_idx" on "feedback" ("user_id");
-- {% endfunc %}
