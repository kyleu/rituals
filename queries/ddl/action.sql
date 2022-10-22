-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func ActionDrop() %}
drop table if exists "action";
-- {% endfunc %}

-- {% func ActionCreate() %}
create table if not exists "action" (
  "id" uuid not null,
  "svc" model_service not null,
  "model_id" uuid not null,
  "user_id" uuid not null,
  "act" text not null,
  "content" text not null,
  "note" text not null,
  "created" timestamp not null default now(),
  foreign key ("user_id") references "user" ("id"),
  primary key ("id")
);

create index if not exists "action__user_id_idx" on "action" ("user_id");
-- {% endfunc %}
