package report

import "fmt"

func (r *Report) PublicWebPath(uslug string) string {
	if uslug == "" {
		uslug = r.StandupID.String()
	}
	return fmt.Sprintf("/standup/%v#modal-report-%v", uslug, r.ID.String())
}
