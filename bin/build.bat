go-embed -input ../npn/npnasset/vendor -output ../npn/npnasset/assets/assets.go
go-embed -input web/assets -output app/assets/assets.go

md build
md build\release

go build -o build/release github.com/kyleu/rituals.dev/cmd/rituals

git checkout app/assets/assets.go
git checkout ../npn/npnasset/assets/assets.go
