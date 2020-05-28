package auth_test

import (
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
)

func TestAuthorize(t *testing.T) {
	c := auth.New("", "")
	c.Authorize("http://127.0.0.1", "wap")
}