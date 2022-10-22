-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func EmailDrop() %}
drop table if exists "email";
-- {% endfunc %}

-- {% func EmailCreate() %}
create table if not exists "email" (
  "id" uuid not null,
  "recipients" jsonb not null,
  "subject" text not null,
  "data" jsonb not null,
  "plain" text not null,
  "html" text not null,
  "user_id" uuid not null,
  "status" text not null,
  "created" timestamp not null default now(),
  foreign key ("user_id") references "user" ("id"),
  primary key ("id")
);

create index if not exists "email__user_id_idx" on "email" ("user_id");
-- {% endfunc %}
