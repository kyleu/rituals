package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/model/transcript"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/gorilla/mux"
)

func TranscriptList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = util.PluralTitle(util.KeyTranscript)
		ctx.Breadcrumbs = adminBC(ctx, util.KeyTranscript, util.Plural(util.KeyTranscript))
		return act.T(admintemplates.TranscriptList(transcript.AllTranscripts, ctx, w))
	})
}

func TranscriptRun(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		t := transcript.FromString(key)
		if t == nil {
			return "", util.IDError(util.KeyTranscript, key)
		}

		param := r.URL.Query().Get("param")

		formatStr := r.URL.Query().Get("format")
		if len(formatStr) == 0 {
			ct := act.GetContentType(r)
			if act.IsContentTypeJSON(ct) {
				formatStr = util.KeyJSON
			} else {
				formatStr = util.KeyHTML
			}
		}
		format := transcript.FormatFromString(formatStr)

		content, err := t.Resolve(ctx.App, ctx.Profile.UserID, param)
		if err != nil {
			return act.EResp(err, "error running transcript ["+key+"]")
		}

		return act.ExportTemplate(t, r.URL.Path, content, format, ctx, w)
	})
}
