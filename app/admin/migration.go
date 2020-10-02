package admin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"emperror.dev/errors"

	"github.com/gorilla/mux"
)

func MigrationList(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Migration List"
		ctx.Breadcrumbs = npncontroller.AdminBC(ctx, npncore.KeyMigration, npncore.Plural(npncore.KeyMigration))

		params := npnweb.ParamSetFromRequest(r)
		migrations := app.Svc(ctx.App).Database.ListMigrations(params.Get(npncore.KeyMigration, ctx.Logger))
		return npncontroller.T(admintemplates.MigrationList(migrations, params, ctx, w))
	})
}

func MigrationDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		migrationIdxStr, ok := mux.Vars(r)[npncore.KeyIdx]
		if !ok {
			return npncontroller.EResp(errors.New("invalid migration id"))
		}
		migrationIdx, _ := strconv.ParseInt(migrationIdxStr, 10, 64)
		e := app.Svc(ctx.App).Database.GetMigrationByIdx(int(migrationIdx))
		if e == nil {
			msg := "can't load migration [" + migrationIdxStr + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(npncore.KeyMigration), w, r, ctx)
		}

		params := npnweb.ParamSetFromRequest(r)

		title := fmt.Sprintf("Migration %v: %v", e.Idx, e.Title)
		ctx.Title = title
		bc := npncontroller.AdminBC(ctx, npncore.KeyMigration, npncore.Plural(npncore.KeyMigration))
		idxStr := fmt.Sprint(e.Idx)
		bc = append(bc, npnweb.BreadcrumbSelf(idxStr))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.MigrationDetail(e, params, ctx, w))
	})
}
