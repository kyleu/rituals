package emember

import (
	"fmt"
)

func (e *EstimateMember) PublicWebPath(eslug string) string {
	if eslug == "" {
		eslug = e.EstimateID.String()
	}
	return fmt.Sprintf("/estimate/%v#modal-member-%v", eslug, e.UserID.String())
}
