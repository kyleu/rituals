-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func CommentDrop() %}
drop table if exists "comment";
-- {% endfunc %}

-- {% func CommentCreate() %}
create table if not exists "comment" (
  "id" uuid not null,
  "svc" model_service not null,
  "model_id" uuid not null,
  "target_type" text not null,
  "target_id" uuid not null,
  "user_id" uuid not null,
  "content" text not null,
  "html" text not null,
  "created" timestamp not null default now(),
  foreign key ("user_id") references "user" ("id"),
  primary key ("id")
);

create index if not exists "comment__user_id_idx" on "comment" ("user_id");
-- {% endfunc %}
