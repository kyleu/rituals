package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/app/web/act"
	"github.com/skip2/go-qrcode"
)

func QRCode(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		q := r.URL.Query()
		path := q.Get("path")
		if len(path) == 0 {
			return act.ENew("must provide path")
		}

		url := ctx.App.Auth.FullURL(path)
		bytes, err := qrcode.Encode(url, qrcode.Medium, 256)
		if err != nil {
			return act.EResp(err, "can't create QR code")
		}
		return act.RespondMIME("", "image/png", "png", bytes, w)
	})
}
