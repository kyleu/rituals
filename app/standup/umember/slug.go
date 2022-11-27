package umember

import (
	"fmt"
)

func (s *StandupMember) PublicWebPath(uslug string) string {
	if uslug == "" {
		uslug = s.StandupID.String()
	}
	return fmt.Sprintf("/standup/%v#modal-member-%v", uslug, s.UserID.String())
}
