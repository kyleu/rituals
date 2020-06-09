package admin

import (
	"fmt"
	"net/http"
	"strconv"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"
)

func MigrationList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "Migration List"
		ctx.Breadcrumbs = adminBC(ctx, util.KeyMigration, util.Plural(util.KeyMigration))

		params := act.ParamSetFromRequest(r)
		migrations := ctx.App.Database.ListMigrations(params.Get(util.KeyMigration, ctx.Logger))
		return act.T(admintemplates.MigrationList(migrations, params, ctx, w))
	})
}

func MigrationDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		migrationIdxStr, ok := mux.Vars(r)[util.KeyIdx]
		if !ok {
			return act.EResp(errors.New("invalid migration id"))
		}
		migrationIdx, _ := strconv.ParseInt(migrationIdxStr, 10, 64)
		e := ctx.App.Database.GetMigrationByIdx(int(migrationIdx))
		if e == nil {
			msg := "can't load migration [" + migrationIdxStr + "]"
			return act.FlashAndRedir(false, msg, util.AdminLink(util.KeyMigration), w, r, ctx)
		}

		params := act.ParamSetFromRequest(r)

		title := fmt.Sprintf("Migration %v: %v", e.Idx, e.Title)
		ctx.Title = title
		bc := adminBC(ctx, util.KeyMigration, util.Plural(util.KeyMigration))
		link := util.AdminLink(util.KeyMigration, util.KeyDetail)
		idxStr := fmt.Sprintf("%v", e.Idx)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyIdx, idxStr), idxStr)...)
		ctx.Breadcrumbs = bc

		return act.T(admintemplates.MigrationDetail(e, params, ctx, w))
	})
}
