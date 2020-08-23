package controllers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/kyleu/npn/npnasset"

	"github.com/kyleu/rituals.dev/app/assets"
)

const assetPath = "web/assets"

func Favicon(w http.ResponseWriter, r *http.Request) {
	data, hash, contentType, err := assets.Asset(assetPath, "/favicon.ico")
	npnasset.ZipResponse(w, r, data, hash, contentType, err)
}

func RobotsTxt(w http.ResponseWriter, r *http.Request) {
	data, hash, contentType, err := assets.Asset(assetPath, "/robots.txt")
	npnasset.ZipResponse(w, r, data, hash, contentType, err)
}

func Static(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(strings.TrimPrefix(r.URL.Path, "/assets"))
	if err == nil {
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		data, hash, contentType, err := assets.Asset(assetPath, path)
		npnasset.ZipResponse(w, r, data, hash, contentType, err)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
