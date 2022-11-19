package story

import (
	"fmt"
)

func (s *Story) PublicWebPath() string {
	return fmt.Sprintf("/estimate/%v#modal-story-%v", s.EstimateID.String(), s.ID.String())
}
