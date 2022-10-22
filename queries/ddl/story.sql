-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func StoryDrop() %}
drop table if exists "story";
-- {% endfunc %}

-- {% func StoryCreate() %}
create table if not exists "story" (
  "id" uuid not null,
  "estimate_id" uuid not null,
  "idx" int not null,
  "user_id" uuid not null,
  "title" text not null,
  "status" session_status not null,
  "final_vote" text not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("estimate_id") references "estimate" ("id"),
  foreign key ("user_id") references "user" ("id"),
  primary key ("id")
);

create index if not exists "story__estimate_id_idx" on "story" ("estimate_id");

create index if not exists "story__user_id_idx" on "story" ("user_id");
-- {% endfunc %}
