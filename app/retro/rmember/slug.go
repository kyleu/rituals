package rmember

import (
	"fmt"
)

func (r *RetroMember) PublicWebPath(rslug string) string {
	if rslug == "" {
		rslug = r.RetroID.String()
	}
	return fmt.Sprintf("/retro/%v#modal-member-%v", rslug, r.UserID.String())
}
