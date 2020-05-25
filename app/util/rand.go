package util

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"math/rand"
	"time"
)

func UUID() uuid.UUID {
	ret, err := uuid.NewV4()
	if err != nil {
		panic(errors.New("unable to create random UUID"))
	}

	return ret
}

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
