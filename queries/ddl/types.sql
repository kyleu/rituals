-- {% func TypesDrop() %}
drop type if exists "session_status";
drop type if exists "model_service";
drop type if exists "member_status";
-- {% endfunc %}

-- {% func TypesCreate() %}
do $$ begin
  create type "member_status" as enum ('owner', 'member', 'observer');
exception
  when duplicate_object then null;
end $$;
do $$ begin
  create type "model_service" as enum ('team', 'sprint', 'estimate', 'standup', 'retro', 'story', 'feedback', 'report');
exception
  when duplicate_object then null;
end $$;
do $$ begin
  create type "session_status" as enum ('new', 'active', 'complete');
exception
  when duplicate_object then null;
end $$;
-- {% endfunc %}
