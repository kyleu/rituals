@ECHO OFF

rem Builds all the templates using quicktemplate

cd %~dp0\..

@ECHO ON
qtc -ext sql -dir "queries"
qtc -ext html -dir "views"
