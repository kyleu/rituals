create table projects (
  key varchar(512) not null primary key,
  title varchar(512) not null,
  description text,
  owner uuid,
  engine varchar(64) not null,
  url text not null,
  username varchar(512),
  password varchar(512)
);

-- <%: func CreateTableProject(w io.Writer) %>
