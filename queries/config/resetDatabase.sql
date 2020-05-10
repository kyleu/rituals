drop table if exists "vote";
drop table if exists "story";
drop table if exists "estimate_member";
drop table if exists "estimate";

drop table if exists "standup_update"; -- Legacy
drop table if exists "report";
drop table if exists "standup_member";
drop table if exists "standup";

drop table if exists "retro_member";
drop table if exists "retro";

drop table if exists "invitation";

drop table if exists "system_user";

drop type if exists "estimate_status";
drop type if exists "invitation_type";
drop type if exists "invitation_status";
drop type if exists "retro_status";
drop type if exists "standup_status";
drop type if exists "story_status";
drop type if exists "system_role";

-- <%: func ResetDatabase(w io.Writer) %>
