package controllers

import (
	"emperror.dev/emperror"
	"emperror.dev/errors"
	"github.com/kyleu/npn/npncontroller"
	"net/http"
	"path/filepath"
	"strings"


	"github.com/kyleu/rituals.dev/app/assets"
)

const assetPath = "web/assets"

func Favicon(w http.ResponseWriter, r *http.Request) {
	data, hash, contentType, err := assets.Asset(assetPath, "/favicon.ico")
	ZipResponse(w, r, data, hash, contentType, err)
}

func RobotsTxt(w http.ResponseWriter, r *http.Request) {
	data, hash, contentType, err := assets.Asset(assetPath, "/robots.txt")
	ZipResponse(w, r, data, hash, contentType, err)
}

func Static(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(strings.TrimPrefix(r.URL.Path, "/assets"))
	if err == nil {
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		data, hash, contentType, err := assets.Asset(assetPath, path)
		ZipResponse(w, r, data, hash, contentType, err)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func ZipResponse(w http.ResponseWriter, r *http.Request, data []byte, hash string, contentType string, err error) {
	if err == nil {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", contentType)
		// w.Header().Add("Cache-Control", "public, max-age=31536000")
		w.Header().Add("ETag", hash)
		if r.Header.Get("If-None-Match") == hash {
			w.WriteHeader(http.StatusNotModified)
		} else {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(data)
			emperror.Panic(errors.Wrap(err, "unable to write to response"))
		}
	} else {
		npncontroller.NotFound(w, r)
	}
}
