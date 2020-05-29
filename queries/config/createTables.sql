-- User
create table if not exists "system_user" (
  "id" uuid not null primary key,
  "name" varchar(2048) not null,
  "role" system_role not null,
  "theme" varchar(32) not null,
  "nav_color" varchar(32) not null,
  "link_color" varchar(32) not null,
  "picture" text not null,
  "locale" varchar(32) not null,
  "created" timestamp not null default now()
);

create table if not exists "auth" (
  "id" uuid not null primary key,
  "user_id" uuid not null references "system_user"("id"),
  "provider" auth_provider not null,
  "provider_id" text not null,
  "user_list_id" varchar(512) not null,
  "user_list_name" varchar(2048) not null,
  "access_token" text not null,
  "expires" timestamp,
  "name" varchar(2048) not null,
  "email" varchar(2048) not null,
  "picture" text not null,
  "created" timestamp not null default now()
);

create index if not exists idx_auth_provider_provider_id on auth(provider, provider_id);

-- Team
create table if not exists "team" (
  "id" uuid not null primary key,
  "slug" varchar(128) not null unique,
  "title" varchar(2048) not null,
  "owner" uuid references "system_user"("id"),
  "created" timestamp not null default now()
);

create unique index if not exists idx_team_slug on team(slug);

create table if not exists "team_member" (
  "team_id" uuid not null references "team"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  primary key ("team_id", "user_id")
);

create table if not exists "team_permission" (
  "team_id" uuid not null references "team"("id"),
  "k" varchar(128) not null,
  "v" varchar(2048) not null,
  "access" member_status not null,
  "created" timestamp not null default now(),
  primary key (team_id, k, v)
);

create index if not exists idx_team_permission_team_id on team_permission(team_id);

-- Sprint
create table if not exists "sprint" (
  "id" uuid not null primary key,
  "slug" varchar(128) not null unique,
  "title" varchar(2048) not null,
  "team_id" uuid references "team"("id"),
  "owner" uuid references "system_user"("id"),
  "start_date" date,
  "end_date" date,
  "created" timestamp not null default now()
);

create unique index if not exists idx_sprint_slug on sprint(slug);

create table if not exists "sprint_member" (
  "sprint_id" uuid not null references "sprint"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  primary key ("sprint_id", "user_id")
);

create table if not exists "sprint_permission" (
  "sprint_id" uuid not null references "sprint"("id"),
  "k" varchar(128) not null,
  "v" varchar(2048) not null,
  "access" member_status not null,
  "created" timestamp not null default now(),
  primary key (sprint_id, k, v)
);

create index if not exists idx_sprint_permission_sprint_id on sprint_permission(sprint_id);

-- Estimate
create table if not exists "estimate" (
  "id" uuid not null primary key,
  "slug" varchar(128) not null unique,
  "title" varchar(2048) not null,
  "team_id" uuid references "team"("id"),
  "sprint_id" uuid references "sprint"("id"),
  "owner" uuid references "system_user"("id"),
  "status" estimate_status not null,
  "choices" varchar(2048)[] not null,
  "created" timestamp not null default now()
);

create unique index if not exists idx_estimate_slug on estimate(slug);

create table if not exists "estimate_member" (
  "estimate_id" uuid not null references "estimate"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  primary key ("estimate_id", "user_id")
);

create table if not exists "estimate_permission" (
  "estimate_id" uuid not null references "estimate"("id"),
  "k" varchar(128) not null,
  "v" varchar(2048) not null,
  "access" member_status not null,
  "created" timestamp not null default now(),
  primary key (estimate_id, k, v)
);

create index if not exists idx_estimate_permission_estimate_id on estimate_permission(estimate_id);

create table if not exists "story" (
  "id" uuid not null primary key,
  "estimate_id" uuid not null references "estimate"("id"),
  "idx" int not null default 0,
  "author_id" uuid not null references "system_user"("id"),
  "title" varchar(2048) not null,
  "status" story_status not null default 'pending',
  "final_vote" varchar(2048) not null,
  "created" timestamp not null default now()
);

create table if not exists "vote" (
  "story_id" uuid not null references "story"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "choice" varchar(256) not null,
  "updated" timestamp not null default now(),
  "created" timestamp not null default now(),
  primary key ("story_id", "user_id")
);

-- Standup
create table if not exists "standup" (
  "id" uuid not null primary key,
  "slug" varchar(128) not null unique,
  "title" varchar(2048) not null,
  "team_id" uuid references "team"("id"),
  "sprint_id" uuid references "sprint"("id"),
  "owner" uuid references "system_user"("id"),
  "status" standup_status not null,
  "created" timestamp not null default now()
);

create unique index if not exists idx_standup_slug on standup(slug);

create table if not exists "standup_member" (
  "standup_id" uuid not null references "standup"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  primary key ("standup_id", "user_id")
);

create table if not exists "standup_permission" (
  "standup_id" uuid not null references "standup"("id"),
  "k" varchar(128) not null,
  "v" varchar(2048) not null,
  "access" member_status not null,
  "created" timestamp not null default now(),
  primary key (standup_id, k, v)
);

create index if not exists idx_standup_permission_standup_id on standup_permission(standup_id);

create table "report" (
  "id" uuid not null primary key,
  "standup_id" uuid not null references "standup"("id"),
  "d" date not null default now()::date,
  "author_id" uuid not null references "system_user"("id"),
  "content" text not null,
  "html" text not null,
  "created" timestamp not null default now()
);

-- Retro
create table if not exists "retro" (
  "id" uuid not null primary key,
  "slug" varchar(128) not null unique,
  "title" varchar(2048) not null,
  "team_id" uuid references "team"("id"),
  "sprint_id" uuid references "sprint"("id"),
  "owner" uuid references "system_user"("id"),
  "status" retro_status not null,
  "categories" varchar(2048)[] not null,
  "created" timestamp not null default now()
);

create unique index if not exists idx_retro_slug on retro(slug);

create table if not exists "retro_member" (
  "retro_id" uuid not null references "retro"("id"),
  "user_id" uuid not null references "system_user"("id"),
  "name" varchar(2048) not null,
  "role" member_status not null,
  "created" timestamp not null default now(),
  primary key ("retro_id", "user_id")
);

create table if not exists "retro_permission" (
  "retro_id" uuid not null references "retro"("id"),
  "k" varchar(128) not null,
  "v" varchar(2048) not null,
  "access" member_status not null,
  "created" timestamp not null default now(),
  primary key (retro_id, k, v)
);

create index if not exists idx_retro_permission_retro_id on retro_permission(retro_id);

create table if not exists "feedback" (
  "id" uuid not null primary key,
  "retro_id" uuid not null references "retro"("id"),
  "idx" int not null default 0,
  "author_id" uuid not null references "system_user"("id"),
  "category" varchar(2048) not null,
  "content" text not null,
  "html" text not null,
  "created" timestamp not null default now()
);

-- Invite
create table if not exists "invitation" (
  "key" varchar(128) not null primary key,
  "k" invitation_type not null,
  "v" uuid not null,
  "src" uuid references "system_user"("id"),
  "tgt" uuid references "system_user"("id"),
  "note" text,
  "status" invitation_status not null,
  "redeemed" timestamp,
  "created" timestamp not null default now()
);

-- Actions
create table if not exists "action" (
  "id" uuid not null primary key,
  "svc" model_service not null,
  "model_id" uuid not null,
  "author_id" uuid references "system_user"("id"),
  "act" varchar(64) not null,
  "content" json,
  "note" text,
  "created" timestamp not null default now()
);

-- Actions
create table if not exists "comment" (
  "id" uuid not null primary key,
  "svc" model_service not null,
  "model_id" uuid not null,
  "author_id" uuid references "system_user"("id"),
  "content" text not null,
  "html" text not null,
  "created" timestamp not null default now()
);

create index if not exists idx_comment_svc_model_id on comment(svc, model_id);

-- <%: func CreateTables(w io.Writer) %>
