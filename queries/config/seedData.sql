-- User
insert into system_user
  (id, name, role, theme, nav_color, link_color, picture, locale)
values
  ('00000000-0000-0000-0000-000000000000', 'Kyle U', 'admin', 'light', 'bluegrey', 'bluegrey', '', 'en-US'),
  ('00000001-0000-0000-0000-000000000000', 'Katie', 'guest', 'light', 'bluegrey', 'bluegrey', '', 'en-US'),
  ('00000002-0000-0000-0000-000000000000', 'Dan', 'guest', 'light', 'bluegrey', 'bluegrey', '', 'en-US'),
  ('00000003-0000-0000-0000-000000000000', 'Janet', 'guest', 'light', 'bluegrey', 'bluegrey', '', 'en-US')
;

-- Auth
insert into auth
  (id, user_id, provider, provider_id, expires, name, email, picture)
values
  ('03000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'google', '105933957667429955106', now(), 'Kyle Unverferth', 'kyle@kyleu.com', 'https://lh3.googleusercontent.com/a-/AOh14GiTXL-FkFruVKRTqLxdDmM7RqKe2CZxWRR57xLC27GP9YozCdan2ZUli-y6VoqruMMo2p3-AMNgNnh07l22NtQtRtJ1Y8nm1yi8C4udJft1vp90XwjULbfCT-e8yJJkpkqrb5BcsS_c2u3FI8TmL5zglH_IeUAJ5GPcFh8wV4-n0Ljf9IRLsNb0iinEPjYQKOifp_OpbexnQU1dn7SN3b0I2ygl9JWOZMIZeIP8dDY5JzUO0DZniYGgrX6BuWQpOvydcwxPw9YkagvSEez9dM_OZGget-cXm58nNOMytABeJu7GUbxO60MJm0fM7nYHFwJXzLazYQZNdLQDCnAMvI_HquWL-kzOvTeEt7jQXFTiSddXTeqx-YW4avE-fYvDC71pz0vK9UE8mE-5FIH0WqrFI7xOwTEif4oxAeDWnNLU6MPbSeusNw4Rxn-eIB-TShP-ZHYoUNJ0pm3n2HeRfISXv8STyvLl80SbqeumnS7yPluwseZBXyBDMXv9C1cL3JXS6HhNoTUrGC8sXU2WN5FVbBrh52ZZQEwcvWefQylTJbNKhYh46plQfqbPXByTWzU71YzLRRhhO1uqcVJOy1U1mGPj3X6pHKw3xxAtHQEc4C9u1Z-fayOHZormz6jb1hAKOJHgynZrgi9pb0OKysrhDx3zpgvLbtSpZnNHKuAywUU-gWVOZth0GjbQnVhEHhxRFcZTp-VNIlAwsRrr5Q36vDcqeQlm8nJPKHc7SklCjxG-TV2I6nSLzd0UeKZN80H_ng'),
  ('03000001-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'slack', 'kyle@kyleu.com', now(), 'Kyle Unverferth', 'kyle@kyleu.com', 'https://secure.gravatar.com/avatar/c693fdc9e93d89bba0bfb012fa4a1d76.jpg?s=192&d=https%3A%2F%2Fa.slack-edge.com%2Fdf10d%2Fimg%2Favatars%2Fava_0006-192.png')
;

-- Team
insert into team
  (id, slug, title, owner)
values
  ('10000000-0000-0000-0000-000000000000', 'team-a', 'Team A', '00000000-0000-0000-0000-000000000000'),
  ('10000001-0000-0000-0000-000000000000', 'team-b', 'Team B', '00000001-0000-0000-0000-000000000000')
;

insert into team_permission
  (team_id, k, v, access)
values
  ('10000000-0000-0000-0000-000000000000', 'google', '@kyleu.com', 'member'),
  ('10000000-0000-0000-0000-000000000000', 'slack', '@kyleu.com', 'member')
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
  ('20000000-0000-0000-0000-000000000000', 'google', '@kyleu.com', 'member'),
  ('20000000-0000-0000-0000-000000000000', 'slack', '@kyleu.com', 'member')
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
  ('30000000-0000-0000-0000-000000000000', 'slack', '@kyleu.com', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'team', '', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'sprint', '', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'github', '@kyleu.com', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'google', '@kyleu.com', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'google', '@gmail.com', 'member'),
  ('30000001-0000-0000-0000-000000000000', 'slack', '@kyleu.com', 'member')
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
  (id, estimate_id, idx, author_id, title, status, final_vote)
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
  ('40000000-0000-0000-0000-000000000000', 'google', '@kyleu.com', 'member'),
  ('40000000-0000-0000-0000-000000000000', 'slack', '@kyleu.com', 'member')
;

insert into standup_member
  (standup_id, user_id, name, role)
values
  ('40000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'owner'),
  ('40000000-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'owner')
;

insert into report
  (id, standup_id, d, author_id, content, html)
values
  ('41000000-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2020-05-08', '00000000-0000-0000-0000-000000000000', 'Did some stuff', 'Report A'),
  ('41000001-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2020-05-09', '00000000-0000-0000-0000-000000000000', 'Performed maintenance', 'Report B'),
  ('41000002-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2020-05-09', '00000001-0000-0000-0000-000000000000', 'Completed a lot of stories', 'Report C'),
  ('41000003-0000-0000-0000-000000000000', '40000000-0000-0000-0000-000000000000', '2020-05-10', '00000001-0000-0000-0000-000000000000', 'Didn''t do much tbh', 'Report D')
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
  ('50000000-0000-0000-0000-000000000000', 'google', '@kyleu.com', 'member'),
  ('50000000-0000-0000-0000-000000000000', 'slack', '@kyleu.com', 'member')
;

insert into retro_member
  (retro_id, user_id, name, role)
values
  ('50000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'owner'),
  ('50000000-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'owner')
;

insert into feedback
  (id, retro_id, idx, author_id, category, content, html)
values
  ('51000000-0000-0000-0000-000000000000', '50000000-0000-0000-0000-000000000000', 0, '00000000-0000-0000-0000-000000000000', 'bad', 'Servers are slow', 'Bad!'),
  ('51000001-0000-0000-0000-000000000000', '50000000-0000-0000-0000-000000000000', 1, '00000000-0000-0000-0000-000000000000', 'good', 'New leader is doing great', 'Good A'),
  ('51000002-0000-0000-0000-000000000000', '50000000-0000-0000-0000-000000000000', 2, '00000001-0000-0000-0000-000000000000', 'good', 'Documentation has improved', 'Good B'),
  ('51000003-0000-0000-0000-000000000000', '50000000-0000-0000-0000-000000000000', 3, '00000001-0000-0000-0000-000000000000', 'improve', 'Fix the build', 'Improve stuff')
;


-- Action
insert into action
  (id, svc, model_id, author_id, act, content, note)
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

-- <%: func SeedData(w io.Writer) %>
