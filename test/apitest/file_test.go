package apitest

import (
	"NetworkDisk/utils/httptestutils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileStatus(t *testing.T) {
	result := httptestutils.Get("/status/file", map[string]string{"md5": "1", "sha1": "1"})
	fmt.Println(result.Body.String())
	assert.Equal(t, 404, result.Code)
}
