insert into projects
(
  key,
  title,
  description,
  owner,
  engine,
  url,
  username,
  password
)
values
(
  'test',
  'Test',
  'Test database',
  '00000000-0000-0000-0000-000000000000',
  'pgx',
  'postgres://127.0.0.1:5432/dbui?sslmode=disable',
  null,
  null
),
(
  'mysql',
  'MySQL',
  'MySQL test database',
  '00000000-0000-0000-0000-000000000000',
  'mysql',
  'dbui:password@tcp(localhost)/dbui',
  'dbui',
  'password'
);

-- <%: func InsertDataProject(w io.Writer) %>
