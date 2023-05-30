package verifyuser

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"

	"github.com/goccy/go-json"
)

func EncodeUser(message UserMessage) []byte {
	cry, err := json.Marshal(message)
	if err != nil {
		log.Panic("用户加密信息解析错误", err)
	}
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, PublicKey, cry, nil)
	if err != nil {
		log.Panic("用户信息加密错误", err)
	}
	return encryptedBytes
}
