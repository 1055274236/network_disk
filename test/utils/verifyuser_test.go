package text_utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	"NetworkDisk/utils/verifyuser"
)

func TestEncodeUser(t *testing.T) {
	result := verifyuser.EncodeUser(verifyuser.UserMessage{Id: 123, Account: "asd", Ip: "sa", Ext: 456})
	fmt.Println(base64.StdEncoding.EncodeToString(result))
}

func TestDecodeUser(t *testing.T) {
	encryptedBytes := verifyuser.EncodeUser(verifyuser.UserMessage{Id: 123, Account: "asd", Ip: "sa", Ext: 456})
	result := verifyuser.DecodeUser(encryptedBytes)
	value, _ := json.Marshal(result)
	fmt.Println(string(value))
}
