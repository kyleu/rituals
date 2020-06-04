package admin

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/kyleu/rituals.dev/gen/admintemplates"
	"net/http"
	"strconv"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

)

func MigrationList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Migration List"
		ctx.Breadcrumbs = adminBC(ctx, util.KeyMigration, util.Plural(util.KeyMigration))

		params := act.ParamSetFromRequest(r)
		migrations := ctx.App.Database.ListMigrations(params.Get(util.KeyMigration, ctx.Logger))
		return tmpl(admintemplates.MigrationList(migrations, params, ctx, w))
	})
}

func MigrationDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		migrationIdxStr, ok := mux.Vars(r)[util.KeyIdx]
		if !ok {
			return eresp(errors.New("invalid migration id"), "")
		}
		migrationIdx, _ := strconv.ParseInt(migrationIdxStr, 10, 64)
		e := ctx.App.Database.GetMigrationByIdx(int(migrationIdx))
		if e == nil {
			ctx.Session.AddFlash("error:Can't load migration [" + migrationIdxStr + "]")
			act.SaveSession(w, r, &ctx)
			return ctx.Route(util.AdminLink(util.KeyMigration)), nil
		}

		params := act.ParamSetFromRequest(r)

		title := fmt.Sprintf("Migration %v: %v", e.Idx, e.Title)
		ctx.Title = title
		bc := adminBC(ctx, util.KeyMigration, util.Plural(util.KeyMigration))
		link := util.AdminLink(util.KeyMigration, util.KeyDetail)
		idxStr := fmt.Sprintf("%v", e.Idx)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyIdx, idxStr), idxStr)...)
		ctx.Breadcrumbs = bc

		return tmpl(admintemplates.MigrationDetail(e, params, ctx, w))
	})
}
