create type system_role as enum ('guest', 'user', 'admin');
create type model_service as enum ('team', 'sprint', 'estimate', 'standup', 'retro', 'story', 'feedback', 'report');

create type auth_provider as enum ('team', 'sprint', 'github', 'google', 'slack', 'facebook', 'amazon', 'microsoft');
create type member_status as enum ('owner', 'member', 'observer');

create type session_status as enum ('new', 'active', 'complete', 'deleted');

-- <%: func CreateTypes(w io.Writer) %>
