package form

import (
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"logur.dev/logur"
	"net/http"
	"strings"
)

type ProfileForm struct {
	Username  string `mapstructure:"username"`
	Theme     string `mapstructure:"theme"`
	LinkColor string `mapstructure:"linkColor"`
	NavColor  string `mapstructure:"navColor"`
}

type ConnectionForm struct {
	Svc   string `mapstructure:"svc"`
	Cmd   string `mapstructure:"cmd"`
	Param string `mapstructure:"param"`
}

func Decode(r *http.Request, tgt interface{}, logger logur.Logger) error {
	_ = r.ParseForm()
	frm := make(map[string]interface{}, len(r.Form))
	for k, v := range r.Form {
		frm[k] = strings.Join(v, "||")
	}
	md := &mapstructure.Metadata{}
	err := mapstructure.DecodeMetadata(frm, tgt, md)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to parse [%T] form", tgt))
	}

	if logger != nil {
		msg := fmt.Sprintf("Parsed [%T] form with unused keys [%v]", tgt, strings.Join(md.Unused, ", "))
		logger.Warn(msg)
		bytes, _ := json.Marshal(tgt)
		logger.Warn(string(bytes))
	}
	return nil
}
