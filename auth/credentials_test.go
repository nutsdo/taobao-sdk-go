package auth_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
)

func TestCredentials(t *testing.T) {
	c := auth.New("", "")
	u := c.AuthorizeUrlBuilder("http://127.0.0.1", "wap")
	fmt.Println(u)
}
