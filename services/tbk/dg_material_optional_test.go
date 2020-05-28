package tbk_test

import (
	"fmt"
	"testing"

	"github.com/nutsdo/taobao-sdk-go/auth"
	"github.com/nutsdo/taobao-sdk-go/services/tbk"
)

func TestDgMaterialOptional(t *testing.T) {

	credential := &auth.Credentials{AppKey: "", AppSecret: []byte("")}
	c := tbk.NewClientWith(credential)

	req := tbk.NewDgMaterialOptionalRequest()
	req.SetAdzoneId("52280450373")
	req.SetPageSize("5")
	req.SetQ("https://detail.tmall.com/item.htm?id=612887993968")

	fmt.Printf("values: %#v\n", req.GetValues())
	resp,err := c.DgMaterialOptional(req)
	if err != nil {
		fmt.Printf("错误响应:%v", err)

		return
	}

	fmt.Printf("正常响应:%#v", resp)

}
