package helpers_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/services/helpers"
)

func TestIdentifyTpwd(t *testing.T) {

	tkl, err := helpers.IdentifyTpwd("$KEX21cd6xqg₴")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("识别后的淘口令:", tkl)
}
