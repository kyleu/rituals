package controllers

import (
	"github.com/kyleu/rituals.dev/app/model/transcript/pdf"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/model/transcript"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/gen/templates"
	"github.com/kyleu/rituals.dev/gen/transcripttemplates"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
)

func ExportAct(svc util.Service, f func(string, *web.RequestContext) (*uuid.UUID, string, string, *permission.Service), w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		_, key, format := exportQueryParams(svc, r)

		modelID, slug, title, permSvc := f(key, ctx)
		if modelID == nil {
			return exportErr(&svc, key, ctx, w, r)
		}

		params := &PermissionParams{Svc: svc, ModelID: *modelID, Slug: slug, Title: title}
		errRsp, err := exportCheck(params, format, ctx, permSvc, w)
		if err != nil {
			return errRsp, err
		}

		t := transcript.FromString(svc.Key)
		rsp, err := t.Resolve(ctx.App, ctx.Profile.UserID, slug)
		if err != nil {
			return eresp(err, "cannot load transcript")
		}

		return exportTemplate(t, rsp, format, ctx, w)
	})
}

func exportQueryParams(svc util.Service, r *http.Request) (util.Service, string, transcript.Format) {
	key := mux.Vars(r)[util.KeyKey]
	fmtStr := mux.Vars(r)[util.KeyFmt]
	format := transcript.FormatFromString(fmtStr)
	return svc, key, format
}

func exportErr(svc *util.Service, key string, ctx *web.RequestContext, w http.ResponseWriter, r *http.Request) (string, error) {
	msg := "can't load " + svc.Key + " [" + key + "]"
	return act.FlashAndRedir(false, msg, svc.Key+".list", w, r, ctx)
}

func exportCheck(params *PermissionParams, fmt transcript.Format, ctx *web.RequestContext, permSvc *permission.Service, w http.ResponseWriter) (string, error) {
	auths, permErrors, bc := check(ctx, permSvc, params)
	url := ctx.Route(params.Svc.Key+"."+util.KeyExport, util.KeyKey, params.Slug, util.KeyFmt, fmt.Key)
	bc = append(bc, web.BreadcrumbsSimple(url, fmt.Key)...)
	ctx.Breadcrumbs = bc
	if len(permErrors) > 0 {
		return permErrorTemplate(params.Svc, permErrors, auths, ctx, w)
	}
	ctx.Title = params.Title
	return "", nil
}

func exportTemplate(t *transcript.Transcript, rsp interface{}, fmt transcript.Format, ctx *web.RequestContext, w http.ResponseWriter) (string, error) {
	switch fmt {
	case transcript.FormatJSON:
		return act.RespondJSON(w, rsp, ctx.Logger)
	case transcript.FormatPDF:
		ba, err := pdf.Render(rsp)
		if err != nil {
			return eresp(err, "unable to render pdf")
		}
		return act.RespondPDF(w, ba)
	case transcript.FormatPrint:
		tx := &transcript.Context{UserID: ctx.Profile.UserID, App: ctx.App, Logger: ctx.App.Logger, Routes: ctx.Routes}
		return tmpl(transcripttemplates.Print(t, rsp, tx, w))
	default:
		return tmpl(templates.StaticMessage("TODO", ctx, w))
	}
}
