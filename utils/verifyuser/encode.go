package verifyuser

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	"github.com/goccy/go-json"
)

func EncodeUser(message UserMessage) ([]byte, error) {
	cry, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, PublicKey, cry, nil)
	if err != nil {
		return nil, err
	}
	return encryptedBytes, nil
}
