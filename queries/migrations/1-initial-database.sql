-- $PF_IGNORE$
-- {% import "github.com/kyleu/rituals/queries/ddl" %}
-- {% func Migration1InitialDatabase(debug bool) %}

-- {%- if debug -%}
-- {%= ddl.DropAll() %}
-- {%- endif -%}

-- {%= ddl.CreateAll() %}
-- {% endfunc %}
