package act

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/model/transcript"
	"github.com/kyleu/rituals.dev/app/model/transcript/pdf"
	"github.com/kyleu/rituals.dev/app/model/transcript/xls"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/gen/templates"
	"github.com/kyleu/rituals.dev/gen/transcripttemplates"
	"net/http"
	"strings"
)

type ExportParams struct {
	ModelID *uuid.UUID
	Slug    string
	Title   string
	Path    string
	PermSvc *permission.Service
}

func ExportAct(svc util.Service, f func(string, *web.RequestContext) ExportParams, w http.ResponseWriter, r *http.Request) {
	Act(w, r, func(ctx *web.RequestContext) (string, error) {
		_, key, format := exportQueryParams(svc, r)

		p := f(key, ctx)
		if p.ModelID == nil {
			return exportErr(&svc, key, ctx, w, r)
		}

		params := &PermissionParams{Svc: svc, ModelID: *p.ModelID, Slug: p.Slug, Title: p.Title}
		errRsp, err := exportCheck(params, format, ctx, p.PermSvc, w)
		if err != nil {
			return errRsp, err
		}

		t := transcript.FromString(svc.Key)
		rsp, err := t.Resolve(ctx.App, ctx.Profile.UserID, p.Slug)
		if err != nil {
			return EResp(err, "cannot load transcript")
		}

		return ExportTemplate(t, p.Path, rsp, format, ctx, w)
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
	return FlashAndRedir(false, msg, svc.Key+".list", w, r, ctx)
}

func ExportTemplate(t *transcript.Transcript, path string, rsp interface{}, fmt transcript.Format, ctx *web.RequestContext, w http.ResponseWriter) (string, error) {
	switch fmt {
	case transcript.FormatJSON:
		fn := strings.Split(path, "/")
		return RespondJSON(w, fn[len(fn) - 1], rsp, ctx.Logger)
	case transcript.FormatPDF:
		filename, ba, err := pdf.Render(rsp, ctx.App.Auth.FullURL(path))
		if err != nil {
			return "", errors.Wrap(err, "unable to render pdf")
		}
		return RespondMIME(filename, "application/pdf", "pdf", ba, w)
	case transcript.FormatExcel:
		filename, ba, err := xls.Render(rsp, ctx.App.Auth.FullURL(path))
		if err != nil {
			return "", errors.Wrap(err, "unable to render Excel")
		}
		return RespondMIME(filename, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "xlsx", ba, w)
	case transcript.FormatPrint:
		tx := &transcript.Context{UserID: ctx.Profile.UserID, App: ctx.App, Logger: ctx.App.Logger, Routes: ctx.Routes}
		return T(transcripttemplates.Print(t, rsp, tx, w))
	default:
		return T(templates.StaticMessage("TODO", ctx, w))
	}
}

func exportCheck(params *PermissionParams, fmt transcript.Format, ctx *web.RequestContext, permSvc *permission.Service, w http.ResponseWriter) (string, error) {
	auths, permErrors, bc := CheckPerms(ctx, permSvc, params)
	url := ctx.Route(params.Svc.Key+"."+util.KeyExport, util.KeyKey, params.Slug, util.KeyFmt, fmt.Key)
	bc = append(bc, web.BreadcrumbsSimple(url, fmt.Key)...)
	ctx.Breadcrumbs = bc
	if len(permErrors) > 0 {
		return PermErrorTemplate(params.Svc, permErrors, auths, ctx, w)
	}
	ctx.Title = params.Title
	return "", nil
}
