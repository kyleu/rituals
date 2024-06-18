-- {% func VoteDrop() %}
drop table if exists "vote";
-- {% endfunc %}

-- {% func VoteCreate() %}
create table if not exists "vote" (
  "story_id" uuid not null,
  "user_id" uuid not null,
  "choice" text not null,
  "created" timestamp not null default now(),
  "updated" timestamp default now(),
  foreign key ("story_id") references "story" ("id"),
  foreign key ("user_id") references "user" ("id"),
  primary key ("story_id", "user_id")
);

create index if not exists "vote__story_id_idx" on "vote" ("story_id");

create index if not exists "vote__user_id_idx" on "vote" ("user_id");
-- {% endfunc %}
