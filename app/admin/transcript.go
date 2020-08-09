package admin

import (
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app/web"
	"net/http"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/transcript"

	"github.com/gorilla/mux"
)

func TranscriptList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = npncore.PluralTitle(util.KeyTranscript)
		ctx.Breadcrumbs = adminBC(ctx, util.KeyTranscript, npncore.Plural(util.KeyTranscript))
		return npncontroller.T(admintemplates.TranscriptList(transcript.AllTranscripts, ctx, w))
	})
}

func TranscriptRun(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		t := transcript.FromString(key)
		if t == nil {
			return "", npncore.IDError(util.KeyTranscript, key)
		}

		param := r.URL.Query().Get("param")

		formatStr := r.URL.Query().Get("format")
		if len(formatStr) == 0 {
			ct := npncontroller.GetContentType(r)
			if npncontroller.IsContentTypeJSON(ct) {
				formatStr = npncore.KeyJSON
			} else {
				formatStr = npncore.KeyHTML
			}
		}
		format := transcript.FormatFromString(formatStr)

		content, err := t.Resolve(ctx.App, ctx.Profile.UserID, param)
		if err != nil {
			return npncontroller.EResp(err, "error running transcript ["+key+"]")
		}

		return web.ExportTemplate(t, r.URL.Path, content, format, ctx, w)
	})
}
