delete from comment where id = '70000099-0000-0000-0000-000000000000';

insert into comment
  (id, svc, model_id, target_type, target_id, user_id, content, html)
values
  ('70000099-0000-0000-0000-000000000000', 'team', '10000000-0000-0000-0000-000000000000', '', null, 'F0000000-0000-0000-0000-000000000000', 'migration comment', '<div>migration comment</div>')
;
-- <%: func Migration1(w io.Writer) %>
