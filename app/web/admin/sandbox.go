package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/model/sandbox"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/gorilla/mux"
)

func SandboxList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = util.PluralTitle(util.KeySandbox)
		ctx.Breadcrumbs = adminBC(ctx, util.KeySandbox, util.Plural(util.KeySandbox))
		return tmpl(admintemplates.SandboxList(sandbox.AllSandboxes, ctx, w))
	})
}

func SandboxRun(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sb := sandbox.FromString(key)
		if sb == nil {
			return "", util.IDError(util.KeySandbox, key)
		}
		content, rsp, err := sb.Resolve(ctx)
		if err != nil {
			return eresp(err, "error running sandbox ["+key+"]")
		}

		ctx.Title = sb.Title + " Sandbox"
		bc := adminBC(ctx, util.KeySandbox, util.Plural(util.KeySandbox))
		bc = append(bc, web.Breadcrumb{Path: ctx.Route(util.AdminLink(util.KeySandbox+".run"), util.KeyKey, key), Title: key})
		ctx.Breadcrumbs = bc

		return tmpl(admintemplates.SandboxRun(sb, content, util.ToJSON(rsp, ctx.Logger), ctx, w))
	})
}
