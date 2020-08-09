package controllers

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnweb"

	"github.com/skip2/go-qrcode"
)

func QRCode(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		q := r.URL.Query()
		path := q.Get("path")
		if len(path) == 0 {
			return npncontroller.ENew("must provide path")
		}

		url := ctx.App.Auth().FullURL(path)
		bytes, err := qrcode.Encode(url, qrcode.Medium, 256)
		if err != nil {
			return npncontroller.EResp(err, "can't create QR code")
		}
		return npncontroller.RespondMIME("", "image/png", "png", bytes, w)
	})
}
