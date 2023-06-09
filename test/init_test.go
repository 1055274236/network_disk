package test

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"testing"
	"time"
)

func TestT(t *testing.T) {
	p := path.Join("file", "path", strconv.FormatInt(time.Now().UnixMilli(), 10))
	err := os.MkdirAll(p, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
