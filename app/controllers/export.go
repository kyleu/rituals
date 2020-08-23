package controllers

import (
	"net/http"
	"strings"

	npnxls "github.com/kyleu/npn/npnexport/xls"
	"github.com/kyleu/rituals.dev/app/transcript/pdf"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	npnpdf "github.com/kyleu/npn/npnexport/pdf"
	"github.com/kyleu/npn/npntemplate/gen/npntemplate"
	"github.com/kyleu/npn/npnweb"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/app/transcript/xls"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/gen/transcripttemplates"
)

type ExportParams struct {
	ModelID *uuid.UUID
	Slug    string
	Title   string
	Path    string
	PermSvc *permission.Service
}

func ExportAct(svc util.Service, f func(string, *npnweb.RequestContext) ExportParams, w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
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
			return npncontroller.EResp(err, "cannot load transcript")
		}

		return ExportTemplate(t, p.Path, rsp, format, ctx, w)
	})
}

func exportQueryParams(svc util.Service, r *http.Request) (util.Service, string, transcript.Format) {
	key := mux.Vars(r)[npncore.KeyKey]
	fmtStr := mux.Vars(r)[npncore.KeyFmt]
	format := transcript.FormatFromString(fmtStr)
	return svc, key, format
}

func exportErr(svc *util.Service, key string, ctx *npnweb.RequestContext, w http.ResponseWriter, r *http.Request) (string, error) {
	msg := "can't load " + svc.Key + " [" + key + "]"
	return npncontroller.FlashAndRedir(false, msg, svc.Key+".list", w, r, ctx)
}

func ExportTemplate(t *transcript.Transcript, path string, rsp interface{}, fmt transcript.Format, ctx *npnweb.RequestContext, w http.ResponseWriter) (string, error) {
	switch fmt {
	case transcript.FormatJSON:
		fn := strings.Split(path, "/")
		return npncontroller.RespondJSON(w, fn[len(fn)-1], rsp, ctx.Logger)
	case transcript.FormatPDF:
		filename, ba, err := npnpdf.Render(rsp, ctx.App.Auth().FullURL(path), pdf.RenderCallback)
		if err != nil {
			return "", errors.Wrap(err, "unable to render pdf")
		}
		return npncontroller.RespondMIME(filename, "application/pdf", "pdf", ba, w)
	case transcript.FormatExcel:
		filename, ba, err := npnxls.Render(rsp, ctx.App.Auth().FullURL(path), xls.RenderResponse)
		if err != nil {
			return "", errors.Wrap(err, "unable to render Excel")
		}
		return npncontroller.RespondMIME(filename, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "xlsx", ba, w)
	case transcript.FormatPrint:
		tx := &transcript.Context{UserID: ctx.Profile.UserID, App: ctx.App, Logger: ctx.App.Logger(), Routes: ctx.Routes}
		return npncontroller.T(transcripttemplates.Print(t, rsp, tx, w))
	default:
		return npncontroller.T(npntemplate.StaticMessage("TODO", ctx, w))
	}
}

func exportCheck(params *PermissionParams, fmt transcript.Format, ctx *npnweb.RequestContext, permSvc *permission.Service, w http.ResponseWriter) (string, error) {
	auths, permErrors, bc := CheckPerms(ctx, permSvc, params)
	url := ctx.Route(params.Svc.Key+"."+npncore.KeyExport, npncore.KeyKey, params.Slug, npncore.KeyFmt, fmt.Key)
	bc = append(bc, npnweb.BreadcrumbsSimple(url, fmt.Key)...)
	ctx.Breadcrumbs = bc
	if len(permErrors) > 0 {
		return PermErrorTemplate(params.Svc, permErrors, auths, ctx, w)
	}
	ctx.Title = params.Title
	return "", nil
}
