package verifyuser

import (
	"crypto/rand"
	"crypto/rsa"
)

type UserMessage struct {
	Id      int    `json:"id"`
	Account string `json:"account"`
	Ip      string `json:"ip"`
	Ext     int64  `json:"ext"`
}

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func init() {
	var err error
	PrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// The public key is a part of the *rsa.PrivateKey struct
	PublicKey = &PrivateKey.PublicKey
}
