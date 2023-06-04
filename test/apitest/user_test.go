package apitest

import (
	"NetworkDisk/utils/httptestutils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginApi(t *testing.T) {

	result := httptestutils.PostForm("/login", map[string]string{"account": "a1", "password": "123456"})
	fmt.Println(result.Body.String())
	assert.Equal(t, 200, result.Code)
}

func TestSignInApi(t *testing.T) {
	result := httptestutils.PostForm("/signin", map[string]string{"account": "a1", "password": "123456"})
	fmt.Println(result.Body.String())
	assert.Equal(t, 400, result.Code)
}
