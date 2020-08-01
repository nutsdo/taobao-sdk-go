package auth_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
)

func TestAuthorize(t *testing.T) {
	c := auth.New("28144153", "181f90dee3bf06f1c9553f776b52dd97")
	data, _ := c.RefreshToken("62020180e3eb4a28549d010494ZZ4a1dfd26c644cb2e8ef1866605434", "wap")
	fmt.Println(string(data))
}
