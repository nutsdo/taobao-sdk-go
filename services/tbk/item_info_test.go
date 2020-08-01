package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestItemInfo(t *testing.T) {

	c := tbk.NewClientWith(credential)

	req := tbk.NewItemInfoRequest()
	req.SetNumIids("532901398400")

	fmt.Printf("values: %#v\n", req.GetValues())
	resp, err := c.GetItemInfo(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}
	fmt.Printf("正常响应:%#v", resp)

}
