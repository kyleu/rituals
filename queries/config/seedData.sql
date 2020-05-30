-- User
insert into system_user
  (id, name, role, theme, nav_color, link_color, picture, locale)
values
  ('00000000-0000-0000-0000-000000000000', 'Kyle U', 'admin', 'default', 'bluegrey', 'bluegrey', '', 'en-US'),
  ('00000001-0000-0000-0000-000000000000', 'Katie', 'guest', 'default', 'bluegrey', 'bluegrey', '', 'en-US'),
  ('00000002-0000-0000-0000-000000000000', 'Dan', 'guest', 'light', 'bluegrey', 'bluegrey', '', 'en-US'),
  ('00000003-0000-0000-0000-000000000000', 'Janet', 'guest', 'dark', 'bluegrey', 'bluegrey', '', 'en-US')
;

-- Auth
insert into auth
  (id, user_id, provider, provider_id, user_list_id, user_list_name, access_token, expires, name, email, picture)
values
  ('01000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'google', '105933957667429955106', '', '', 'seed-data', now(), 'Kyle Unverferth', 'kyle@kyleu.com', 'https://placekitten.com/100/100'),
  ('01000001-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'slack', 'kyle@kyleu.com', ' T013NH3CSCB', 'rituals.dev', 'xoxp-1124581434419-1148209542112-1109658115543-c510da45e8c8481be7f2b06b7b04122f', now(), 'Kyle Unverferth', 'kyle@kyleu.com', 'https://placekitten.com/100/100')
;

-- Team
insert into team
  (id, slug, title, owner)
values
  ('10000000-0000-0000-0000-000000000000', 'team-a', 'Team A', '00000000-0000-0000-0000-000000000000'),
  ('10000001-0000-0000-0000-000000000000', 'team-b', 'Team B', '00000001-0000-0000-0000-000000000000')
;

insert into team_history
  (slug, model_id, model_name)
values
  ('team-old', '10000000-0000-0000-0000-000000000000', 'Team Old')
;

insert into team_permission
  (team_id, k, v, access)
values
  ('10000000-0000-0000-0000-000000000000', 'google', '@kyleu.com', 'member')
;

insert into team_member
  (team_id, user_id, name, role)
values
  ('10000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'owner'),
  ('10000000-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'member'),
  ('10000001-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'owner'),
  ('10000001-0000-0000-0000-000000000000', '00000002-0000-0000-0000-000000000000', '', 'member')
;

-- Sprint
insert into sprint
  (id, slug, title, team_id, owner, start_date, end_date)
values
  ('20000000-0000-0000-0000-000000000000', 'sprint-a', 'Sprint 4', '10000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '2020-01-01', '2030-01-01'),
  ('20000001-0000-0000-0000-000000000000', 'sprint-b', 'Ad-hoc Sprint', null, '00000001-0000-0000-0000-000000000000', null, '2010-01-01'),
  ('20000002-0000-0000-0000-000000000000', 'sprint-c', 'Sprint 2020-09', null, '00000003-0000-0000-0000-000000000000', null, null)
;

insert into sprint_permission
  (sprint_id, k, v, access)
values
  ('20000000-0000-0000-0000-000000000000', 'team', '', 'member'),
  ('20000000-0000-0000-0000-000000000000', 'google', '@kyleu.com', 'member')
;

insert into sprint_history
  (slug, model_id, model_name)
values
  ('sprint-old', '20000000-0000-0000-0000-000000000000', 'Sprint Old')
;

insert into sprint_member
  (sprint_id, user_id, name, role)
values
  ('20000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'owner'),
  ('20000000-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'member'),
  ('20000001-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'owner'),
  ('20000002-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'member'),
  ('20000002-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'member'),
  ('20000002-0000-0000-0000-000000000000', '00000002-0000-0000-0000-000000000000', '', 'member')
;

-- Estimate
insert into estimate
  (id, slug, title, team_id, sprint_id, owner, status, choices)
values
  ('30000000-0000-0000-0000-000000000000', 'estimate-a', 'Estimation A', '10000000-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'new', '{}'),
  ('30000001-0000-0000-0000-000000000000', 'estimate-b', 'Secure Session', '10000001-0000-0000-0000-000000000000', '20000001-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', 'new', '{}'),
  ('30000002-0000-0000-0000-000000000000', 'estimate-c', 'Ad-hoc Estimates', null, null, '00000003-0000-0000-0000-000000000000', 'new', '{1,2,3,4,5}')
;

insert into estimate_permission
  (estimate_id, k, v, access)
values
  ('30000000-0000-0000-0000-000000000000', 'team', '', 'member'),
  ('30000000-0000-0000-0000-000000000000', 'google', '@kyleu.com', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'team', '', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'sprint', '', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'github', '@kyleu.com', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'google', '@kyleu.com', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'google', '@gmail.com', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'slack', '@kyleu.com', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'facebook', '@kyleu.com', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'amazon', '@gmail.com', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'microsoft', '@gmail.com', 'member')
;

insert into estimate_history
  (slug, model_id, model_name)
values
  ('estimate-old', '30000000-0000-0000-0000-000000000000', 'Estimate Old')
;

insert into estimate_member
  (estimate_id, user_id, name, role)
values
  ('30000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'owner'),
  ('30000000-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'member'),
  ('30000001-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'Kyle!', 'member'),
  ('30000001-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'owner'),
  ('30000002-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'member'),
  ('30000002-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'member'),
  ('30000002-0000-0000-0000-000000000000', '00000002-0000-0000-0000-000000000000', '', 'member')
;

insert into story
  (id, estimate_id, idx, user_id, title, status, final_vote)
values
  ('31000000-0000-0000-0000-000000000000', '30000000-0000-0000-0000-000000000000', 0, '00000000-0000-0000-0000-000000000000', 'Design the new widget', 'pending', ''),
  ('31000001-0000-0000-0000-000000000000', '30000000-0000-0000-0000-000000000000', 1, '00000000-0000-0000-0000-000000000000', 'Complete product testing', 'pending', ''),
  ('31000002-0000-0000-0000-000000000000', '30000000-0000-0000-0000-000000000000', 2, '00000000-0000-0000-0000-000000000000', 'Deploy to production cluster', 'active', ''),
  ('31000003-0000-0000-0000-000000000000', '30000000-0000-0000-0000-000000000000', 3, '00000000-0000-0000-0000-000000000000', 'Give up, find new jobs', 'complete', '4')
;

insert into vote
  (story_id, user_id, choice)
values
  ('31000001-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '1'),
  ('31000001-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '2'),
  ('31000002-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '3'),
  ('31000003-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '4')
;


-- Standup
insert into standup
  (id, slug, title, team_id, sprint_id, owner, status)
values
  ('40000000-0000-0000-0000-000000000000', 'standup-a', 'Standup A', '10000000-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'new'),
  ('40000001-0000-0000-0000-000000000000', 'standup-b', 'Team Standup', null, null, '00000000-0000-0000-0000-000000000000', 'new')
;

insert into standup_permission
  (standup_id, k, v, access)
values
  ('40000000-0000-0000-0000-000000000000', 'team', '', 'member'),
  ('40000000-0000-0000-0000-000000000000', 'google', '@kyleu.com', 'member')
;

insert into standup_member
  (standup_id, user_id, name, role)
values
  ('40000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'owner'),
  ('40000000-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'owner')
;

insert into report
  (id, standup_id, d, user_id, content, html)
values
  ('41000000-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2020-05-08', '00000000-0000-0000-0000-000000000000', 'Did some stuff', 'Did some stuff'),
  ('41000001-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2020-05-09', '00000000-0000-0000-0000-000000000000', 'Performed maintenance', 'Performed maintenance'),
  ('41000002-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2020-05-09', '00000001-0000-0000-0000-000000000000', 'Completed a lot of stories', 'Completed a lot of stories'),
  ('41000003-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2020-05-10', '00000001-0000-0000-0000-000000000000', 'Didn''t do much tbh', 'Didn''t do much tbh')
;


-- Retro
insert into retro
  (id, slug, title, team_id, sprint_id, owner, status, categories)
values
  ('50000000-0000-0000-0000-000000000000', 'retro-a', 'Retro A', '10000000-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'new', '{}'),
  ('50000001-0000-0000-0000-000000000000', 'retro-b', 'Team Retro', null, null, '00000000-0000-0000-0000-000000000000', 'new', '{}')
;

insert into retro_permission
  (retro_id, k, v, access)
values
  ('50000000-0000-0000-0000-000000000000', 'team', '', 'member'),
  ('50000000-0000-0000-0000-000000000000', 'google', '@kyleu.com', 'member')
;

insert into retro_member
  (retro_id, user_id, name, role)
values
  ('50000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'owner'),
  ('50000000-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'owner')
;

insert into feedback
  (id, retro_id, idx, user_id, category, content, html)
values
  ('51000000-0000-0000-0000-000000000000', '50000000-0000-0000-0000-000000000000', 0, '00000000-0000-0000-0000-000000000000', 'bad', 'Servers are slow', 'Bad!'),
  ('51000001-0000-0000-0000-000000000000', '50000000-0000-0000-0000-000000000000', 1, '00000000-0000-0000-0000-000000000000', 'good', 'New leader is doing great', 'Good A'),
  ('51000002-0000-0000-0000-000000000000', '50000000-0000-0000-0000-000000000000', 2, '00000001-0000-0000-0000-000000000000', 'good', 'Documentation has improved', 'Good B'),
  ('51000003-0000-0000-0000-000000000000', '50000000-0000-0000-0000-000000000000', 3, '00000001-0000-0000-0000-000000000000', 'improve', 'Fix the build', 'Improve stuff')
;


-- Action
insert into action
  (id, svc, model_id, user_id, act, content, note)
values
  ('60000000-0000-0000-0000-000000000000', 'estimate', '30000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'add-member', 'null', 'Action A'),
  ('60000001-0000-0000-0000-000000000000', 'estimate', '30000001-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', 'add-member', 'null', 'Action B'),
  ('60000002-0000-0000-0000-000000000000', 'standup', '40000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'add-member', 'null', 'Action C'),
  ('60000003-0000-0000-0000-000000000000', 'retro', '50000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'add-member', 'null', 'Action D')
;


-- Invite
insert into invitation
  (key, k, v, src, tgt, note, status)
values
  ('private', 'estimate', '30000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'Targeted Invite', 'pending'),
  ('public', 'estimate', '30000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000000', null, 'Open Invite', 'pending')
;


-- Comment
insert into comment
  (id, svc, model_id, target_type, target_id, user_id, content, html)
values
  ('70000000-0000-0000-0000-000000000000', 'team', '10000000-0000-0000-0000-000000000000', '', null, '00000000-0000-0000-0000-000000000000', 'team comment', '<div>team comment</div>'),
  ('70000001-0000-0000-0000-000000000000', 'sprint', '20000000-0000-0000-0000-000000000000', '', null, '00000000-0000-0000-0000-000000000000', 'sprint comment', '<div>sprint comment</div>'),
  ('70000002-0000-0000-0000-000000000000', 'estimate', '30000000-0000-0000-0000-000000000000', '', null, '00000000-0000-0000-0000-000000000000', 'estimate comment 1', '<div>estimate comment 1</div>'),
  ('70000003-0000-0000-0000-000000000000', 'estimate', '30000000-0000-0000-0000-000000000000', '', null, '00000000-0000-0000-0000-000000000000', 'estimate comment 2', '<div>estimate comment 2</div>'),
  ('70000004-0000-0000-0000-000000000000', 'estimate', '30000000-0000-0000-0000-000000000000', 'story', '31000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'story comment', '<div>story comment</div>'),
  ('70000005-0000-0000-0000-000000000000', 'standup', '40000000-0000-0000-0000-000000000000', '', null, '00000000-0000-0000-0000-000000000000', 'standup comment', '<div>standup comment</div>'),
  ('70000006-0000-0000-0000-000000000000', 'standup', '40000000-0000-0000-0000-000000000000', 'report', '41000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'report comment', '<div>report comment</div>'),
  ('70000007-0000-0000-0000-000000000000', 'retro', '50000000-0000-0000-0000-000000000000', '', null, '00000000-0000-0000-0000-000000000000', 'retro comment', '<div>retro comment</div>'),
  ('70000008-0000-0000-0000-000000000000', 'retro', '50000000-0000-0000-0000-000000000000', 'feedback', '51000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'feedback comment', '<div>feedback comment</div>')
;


-- <%: func SeedData(w io.Writer) %>
