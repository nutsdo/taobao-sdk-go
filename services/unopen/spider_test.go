package unopen_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/nutsdo/taobao-sdk-go/services/unopen"
)

func TestParseTpwd(t *testing.T) {
	fmt.Println(time.Now())
	unopen.ParseTpwd("￥CsJO1hNkyjR￥")
}