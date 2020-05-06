insert into system_user
  (id, name, role, theme, nav_color, link_color, locale)
values
  ('00000000-0000-0000-0000-000000000000', 'Default User', 'guest', 'light', 'bluegrey', 'bluegrey', 'en-US'),
  ('FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF', 'System Admin', 'admin', 'light', 'bluegrey', 'bluegrey', 'en-US')
;

insert into estimate
  (id, slug, title, owner, status, choices, options)
values
  ('10000000-0000-0000-0000-000000000000', 'estimate-a', 'Estimation Session A', '00000000-0000-0000-0000-000000000000', 'new', '{}', '{}'),
  ('10000001-0000-0000-0000-000000000000', 'estimate-b', 'Estimation Session B', '00000000-0000-0000-0000-000000000000', 'new', '{}', '{}')
;

insert into estimate_member
  (estimate_id, user_id, name, role)
values
  ('10000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '', 'owner'),
  ('10000000-0000-0000-0000-000000000000', 'FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF', '', 'member')
;

insert into poll
  (id, estimate_id, idx, author_id, title, status, final_vote)
values
  ('11000000-0000-0000-0000-000000000000', '10000000-0000-0000-0000-000000000000', 0, '00000000-0000-0000-0000-000000000000', 'Poll A', 'pending', ''),
  ('11000001-0000-0000-0000-000000000000', '10000000-0000-0000-0000-000000000000', 1, '00000000-0000-0000-0000-000000000000', 'Poll B', 'pending', '1'),
  ('11000002-0000-0000-0000-000000000000', '10000000-0000-0000-0000-000000000000', 2, '00000000-0000-0000-0000-000000000000', 'Poll C', 'pending', '2'),
  ('11000003-0000-0000-0000-000000000000', '10000000-0000-0000-0000-000000000000', 3, '00000000-0000-0000-0000-000000000000', 'Poll D', 'pending', '3.5')
;

insert into vote
  (poll_id, user_id, choice)
values
  ('11000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '1'),
  ('11000001-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '2'),
  ('11000002-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '3'),
  ('11000003-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '4')
;

insert into standup
  (id, slug, title, owner, status)
values
  ('20000000-0000-0000-0000-000000000000', 'standup-a', 'Daily Standup A', '00000000-0000-0000-0000-000000000000', 'new'),
  ('20000001-0000-0000-0000-000000000000', 'standup-b', 'Daily Standup B', '00000000-0000-0000-0000-000000000000', 'new')
;

insert into standup_member
  (standup_id, user_id, name, role)
values
  ('20000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'AliasA', 'owner'),
  ('20000000-0000-0000-0000-000000000000', 'FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF', 'AliasB', 'owner')
;

insert into retro
  (id, slug, title, owner, status)
values
  ('30000000-0000-0000-0000-000000000000', 'retro-a', 'Retrospective A', '00000000-0000-0000-0000-000000000000', 'new'),
  ('30000001-0000-0000-0000-000000000000', 'retro-b', 'Retrospective B', '00000000-0000-0000-0000-000000000000', 'new')
;

insert into retro_member
  (retro_id, user_id, name, role)
values
  ('30000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'AliasA', 'owner'),
  ('30000000-0000-0000-0000-000000000000', 'FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF', 'AliasB', 'owner')
;

insert into invitation
  (key, k, v, src, tgt, note, status)
values
  ('private', 'estimate', '10000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'Targeted Invite', 'pending'),
  ('public', 'estimate', '10000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000000', null, 'Open Invite', 'pending')
;

-- <%: func SeedData(w io.Writer) %>
