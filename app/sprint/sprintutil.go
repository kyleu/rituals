package sprint

import "github.com/google/uuid"

func (s Sprints) TitleFor(id *uuid.UUID) string {
	if id == nil {
		return "-"
	}
	for _, x := range s {
		if x.ID == *id {
			return x.TitleString()
		}
	}
	return id.String()
}
