module github.com/kyleu/rituals.dev

go 1.14

require (
	emperror.dev/emperror v0.32.0
	emperror.dev/errors v0.7.0
	emperror.dev/handler/logur v0.4.0
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/websocket v1.4.1
	github.com/graphql-go/graphql v0.7.9
	github.com/jackc/pgx v3.6.2+incompatible // indirect
	github.com/johnfercher/maroto v0.27.0
	github.com/jonboulle/clockwork v0.2.0 // indirect
	github.com/kyleu/npn/npnasset v1.0.0
	github.com/kyleu/npn/npnconnection v1.0.0
	github.com/kyleu/npn/npncontroller v1.0.0
	github.com/kyleu/npn/npncore v1.0.0
	github.com/kyleu/npn/npndatabase v1.0.0
	github.com/kyleu/npn/npnexport v1.0.0
	github.com/kyleu/npn/npngraphql v1.0.0
	github.com/kyleu/npn/npnservice v1.0.0
	github.com/kyleu/npn/npntemplate v1.0.0
	github.com/kyleu/npn/npnuser v1.0.0
	github.com/kyleu/npn/npnweb v1.0.0
	github.com/mattn/go-runewidth v0.0.8 // indirect
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/russross/blackfriday v2.0.0+incompatible
	github.com/sagikazarmark/ocmux v0.2.0
	github.com/shiyanhui/hero v0.0.2
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/skip2/go-qrcode v0.0.0-20200526175731-7ac0b40b2038
	github.com/spf13/cobra v0.0.5
	golang.org/x/mobile v0.0.0-20200212152714-2b26a4705d24 // indirect
	logur.dev/logur v0.16.2
)

replace github.com/kyleu/npn/npnasset => ./npn/npnasset

replace github.com/kyleu/npn/npnconnection => ./npn/npnconnection

replace github.com/kyleu/npn/npncontroller => ./npn/npncontroller

replace github.com/kyleu/npn/npncore => ./npn/npncore

replace github.com/kyleu/npn/npndatabase => ./npn/npndatabase

replace github.com/kyleu/npn/npnexport => ./npn/npnexport

replace github.com/kyleu/npn/npngraphql => ./npn/npngraphql

replace github.com/kyleu/npn/npnservice => ./npn/npnservice

replace github.com/kyleu/npn/npntemplate => ./npn/npntemplate

replace github.com/kyleu/npn/npnuser => ./npn/npnuser

replace github.com/kyleu/npn/npnweb => ./npn/npnweb
