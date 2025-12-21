package smember

import "fmt"

func (s *SprintMember) PublicWebPath(sslug string) string {
	if sslug == "" {
		sslug = s.SprintID.String()
	}
	return fmt.Sprintf("/sprint/%v#modal-member-%v", sslug, s.UserID.String())
}
