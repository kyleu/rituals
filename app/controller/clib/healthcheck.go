package clib

import (
	"net/http"

	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
)

func Healthcheck(w http.ResponseWriter, _ *http.Request) {
	x := util.ValueMap{"status": util.OK}
	_, _ = cutil.RespondJSON(cutil.NewWriteCounter(w), "", x)
}
