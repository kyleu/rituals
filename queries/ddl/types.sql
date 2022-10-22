-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func TypesDrop() %}
drop type if exists "session_status";
drop type if exists "model_service";
drop type if exists "member_status";
-- {% endfunc %}

-- {% func TypesCreate() %}
create type "member_status" as enum ('owner', 'member', 'observer');
create type "model_service" as enum ('team', 'sprint', 'estimate', 'standup', 'retro', 'story', 'feedback', 'report');
create type "session_status" as enum ('new', 'active', 'complete', 'deleted');
-- {% endfunc %}
