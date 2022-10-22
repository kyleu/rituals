-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func CreateDatabase() %}
create role "rituals" with login password 'rituals';

create database "rituals";
alter database "rituals" set timezone to 'utc';
grant all privileges on database "rituals" to "rituals";
-- {% endfunc %}
