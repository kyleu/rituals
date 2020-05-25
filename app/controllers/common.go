package controllers

import (
	"github.com/gofrs/uuid"
	"net/url"
)

func tmpl(_ int, err error) (string, error) {
	return "", err
}

func getUUID(m url.Values, key string) *uuid.UUID {
	retString := m.Get(key)
	var retID *uuid.UUID
	if retString != "" {
		s, err := uuid.FromString(retString)
		if err == nil {
			retID = &s
		}
	}
	return retID
}
