package test

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	value := map[string]string{
		"asd": "qwe",
		"qwe": "zxc",
	}
	for index, item := range value {
		fmt.Println(index, item)
	}
}
