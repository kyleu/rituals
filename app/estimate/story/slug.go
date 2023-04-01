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

func (s Stories) Replace(st *Story) {
	for idx, x := range s {
		if x.ID == st.ID {
			s[idx] = st
			return
		}
	}
}
