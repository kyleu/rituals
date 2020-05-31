create database "rituals.dev";
create role "rituals.dev" with login password 'rituals.dev';
grant all privileges on database "rituals.dev" to "rituals.dev";
alter database "rituals.dev" set timezone to 'utc';
