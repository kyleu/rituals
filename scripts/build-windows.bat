REM https://git-scm.com/download/win
REM  https://golang.org/doc/install
REM  http://tdm-gcc.tdragon.net/download

hero -extensions .html -source web/templates -pkgname templates -dest internal/gen/templates
hero -extensions .sql -pkgname queries -source queries -dest internal/gen/queries

go-embed -input web/assets -output internal/app/controllers/assets/assets.go

go build -o build/ ./cmd/...
