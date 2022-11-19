-- $PF_IGNORE$
-- {% import "github.com/kyleu/rituals/queries/ddl" %}
-- {% import "github.com/kyleu/rituals/queries/seeddata" %}

-- {% func Migration1InitialDatabase(debug bool) %}

-- {%- if debug -%}
-- {%= ddl.DropAll() %}
-- {%- endif -%}

-- {%= ddl.CreateAll() %}
-- {%= seeddata.SeedDataAll() %}
-- {% endfunc %}
