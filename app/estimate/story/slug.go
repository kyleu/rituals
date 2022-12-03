package story

import (
	"fmt"
)

func (s *Story) PublicWebPath(eslug string) string {
	if eslug == "" {
		eslug = s.EstimateID.String()
	}
	return fmt.Sprintf("/estimate/%v#modal-story-%v", eslug, s.ID.String())
}

func (s *Story) FinalVoteSafe() string {
	if s.FinalVote != "" {
		return s.FinalVote
	}
	return "-"
}
