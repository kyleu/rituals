package pdf

import (
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/permission"
)

func renderPermissionList(permissions permission.Permissions, m pdfgen.Maroto) {
	if len(permissions) > 0 {
		for _, p := range permissions {
			tp := p
			tr(func() {
				td(tp.Message(), 12, m)
			}, 6, m)
		}
	}
}
