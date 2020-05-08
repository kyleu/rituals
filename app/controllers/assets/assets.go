package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"mime"
	"path/filepath"

	"emperror.dev/errors"
)

func init() {
	_ = mime.AddExtensionType(".ico", "image/x-icon")
	_ = mime.AddExtensionType(".eot", "font/eot")
	_ = mime.AddExtensionType(".tff", "font/tff")
	_ = mime.AddExtensionType(".woff", "application/font-woff")
	_ = mime.AddExtensionType(".woff2", "application/font-woff")
}

func Asset(base, path string) ([]byte, string, string, error) {
	var b bytes.Buffer
	file := base + path
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, "", "", errors.WithStack(errors.Wrap(err, "error reading asset ["+path+"]"))
	}
	if data != nil {
		w := gzip.NewWriter(&b)
		_, _ = w.Write(data)
		_ = w.Close()
		data = b.Bytes()
	}
	sum := md5.Sum(data)
	return data, hex.EncodeToString(sum[1:]), mime.TypeByExtension(filepath.Ext(file)), nil
}
