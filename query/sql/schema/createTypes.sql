create type system_role as enum ('guest', 'user', 'admin');
create type model_service as enum ('team', 'sprint', 'estimate', 'standup', 'retro', 'story', 'feedback', 'report');

create type auth_provider as enum ('team', 'sprint', 'github', 'google', 'slack', 'facebook', 'amazon', 'microsoft');
create type member_status as enum ('owner', 'member', 'observer');

create type estimate_status as enum ('new', 'active', 'complete', 'deleted');
create type story_status as enum('pending', 'active', 'complete', 'deleted');

create type standup_status as enum ('new', 'deleted');

create type retro_status as enum ('new', 'deleted');

-- <%: func CreateTypes(w io.Writer) %>
