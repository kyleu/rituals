package team

import "github.com/google/uuid"

func (t Teams) TitleFor(id *uuid.UUID) string {
	if id == nil {
		return "-"
	}
	for _, x := range t {
		if x.ID == *id {
			return x.TitleString()
		}
	}
	return id.String()
}
