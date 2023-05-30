package verifyuser

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"log"
)

func DecodeUser(encryptedBytes []byte) UserMessage {
	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, PrivateKey, encryptedBytes, nil)
	if err != nil {
		log.Panic("用户数据解密错误", err)
	}
	var result UserMessage
	err = json.Unmarshal(decryptedBytes, &result)
	if err != nil {
		log.Panic("用户解密数据解析错误", err)
	}
	return result
}
