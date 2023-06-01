package verifyuser

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
)

func DecodeUser(encryptedBytes []byte) (UserMessage, error) {
	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, PrivateKey, encryptedBytes, nil)
	if err != nil {
		return UserMessage{}, err
	}
	var result UserMessage
	err = json.Unmarshal(decryptedBytes, &result)
	if err != nil {
		return UserMessage{}, err
	}
	return result, nil
}
