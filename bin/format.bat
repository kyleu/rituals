@ECHO OFF
rem Content managed by Project Forge, see [projectforge.md] for details.

rem Formatting code from all projects

cd %~dpnx0\..

@ECHO ON
echo "=== formatting ==="
gofumpt -w $(find . -type f -name "*.go" | grep -v .html.go | grep -v .sql.go)
