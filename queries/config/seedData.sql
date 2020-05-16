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
  ('01000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'google', '105933957667429955106', now(), 'Kyle Unverferth', 'kyle@kyleu.com', 'https://lh3.googleusercontent.com/a-/AOh14GiTXL-FkFruVKRTqLxdDmM7RqKe2CZxWRR57xLC27GP9YozCdan2ZUli-y6VoqruMMo2p3-AMNgNnh07l22NtQtRtJ1Y8nm1yi8C4udJft1vp90XwjULbfCT-e8yJJkpkqrb5BcsS_c2u3FI8TmL5zglH_IeUAJ5GPcFh8wV4-n0Ljf9IRLsNb0iinEPjYQKOifp_OpbexnQU1dn7SN3b0I2ygl9JWOZMIZeIP8dDY5JzUO0DZniYGgrX6BuWQpOvydcwxPw9YkagvSEez9dM_OZGget-cXm58nNOMytABeJu7GUbxO60MJm0fM7nYHFwJXzLazYQZNdLQDCnAMvI_HquWL-kzOvTeEt7jQXFTiSddXTeqx-YW4avE-fYvDC71pz0vK9UE8mE-5FIH0WqrFI7xOwTEif4oxAeDWnNLU6MPbSeusNw4Rxn-eIB-TShP-ZHYoUNJ0pm3n2HeRfISXv8STyvLl80SbqeumnS7yPluwseZBXyBDMXv9C1cL3JXS6HhNoTUrGC8sXU2WN5FVbBrh52ZZQEwcvWefQylTJbNKhYh46plQfqbPXByTWzU71YzLRRhhO1uqcVJOy1U1mGPj3X6pHKw3xxAtHQEc4C9u1Z-fayOHZormz6jb1hAKOJHgynZrgi9pb0OKysrhDx3zpgvLbtSpZnNHKuAywUU-gWVOZth0GjbQnVhEHhxRFcZTp-VNIlAwsRrr5Q36vDcqeQlm8nJPKHc7SklCjxG-TV2I6nSLzd0UeKZN80H_ng'),
  ('01000001-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'slack', 'kyle@kyleu.com', now(), 'Kyle Unverferth', 'kyle@kyleu.com', 'https://secure.gravatar.com/avatar/c693fdc9e93d89bba0bfb012fa4a1d76.jpg?s=192&d=https%3A%2F%2Fa.slack-edge.com%2Fdf10d%2Fimg%2Favatars%2Fava_0006-192.png')
;

-- Sprint
insert into sprint
  (id, slug, title, owner, end_date)
values
  ('90000000-0000-0000-0000-000000000000', 'sprint-a', 'Sprint 4', '00000000-0000-0000-0000-000000000000', '2030-01-01'),
  ('90000001-0000-0000-0000-000000000000', 'sprint-b', 'July Sprint', '00000001-0000-0000-0000-000000000000', '2010-01-01'),
  ('90000002-0000-0000-0000-000000000000', 'sprint-c', 'Sprint 2020-09', '00000003-0000-0000-0000-000000000000', null)
;

insert into sprint_member
  (sprint_id, user_id, name, role)
values
  ('90000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'owner'),
  ('90000000-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'member'),
  ('90000001-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'Kyle!', 'member'),
  ('90000001-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'owner'),
  ('90000002-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'member'),
  ('90000002-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'member'),
  ('90000002-0000-0000-0000-000000000000', '00000002-0000-0000-0000-000000000000', '', 'member')
;

-- Estimate
insert into estimate
  (id, slug, title, sprint_id, owner, status, choices, options)
values
  ('10000000-0000-0000-0000-000000000000', 'estimate-a', 'Estimation A', '90000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'new', '{}', '{}'),
  ('10000001-0000-0000-0000-000000000000', 'estimate-b', 'Ad-hoc Session', null, '00000001-0000-0000-0000-000000000000', 'new', '{}', '{}'),
  ('10000002-0000-0000-0000-000000000000', 'estimate-c', 'April Estimates', null, '00000003-0000-0000-0000-000000000000', 'new', '{}', '{}')
;

insert into estimate_member
  (estimate_id, user_id, name, role)
values
  ('10000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'owner'),
  ('10000000-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'member'),
  ('10000001-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'Kyle!', 'member'),
  ('10000001-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'owner'),
  ('10000002-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'member'),
  ('10000002-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'member'),
  ('10000002-0000-0000-0000-000000000000', '00000002-0000-0000-0000-000000000000', '', 'member')
;

insert into story
  (id, estimate_id, idx, author_id, title, status, final_vote)
values
  ('11000000-0000-0000-0000-000000000000', '10000000-0000-0000-0000-000000000000', 0, '00000000-0000-0000-0000-000000000000', 'Design the new widget', 'pending', ''),
  ('11000001-0000-0000-0000-000000000000', '10000000-0000-0000-0000-000000000000', 1, '00000000-0000-0000-0000-000000000000', 'Complete product testing', 'pending', ''),
  ('11000002-0000-0000-0000-000000000000', '10000000-0000-0000-0000-000000000000', 2, '00000000-0000-0000-0000-000000000000', 'Deploy to production cluster', 'active', ''),
  ('11000003-0000-0000-0000-000000000000', '10000000-0000-0000-0000-000000000000', 3, '00000000-0000-0000-0000-000000000000', 'Give up, find new jobs', 'complete', '4')
;

insert into vote
  (story_id, user_id, choice)
values
  ('11000001-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '1'),
  ('11000001-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '2'),
  ('11000002-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '3'),
  ('11000003-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '4')
;


-- Standup
insert into standup
  (id, slug, title, sprint_id, owner, status)
values
  ('20000000-0000-0000-0000-000000000000', 'standup-a', 'Standup A', '90000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'new'),
  ('20000001-0000-0000-0000-000000000000', 'standup-b', 'Team Standup', null, '00000000-0000-0000-0000-000000000000', 'new')
;

insert into standup_member
  (standup_id, user_id, name, role)
values
  ('20000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'owner'),
  ('20000000-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'owner')
;

insert into report
  (id, standup_id, d, author_id, content, html)
values
  ('21000000-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '2020-05-08', '00000000-0000-0000-0000-000000000000', 'Did some stuff', 'Report A'),
  ('21000001-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '2020-05-09', '00000000-0000-0000-0000-000000000000', 'Performed maintenance', 'Report B'),
  ('21000002-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '2020-05-09', '00000001-0000-0000-0000-000000000000', 'Completed a lot of stories', 'Report C'),
  ('21000003-0000-0000-0000-000000000000', '20000000-0000-0000-0000-000000000000', '2020-05-10', '00000001-0000-0000-0000-000000000000', 'Didn''t do much tbh', 'Report D')
;


-- Retro
insert into retro
  (id, slug, title, sprint_id, owner, status, categories)
values
  ('30000000-0000-0000-0000-000000000000', 'retro-a', 'Retro A', '90000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'new', '{}'),
  ('30000001-0000-0000-0000-000000000000', 'retro-b', 'Team Retro', null, '00000000-0000-0000-0000-000000000000', 'new', '{}')
;

insert into retro_member
  (retro_id, user_id, name, role)
values
  ('30000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'owner'),
  ('30000000-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', '', 'owner')
;

insert into feedback
  (id, retro_id, idx, author_id, category, content, html)
values
  ('31000000-0000-0000-0000-000000000000', '30000000-0000-0000-0000-000000000000', 0, '00000000-0000-0000-0000-000000000000', 'bad', 'Servers are slow', 'Bad!'),
  ('31000001-0000-0000-0000-000000000000', '30000000-0000-0000-0000-000000000000', 1, '00000000-0000-0000-0000-000000000000', 'good', 'New leader is doing great', 'Good A'),
  ('31000002-0000-0000-0000-000000000000', '30000000-0000-0000-0000-000000000000', 2, '00000001-0000-0000-0000-000000000000', 'good', 'Documentation has improved', 'Good B'),
  ('31000003-0000-0000-0000-000000000000', '30000000-0000-0000-0000-000000000000', 3, '00000001-0000-0000-0000-000000000000', 'improve', 'Fix the build', 'Improve stuff')
;


-- Action
insert into action
  (id, svc, model_id, author_id, act, content, note)
values
  ('40000000-0000-0000-0000-000000000000', 'estimate', '10000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'add-member', 'null', 'Action A'),
  ('40000001-0000-0000-0000-000000000000', 'estimate', '10000001-0000-0000-0000-000000000000', '00000001-0000-0000-0000-000000000000', 'add-member', 'null', 'Action B'),
  ('40000002-0000-0000-0000-000000000000', 'standup', '20000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'add-member', 'null', 'Action C'),
  ('40000003-0000-0000-0000-000000000000', 'retro', '30000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'add-member', 'null', 'Action D')
;


-- Invite
insert into invitation
  (key, k, v, src, tgt, note, status)
values
  ('private', 'estimate', '10000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'Targeted Invite', 'pending'),
  ('public', 'estimate', '10000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000000', null, 'Open Invite', 'pending')
;

-- <%: func SeedData(w io.Writer) %>
