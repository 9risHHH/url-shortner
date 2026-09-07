package storage

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateShortID() (string, error) {
	b := make([]byte, 5)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	id := base64.URLEncoding.EncodeToString(b)
	return id[:6], nil
}
