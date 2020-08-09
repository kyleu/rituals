package admin

import (
	"fmt"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"net/http"
	"strconv"

	"emperror.dev/errors"

	"github.com/gorilla/mux"
)

func MigrationList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Migration List"
		ctx.Breadcrumbs = adminBC(ctx, npncore.KeyMigration, npncore.Plural(npncore.KeyMigration))

		// params := npnweb.ParamSetFromRequest(r)
		// migrations := app.Database(ctx.App).ListMigrations(params.Get(npncore.KeyMigration, ctx.Logger))
		// return npncontroller.T(admintemplates.MigrationList(migrations, params, ctx, w))

		// TODO
		return "TODO", nil
	})
}

func MigrationDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		migrationIdxStr, ok := mux.Vars(r)[npncore.KeyIdx]
		if !ok {
			return npncontroller.EResp(errors.New("invalid migration id"))
		}
		migrationIdx, _ := strconv.ParseInt(migrationIdxStr, 10, 64)
		e := app.Database(ctx.App).GetMigrationByIdx(int(migrationIdx))
		if e == nil {
			msg := "can't load migration [" + migrationIdxStr + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(npncore.KeyMigration), w, r, ctx)
		}

		// params := npnweb.ParamSetFromRequest(r)

		title := fmt.Sprintf("Migration %v: %v", e.Idx, e.Title)
		ctx.Title = title
		bc := adminBC(ctx, npncore.KeyMigration, npncore.Plural(npncore.KeyMigration))
		idxStr := fmt.Sprintf("%v", e.Idx)
		bc = append(bc, npnweb.BreadcrumbSelf(idxStr))
		ctx.Breadcrumbs = bc

		// return npncontroller.T(admintemplates.MigrationDetail(e, params, ctx, w))

		// TODO
		return "TODO", nil
	})
}
