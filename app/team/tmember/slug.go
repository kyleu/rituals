package tmember

import "fmt"

func (t *TeamMember) PublicWebPath(tslug string) string {
	if tslug == "" {
		tslug = t.TeamID.String()
	}
	return fmt.Sprintf("/team/%v#modal-member-%v", tslug, t.UserID.String())
}
