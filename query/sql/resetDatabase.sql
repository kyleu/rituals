drop table if exists "vote";
drop table if exists "story";
drop table if exists "estimate_permission";
drop table if exists "estimate_history";
drop table if exists "estimate_member";
drop table if exists "estimate";

drop table if exists "report";
drop table if exists "standup_permission";
drop table if exists "standup_history";
drop table if exists "standup_member";
drop table if exists "standup";

drop table if exists "feedback";
drop table if exists "retro_permission";
drop table if exists "retro_history";
drop table if exists "retro_member";
drop table if exists "retro";

drop table if exists "sprint_permission";
drop table if exists "sprint_history";
drop table if exists "sprint_member";
drop table if exists "sprint";

drop table if exists "team_permission";
drop table if exists "team_history";
drop table if exists "team_member";
drop table if exists "team";

drop table if exists "action";
drop table if exists "comment";

drop table if exists "auth";
drop table if exists "system_user";

drop type if exists "auth_provider";
drop type if exists "estimate_status";
drop type if exists "member_status";
drop type if exists "model_service";
drop type if exists "retro_status";
drop type if exists "standup_status";
drop type if exists "story_status";
drop type if exists "system_role";

-- <%: func ResetDatabase(w io.Writer) %>
