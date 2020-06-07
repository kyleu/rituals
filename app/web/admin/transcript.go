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
		return tmpl(admintemplates.TranscriptList(transcript.AllTranscripts, ctx, w))
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
			return eresp(err, "error running transcript ["+key+"]")
		}

		switch format {
		case transcript.FormatJSON:
			return act.RespondJSON(w, content, ctx.Logger)
		default:
			ctx.Title = t.Title + " Transcript"
			bc := adminBC(ctx, util.KeyTranscript, util.Plural(util.KeyTranscript))
			bc = append(bc, web.Breadcrumb{Path: ctx.Route(util.AdminLink(util.KeyTranscript+".run"), util.KeyKey, key), Title: key})
			ctx.Breadcrumbs = bc

			return tmpl(admintemplates.TranscriptRun(t, param, content, ctx, w))
		}
	})
}
