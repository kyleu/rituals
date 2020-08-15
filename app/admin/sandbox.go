package admin

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/sandbox"

	"github.com/gorilla/mux"
)

func SandboxList(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = npncore.PluralTitle(npncore.KeySandbox)
		ctx.Breadcrumbs = npncontroller.AdminBC(ctx, npncore.KeySandbox, npncore.Plural(npncore.KeySandbox))
		return npncontroller.T(admintemplates.SandboxList(sandbox.AllSandboxes, ctx, w))
	})
}

func SandboxRun(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		sb := sandbox.FromString(key)
		if sb == nil {
			return "", npncore.IDError(npncore.KeySandbox, key)
		}
		content, rsp, err := sb.Resolve(ctx)
		if err != nil {
			return npncontroller.EResp(err, "error running sandbox ["+key+"]")
		}

		ctx.Title = sb.Title + " Sandbox"
		bc := npncontroller.AdminBC(ctx, npncore.KeySandbox, npncore.Plural(npncore.KeySandbox))
		bc = append(bc, npnweb.BreadcrumbSelf(key))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.SandboxRun(sb, content, npncore.ToJSON(rsp, ctx.Logger), ctx, w))
	})
}
