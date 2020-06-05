#!/bin/sh

## Builds all the templates using hero

echo "updating [query/sql] templates"
rm -rf gen/query
hero -extensions .sql -source "query/sql" -pkgname query -dest gen/query

echo "updating [web/components] templates"
rm -rf gen/components
hero -extensions .html -source "web/components" -pkgname components -dest gen/components

echo "updating [web/transcript] templates"
rm -rf gen/transcripttemplates
hero -extensions .html -source "web/transcript" -pkgname transcripttemplates -dest gen/transcripttemplates

echo "updating [web/templates] templates"
rm -rf gen/templates
hero -extensions .html -source "web/templates" -pkgname templates -dest gen/templates

echo "updating [web/admin] templates"
rm -rf gen/admintemplates
hero -extensions .html -source "web/admin" -pkgname admintemplates -dest gen/admintemplates
