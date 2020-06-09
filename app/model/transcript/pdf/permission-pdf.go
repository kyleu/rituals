package pdf

import (
	pdfgen "github.com/johnfercher/maroto/pkg/pdf"
	"github.com/kyleu/rituals.dev/app/model/permission"
)

func renderPermissionList(permissions permission.Permissions, m pdfgen.Maroto) (string, error) {
	if len(permissions) > 0 {
		for _, p := range permissions {
			tr(func() {
				td(p.Message(), 12, m)
			}, 6, m)
		}
	}
	return "", nil
}
